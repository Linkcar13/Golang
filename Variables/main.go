//Variables script in Golang

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Golang")

	//Variables
	var (
		myName     string = "Carlos"
		myLastname string
	)

	myLastname = "Estrada"
	//Constantes

	const PI float64 = math.Pi
	fmt.Println(myName, myLastname, PI)

	team := "CyberSecurity"
	fmt.Println(team)

	//Zero Values
	var (
		state  bool
		number float64
		age    int64
		club   string
	)
	fmt.Println(state, number, age, club)

	fmt.Printf("Fullname: %s %s\n", myName, myLastname)
	fullname := fmt.Sprintf("Fullname: %s %s\n", myName, myLastname)
	fmt.Println(fullname)
}
