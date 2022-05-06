/*
Write a program which prompts the user to first enter a name, and then enter an address. 
Your program should create a map and add the name and address to the map using 
the keys “name” and “address”, respectively. Your program should use Marshal() to create 
a JSON object from the map, and then your program should print the JSON object.
*/
package main
import (
	"encoding/json"
	"fmt"
	"log"
)
func main() {
	person := make(map[string]string)
	var name string
	var address string
	fmt.Printf("Please, enter your name: ")
	fmt.Scan(&name)
	fmt.Printf("Please, enter your address: ")
	fmt.Scan(&address)

	person["name"] = name
	person["address"] = address
	
	var jsonData []byte
	jsonData, err:= json.MarshalIndent(person, "", "\t")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(jsonData))
}