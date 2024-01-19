package models

// ID        string  `json:"id"      dynamodbav:"id"` // PK
type Record struct {
	ID        string `json:"id"` // PK
	Creator   string `json:"creador"`
	ShapeType string `json:"tipo"` // SK
	A         int    `json:"a"`
	B         int    `json:"b"`
}

type Shape string

const (
	Rectangle Shape = "RECTANGLE"
	Triangle  Shape = "TRIANGLE"
	Ellipse   Shape = "ELLIPSE"
)
