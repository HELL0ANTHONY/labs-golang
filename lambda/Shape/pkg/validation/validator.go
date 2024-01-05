package validation

import (
	"errors"

	"github.com/HELL0ANTHONY/labs-golang/lambda/Shape/pkg/models"
)

type Validation struct{}

type ValidationService interface {
	Validator(models.Request) error
}

func New() ValidationService {
	return Validation{}
}

func (Validation) Validator(req models.Request) error {
	validShapes := map[models.Shape]bool{
		models.Rectangle: true,
		models.Ellipse:   true,
		models.Triangle:  true,
	}

	if !validShapes[models.Shape(req.ShapeType)] {
		return errors.New(req.ShapeType + " it is not a valid option")
	}

	return nil
}
