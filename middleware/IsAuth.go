package middleware

import (
	"context"
	"net/http"

	"github.com/HazimEmam/JWTtutorial/utils"
)

func IsAuthorized(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c , err:= r.Cookie("token")
		if err != nil{
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Middleware: bad request"))
			return
		}
		claims , err:= utils.VerifyToken(c.Value)

		if err != nil{
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Middleware: unauthorized"))
			return
		}

		ctx := context.WithValue(r.Context(),"role" ,claims.Role)
		r = r.WithContext(ctx)

		next(w,r)
	})
}