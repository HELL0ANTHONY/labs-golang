package handler

import (
	"context"
	"log"
)

type Handler struct{}

func New() Handler {
	return Handler{}
}

func (h Handler) Handle(ctx context.Context) error {
	log.Println("Hola Mundo")
	return nil
}
