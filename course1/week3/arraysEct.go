package main
import "fmt"
func main() {
	var x [5]int
	//By default elements initialined to zero value
	x[0] = 2
	var y [3]int = [3]int{1, 2, 3}
	
	z := [...]int{1, 2, 3, 4}
	fmt.Println(x[1])
	/////////////////////
	a := [3]int {1, 3, 4}
	for i:=0, v range x {
		fmt.Printf("ind %d, val %d", i, v)
	}
	////// Slices. You can change the size of the slice, but you cannot do in with array
	/* Pointer indicates the start of the slice
	 * Length is the number of elts in the slice
	 * Capacity is maximum number of elts
	 */
	arr := [...]string{"a", "b", "c", "d", "e", "f", "g"}
	s1 := arr[1:3]//1 and 2
	s2 := arr[2:5]//2,3,4
	///////////
	a1 := [3]string("a", "b", "c")
	sli1 := a1[0:1]
	fmt.Printf(len(sli1), cap(sli1)) //length and capacity
	// len is 1 (only 1 element in), but capacity is 3 (it is maximum for it)
	sli := []int{1, 2, 3}
	////////////
	//////////// Making slices
	////////////
	sli3 = make([]int, 10)//length and capacity together
	sli4 = make([]int, 10, 15)//capacity is bigger than length
	//it can make a new bigger array if it is nessesary
	sli5 = append(sli, 100)//creates slice with no argument
	//Hash Table
	//Constant key/value pairs
	/*SSN/email
	 *GPS cords/adress
	 *In hash tables it is not nessesary to use numbers of slots, you van use key
	 *Map - golang's implementation of hash tables 
	 */
	 var idMap map[string]int //making var
	 idMap = make(map[string]int) //making actual map
	//or
	idMap1 := map[string]int {
		"joe": 123
	}
	//Accessing Maps
	fmt.Println(idMap1["joe"])
	//add a key value inti the map or change existed
	idMap1["jane"] = 456
	delete(idMap1, "joe")
	id, p := idMap1["joe"] // id == value, p == present (true or false)
	fmt.Println(len(idMap1))//length
	for key, val := range idMap1 {
		fmt.Println(key, val)
	}
	// Struct
	// name, adress, phone
	type struct Person {
		name string
		addr string
		phone string
	}
	var p1 Person
	p1.name = "joe"
	x = p1.addr
	p2 := new(Person)
	p3 := Person(
		name: "joe",
		addr: "a st",
		phone: "123"
	)
}