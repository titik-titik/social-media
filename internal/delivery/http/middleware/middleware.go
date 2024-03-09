package middleware

import (
	"net/http"
	"social-media/internal/config"
	"social-media/internal/repository"
	"strings"
	"time"

	"github.com/guregu/null"
)

type AuthMiddleware struct {
	DatabaseConfig    *config.DatabaseConfig
	SessionRepository *repository.SessionRepository
}

func NewAuthMiddleware(sessionRepository *repository.SessionRepository) *AuthMiddleware {
	return &AuthMiddleware{
		SessionRepository: sessionRepository,
	}
}

func (authMiddleware *AuthMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		token = strings.Replace(token, "Bearer ", "", 1)
		if token == "" {
			http.Error(w, "401 - Unauthorized: Missing token", http.StatusUnauthorized)
			return
		}

		tx, err := authMiddleware.DatabaseConfig.CockroachdbDatabase.Connection.Begin()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer tx.Rollback()

		session, err := authMiddleware.SessionRepository.FindOneByToken(tx, token)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if session == nil {
			http.Error(w, "401 - Unauthorized: Invalid Token", http.StatusUnauthorized)
			return
		}
		if session.AccessTokenExpiredAt == null.NewTime(time.Now(), true) {
			http.Error(w, "401 - Unauthorized: Token expired", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
