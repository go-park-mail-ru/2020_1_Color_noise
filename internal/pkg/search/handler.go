package search

import (
	"2020_1_Color_noise/internal/models"
	"2020_1_Color_noise/internal/pkg/comment"
	"2020_1_Color_noise/internal/pkg/error"
	"2020_1_Color_noise/internal/pkg/metric"
	"2020_1_Color_noise/internal/pkg/pin"
	userService "2020_1_Color_noise/internal/pkg/proto/user"
	"2020_1_Color_noise/internal/pkg/response"
	"go.uber.org/zap"
	"google.golang.org/grpc/status"
	"net/http"
	"strconv"
)

type Handler struct {
	commentUsecase comment.IUsecase
	pinUsecase     pin.IUsecase
	us             userService.UserServiceClient
	logger         *zap.SugaredLogger
}

func NewHandler(commentUsecase comment.IUsecase, pinUsecase pin.IUsecase, us userService.UserServiceClient, logger *zap.SugaredLogger) *Handler {
	return &Handler{
		commentUsecase: commentUsecase,
		pinUsecase:     pinUsecase,
		us:             us,
		logger:         logger,
	}
}

func (sh *Handler) Search(w http.ResponseWriter, r *http.Request) {
	metric.Increase()

	reqId := r.Context().Value("ReqId")
	/*
		isAuth := r.Context().Value("IsAuth")
		if isAuth != true {
			err := error.Unauthorized.New("Search: user is unauthorized")
			error.ErrorHandler(w, r, sh.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
			return
		}
	*/
	what := r.URL.Query().Get("what")
	description := r.URL.Query().Get("description")
	start, _ := strconv.Atoi(r.URL.Query().Get("start"))

	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		limit = 10
	}

	switch what {
	case "comment":
		comments, err := sh.commentUsecase.GetByText(description, start, limit)
		if err != nil {
			err = error.Wrap(err, "Error searching comments")
			error.ErrorHandler(w, r, sh.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
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

		response.Respond(w, http.StatusOK, resp)
		return
	case "pin":
		date := r.URL.Query().Get("date")
		if date != "day" && date != "week" && date != "month" {
			date = ""
		}

		descString := r.URL.Query().Get("desc")
		desc := false
		if descString == "true" {
			desc = true
		}

		most := r.URL.Query().Get("most")
		if most != "popular" && most != "comment" {
			most = ""
		}

		pins, err := sh.pinUsecase.GetByName(description, start, limit, date, desc, most)
		if err != nil {
			err = error.Wrap(err, "Error searching pins")
			error.ErrorHandler(w, r, sh.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
			return
		}

		resp := make([]models.ResponsePin, 0)
		for _, pin := range pins {
			resp = append(resp, models.ResponsePin{
				Id:          pin.Id,
				BoardId:     pin.BoardId,
				UserId:      pin.UserId,
				Name:        pin.Name,
				Description: pin.Description,
				Image:       pin.Image,
			})
		}

		response.Respond(w, http.StatusOK, resp)
		return
	case "user":
		users, err := sh.us.Search(r.Context(), &userService.Searching{
			Login: &userService.Login{Login: description},
			Start: int64(start),
			Limit: int64(limit),
		})
		if err != nil {
			e := error.NoType
			errStatus, ok := status.FromError(err)
			msg := "Unknown GRPC error"
			if ok == true {
				e = error.Cast(int(errStatus.Code()))
				msg = errStatus.Message()
			}
			error.ErrorHandler(w, r, sh.logger, reqId, error.Wrapf(e.New(msg), "request id: %s", reqId))
			return
		}

		resp := make([]models.ResponseUser, 0)
		for _, user := range users.Users {
			resp = append(resp, models.ResponseUser{
				Id:            uint(user.Id),
				Login:         user.Login,
				About:         user.About,
				Avatar:        user.Avatar,
				Subscribers:   int(user.Subscribers),
				Subscriptions: int(user.Subscriptions),
			})
		}

		response.Respond(w, http.StatusOK, resp)
		return
	default:
		err = error.WithMessage(error.SearchNotFound.New("Bad id when in during searching"), "Bad parametrs")
		error.ErrorHandler(w, r, sh.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
	}
}
