package middleware

import (
	"context"
	"net/http"
	sessions "pinterest/pkg/session/usecase"
	"time"
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
				ctx = context.WithValue(ctx, "isAuth", false)
			} else {
				if r.Method != http.MethodGet {
					token := r.Header.Get("X-CSRF-Token")
					if token != session.Token {
						ctx = context.WithValue(ctx, "isAuth", false)
					} else {
						ctx = context.WithValue(ctx, "isAuth", true)
						ctx = context.WithValue(ctx, "Id", session.Id)
						newToken := m.sessions.CreateToken()
						m.sessions.UpdateToken(session, newToken)
						token := &http.Cookie{
							Name:    "csrf_token",
							Value:    newToken,
							Expires: time.Now().Add(10 * time.Hour),
							Domain: r.Host,
						}
						http.SetCookie(w, token)
					}
				} else {
					ctx = context.WithValue(ctx, "isAuth", true)
					ctx = context.WithValue(ctx, "Id", session.Id)
					newToken := m.sessions.CreateToken()
					m.sessions.UpdateToken(session, newToken)
					token := &http.Cookie{
						Name:    "csrf_token",
						Value:    newToken,
						Expires: time.Now().Add(10 * time.Hour),
						Domain: r.Host,
					}
					http.SetCookie(w, token)
				}

			}
		}
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (m *Middleware) CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Methods", "POST,PUT,DELETE,GET")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,X-CSRF-Token")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
		if r.Method == http.MethodOptions {
			return
		}
		next.ServeHTTP(w, r)
	})
}