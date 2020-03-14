package geofind

import "encoding/json"

type Finder interface {
	In(point Point) bool
}

type Polygon struct {
	Coordinates [][]float32
	Itterations int
}

// alghoritm @see http://alienryderflex.com/polygon/
func (polygon *Polygon) In(point Point) bool {
	var k_prev int
	result := false

	for k, p := range polygon.Coordinates {
		if k <= 0 {
			k_prev = len(polygon.Coordinates) - 1
		} else {
			k_prev = k - 1
		}

		iflng := p[1] < point.Lng && polygon.Coordinates[k_prev][1] >= point.Lng || polygon.Coordinates[k_prev][1] < point.Lng && p[1] >= point.Lng
		iflat := p[0] <= point.Lat || polygon.Coordinates[k_prev][0] <= point.Lat

		if iflng && iflat {
			if p[0]+(point.Lng-p[1])/(polygon.Coordinates[k_prev][1]-p[1])*(polygon.Coordinates[k_prev][0]-p[0]) < point.Lat {
				result = !result
			}
		}

		polygon.Itterations++
	}

	return result
}

func Init(geobase []byte) (*Polygons, error) {
	polygons := &Polygons{}

	if err := json.Unmarshal(geobase, &polygons); err != nil {
		return nil, err
	}

	return polygons, nil
}
