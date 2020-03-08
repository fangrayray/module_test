package main

import (
	"fmt"

	pc "github.com/fangrayray/module_test/popcount"
	tc "github.com/fangrayray/module_test/tempconv"
)

func main() {
	fmt.Println("==== START ====")
	fmt.Printf("%v", pc.ShowPopCount())

	fmt.Println()
	x := uint64(259)
	fmt.Printf("%v", pc.PopCount(x))
	fmt.Println()
	fmt.Println(x>>(0*8), ",", byte(x>>(0*8)))
	fmt.Println(x>>(1*8), ",", byte(x>>(1*8)))
	fmt.Println(x>>(2*8), ",", byte(x>>(2*8)))
	fmt.Println(x>>(3*8), ",", byte(x>>(3*8)))
	fmt.Println(x>>(4*8), ",", byte(x>>(4*8)))
	fmt.Println(x>>(5*8), ",", byte(x>>(5*8)))
	fmt.Println(x>>(6*8), ",", byte(x>>(6*8)))
	fmt.Println(x>>(7*8), ",", byte(x>>(7*8)))

	// a, b := get1n2()
	// fmt.Println(fmt.Sprintf("%d, %d", a, b))
	// fmt.Println(fmt.Sprintf("%d, %d", a, b))
	// fmt.Println(fmt.Sprintf("%d, %d", a, b))
	// fmt.Println(fmt.Sprintf("%d, %d", a, b))

	// c, b := get3n4()
	// fmt.Println(fmt.Sprintf("%d, %d", c, b))

}

func tryConv() {
	c := tc.AbsoluteZeroC
	fmt.Printf("%g°K", tc.CToK(c))
	fmt.Println()
	c = tc.FreezingC
	fmt.Printf("%g°K", tc.CToK(c))
	fmt.Println()
	c = tc.BoilingC
	fmt.Printf("%g°K", tc.CToK(c))
	fmt.Println()
}

func get1n2() (int, int) {
	return 1, 2
}

func get3n4() (int, int) {
	return 3, 4
}
