package model

type SearchResult struct {
	Driver   Driver  `json:"driver"`
	Distance float64 `json:"distance"` // in km
}
