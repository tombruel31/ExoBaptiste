package etoilelib

import (
	"fmt"
	"math"
)

type Star struct {
	name               string
	diameter           float64
	starType           string
	position           []float64
	averageTemperature float64
}

func CreateNewStar(name string, diameter float64, starType string, position []float64, averageTemperature float64) Star {
	neb := Star{name, diameter, starType, position, averageTemperature}
	return neb
}

func (k Star) ShowElem() {
	fmt.Println("-----------------------------------------------------------------------------------")
	//fmt.Println("STAR/ name:", k.GetName(), ", diameter: ", k.GetDiameter(), ", position: ", k.GetStarType(), ", StarType:", k.GetPosition(), ", Average Temperature:", k.GetAverageTemperature())
	var name string = k.GetName()
	var position []float64 = k.GetPosition()
	var AverageTemperature float64 = k.GetAverageTemperature()
	var Area float64 = k.GetAreaKm2()

	fmt.Println("Elem Star/ name:", name, ", position: ", position, ", Average Temperature:", AverageTemperature, ", Area", Area)

}

func (k Star) GetName() string {
	return k.name
}

func (k Star) GetStarType() string {
	return k.starType
}

func (k Star) GetDiameter() float64 {
	return k.diameter
}

func (k Star) GetPosition() []float64 {
	return k.position
}

func (k Star) GetAverageTemperature() float64 {
	return k.averageTemperature
}

func (k *Star) SetName(value string) {
	k.name = value
}

func (k *Star) SetDiameter(value float64) {
	k.diameter = value
}

func (k *Star) SetStarType(value string) {
	k.starType = value
}

func (k *Star) SetPosition(value []float64) {
	k.position = value
}

func (k *Star) SetAverageTemperature(value []float64) {
	k.position = value
}

func (k *Star) FillPosition(valueX float64, valueY float64, valueZ float64) {
	k.position = append(k.position, valueX)
	k.position = append(k.position, valueY)
	k.position = append(k.position, valueZ)
}

func (k Star) GetAreaKm2() float64 {
	var squareOfRadius float64 = math.Pow((k.GetDiameter() / 2), 2)
	var area float64 = math.Pi * squareOfRadius
	return area
}
