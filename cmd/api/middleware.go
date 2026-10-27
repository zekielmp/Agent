package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	resp "github.com/zekielmp/Agent/internal/helpers"
	"github.com/zekielmp/Agent/internal/utils"
)

func Authorize(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authtoken := r.Header.Get("Authorization")
		if len(authtoken) == 0 {
			_ = resp.ErrorJSON(w, utils.ErrAuthHeader, http.StatusUnauthorized)
			return
		}

		bearer := strings.Split(authtoken, " ")
		if len(bearer) < 2 {
			_ = resp.ErrorJSON(w, utils.ErrInvalidAuthHeader, http.StatusUnauthorized)
			return
		}

		if bearer[0] != "Bearer" {
			_ = resp.ErrorJSON(w, utils.ErrUnsupportedAuthType, http.StatusUnauthorized)
			return
		}

		accessToken := bearer[1]

		payload, err := utils.VerifyToken(accessToken, os.Getenv("JWT_SECRET"))
		if err != nil {
			_ = resp.ErrorJSON(w, fmt.Errorf("%w", err), http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(context.Background(), struct{ id, role string }{"user_id", "user_role"}, struct{ id, role any }{payload.UserID, payload.Role})
		next.ServeHTTP(w, r.Clone(ctx))

	})
}
