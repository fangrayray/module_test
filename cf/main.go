package main

import (
	"fmt"
	"os"
	"strconv"

	tc "github.com/fangrayray/module_test/tempconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v \n", err)
			os.Exit(1)
		}
		k := tc.Kelvin(t)
		c := tc.Celsius(t)
		fmt.Printf("%s = %s, %s = %s \n", k, tc.KToC(k), c, tc.CToK(c))
	}
}
