package inits

import (
	"github.com/dre1080/recover"
	"github.com/rs/cors"
	"log"
	"net/http"
)

func initPanicHandler(handler http.Handler) http.Handler {
	recovery := recover.New(&recover.Options{
		Log: log.Print,
	})
	return recovery(handler)
}

func initCors(handler http.Handler) http.Handler {
	allowLocal := cors.New(cors.Options{
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		AllowedMethods: []string{"GET", "POST", "PATCH", "DELETE"},
		Debug: true,
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
