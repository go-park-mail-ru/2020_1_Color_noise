package grpc

import (
	"2020_1_Color_noise/internal/models"
	. "2020_1_Color_noise/internal/pkg/error"
	authService "2020_1_Color_noise/internal/pkg/proto/session"
	"2020_1_Color_noise/internal/pkg/session"
	"golang.org/x/net/context"
)

type SessionManager struct {
	usecase session.IUsecase
}

func NewSessionManager(usecase session.IUsecase) *SessionManager {
	return &SessionManager{
		usecase,
	}
}

func (sm *SessionManager) Create(ctx context.Context, in *authService.UserID) (*authService.Session, error) {
	s, err := sm.usecase.Create(uint(in.Id))
	if err != nil {
		return &authService.Session{Error: int32(GetType(err))}, Wrapf(err, "GRPC create: Creating session error, id: %d", in.Id)
	}

	sess := &authService.Session{
		Id: int64(s.Id),
		Cookie: s.Cookie,
		Token: s.Token,
	}

	return sess, nil
}

func (sm *SessionManager) GetByCookie(ctx context.Context, in *authService.Cookie) (*authService.Session, error) {
	s, err := sm.usecase.GetByCookie(in.Cookie)
	if err != nil {
		return &authService.Session{Error: int32(GetType(err))}, Wrap(err, "GPRC GetByCoolie: Getting session error")
	}

	sess := &authService.Session{
		Id: int64(s.Id),
		Cookie: s.Cookie,
		Token: s.Token,
	}

	return sess, nil
}

func (sm *SessionManager) Update(ctx context.Context, in *authService.Session) (*authService.Nothing, error) {
	sess := &models.Session{
		Id:     uint(in.Id),
		Cookie: in.Cookie,
		Token:  in.Token,
	}

	err := sm.usecase.Update(sess)
	if err != nil {
		return &authService.Nothing{Error: int32(GetType(err))}, Wrap(err, "GPRC Update: Updating session error")
	}

	return &authService.Nothing{}, nil
}

func (sm *SessionManager) Delete(ctx context.Context, in *authService.Cookie) (*authService.Nothing, error) {
	err := sm.usecase.Delete(in.Cookie)
	if err != nil {
		return &authService.Nothing{Error: int32(GetType(err))}, Wrap(err, "GPRC Delete: Deleting session error")
	}

	return &authService.Nothing{}, nil
}

func (sm *SessionManager) Login(ctx context.Context, in *authService.SignIn) (*authService.Session, error) {
	us := &models.User{
		Id:            uint(in.User.Id),
		Email:         in.User.Email,
		Login:         in.User.Login,
		EncryptedPassword: in.User.EncryptedPassword,
		About:         in.User.About,
		Avatar:        in.User.Avatar,
		Subscribers:   int(in.User.Subscribers),
		Subscriptions: int(in.User.Subscriptions),
	}

	s, err := sm.usecase.Login(us, in.Password)
	if err != nil {
		return &authService.Session{Error: int32(GetType(err))}, Wrapf(err, "GRPC Login: Login error",)
	}


	sess := &authService.Session{
			Id:     int64(s.Id),
			Cookie: s.Cookie,
			Token:  s.Token,
		}

	return sess, nil
}