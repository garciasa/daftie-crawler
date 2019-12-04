package domain

import (
	"fmt"
	"log"
	"time"
)

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func printHouse(house House) {
	if house.NewDevelopment {
		fmt.Printf("%s - New Development\n", house.BrandLink)
	} else {
		fmt.Printf("%s - %s - %s - %s\n", house.BrandLink, house.Date, house.Price, house.Meters)
	}
}
