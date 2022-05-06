/* Write a program which prompts the user to enter a string. 
 * The program searches through the entered string for the characters 
 * ‘i’, ‘a’, and ‘n’. The program should print “Found!” 
 * if the entered string starts with the character ‘i’, ends with 
 * the character ‘n’, and contains the character ‘a’. 
 * The program should print “Not Found!” otherwise. 
 * The program should not be case-sensitive, so it does not matter if 
 * the characters are upper-case or lower-case.
 */
 package main
 import (
	 "fmt"
	 "strings"
 )
 func main() {
	fmt.Printf("Please, enter your input\n")
	var str string
	fmt.Scan(&str)
	if (strings.ContainsAny(str, "a") || strings.ContainsAny(str, "A")) && 
	(strings.HasPrefix(str, "i") || strings.HasPrefix(str, "I")) &&
	(strings.HasSuffix(str, "n") || strings.HasSuffix(str, "N")) {
		fmt.Printf("Found!\n")
	} else {
		fmt.Printf("Not Found\n")
	}
 }