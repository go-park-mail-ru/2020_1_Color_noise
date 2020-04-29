package rpcRepo

import (
	"2020_1_Color_noise/internal/models"
	. "2020_1_Color_noise/internal/pkg/error"
	"2020_1_Color_noise/internal/pkg/session"
	"context"
)

type Repository struct {
	sm session.AuthCheckerClient
}

func NewRepo(sm session.AuthCheckerClient) *Repository {
	return &Repository{
		sm: sm,
	}
}

func (sr *Repository) Add(sess *models.Session) error {
	in := &session.ProtoSession{
		Id:     int64(sess.Id),
		Cookie: sess.Cookie,
		Token:  sess.Token,
	}

	_, err := sr.sm.RPCCreate(context.Background(), in)
	if err != nil {
		return Wrap(err, "RPC: Create session error")
	}

	return nil
}

func (sr *Repository) GetByCookie(cookie string) (*models.Session, error) {
	pc := &session.ProtoCookie{
		Cookie: cookie,
	}
	protoSess, err := sr.sm.RPCGetByCookie(context.Background(), pc)
	if err != nil {
		return nil, Unauthorized.Newf("RPC: Session is not found, cookie: %s", cookie)
	}

	sess := &models.Session{
		Id: uint(protoSess.Id),
		Cookie: protoSess.Cookie,
		Token:  protoSess.Token,
	}

	return sess, nil
}

func (sr *Repository) Update(sess *models.Session) error {
	in := &session.ProtoSession{
		Id:     int64(sess.Id),
		Cookie: sess.Cookie,
		Token:  sess.Token,
	}

	_, err := sr.sm.RPCUpdate(context.Background(), in)
	if err != nil {
		return Wrap(err, "RPC: Update session error")
	}

	return nil
}

func (sr *Repository) Delete(cookie string) error {
	pc := &session.ProtoCookie{
		Cookie: cookie,
	}
	_, err := sr.sm.RPCDelete(context.Background(), pc)
	if err != nil {
		return Unauthorized.Newf("RPC: Session is not found, cookie: %s", cookie)
	}

	return nil
}
