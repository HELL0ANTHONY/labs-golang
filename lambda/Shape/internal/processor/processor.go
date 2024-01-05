package processor

import (
	"log"

	"github.com/HELL0ANTHONY/labs-golang/lambda/Shape/pkg/models"
)

type Processor struct{}

type ProcessorService interface {
	Process(models.Request) error
}

func New() ProcessorService {
	return Processor{}
}

func (p Processor) Process(req models.Request) error {
	log.Printf("PROCESSOR: %+v", req)
	return nil
}
