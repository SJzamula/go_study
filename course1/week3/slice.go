/* Write a program which prompts the user to enter integers and stores 
 the integers in a sorted slice. The program should be written as a loop. 
 Before entering the loop, the program should create an empty integer 
 slice of size (length) 3. During each pass through the loop, 
 the program prompts the user to enter an integer to be added to the slice. 
 The program adds the integer to the slice, sorts the slice, and prints 
 the contents of the slice in sorted order. The slice must grow in size 
 to accommodate any number of integers which the user decides to enter. 
 The program should only quit (exiting the loop) when the user 
 enters the character ‘X’ instead of an integer. */

 package main
import "fmt"

func main() {
	sli := make([]int, 1, 3)
	fmt.Printf("It is the program, whih will add your input integer " +
			    "and sort from the smallest number to biggest. " +
				"It will output new slice with each iteration. " +
				"To stop iterating, input 'x', to continue, press any letter.\n")
	var cont string
	var num int
	var i int = 0
	for cont != "X"{
		if i == len(sli) {
			newSli := make([]int, len(sli)+1)
			copy(newSli, sli)
			sli = newSli
		}
		fmt.Printf("Input: ")
		fmt.Scanln(&num)
		sli[i] = num
		var j int = 0
		for j<len(sli)-1 {
			
			var tmp int
			if sli[len(sli)-1-j] < sli[len(sli)-2-j]  {
				tmp = sli[len(sli)-1-j]
				sli[len(sli)-1-j] = sli[len(sli)-2-j]
				sli[len(sli)-2-j] = tmp
			}
			j++
		}
		fmt.Print(sli)

		fmt.Printf("\nShould we continue?(y/X)")
		fmt.Scan(&cont)
		i++
	}
	fmt.Printf("End of slice\n")
}