package processor

import (
	"log"
	"math"

	"github.com/HELL0ANTHONY/labs-golang/lambda/Shape/pkg/models"
)

type Processor struct{}

type ProcessorService interface {
	Process(models.Request) (float64, error)
}

func New() ProcessorService {
	return Processor{}
}

func (p Processor) Process(req models.Request) (float64, error) {
	log.Printf("PROCESSOR: %+v", req)

	if req.ShapeType == string(models.Ellipse) {
		return p.ellipse(req.A, req.B), nil
	}

	if req.ShapeType == string(models.Rectangle) {
		return p.rectangle(req.A, req.B), nil
	}

	if req.ShapeType == string(models.Triangle) {
		return p.triangle(req.A, req.B), nil
	}

	return 0.0, nil
}

func (Processor) triangle(a, b float64) float64 {
	return (a * b) / 2
}

func (Processor) ellipse(a, b float64) float64 {
	return a * b * math.Pi
}

func (Processor) rectangle(a, b float64) float64 {
	return a * b
}
