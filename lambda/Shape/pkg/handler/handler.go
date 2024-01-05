package handler

import (
	"context"
	"log"

	"github.com/HELL0ANTHONY/labs-golang/lambda/Shape/internal/processor"
	"github.com/HELL0ANTHONY/labs-golang/lambda/Shape/pkg/models"
	"github.com/HELL0ANTHONY/labs-golang/lambda/Shape/pkg/validation"
)

type Handler struct {
	p processor.ProcessorService
	v validation.ValidationService
}

func New(p processor.ProcessorService, v validation.ValidationService) Handler {
	return Handler{
		p: p,
		v: v,
	}
}

func (h Handler) Handle(ctx context.Context, req models.Request) (float64, error) {
	log.Printf("%#v", req)

	if err := h.v.Validator(req); err != nil {
		log.Println(err)
		return 0.0, err
	}

	return h.p.Process(req)
}
