package middleware

import (
	"2020_1_Color_noise/internal/pkg/error"
	"2020_1_Color_noise/internal/pkg/response"
	sessions "2020_1_Color_noise/internal/pkg/session/usecase"
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"regexp"
)

type Middleware struct {
	sessions *sessions.SessionUsecase
}

func NewMiddleware(su *sessions.SessionUsecase) Middleware {
	return Middleware{
		sessions: su,
	}
}

var authMethods = map[string][]string{
	"/api/auth":           []string{"DELETE"},
	"/api/user":           []string{"GET", "PUT"},
	"/api/user/following": []string{"POST", "DELETE"},
	"/api/pin":            []string{"POST", "GET"},
	"/api/board":          []string{"POST", "GET"},
	"/api/comment":        []string{"POST", "GET"},
	"/api/search":         []string{"POST", "GET"},
}

var unauthMethods = map[string][]string{
	"/api/auth":           []string{"POST"},
	"/api/user":           []string{"POST"},
}

func (m *Middleware) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		reqId := fmt.Sprintf("%016x", rand.Int())[:10]
		ctx = context.WithValue(ctx, "ReqId", reqId)

		cookie, err := r.Cookie("session_id")
		if err != nil {
			unauthHandler(next, w, r.WithContext(ctx))
			return
		} else {
			session, err := m.sessions.GetByCookie(cookie.Value)
			if err != nil {
				unauthHandler(next, w, r.WithContext(ctx))
				return
			} else {
				ctx = context.WithValue(ctx, "Id", session.Id)
				authHandler(next, w, r.WithContext(ctx))
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
		//next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func unauthHandler(next http.Handler, w http.ResponseWriter, r *http.Request) {
	for path, methods := range authMethods {
		ok, err := regexp.MatchString(path, r.URL.Path)
		if err != nil {
			err := error.Wrap(err,"Error in AuthMiddleware")
			reqId:= r.Context().Value("ReqId")
			error.ErrorHandler(w, error.Wrapf(err, "request id: %s", reqId))
			return
		}

		if ok {
			for _, method := range  methods {
				if method == r.Method {
					err := error.Unauthorized.New("User is unauthorized")
					reqId:= r.Context().Value("ReqId")
					error.ErrorHandler(w, error.Wrapf(err, "request id: %s", reqId))
					return
				}
			}
		}
	}
	next.ServeHTTP(w, r)
}

func authHandler(next http.Handler, w http.ResponseWriter, r *http.Request) {
	for path,  methods := range unauthMethods{
		if path == r.URL.Path {
			for _,  method := range methods {
				if method == r.Method {
					response.Respond(w, http.StatusOK, map[string]string{
						"message": "Ok",
					})
					return
				}
			}
		}
	}
	next.ServeHTTP(w, r)
}

func (m *Middleware) CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Methods", "POST,PUT,DELETE,GET")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,X-CSRF-Token")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Origin", "*"/*r.Header.Get("Origin")*/)
		if r.Method == http.MethodOptions {
			return
		}
		next.ServeHTTP(w, r)
	})
}