package main
import (
	"fmt"
	"math"
)
/*  FUNCTIONS AS RETURN VALUES
	1. Functions can return functions
	2. Might create a function with controllable parametrs
		* Takes a point (x, y, coordinates)
		* Return distance to origin
	3. What if I want to change the origin
		* Option 1: Pass origin as argument
		* Option 2: Define a function with new origin

	ENVIRONMENT OF A FUNCTION
	1. Set names that are valid inside a function
	2. Names defined locally, tn the function
	3. Lexical Scoping

	CLOSURE = Function + its environment
	1. When functions are passed/turnes, their environment comes with them
*/
func MakeDistOrigin (o_x, o_y float64) func (float64, float64) float64 {
		fn := func (x, y float64) float64 {
			return math.Sqrt(math.Pow(x - o_x, 2) + math.Pow(y - o_y, 2))
		}
	return fn
}
func main() {
	Dist1 := MakeDistOrigin(0,0)
	Dist2 := MakeDistOrigin(2,2)
	fmt.Println(Dist1(2,2))
	fmt.Println(Dist2(2,2))
}