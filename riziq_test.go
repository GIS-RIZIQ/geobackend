package geobackend

import (
	"fmt"
	"testing"
)

var dbname = "GIS"
var collname = "GIS"

// func TestGeoIntersects(t *testing.T) {
// 	mconn := SetConnection("MONGOSTRING", dbname)
// 	coordinates := Polygon{
// 		Coordinates: [][][]float64{
// 			{
// 				{100.5918079639361, -0.45832970470556234},
// 				{100.59180519679398, -0.4584272433564678},
// 				{100.5918916699963, -0.4584327774636705},
// 				{100.59189305356739, -0.4583317799957598},
// 				{100.5918079639361, -0.45832970470556234},
// 			},
// 		},
// 	}
// 	datagedung := GeoIntersects(mconn, collname, coordinates)
// 	fmt.Println(datagedung)
// }

func TestGeoWithin(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", dbname)
	coordinates := Polygon{
		Coordinates: [][][]float64{
			{
				{100.5918079639361, -0.45832970470556234},
				{100.59180519679398, -0.4584272433564678},
				{100.5918916699963, -0.4584327774636705},
				{100.59189305356739, -0.4583317799957598},
				{100.5918079639361, -0.45832970470556234},
			},
		},
	}
	datagedung := GeoWithin(mconn, collname, coordinates)
	fmt.Println(datagedung)
}

func TestNear(t *testing.T) {
	mconn := SetConnection2dsphere("MONGOSTRING", "GIS", "GIS")
	coordinates := Point{
		Coordinates: []float64{
			100.59211393336687, -0.45839020260632424,
		},
	}
	datagedung := Near(mconn, "GIS", coordinates)
	fmt.Println(datagedung)
}

func TestNearSphere(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "GIS")
	coordinates := Point{
		Coordinates: []float64{
			100.59203594971524, -0.458401712739942,
		},
	}
	datagedung := NearSphere(mconn, "GIS", coordinates)
	fmt.Println(datagedung)
}

func TestBox(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "GIS")
	coordinates := Polyline{
		Coordinates: [][]float64{
			{100.58376081681058, -0.4665724708199832},
			{100.58461441702951, -0.466613772686415},
		},
	}
	datagedung := Box(mconn, collname, coordinates)
	fmt.Println(datagedung)
}
