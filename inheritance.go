package main

import "fmt"

type Car struct {
    wheelCount int 
}

// define a behavior for Car
func (car Car) numberOfWheels() int {
    return car.wheelCount
}

type Ferrari struct {
    Car //anonymous field Car
}

func main() {
    f := Ferrari{Car{4}}
    fmt.Println("A Ferrari has this many wheels: ", f.numberOfWheels()) //no method defined for Ferrari, but we have the same behavior as Car.
}