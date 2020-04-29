package rpc

import (
	"2020_1_Color_noise/internal/models"
	. "2020_1_Color_noise/internal/pkg/error"
	"2020_1_Color_noise/internal/pkg/session"
	"golang.org/x/net/context"
)

type SessionManager struct {
	repo session.IRepository
}

func NewSessionManager(repo session.IRepository) *SessionManager {
	return &SessionManager{
		repo,
	}
}

func (sm *SessionManager) RPCCreate(ctx context.Context, in *session.ProtoSession) (*session.Nothing, error) {
	sess := &models.Session{
		Id:     uint(in.Id),
		Cookie: in.Cookie,
		Token:  in.Token,

	}

	err := sm.repo.Add(sess)
	if err != nil {
		return &session.Nothing{Error: true}, Wrapf(err, "Creating session error, id: %d", in.Id)
	}

	return &session.Nothing{Error: false}, nil
}

func (sm *SessionManager) RPCGetByCookie(ctx context.Context, in *session.ProtoCookie) (*session.ProtoSession, error) {
	sess, err := sm.repo.GetByCookie(in.Cookie)
	if err != nil {
		return &session.ProtoSession{}, Wrap(err, "Getting session error")
	}

	protoSession := &session.ProtoSession{
		Id: int64(sess.Id),
		Cookie: sess.Cookie,
		Token:  sess.Token,
	}

	return protoSession, nil
}

func (sm *SessionManager) RPCUpdate(ctx context.Context, in *session.ProtoSession) (*session.Nothing, error) {
	sess := &models.Session{
		Id:     uint(in.Id),
		Cookie: in.Cookie,
		Token:  in.Token,
	}

	err := sm.repo.Update(sess)
	if err != nil {
		return &session.Nothing{Error: true}, Wrap(err, "Updating session error")
	}

	return &session.Nothing{Error: false}, nil
}

func (sm *SessionManager) RPCDelete(ctx context.Context, in *session.ProtoCookie) (*session.Nothing, error) {
	err := sm.repo.Delete(in.Cookie)
	if err != nil {
		return &session.Nothing{Error: true}, Wrap(err, "Deleting session error")
	}

	return &session.Nothing{Error: false}, nil
}