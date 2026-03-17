package middleware

import (
	"context"
	"net/http"
	"strings"
	"taskify/internal/domain"
	"taskify/internal/helpers"
	"taskify/pkg/errors"
)

const UserIDContextKey string = "UserID"

type AuthMiddleware struct {
	userRepository domain.UserRepository
}

func NewAuthMiddleware(userRepository domain.UserRepository) *AuthMiddleware {
	return &AuthMiddleware{
		userRepository: userRepository,
	}
}

func (m *AuthMiddleware) VerifyAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// HTTP header
		// Authorization: Bearer .....
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			helpers.ResponseWithError(w, http.StatusUnauthorized, errors.ErrInvalidAuthorization.Error())
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		if token == "" {
			helpers.ResponseWithError(w, http.StatusUnauthorized, errors.ErrInvalidAuthorization.Error())
			return
		}

		authenticatedUser, err := m.userRepository.GetById(token)

		if err != nil {
			helpers.ResponseWithError(w, http.StatusUnauthorized, errors.ErrInvalidToken.Error())
			return
		}

		ctx := context.WithValue(r.Context(), UserIDContextKey, authenticatedUser.Id)

		next(w, r.WithContext(ctx))
	}
}

func GetUserIDFromContext(ctx context.Context) (string, error) {
	userID, ok := ctx.Value(UserIDContextKey).(string)

	if !ok {
		return "", errors.ErrInvalidCredentials
	}

	return userID, nil
}
