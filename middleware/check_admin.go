package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/Ananth1082/Lab_Manual/config"
)

func CheckAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if r.Method == "GET" {
				next.ServeHTTP(w, r)
				return
			}
			ctx := context.Background()
			// Manual parsing of the section from the URL path
			parts := strings.Split(r.URL.Path, "/")
			var section string
			if len(parts) > 1 {
				section = parts[1]
			}
			userAuthDetails := r.Header["Authorization"]
			if len(userAuthDetails) == 0 {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("Enter admin keys"))
				return
			}

			userToken := strings.Split(userAuthDetails[0], " ")[1]
			docSnap, err := config.Firebase.Fs.Collection("sections").Doc(section).Get(ctx)

			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("Invalid section" + err.Error()))
				return
			}

			data, _ := docSnap.DataAt("admin_tokens")
			adminTokens := data.([]interface{})
			for _, adminToken := range adminTokens {
				token := adminToken.(string)
				if userToken == token {
					next.ServeHTTP(w, r)
					return
				}
			}
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("UNAUTHORIZED"))
		})
}
