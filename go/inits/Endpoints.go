package inits

import (
	"github.com/wobenshain/fullstack-developer-challenge/go/endpoints"
	"github.com/wobenshain/fullstack-developer-challenge/go/restapi/operations"
	"github.com/wobenshain/fullstack-developer-challenge/go/types"
)

func InitEndpoints(api *operations.BelleseChallengeAPI, rs types.Repositories) {
	api.GetItemHandler = endpoints.GetItem(rs.ItemRepository)
	api.GetItemIDHandler = endpoints.GetItemID(rs.ItemRepository)
	api.DeleteItemIDHandler = endpoints.DeleteItemID(rs.ItemRepository)
	api.PatchItemIDHandler = endpoints.PatchItemID(rs.ItemRepository)
	api.PostItemHandler = endpoints.PostItem(rs.ItemRepository)
}
