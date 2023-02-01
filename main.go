package main

import (
	"errors"
	"log"
	"math"
)

func areaOfCircle(radius float64) (float64, error){

	if radius < 0	{
		err := errors.New("radius can not be negative")
		return -1, err
	}

	pi := math.Pi

	area := pi * math.Pow(radius,2)

	return area, nil
}

func main(){
	log.Println("worlking")

	radius := 4.5

	area, err := areaOfCircle(radius)

	if err != nil{
		log.Fatal(err)
	}

	log.Printf("THe area of the circle of radius %.2f is %.2f", radius, area)


}