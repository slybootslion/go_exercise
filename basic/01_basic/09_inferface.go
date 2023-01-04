package main

import (
	"fmt"
	"math"
)

func main() {
	truckSize := 0.0
	var materials []interface{}
	materials = append(materials, cube{12.5})
	materials = append(materials, cuboid{25, 13, 60})
	materials = append(materials, cylinder{5, 25.3})
	for _, singleMaterial := range materials {
		truckSize += calcSize(singleMaterial)
	}
	fmt.Println(truckSize)
}

type cube struct {
	length float64
}

func (c *cube) cubeVolume() float64 {
	return c.length * c.length * c.length
}

type cuboid struct {
	length float64
	width  float64
	height float64
}

func (c *cuboid) cuboidVolume() float64 {
	return c.length * c.width * c.height
}

type cylinder struct {
	diameter float64
	height   float64
}

func (c cylinder) cylinderVolume() float64 {
	return math.Pi * (c.diameter / 2) * (c.diameter / 2) * c.height
}

func calcSize(material interface{}) float64 {
	cubeMaterial, cubeOk := material.(cube)
	if cubeOk {
		return cubeMaterial.cubeVolume()
	}
	cuboidMaterial, cuboidOk := material.(cuboid)
	if cuboidOk {
		return cuboidMaterial.cuboidVolume()
	}
	cylinderMaterial, cylinderOk := material.(cylinder)
	if cylinderOk {
		return cylinderMaterial.cylinderVolume()
	}
	return 0
}
