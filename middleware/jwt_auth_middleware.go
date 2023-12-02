package middleware

import (
	"net/http"
	"time"

	"github.com/NurFirdausR/go-pos/config"
	"github.com/NurFirdausR/go-pos/helper"
	"github.com/golang-jwt/jwt/v5"
	"github.com/julienschmidt/httprouter"
)

func JWTMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		cookie, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				helper.PanicIfError(err)
			}
			helper.PanicIfError(err)
		}
		// ambil token value
		tokenString := cookie.Value
		claims := &config.JWTClaim{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return config.JWT_KEY, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				helper.PanicIfError(err)
			}
			if claims.ExpiresAt != nil {
				expirationTime := claims.ExpiresAt
				currentTime := jwt.NewNumericDate(time.Now())
				if currentTime.Unix() > expirationTime.Unix() {
					panic("token has expired")
				}
			}
			helper.PanicIfError(err)
		}

		if !token.Valid {
			helper.PanicIfError(err)
		}

		next(w, r, ps)
	}
}
