package models

type DynamoModel struct {
	ID        string  `json:"id"      dynamodbav:"id"`
	Creator   string  `json:"creador" dynamodbav:"creador"`
	ShapeType string  `json:"tipo"    dynamodbav:"tipo"`
	A         float64 `json:"a"       dynamodbav:"a"`
	B         float64 `json:"b"       dynamodbav:"b"`
}

type Shape string

const (
	Rectangle Shape = "RECTANGLE"
	Triangle  Shape = "TRIANGLE"
	Ellipse   Shape = "ELLIPSE"
)
