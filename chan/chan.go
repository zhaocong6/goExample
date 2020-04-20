package main

import "fmt"

type Trip string

type Driver struct {
	trips []Trip
}

func (d *Driver) SetTrips(trips []Trip) {
	//d.trips = trips
	d.trips = make([]Trip, len(trips))
	copy(d.trips, trips)
}

func main() {
	d1 := &Driver{trips: make([]Trip, 2)}

	trips := make([]Trip, 2)

	d1.SetTrips(trips)

	trips[1] = "1"
	fmt.Println(d1)
	fmt.Println(trips)
}
