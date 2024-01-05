package models

type Shape struct {
	Creator   string `json:"creador"`
	ShapeType string `json:"tipo"`
	A         int64  `json:"a"`
	B         int64  `json:"b"`
}
