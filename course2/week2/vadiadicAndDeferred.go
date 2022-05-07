package main
import "fmt"
//VARIADIC AND DEFERRED

/*  VARIABLE ARGUMENT NUMBER
	1. Functions can take a variable number of arguments
	2. Use ellipsis ... to specify
	3. Treated as a slice inside function

	VARIADIC SLICE ARGUMENT
	1. Can pass a slice to a variadic function
	2. Need the ... sufix

	DEFERRED FUNCTION CALLS
	1. Call can be deferred untill the surounding function completes
	2. Typically used for cleanup activities
	3. Arguments of a deferred call are evaluated immediately, (but the call late)
*/
func getMax(vals ...int) int { // is treated like slice of integers
	maxV := -1
	for _, v := range vals {
		if v > maxV {
			maxV = v
		}
	}
	return maxV
}
func main() {
	defer fmt.Println ("\nBye!")
	fmt.Println ("Hello!\n")

	fmt.Println(getMax(1, 6, 3, 4))
	vslice := []int{1, 7, 6, 2}
	fmt.Println(getMax(vslice...))
}

