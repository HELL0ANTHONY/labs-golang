package processor

import (
	"log"
	"math"
	"strings"

	"github.com/HELL0ANTHONY/labs-golang/lambda/Shape/internal/repository"
	"github.com/HELL0ANTHONY/labs-golang/lambda/Shape/pkg/models"
)

type Processor struct {
	d repository.DynamoDBService
}

type ProcessorService interface {
	Process(models.Request) (float64, error)
}

func New(d repository.DynamoDBService) ProcessorService {
	return Processor{
		d: d,
	}
}

func (p Processor) Process(req models.Request) (float64, error) {
	log.Printf("PROCESSOR: %+v", req)

	id := req.ID
	if id != "" {
		m, err := p.d.FindClientBy(id)
		if err != nil {
			return 0, err
		}
		log.Printf("%#v", m)
	}

	a := req.A
	b := req.B
	s := strings.ToUpper(req.ShapeType)

	if s == string(models.Ellipse) {
		return p.ellipse(a, b), nil
	}

	if s == string(models.Rectangle) {
		return p.rectangle(a, b), nil
	}

	if s == string(models.Triangle) {
		return p.triangle(a, b), nil
	}

	return 0, nil
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
