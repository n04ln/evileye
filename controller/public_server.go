package controller

import "github.com/NoahOrberg/evileye/usecase"

type PublicServerHandler struct {
	UUsecase usecase.ServerUserUsecase
}

func NewPublicServerHandler(su usecase.ServerUserUsecase) *PublicServerHandler {
	return &PublicServerHandler{
		UUsecase: su,
	}
}
