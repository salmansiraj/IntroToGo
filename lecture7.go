package main

import (
	"fmt"
)

func birthDate(year int) int {
	lastYear := year
	leap := 1
	days := 0
	for lastYear != 2018 {
		if lastYear != 2018 {
			if leap != 4 {
				lastYear++
				leap++
				days += 365
			} else {
				lastYear++
				leap -= 3
				days += 366
			}
		}
	}
	return days
}

func birthDate2(birthYear int, presentYear int) int {
	return (((presentYear - birthYear - (presentYear / birthYear)) * 365) + ((presentYear / birthYear) * 366))
}

func main() {
	birthYear := 1999
	fmt.Println(birthDate(birthYear))
	fmt.Println(birthDate2(birthYear, 2018))
}
