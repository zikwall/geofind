package geofind

type Point struct {
	Lat float32
	Lng float32
}

// This is a common framework adopted by Yandex and Google.
type Feature struct {
	Type       string     `json:"type"`
	Geometry   Geometry   `json:"geometry"`
	Properties Properties `json:"properties"`
}

// Some regions/countries/arbitrary regions may have separate polygons, divided by the main.
// for example, enclaves, etc.
func (feature *Feature) IsMultiPolygonal() bool {
	return feature.Geometry.Type == "MultiPolygon"
}

func (feature *Feature) GetSinglePolygon() [][]float32 {
	// All polygons are in nested elements of arrays, so the structure is always three-level
	return feature.Geometry.Coordinates[0] // [0] In any case, there is only 1 and only 1 element!
}

func (feature *Feature) GetAllPolygons() [][][]float32 {
	return feature.Geometry.Coordinates
}

type Properties struct {
	Iso3166   string   `json:"iso_3166"`
	Name      string   `json:"name"`
	Level     string   `json:"level"`
	Parents   []Parent `json:"parents"`
	Neighbors []string `json:"neighbors"`
}

type Geometry struct {
	Type        string        `json:"type"`
	Coordinates [][][]float32 `json:"coordinates"`
}

type Parent struct {
	Delta   int    `json:"delta"`
	Iso3166 string `json:"iso_3166"`
}

type Polygons struct {
	Type     string    `json:"type"`
	Features []Feature `json:"features"`
}
