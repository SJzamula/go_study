/*
	ENCAPSULATION and keep private data
	Controlling Access (You will control how much people can use your data)
	* Can define public functions to allow access to hidden data

	Controlling Access to Structs
	* Hide fields of structs by starting field name with a lower-case letter
	* Define public methods which access hidden data
*/
package data
type Point struct {
	x float64
	y float64
}
func (p *Point) InotMe(xn, yn float64) {
	p.x = xn
	p.y = yn
}
func (p *Point) Scale(v float64) {
	p.x = p.x * v 
	p.y = p.y * v 
}
func (p *Point) PrintMe() {
	fmt.Println(p.x, p.y)
}

var x int = 1
func PrintX() {
	fmt.Println(x)
	//data can access PrincX, even if cannot access directly x
	//you can access x, only using PrintX()

}