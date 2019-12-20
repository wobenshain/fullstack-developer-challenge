package endpoints

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/wobenshain/fullstack-developer-challenge/go/errors"
	"github.com/wobenshain/fullstack-developer-challenge/go/restapi/models"
	"github.com/wobenshain/fullstack-developer-challenge/go/restapi/operations"
	"github.com/wobenshain/fullstack-developer-challenge/go/types"
	"net/http"
)

func GetItem(repo types.ItemRepository) operations.GetItemHandlerFunc {
	return func(params operations.GetItemParams) middleware.Responder {
		items, err := repo.GetAll()
		if err != nil {
			return errors.GetItem(http.StatusInternalServerError, err.Error())
		}
		itemsPayload := make([]*models.Item, 0)
		for i := range items {
			itemsPayload = append(itemsPayload, &items[i])
		}

		return operations.NewGetItemOK().WithPayload(itemsPayload)
	}
}

func GetItemID(repo types.ItemRepository) operations.GetItemIDHandlerFunc {
	return func(params operations.GetItemIDParams) middleware.Responder {
		id := int(params.ID)
		if float64(id) != params.ID {
			return errors.GetItemID(http.StatusBadRequest, "ID must be an integer")
		}
		item, err := repo.Get(id)
		if err != nil {
			return errors.GetItemID(http.StatusNotFound, err.Error())
		}
		return operations.NewGetItemIDOK().WithPayload(&item)
	}
}

func DeleteItemID(repo types.ItemRepository) operations.DeleteItemIDHandlerFunc {
	return func(params operations.DeleteItemIDParams) middleware.Responder {
		id := int(params.ID)
		if float64(id) != params.ID {
			return errors.DeleteItemID(http.StatusBadRequest, "ID must be an integer")
		}
		item, err := repo.Delete(id)
		if err != nil {
			return errors.DeleteItemID(http.StatusInternalServerError, err.Error())
		}
		return operations.NewGetItemIDOK().WithPayload(&item)
	}
}

func PatchItemID(repo types.ItemRepository) operations.PatchItemIDHandlerFunc {
	return func(params operations.PatchItemIDParams) middleware.Responder {
		if params.Item == nil {
			return errors.PatchItemID(http.StatusBadRequest, "Request body must contain an item")
		}
		if params.ID != float64(params.Item.ID) {
			return errors.PatchItemID(http.StatusBadRequest, "Request ID must match Path ID")
		}
		item, err := repo.Update(*params.Item)
		if err != nil {
			return errors.PatchItemID(http.StatusInternalServerError, err.Error())
		}
		return operations.NewPatchItemIDOK().WithPayload(&item)
	}
}

func PostItem(repo types.ItemRepository) operations.PostItemHandlerFunc {
	return func(params operations.PostItemParams) middleware.Responder {
		if params.Item == nil {
			return errors.PostItem(http.StatusBadRequest, "Request body must contain an item")
		}
		if params.Item.ID != 0 {
			return errors.PostItem(http.StatusBadRequest, "Request ID must not be set")
		}
		item, err := repo.Create(*params.Item)
		if err != nil {
			return errors.PostItem(http.StatusInternalServerError, err.Error())
		}
		return operations.NewPostItemOK().WithPayload(&item)
	}
}
