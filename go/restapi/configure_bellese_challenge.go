// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"github.com/wobenshain/fullstack-developer-challenge/go/restapi/models"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/wobenshain/fullstack-developer-challenge/go/inits"
	"github.com/wobenshain/fullstack-developer-challenge/go/restapi/operations"
)

//go:generate swagger generate server --target ..\..\go --name BelleseChallenge --spec ..\..\swagger.yml --model-package restapi/models

func configureFlags(api *operations.BelleseChallengeAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.BelleseChallengeAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()
	api.JSONProducer = runtime.JSONProducer()

	repos := inits.InitRepositories()

	api.DeleteItemIDHandler = operations.DeleteItemIDHandlerFunc(func(params operations.DeleteItemIDParams) middleware.Responder {
		return middleware.NotImplemented("operation .DeleteItemID has not yet been implemented")
	})
	api.GetItemHandler = operations.GetItemHandlerFunc(func(params operations.GetItemParams) middleware.Responder {
		items, err := repos.ItemRepository.GetAll()
		if err != nil {
			msg := err.Error()
			return operations.NewGetItemDefault(500).WithPayload(&models.Error{
				Message: &msg,
			})
		}
		itemsPayload := make([]*models.Item, 0)
		for _, i := range items {
			itemsPayload = append(itemsPayload, &i)
		}

		return operations.NewGetItemOK().WithPayload(itemsPayload)
	})
	api.GetItemIDHandler = operations.GetItemIDHandlerFunc(func(params operations.GetItemIDParams) middleware.Responder {
		return middleware.NotImplemented("operation .GetItemID has not yet been implemented")
	})
	api.PatchItemIDHandler = operations.PatchItemIDHandlerFunc(func(params operations.PatchItemIDParams) middleware.Responder {
		return middleware.NotImplemented("operation .PatchItemID has not yet been implemented")
	})
	api.PostItemHandler = operations.PostItemHandlerFunc(func(params operations.PostItemParams) middleware.Responder {
		return middleware.NotImplemented("operation .PostItem has not yet been implemented")
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
