<div align="center">
  <h1>GeoFind</h1>
  <h4>GoLang implementation Point-In-Polygon Algorithm</h4>
</div>

### Installation

`go get github.com/zikwall/geofind`

### Example usage

```go

package main

import (
	"fmt"
	"github.com/zikwall/geofind"
	"io/ioutil"
)

func main() {
	countryJson, err := ioutil.ReadFile("example/data/countries/RU.json")
	if err != nil {
		panic(err)
	}

	//Load countries database to memory, example Russian
	countryPolygons, err := geofind.Init(countryJson)

	if err != nil {
		panic(err)
	}

	found := false
	in := geofind.Properties{}

	// Loop all features of country
	for _, feature := range countryPolygons.Features {
		// read method comment for help
		if feature.IsMultiPolygonal() {
			for _, polygon := range feature.GetAllPolygons() {
				finder := &geofind.Polygon{polygon, 0}
				found = finder.In(geofind.Point{53.671362646, 82.18341123})
				in = feature.Properties
				break

			}
		} else {
			finder := &geofind.Polygon{feature.GetSinglePolygon(), 0}
			found = finder.In(geofind.Point{53.671362646, 82.18341123})
			in = feature.Properties
			break
		}
	}

	fmt.Println(found)
	fmt.Println(in)
}

```

### Theory

_ | _ |
--- | --- | 
![image](/images/Diagram_1.gif) | Figure 1 demonstrates a typical case of a severely concave polygon with 14 sides. The red dot is a point which needs to be tested, to determine if it lies inside the polygon. The solution is to compare each side of the polygon to the Y (vertical) coordinate of the test point, and compile a list of nodes, where each node is a point where one side crosses the Y threshold of the test point. In this example, eight sides of the polygon cross the Y threshold, while the other six sides do not. Then, if there are an odd number of nodes on each side of the test point, then it is inside the polygon; if there are an even number of nodes on each side of the test point, then it is outside the polygon. In our example, there are five nodes to the left of the test point, and three nodes to the right. Since five and three are odd numbers, our test point is inside the polygon.(Note: This algorithm does not care whether the polygon is traced in clockwise or counterclockwise fashion.) 
![image](/images/Diagram_2.gif) | Figure 2 shows what happens if the polygon crosses itself. In this example, a ten-sided polygon has lines which cross each other. The effect is much like “exclusive or,” or XOR as it is known to assembly-language programmers. The portions of the polygon which overlap cancel each other out. So, the test point is outside the polygon, as indicated by the even number of nodes (two and two) on either side of it.
![image](/images/Diagram_3.gif) | In Figure 3, the six-sided polygon does not overlap itself, but it does have lines that cross. This is not a problem; the algorithm still works fine. 
![image](/images/Diagram_4.gif) | Figure 4 demonstrates the problem that results when a vertex of the polygon falls directly on the Y threshold.  Since sides a and b both touch the threshold, should they both generate a node? No, because then there would be two nodes on each side of the test point and so the test would say it was outside of the polygon, when it clearly is not! The solution to this situation is simple. Points which are exactly on the Y threshold must be considered to belong to one side of the threshold. Let’s say we arbitrarily decide that points on the Y threshold will belong to the “above” side of the threshold. Then, side a generates a node, since it has one endpoint below the threshold and its other endpoint on-or-above the threshold. Side b does not generate a node, because both of its endpoints are on-or-above the threshold, so it is not considered to be a threshold-crossing side.
![image](/images/Diagram_4.gif) | Figure 5 shows the case of a polygon in which one of its sides lies entirely on the threshold. Simply follow the rule as described concerning Figure 4. Side c generates a node, because it has one endpoint below the threshold, and its other endpoint on-or-above the threshold. Side d does not generate a node, because it has both endpoints on-or-above the threshold. And side e also does not generate a node, because it has both endpoints on-or-above the threshold.

#### Note

If the test point is on the border of the polygon, this algorithm will deliver unpredictable results; i.e. the result may be “inside” or “outside” depending on arbitrary factors such as how the polygon is oriented with respect to the coordinate system. (That is not generally a problem, since the edge of the polygon is infinitely thin anyway, and points that fall right on the edge can go either way without hurting the look of the polygon.)