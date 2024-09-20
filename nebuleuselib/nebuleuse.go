package nebuleuselib

import (
	"fmt"
	"math"
)

type Nebula struct {
	name               string
	diametre           float64
	gazType            string
	position           []float64
	averageTemperature float64
}

func CreateNewNebula(name string, diametre float64, typeGaz string, position []float64, averageTemperature float64) Nebula {
	neb := Nebula{name, diametre, typeGaz, position, averageTemperature}
	return neb
}

func (k Nebula) ShowElem() {
	fmt.Println("-----------------------------------------------------------------------------------")
	//fmt.Println("NEBULA/ nom:", k.GetName(), ", diametre: ", k.GetDiameter(), ", position: ", k.GetPosition(), ", gazType: ", k.GetGazType(), ", Average Temperature: ", k.GetAverageTemperature())
	var name string = k.GetName()
	var position []float64 = k.GetPosition()
	var AverageTemperature float64 = k.GetAverageTemperature()
	var Area float64 = k.GetAreaKm2()

	fmt.Println("Elem Nebula/ name:", name, ", position: ", position, ", Average Temperature:", AverageTemperature, ", Area", Area)

}

func (k Nebula) GetName() string {
	return k.name
}

func (k Nebula) GetDiameter() float64 {
	return k.diametre
}

func (k Nebula) GetGazType() string {
	return k.gazType
}

func (k Nebula) GetPosition() []float64 {
	return k.position
}

func (k Nebula) GetAverageTemperature() float64 {
	return k.averageTemperature
}

func (k *Nebula) SetName(value string) {
	k.name = value
}

func (k *Nebula) SetDiametre(value float64) {
	k.diametre = value
}

func (k *Nebula) SetGazType(value string) {
	k.gazType = value
}

func (k *Nebula) SetPosition(value []float64) {
	k.position = value
}

func (k *Nebula) SetAverageTemperature(value float64) {
	k.averageTemperature = value
}

func (k *Nebula) FillPosition(valueX float64, valueY float64, valueZ float64) {
	k.position = append(k.position, valueX)
	k.position = append(k.position, valueY)
	k.position = append(k.position, valueZ)
}

func (k Nebula) GetAreaKm2() float64 {
	var squareOfRadius float64 = math.Pow((k.GetDiameter() / 2), 2)
	var area float64 = math.Pi * squareOfRadius
	return area
}
