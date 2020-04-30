package http

import (
	"2020_1_Color_noise/internal/models"
	"2020_1_Color_noise/internal/pkg/comment"
	"2020_1_Color_noise/internal/pkg/error"
	"2020_1_Color_noise/internal/pkg/metric"
	"2020_1_Color_noise/internal/pkg/response"
	"github.com/asaskevich/govalidator"
	"github.com/gorilla/mux"
	"github.com/mailru/easyjson"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type Handler struct {
	commentUsecase comment.IUsecase
	logger         *zap.SugaredLogger
}

func NewHandler(usecase comment.IUsecase , logger *zap.SugaredLogger) *Handler {
	return &Handler{
		commentUsecase: usecase,
		logger:			logger,
	}
}

func (ch *Handler) Create(w http.ResponseWriter, r *http.Request) {
	path := "/api/comment/post"
	metric.Increase()
	reqId := r.Context().Value("ReqId")

	isAuth := r.Context().Value("IsAuth")
	if isAuth != true {
		err := error.Unauthorized.New("Create comment: user is unauthorized")
		error.ErrorHandler(w, ch.logger, reqId, error.Wrapf(err, "request id: %s", reqId), path)
		return
	}

	userId, ok := r.Context().Value("Id").(uint)
	if !ok {
		err := error.NoType.New("Received bad id from context")
		error.ErrorHandler(w, ch.logger, reqId, error.Wrapf(err, "request id: %s", reqId), path)
		return
	}

	input := &models.InputComment{}

	err := easyjson.UnmarshalFromReader(r.Body, input)
	if err != nil {
		err = error.WithMessage(error.BadRequest.Wrap(err, "Decoding error during creation comment"), "Wrong body of request")
		error.ErrorHandler(w, ch.logger, reqId, error.Wrapf(err, "request id: %s", reqId), path)
		return
	}

	_, err = govalidator.ValidateStruct(input)
	if err != nil {
		err = error.WithMessage(error.BadRequest.Wrapf(err, "request id: %s", "5"),
			"Text shouldn't be empty and longer 2000 characters.")
		error.ErrorHandler(w, ch.logger, reqId, error.Wrapf(err, "request id: %s", reqId), path)
		return
	}

	id, err := ch.commentUsecase.Create(input, userId)
	if err != nil {
		error.ErrorHandler(w, ch.logger, reqId, error.Wrapf(err, "request id: %s", reqId), path)
		return
	}

	resp := &models.ResponseComment{
		Id: id,
	}

	response.Respond(w, http.StatusCreated, resp, path)
}

func (ch *Handler) GetComment(w http.ResponseWriter, r *http.Request) {
	path := "/api/comment/get"
	metric.Increase()
	reqId := r.Context().Value("ReqId")

	isAuth := r.Context().Value("IsAuth")
	if isAuth != true {
		err := error.Unauthorized.New("Get comment: user is unauthorized")
		error.ErrorHandler(w, ch.logger, reqId, error.Wrapf(err, "request id: %s", reqId), path)
		return
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		err = error.WithMessage(error.BadRequest.Wrap(err, "Bad id in during getting a comment"), "Bad id")
		error.ErrorHandler(w, ch.logger, reqId, error.Wrapf(err, "request id: %s", reqId), path)
		return
	}

	comment, err := ch.commentUsecase.GetById(uint(id))
	if err != nil {
		error.ErrorHandler(w, ch.logger, reqId, error.Wrapf(err, "request id: %s", reqId), path)
		return
	}

	resp := &models.ResponseComment{
		Id:        comment.Id,
		UserId:    comment.UserId,
		PindId:    comment.PinId,
		Text:      comment.Text,
		CreatedAt: &comment.CreatedAt,
	}

	response.Respond(w, http.StatusOK, resp, path)
}

func (ch *Handler) Fetch(w http.ResponseWriter, r *http.Request) {
	path := "/api/comment/pin/get"
	metric.Increase()
	reqId := r.Context().Value("ReqId")

	isAuth := r.Context().Value("IsAuth")
	if isAuth != true {
		err := error.Unauthorized.New("Fetch comment: user is unauthorized")
		error.ErrorHandler(w, ch.logger, reqId, error.Wrapf(err, "request id: %s", reqId), path)
		return
	}

	vars := mux.Vars(r)
	pinId, err := strconv.Atoi(vars["id"])
	if err != nil {
		err = error.WithMessage(error.BadRequest.Wrap(err, "Bad id when in during getting comments for pin"), "Bad id")
		error.ErrorHandler(w, ch.logger, reqId, error.Wrapf(err, "request id: %s", reqId), path)
		return
	}

	start, _ := strconv.Atoi(r.URL.Query().Get("start"))

	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		limit = 10
	}

	comments, err := ch.commentUsecase.GetByPinId(uint(pinId), start, limit)
	if err != nil {
		error.ErrorHandler(w, ch.logger, reqId, error.Wrapf(err, "request id: %s", reqId), path)
		return
	}

	resp := make([]models.ResponseComment, 0)

	for _, comment := range comments {
		resp = append(resp, models.ResponseComment{
			Id:        comment.Id,
			UserId:    comment.UserId,
			PindId:    comment.PinId,
			Text:      comment.Text,
			CreatedAt: &comment.CreatedAt,
		})
	}

	response.Respond(w, http.StatusOK, resp, path)
}
