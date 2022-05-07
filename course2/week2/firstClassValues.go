//first class values
/*  FUNCTIONS ARE FIRST CLASS
	1. Functions can be treated like other types
		* Variables van be declared with a function type
		* Can be created dynamically
		* Can be passed as argument and returnes as values
		* Can be stored in data structures
	2. Variables as functions
		* Declare a variable as a func
*/
var funcVar func(int) int
func incFn (x int) int {
	return x+1
}
func main() {
	func = incFn
	fmt.Print(funcVar(1))
}
/*
		* Function is on ringht-hand side, without ()
	3. Functions as arguments
		* Function can be passed to another function as an argument
*/
func applyIt(afunct func (int) int, val int) int {
	return afunct(val)
}
func incFn(x int) int {return x + 1}
func decFn(x int) int {return x - 1}
func main() {
	fmt.Println(applyIt(incFn, 2))
	fmt.Println(applyIt(decFn, 2))
	// anonymus
	v := applyIt(func (x int) int {return x + 1}, 2) // has declaration, but without name
}
/*
	4. Anonymous Functions
		* Don't need to name a function
*/