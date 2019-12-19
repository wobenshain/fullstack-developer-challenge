package inits

import (
	"github.com/wobenshain/fullstack-developer-challenge/go/repositories/inmem"
	"github.com/wobenshain/fullstack-developer-challenge/go/types"
)

func InitRepositories() types.Repositories {
	return types.Repositories{
		ItemRepository: inmem.InitItemRepository(),
	}
}
