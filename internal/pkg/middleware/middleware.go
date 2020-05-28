package middleware

import (
	"2020_1_Color_noise/internal/pkg/metric"
	"2020_1_Color_noise/internal/pkg/error"
	authService "2020_1_Color_noise/internal/pkg/proto/session"
	"context"
	"fmt"
	"go.uber.org/zap"
	"math/rand"
	"net/http"
	"time"
)

type Middleware struct {
	as     authService.AuthSeviceClient
	logger *zap.SugaredLogger
}

func NewMiddleware(as authService.AuthSeviceClient, logger *zap.SugaredLogger) Middleware {
	return Middleware{
		as:     as,
		logger: logger,
	}
}

func (m *Middleware) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqId := r.Context().Value("ReqId")

		ctx := r.Context()

		cookie, err := r.Cookie("session_id")
		if err != nil {
			ctx = context.WithValue(ctx, "IsAuth", false)
		} else {
			session, err := m.as.GetByCookie(context.Background(),
				&authService.Cookie{Cookie: cookie.Value})
			if err != nil {
				ctx = context.WithValue(ctx, "IsAuth", false)
				m.logger.Info(r.URL.Path,
					zap.String("reqId:", fmt.Sprintf("%v", reqId)),
					zap.String("userId:", "unknown"),
				)
			} else {
				m.logger.Info(
					zap.String("reqId:", fmt.Sprintf("%v", reqId)),
					zap.Uint("userId:", uint(session.Id)),
				)
				ctx = context.WithValue(ctx, "IsAuth", true)
				ctx = context.WithValue(ctx, "Id", uint(session.Id))
				//newToken := m.sessions.CreateToken()
				//m.sessions.UpdateToken(session, newToken)

				//http.SetCookie(w, token)

				/*
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
					}*/

			}
		}
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
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

//Добавлен мониторинг
func (m *Middleware) AccessLogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		reqId := fmt.Sprintf("%016x", rand.Int())[:10]
		ctx = context.WithValue(ctx, "ReqId", reqId)

		start := time.Now()
		next.ServeHTTP(w, r.WithContext(ctx))
		metric.Increase()
		metric.WorkTime(r.Method, r.URL.Path, time.Since(start))
		metric.Errors(w.Header().Get("Real-Status"), r.URL.Path)

		m.logger.Info(r.URL.Path,
			zap.String("reqId:", reqId),
			zap.String("method", r.Method),
			zap.String("remote_addr", r.RemoteAddr),
			zap.String("url", r.URL.Path),
			zap.Time("start", start),
			zap.Duration("work_time", time.Since(start)),
		)
	})
}

func (m *Middleware) PanicMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			reqId := r.Context().Value("ReqId")
			if err := recover(); err != nil {
				e := error.NoType.Newf("recovered err: %s", err)
				error.ErrorHandler(w, r, m.logger, reqId, e)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
