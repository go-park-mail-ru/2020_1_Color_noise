package middleware

import (
	"context"
	"fmt"
	"net/http"
	sessions "pinterest/pkg/session/usecase"
)

type Middleware struct {
	sessions *sessions.SessionUsecase
}

func NewMiddleware(su *sessions.SessionUsecase) Middleware {
	return Middleware{
		sessions: su,
	}
}

func (m *Middleware) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		cookie, err := r.Cookie("session_id")
		if err != nil {
			ctx = context.WithValue(ctx, "isAuth", false)
		} else {
			session, err := m.sessions.GetByCookie(cookie.Value)
			if err != nil {
				fmt.Println(err,"2222")
				ctx = context.WithValue(ctx, "isAuth", false)
			} else {

				ctx = context.WithValue(ctx, "isAuth", true)
				ctx = context.WithValue(ctx, "Id", session.Id)
			}
		}
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (m *Middleware) CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Methods", "POST,PUT,DELETE,GET")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
		if r.Method == http.MethodOptions {
			return
		}
		next.ServeHTTP(w, r)
	})
}