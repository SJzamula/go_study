/*
Write a Bubble Sort program in Go. The program
should prompt the user to type in a sequence of up to 10 integers. The program
should print the integers out on one line, in sorted order, from least to
greatest. Use your favorite search tool to find a description of how the bubble
sort algorithm works.

As part of this program, you should write a
function called BubbleSort() which
takes a slice of integers as an argument and returns nothing. 
The BubbleSort() function should modify the slice so that the elements are in sorted
order.

A recurring operation in the bubble sort algorithm is
the Swap operation which swaps the position of two adjacent elements in the
slice. You should write a Swap() function which performs this operation. Your Swap()
function should take two arguments, a slice of integers and an index value i which
indicates a position in the slice. The Swap() function should return nothing, but it should swap
the contents of the slice in position i with the contents in position i+1.
*/

package main
import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"strconv"
)
func Swap(slice[]int, pos int) {
	tmp := slice[pos]
	slice[pos] = slice[pos+1]
	slice[pos+1] = tmp
}
func BubbleSort(slice[]int){
	var i, j int
	for i = 0; i < len(slice)-1; i++ {
		for j = 0; j < len(slice)-i-1; j++ {
			if slice[j] > slice[j+1] {
				Swap(slice, j)
			}
		}
	}
}
func main() {
	var userInput string
	in := bufio.NewScanner(os.Stdin)
	fmt.Printf("Input your data(numbers divided by ' '): ")
	in.Scan()
	userInput = in.Text()

	sliStr := strings.Split(userInput, " ")
	sliNums := make([]int, len(sliStr))
	var i int
	for i = 0; i < len(sliStr); i++ {
		num, err := strconv.Atoi(sliStr[i])
    	if err != nil {
        	fmt.Println(err)
    	}
		sliNums[i] = num
	}

	BubbleSort(sliNums)
	fmt.Print("\nSorted numbers\t", sliNums)
}