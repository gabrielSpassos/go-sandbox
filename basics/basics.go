package main

import (
	"fmt"
	"math"
)

func main() {
	numbersTypes()
	operators()
	stringsTypes()
	booleansTypes()
	variables()
	loopsAndConditions()
	collections()
	structs()
	embeddedStructs()
}

func numbersTypes() {
	var a int = 1
	fmt.Println("int", a)

	var b uint8 = 4
	fmt.Println("uint8", b)

	var x byte = 4
	fmt.Println("byte", x)

	var c uint16 = 43545
	fmt.Println("uint16", c)

	var d uint32 = 434766734
	fmt.Println("uint32", d)

	var e uint64 = 545454523342432
	fmt.Println("uint64", e)

	var f int8 = -4
	fmt.Println("int8", f)

	var g int16 = -4354
	fmt.Println("int16", g)

	var h int32 = -435443223
	fmt.Println("int32", h)

	var j rune = -435443223
	fmt.Println("rune", j)

	var k int64 = -4547674712113213213
	fmt.Println("int64", k)

	var l float32 = -44545.656
	fmt.Println("float32", l)

	var m float64 = -43424242.43222
	fmt.Println("float64", m)
}

func operators() {
	fmt.Println("1 + 1 =", 1+1)
	fmt.Println("4.5 - 3 =", 4.5-3)
	fmt.Println("8 * 7 =", 8*7)
	fmt.Println("4 / 2 =", 4/2)
	fmt.Println("4 % 2 =", 4%2)
}

func stringsTypes() {
	fmt.Print("\tThis is a \n")
	fmt.Print(`\tTest \n`)
	fmt.Println("\t| Lenght of `Gabriel` is:", len("Gabriel"))
}

func booleansTypes() {
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}

func variables() {
	var name string = "Gabriel"
	age := 26
	fmt.Println("Hello", name, "age", age)

	const pi float32 = 3.14
	fmt.Println("PI value:", pi)

	var (
		a = 5
		b = 10
		c = 15
	)
	fmt.Println("A:", a)
	fmt.Println("B:", b)
	fmt.Println("C:", c)
}

func loopsAndConditions() {
	for i := 1; i <= 5; i++ {
		if i%2 == 0 {
			fmt.Println(i, "even")
		} else {
			fmt.Println(i, "odd")
		}
	}

	j := 10
	for j <= 15 {
		if j%2 == 0 {
			println(j, "divisible by 2")
		} else if j%3 == 0 {
			println(j, "divisible by 3")
		} else if j%4 == 0 {
			println(j, "divisible by 4")
		}

		j = j + 1
	}

	for i := 0; i <= 6; i++ {
		switch i {
		case 0:
			fmt.Println("Zero")
		case 1:
			fmt.Println("One")
		case 2:
			fmt.Println("Two")
		case 3:
			fmt.Println("Three")
		case 4:
			fmt.Println("Four")
		case 5:
			fmt.Println("Five")
		default:
			fmt.Println("Unknown Number")
		}
	}
}

func collections() {
	// array
	var grades [5]float64
	grades[0] = 100
	grades[1] = 54.6
	grades[2] = 79.6
	grades[3] = 61
	grades[4] = 43.9

	var total float64 = 0
	for i := 0; i < len(grades); i++ {
		total += grades[i]
	}
	var average float64 = total / float64(len(grades))
	fmt.Println("Grades average:", average)

	x := [5]float64{98, 93, 77, 82, 83}
	y := [5]float64{
		98,
		93,
		77,
		82,
		83,
	}
	fmt.Println(x, y)

	// slice
	var a []float64
	b := []int{1, 2, 3}
	c := append(b, 4, 5)
	d := make([]int, 2)
	fmt.Println(a, b, c, d)
	copy(d, b)
	fmt.Println(a, b, c, d)

	// maps
	var keyValue map[string]int = make(map[string]int)
	keyValue["key"] = 12
	fmt.Println(keyValue)
	value, ok := keyValue["foo"]
	fmt.Println(value, ok)
}

type Circle struct {
	x float64
	y float64
	r float64
}

type Shape interface {
	area() float64
}

type Retangle struct {
	x1, y1, x2, y2 float64
}

type Triangle struct {
	a, b, c float64
}

func structs() {
	var c Circle
	fmt.Println(c)

	c2 := new(Circle)
	fmt.Println(c2)

	c3 := Circle{x: 0, y: 0, r: 5}
	fmt.Println(c3)

	c4 := Circle{0, 0, 5}
	fmt.Println(c4)

	fmt.Println(calculateCircleArea(c4))

	r := Retangle{0, 0, 10, 8}
	fmt.Println(r.area())

	t := Triangle{8, 0, 12}
	fmt.Println(t.area())

	fmt.Println(totalArea(&r, &t))
}

func calculateCircleArea(c Circle) float64 {
	return math.Pi * c.r * c.r
}

func (r *Retangle) area() float64 {
	base := r.x2 - r.x1
	height := r.y2 - r.y1

	return base * height
}

func (t *Triangle) area() float64 {
	base := t.c - t.b
	height := t.a

	return (base * height) / 2
}

func totalArea(shapes ...Shape) float64 {
	var totalArea float64
	for _, shape := range shapes {
		totalArea += shape.area()
	}
	return totalArea
}

type Person struct {
	Name string
}

type Android struct {
	Person
	Model string
}

func embeddedStructs() {
	p := Person{"slim shady"}
	p.talk()

	a := Android{Person: p, Model: "2050"}
	a.talk()
}

func (p *Person) talk() {
	fmt.Println("Hi my name is", p.Name)
}
