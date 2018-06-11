package main

import "fmt"

type shape interface {
	getArea() float64
}

//declaration of type struct with fields
type triangle struct {
	base   float64
	height float64
}
type square struct {
	sideLength float64
}

func main() {
	//need to pass field values
	t := triangle{base: 10, height: 10}
	sq := square{sideLength: 10}

	printArea(t)
	printArea(sq)

}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (t triangle) getArea() float64 {
	fmt.Println("Area of Triangle:")
	return 0.5 * t.base * t.height
}

//getArea does not take any argument
func (sq square) getArea() float64 {
	fmt.Println("Area of Square:")
	//we care about receiving value "sq"
	return sq.sideLength * sq.sideLength
}
