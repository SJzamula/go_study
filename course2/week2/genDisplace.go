package main
import (
	"fmt"
	"bufio"
	"strconv"
	"os"
	"strings"
	"math"
)

func GenDisplaceFn(a, v, s float64) func (float64) float64 {
	fn := func (t float64) float64 {
		return 0.5 * a * (math.Pow(t, 2)) + v * t + s
	}
	return fn
}
func main(){
	// Scan values of acceleration, initial velocity, and initial displacement
	input  := bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter acceleration, initial velocity, and initial displacement: ")
	input.Scan()
	userInput := input.Text()

	sliStr := strings.Split(userInput, " ")
	accel, err := strconv.ParseFloat(sliStr[0], 64)
	velocity, err := strconv.ParseFloat(sliStr[1], 64)
	displ, err := strconv.ParseFloat(sliStr[2], 64)
    if err != nil {
       	fmt.Println(err)
    }

	// Scan the value of time
	var time float64
	fmt.Printf("Enter time: ")
	fmt.Scan(&time)

	funcWithTime := GenDisplaceFn(accel, velocity, displ)
	fmt.Println("The result: ", funcWithTime(time))
}