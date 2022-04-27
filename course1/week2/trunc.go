package main
import "fmt"
func main() {
	fmt.Printf("Please, enter a floating point\n")
	var floatingPoint float32
	var integerPoint int
	fmt.Scan(&floatingPoint)
	integerPoint = int(floatingPoint)
	fmt.Println(integerPoint)
}