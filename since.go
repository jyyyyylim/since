package main

import (
	"fmt"
	"os"
	"time"
)

//intent: cmdln tool that returns duration in hours since present
//intended: 24h format

func main() {

	var dh, dm int

	if len(os.Args) > 2 {
		dh, dm = tformat(os.Args[1]+" "+os.Args[2], int16(len(os.Args)))
	} else {
		dh, dm = tformat(os.Args[1], int16(len(os.Args)))
	}
	//prints to stdout

	fmt.Println(diff(dh, dm).Format("15h 04m"))

	return
}



func tformat(startingtime string, arglen int16) (int, int) {

	//why wont golang just adopt strftime?
	var layout string

	if arglen > 2 {
		layout = "15 04"
	} else {
		layout = "15:04"
	}

	fstarting, _ := time.Parse(layout, startingtime)
	dh, dm, _ := fstarting.Clock()

	return dh, dm
}



func diff(dh, dm int) time.Time {

	elapsed := time.Now().Add(-time.Minute * time.Duration(dm+(dh*60)))

	return elapsed
}
