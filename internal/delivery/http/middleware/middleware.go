package middleware

import (
	"net/http"
	"social-media/internal/config"
	"social-media/internal/model/response"
	"social-media/internal/repository"
	"strings"
	"time"

	"github.com/guregu/null"
)

type AuthMiddleware struct {
	SessionRepository *repository.SessionRepository
	DatabaseConfig    *config.DatabaseConfig
}

func NewAuthMiddleware(sessionRepository repository.SessionRepository, databaseConfig *config.DatabaseConfig) *AuthMiddleware {
	return &AuthMiddleware{
		SessionRepository: &sessionRepository,
		DatabaseConfig:    databaseConfig,
	}
}

func (authMiddleware *AuthMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		token = strings.Replace(token, "Bearer ", "", 1)
		if token == "" {
			result := &response.Response[interface{}]{
				Code:    http.StatusUnauthorized,
				Message: "Unauthorized: Missing token",
			}
			response.NewResponse(w, result)
			return
		}

		tx, err := authMiddleware.DatabaseConfig.CockroachdbDatabase.Connection.Begin()
		if err != nil {
			result := &response.Response[interface{}]{
				Code:    http.StatusInternalServerError,
				Message: "transaction error",
			}
			response.NewResponse(w, result)
			return
		}

		session, err := authMiddleware.SessionRepository.FindOneByAccToken(tx, token)
		if err != nil {
			result := &response.Response[interface{}]{
				Code:    http.StatusUnauthorized,
				Message: "Unauthorized: token not found",
			}
			response.NewResponse(w, result)
			return
		}
		if session == nil {
			result := &response.Response[interface{}]{
				Code:    http.StatusUnauthorized,
				Message: "Unauthorized: Invalid Token",
			}
			response.NewResponse(w, result)
			return
		}
		if session.AccessTokenExpiredAt == null.NewTime(time.Now(), true) {
			result := &response.Response[interface{}]{
				Code:    http.StatusUnauthorized,
				Message: "Unauthorized: Token expired",
			}
			response.NewResponse(w, result)
			return
		}
		next.ServeHTTP(w, r)
	})
}
