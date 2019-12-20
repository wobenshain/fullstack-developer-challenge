package errors

import (
	"github.com/wobenshain/fullstack-developer-challenge/go/restapi/models"
	"github.com/wobenshain/fullstack-developer-challenge/go/restapi/operations"
)

func GetItem(code int, msg string) *operations.GetItemDefault {
	return operations.NewGetItemDefault(code).WithPayload(&models.Error{
		Code: int64(code),
		Message: &msg,
	})
}

func GetItemID(code int, msg string) *operations.GetItemIDDefault {
	return operations.NewGetItemIDDefault(code).WithPayload(&models.Error{
		Code: int64(code),
		Message: &msg,
	})
}

func DeleteItemID(code int, msg string) *operations.DeleteItemIDDefault {
	return operations.NewDeleteItemIDDefault(code).WithPayload(&models.Error{
		Code: int64(code),
		Message: &msg,
	})
}

func PatchItemID(code int, msg string) *operations.PatchItemIDDefault {
	return operations.NewPatchItemIDDefault(code).WithPayload(&models.Error{
		Code: int64(code),
		Message: &msg,
	})
}

func PostItem(code int, msg string) *operations.PostItemDefault {
	return operations.NewPostItemDefault(code).WithPayload(&models.Error{
		Code: int64(code),
		Message: &msg,
	})
}
