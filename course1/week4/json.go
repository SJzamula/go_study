package main
import "fmt"
func main() {
	p1 := Person(
		name: "joe",
		addr: "a st.",
		phone: "123"
	)

	barr, err := json.Marshal(p1) // Marshal() returns JSON representation as []byte
	var p2 Person
	err := json.Unmarshal(barr, &p2) // Object must "fit" JSON [byte]

	dat, e:= ioutil.ReadFile("test.txt")

	dat1 = "Hello, world"
	err := ioutil.WriteFile("outfile.txt", dat, 0777)

	/*
	os.Open()
	os.Close()
	os.Read()
	os.Write()
	*/
	// opening and reading
	f, err := os.Open("dt.txt")
	barr := make([]byte, 10) // now i have empty bite array
	nd, err := f.Read(barr) // if i want to read again, it will read netx 10, so i need to close func after executing 
	f.Close()

	f, err := os.Create("outfile.txt")
	barr := []byte{1,2,3}
	nb, err := f.Write(barr)
	nb, err := f.WriteString("Hi")
}