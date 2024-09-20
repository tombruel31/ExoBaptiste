package main

import (
	"ex11/etoilelib"
	"ex11/galaxielib"
	"ex11/nebuleuselib"
	"math"
)

const (
	emptyLightCelerity = 299792458
	lightYearMeter     = 9460730473000000 //9,460730473.1015 m
)

type CelestialOb interface {
	GetName() string
	GetPosition() []float64
	GetAverageTemperature() float64 // Donne la temperature moyenne de l'objet
	GetAreaKm2() float64
	ShowElem() // petit ick, le fait de faire une fonction commune a chaque classe, fait que l'on filtre certains parametres cool
	//alternative, faire une autre methode show pour chauqe class comme au début

}

func main() {
	////////////////////////////////////////////////////////////////Struct/Class Part/////////////////////////////////////////////////////////////////////////////////
	var nameStar1 string = "Sun"
	var diameterStar1 float64 = 1.3927 * math.Pow(10, 4)
	var typeStar1 string = "Yellow Dwarf"
	var positionStar1 = []float64{}
	var averageTemperatureStar1 float64 = 5500
	var star1 etoilelib.Star = etoilelib.CreateNewStar(nameStar1, diameterStar1, typeStar1, positionStar1, averageTemperatureStar1)
	var valueXStar1, valueYStar1, valueZStar1 float64 = 1.0, 4.9, 99.1
	star1.FillPosition(valueXStar1, valueYStar1, valueZStar1)
	star1.ShowElem()

	var nameNebula1 string = "Lyra Nebula" //M57
	var diameterNebula1 float64 = 1.3 * lightYearMeter
	var gazType1 string = "hydrogen"
	var positionNebula1 = []float64{}
	var averageTemperatureNebula1 float64 = 0.0 //pas possible de trouver dans ce cas
	var nebula1 nebuleuselib.Nebula = nebuleuselib.CreateNewNebula(nameNebula1, diameterNebula1, gazType1, positionNebula1, averageTemperatureNebula1)
	var valueXNebula1, valueYNebula1, valueZNebula1 float64 = 99769854327219.1, 4867494786878.9, 1.46587568888888888870
	nebula1.FillPosition(valueXNebula1, valueYNebula1, valueZNebula1)
	nebula1.ShowElem()

	var nameGalaxy1 string = "Andromeda"
	var sliceNameStarGalaxy1 = []etoilelib.Star{}
	var sliceNameStarNebula1 = []nebuleuselib.Nebula{}
	var valueXGalaxy1, valueYGalaxy1, valueZGalaxy1 float64 = 86946854543354356.0, 468448548687367354785.0, 32548456666664597987987.0
	var positionGalaxy1 = []float64{valueXGalaxy1, valueYGalaxy1, valueZGalaxy1}
	var diameterGalaxie1 float64 = 1.3 * lightYearMeter
	var averageTemperatureGalaxy1 float64 = 0.0 //pas possible de trouver dans ce cas
	var galaxie1 galaxielib.Galaxy = galaxielib.CreateNewGalaxy(nameGalaxy1, sliceNameStarGalaxy1, sliceNameStarNebula1, positionGalaxy1, averageTemperatureGalaxy1, diameterGalaxie1)
	galaxie1.ShowElem()
	galaxie1.FillStarSlice(star1)
	galaxie1.FillNebulaSlice(nebula1)
	galaxie1.ShowElem()

	////////////////////////////////////////////////////////////////Interface Part/////////////////////////////////////////////////////////////////////////////////
	var diameterGalaxie2 float64 = 1.3 * lightYearMeter
	var nameGalaxy2 string = "Black Eye "
	var sliceNameStarGalaxy2 = []etoilelib.Star{}
	var sliceNameStarNebula2 = []nebuleuselib.Nebula{}
	var valueXGalaxy2, valueYGalaxy2, valueZGalaxy2 float64 = 777777777777777777777, 333333333333333333333333333333, 99999
	var positionGalaxy2 = []float64{valueXGalaxy2, valueYGalaxy2, valueZGalaxy2}
	var averageTemperatureGalaxy2 float64 = 0.0 //pas possible de trouver dans ce cas

	var Obj CelestialOb
	var someGalaxy galaxielib.Galaxy = galaxielib.CreateNewGalaxy(nameGalaxy2, sliceNameStarGalaxy2, sliceNameStarNebula2, positionGalaxy2, averageTemperatureGalaxy2, diameterGalaxie2) // Assuming you have a Galaxy object created elsewhere
	Obj = someGalaxy

	Obj.ShowElem()
}
