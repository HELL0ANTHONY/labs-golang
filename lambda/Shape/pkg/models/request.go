package models

type Request struct {
	ID        string  `json:"id"`
	Creator   string  `json:"creador"`
	ShapeType string  `json:"tipo"`
	A         float64 `json:"a"`
	B         float64 `json:"b"`
}
