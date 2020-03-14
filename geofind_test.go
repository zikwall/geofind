package geofind

import (
	"encoding/json"
	"github.com/zikwall/geofind"
	"testing"
)

func TestInclude(t *testing.T) {
	testJsonGeometry := `[
   [
      -72.2802114,
      42.9275749
   ],
   [
      -72.2809839,
      42.9266243
   ],
   [
      -72.2817242,
      42.9257444
   ],
   [
      -72.2805333,
      42.9249431
   ],
   [
      -72.2792351,
      42.9239925
   ],
   [
      -72.276628,
      42.9246996
   ],
   [
      -72.2760272,
      42.9255323
   ],
   [
      -72.2767353,
      42.9260351
   ],
   [
      -72.2772717,
      42.9268521
   ],
   [
      -72.2758126,
      42.927677
   ],
   [
      -72.2762632,
      42.9284311
   ],
   [
      -72.2777116,
      42.9287768
   ],
   [
      -72.2785592,
      42.928549
   ],
   [
      -72.2793317,
      42.9279755
   ],
   [
      -72.2799325,
      42.927622
   ],
   [
      -72.2802114,
      42.9275984
   ]
]`
	polygon := struct {
		coordinates [][]float32
	}{}

	if err := json.Unmarshal([]byte(testJsonGeometry), &polygon.coordinates); err != nil {
		t.Error("Failed load polygon!")
	}

	finder := &geofind.Polygon{Coordinates: polygon.coordinates, Itterations: 0}
	found := finder.In(geofind.Point{-72.278, 42.925})

	if found == false {
		t.Error("Not includes coordinate in polygon!")
	}
}
