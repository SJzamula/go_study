package main
import "fmt"
func main() {
	var x int = 10
	var y int 
	var ip *int // ip is pointer to int

	ip = &x
	y = *ip
	/////////////
	var a int32 = 1
	var b int16 = 2
	a = int32(b) // I cannot put a = b, because of diffrent types
	fmt.Println(y)
	fmt.Println(a)
	/* float32 6 digits of precision
	 * float64 15 digits of precision
	 */
	// Strings. When you do actions with strings, it does't modify string, it returns a new string

	const ab = 1.2
	const (
		yz = "Hi"
		name = 5
	)
	// constant definition must be diffrent
	const (
		A int = iota // makes marks diffrent, but you dont have to give value to it now
		B 
		C 
		D 
		F 
	)

	if x == y {
		fmt.Printf("You are ringt!\n")
	}
	for i := 0; i < x; i++ {
		fmt.Println(a)
	}

	var z = 1
	switch z {
	case 1:
		fmt.Printf("case 1\n")
	case 2:
		fmt.Printf("case 2\n")
	default:
		fmt. Printf("nocase\n")
	}

	switch {
	case x > 1:
		fmt.Printf("case 1\n")
	case x < -1:
		fmt.Printf("case 2\n")
	default:
		fmt. Printf("nocase\n")
	}

	i:=0
	for i < 10 {
		i++
		if i == 5 { break }
		fmt.Printf("hi ")
	}
	//Scan
	var appleNum int

	fmt.Printf("Nimber of apples?")
	fmt.Scan(&appleNum)
	fmt.Println(appleNum)
}