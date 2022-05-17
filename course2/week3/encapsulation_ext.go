package main
import "data"
func main() {
	data.PrintX()
	var p data.Point
	p.InotMe(3, 4)
	p.Scale(2)
	p.PrintMe()
	//Access to hidden fields only through public methods
	
}