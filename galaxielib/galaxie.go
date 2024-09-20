package galaxielib

import (
	"ex11/etoilelib"
	"ex11/nebuleuselib"
	"fmt"
	"math"
)

//"ex11/astre2lib"

type Galaxy struct {
	name               string
	starSlice          []etoilelib.Star
	sliceNebula        []nebuleuselib.Nebula
	position           []float64
	averageTemperature float64
	diameter           float64
}

func CreateNewGalaxy(name string, starSlice []etoilelib.Star, sliceNebula []nebuleuselib.Nebula, position []float64, averageTemperature float64, diameter float64) Galaxy {
	neb := Galaxy{name, starSlice, sliceNebula, position, averageTemperature, diameter}
	return neb
}

func (k Galaxy) ShowElem() {
	fmt.Println("-----------------------------------------------------------------------------------")
	//fmt.Println("GALAXY/ name:", k.GetName(), ", starSlice: ", k.GetStarSlice(), ", sliceNebula: ", k.GetNebulaSlice(), ", position:", k.GetPosition(), ", diameter:", k.GetDiameter())
	var name string = k.GetName()
	var position []float64 = k.GetPosition()
	var AverageTemperature float64 = k.GetAverageTemperature()
	var Area float64 = k.GetAreaKm2()

	fmt.Println("Elem Galaxy/ name:", name, ", position: ", position, ", Average Temperature:", AverageTemperature, ", Area", Area)
}

func (k Galaxy) GetName() string {
	return k.name
}

func (k Galaxy) GetStarSlice() []etoilelib.Star {
	return k.starSlice
}

func (k Galaxy) GetNebulaSlice() []nebuleuselib.Nebula {
	return k.sliceNebula
}

func (k Galaxy) GetPosition() []float64 {
	return k.position
}

func (k Galaxy) GetAverageTemperature() float64 {
	return k.averageTemperature
}

func (k Galaxy) GetDiameter() float64 {
	return k.diameter
}

func (k *Galaxy) SetNomEtoile(value string) {
	k.name = value
}

func (k *Galaxy) SetSliceStar(value []etoilelib.Star) {
	k.starSlice = value
}

func (k *Galaxy) SetSliceNebula(value []nebuleuselib.Nebula) {
	k.sliceNebula = value
}

func (k *Galaxy) SetAverageTemperature(value float64) {
	k.averageTemperature = value
}

func (k *Galaxy) SetDiameter(value float64) {
	k.diameter = value
}

func (k *Galaxy) FillPosition(valueX float64, valueY float64, valueZ float64) {
	k.position = append(k.position, valueX)
	k.position = append(k.position, valueY)
	k.position = append(k.position, valueZ)
}

func (k *Galaxy) FillNebulaSlice(value nebuleuselib.Nebula) {
	k.sliceNebula = append(k.sliceNebula, value)
}

func (k *Galaxy) FillStarSlice(value etoilelib.Star) {
	k.starSlice = append(k.starSlice, value)
}

func (k Galaxy) GetAreaKm2() float64 {
	var squareOfRadius float64 = math.Pow((k.GetDiameter() / 2), 2)
	var area float64 = math.Pi * squareOfRadius
	return area
}
