package models

type DynamoModel struct {
	ID        string `json:"id"      dynamodbav:"id"`
	Creator   string `json:"creador" dynamodbav:"creador"`
	ShapeType string `json:"tipo"    dynamodbav:"tipo"`
	A         int64  `json:"a"       dynamodbav:"a"`
	B         int64  `json:"b"       dynamodbav:"b"`
}
