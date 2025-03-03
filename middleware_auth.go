package main

import (
	"fmt"
	"net/http"

	"github.com/ponzaa555/rssagg/internal/auth"
	"github.com/ponzaa555/rssagg/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middleWareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			responseWithError(w, 403, fmt.Sprintf("Auth  error: %v", err))
			return
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			responseWithError(w, 403, fmt.Sprintf("Clound't found user: %v", err))
			return
		}

		handler(w, r, user)
	}
}
