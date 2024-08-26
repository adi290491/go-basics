package main

import (
	"flag"
	"fmt"
	"strconv"
	"temp-conv/lenconv"
	"temp-conv/tempconv"
	"temp-conv/weightconv"
)

func main() {
	flag.Parse()
	t, err := strconv.ParseFloat(flag.Arg(0), 64)

	if err != nil {
		fmt.Println("Enter a value: ")
		fmt.Scan(&t)
	}


	c := tempconv.Celsius(t)
	f := tempconv.Fahrenheit(t)

	fmt.Printf("Temperature in F: %.2f\u00B0F\n", tempconv.CToF(c))
	fmt.Printf("Temperature in C: %.2f\u00B0F\n", tempconv.FToC(f))

	ft := lenconv.Feet(t)
	m := lenconv.Meter(t)

	fmt.Printf("Length in feet: %.2fFt\n", lenconv.MtoFt(m))
	fmt.Printf("Length in meter: %.2fm\n", lenconv.FtToM(ft))

	kg := weightconv.Kilogram(t)
	lb := weightconv.Pound(t)

	fmt.Printf("Weight in Kg: %.2f Kg\n", weightconv.KgToLb(kg))
	fmt.Printf("Weight in Pounds: %.2f lbs\n", weightconv.LbToKg(lb))
}
