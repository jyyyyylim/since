package main

import (
	"fmt"
	"os"
	"time"
)

//intent: cmdln tool that returns duration in hours since present
//intended: 24h format

func main() {
	dh, dm, init := format(os.Args[1])
	//prints to stdout
	fmt.Println(diff(dh, dm, init).Format("15h 04m"))
}

func format(startingtime string) (int, int, time.Time) {
	//why wont golang just adopt strftime?
	const layout string = "15:04"
	fstarting, _ := time.Parse(layout, startingtime)
	dh, dm, _ := fstarting.Clock()
	return dh, dm, fstarting
}

func diff(dh, dm int, fstarting time.Time) time.Time {
	elapsed := time.Now().Add(-time.Minute * time.Duration(dm+(dh*60)))
	return elapsed
}
