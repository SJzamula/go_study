/*  CLASSES
	Collection of data fields and functions that share a well-defined responsibility
	Example: Point class
		- used in a geometry program
		- Data: x, y coordinates
		- Functions
			* DistToOrigin(), Quadrant()
			* AddXOffset(), AddYOffset()
			* SetX(), SetY()
	Classes are a template
	Contain data fields, not data

	OBJECT
	Instance of a class
	Contains real data

	ENCAPSULATION
	Data can be protected from the programmer
	Data can be accepted only using methods
	Maybe we don't trust the programmer to keep data consistent
	Example: Double distance to origin
	- option 1: Make method DoubleDist()
	- option 2: Trust programmer to double X and Y directly

	SUPPORT FOR CLASSES
	NO "Class" Keyword
	* Most languages have a class keyword
	* Data fields and methods aer defined inside a class block
	Associating Methods with Data
	* Method has a reciever type that it is associated with
	* Use dot notation to call the method

	SUPPORT FOR CLASSES (2)
	Structs
	* Struct types compose data fields
	Struncts with Methods
	* Structs and methods together allow arbitary data and functions to be composed

*/
func (p Point) DistToOrigin() {
	t := math.Pow(p.x, 2) + math.Pow(p.y, 2)
	return math.Sqrt(t)
}

type Point struct {
	x float64
	y float64
}

type MyInt int //type has to be in the same package

func (mi MyInt) Double () int {
	return int(mi * 2)
}

func main() {
	v: MyInt(3)
	fmt.Println(v.Double())

	p1 := Point(3,4)
	fmt.Println(p1.DistToOrigin())
}