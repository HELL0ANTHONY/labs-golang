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
	a := req.A
	b := req.B

	if req.ShapeType == string(models.Ellipse) {
		return p.ellipse(a, b), nil
	}

	if req.ShapeType == string(models.Rectangle) {
		return p.rectangle(a, b), nil
	}

	if req.ShapeType == string(models.Triangle) {
		return p.triangle(a, b), nil
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
