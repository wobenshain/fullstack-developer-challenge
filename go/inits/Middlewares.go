package inits

import (
	"encoding/json"
	"fmt"
	"github.com/rs/cors"
	"net/http"
)

type jsonError struct{
	Message interface{} `json:"message"`
}

func initPanicHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				w.Header().Set(http.CanonicalHeaderKey("Content-Type"), "application/json; charset=utf-8")
				e := jsonError{Message: "Internal Server Error"}
				err := json.NewEncoder(w).Encode(e)
				if err != nil {
					fmt.Println(err)
				}
			}
		}()
		handler.ServeHTTP(w, r)
	})
}

func initCors(handler http.Handler) http.Handler {
	allowLocal := cors.New(cors.Options{
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		AllowedMethods: []string{"GET", "POST", "PATCH", "DELETE"},
	})
	return allowLocal.Handler(handler)
}

func InitLocalMiddlewares(handler http.Handler) http.Handler {
	// middlewares are listed in the order they will execute, so top level should go on the bottom
	returnHandler := handler
	return returnHandler
}

func InitGlobalMiddlewares(handler http.Handler) http.Handler {
	// middlewares are listed in the order they will execute, so top level should go on the bottom
	returnHandler := handler
	returnHandler = initCors(returnHandler)
	returnHandler = initPanicHandler(returnHandler)
	return returnHandler
}
