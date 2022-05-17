/*
	Limitation of Methods
	* Recriver is passed implicity as an argument to the method
	* Method cannot modify the data inside the receiver
	* Example OffsetX() should increase coordinate

	Large Receivers
	* If receiver is large, lots of copying is required

	No Need to Dereference
	Dereferencing is automatic
*/
type Image [100][100] int
func (p *Point) OffsetX (v float64) { // with *
	p.x = p.x + v // doesn't need * one more time
}

func main() {
	p1:= Point(3, 4)
	p1.OffsetX(5) //don't need reference when calling the method
	//doesnt change x, but change a copy of x

	i1 := GrabImage()
	i1.BlurImage()
	// 10,000 ints copied to BlurImage
	// waste a lot of time to copy everyting
}