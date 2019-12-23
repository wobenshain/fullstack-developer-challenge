package inits

import (
	"github.com/wobenshain/fullstack-developer-challenge/go/repositories/postgres"
	"github.com/wobenshain/fullstack-developer-challenge/go/types"
)

func InitRepositories() types.Repositories {
	return types.Repositories{
		ItemRepository: postgres.InitItemRepository(),
	}
}
