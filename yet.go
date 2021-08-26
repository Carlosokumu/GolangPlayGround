package main

import "fmt"

func Areacalculator(b float64) float64 {
	pie := 3.142
	area := pie * b * 2
	return area
}

//Variadic functions
func add(args ...int) int {
	total := 0
	for _, val := range args {
		total += val
	}
	return total
}
func printer1() {
	fmt.Println("first")
}
func printer2() {
	fmt.Println("second")
}
func Perimeter(l, w int) int {
	return l * w
}

func MarkCalculater(marks []float64) float64 {
	total := 0.0
	for _, value := range marks {
		total += value
		fmt.Println(value)
	}
	return total / float64(len(marks))
}

//pointers
type Circle struct {
	r, d float64
}

func main() {
	informer := "perimeter"
	fmt.Println(informer + string(Perimeter(10, 8)))
	var x string = "carlos"
	fmt.Println(x)
	x = "okumu"
	fmt.Println(x)
	//for loops
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "even")
		} else {
			fmt.Println(i, "odd")
		}
	}
	xs := []float64{98, 93, 77, 82, 83}
	y := 0.0
	for i, value := range xs {
		y += value
		fmt.Println(i)

	}
	fmt.Println(y)
	// returns  numbers divible by 3 btn 1-100 and counts the number of elements
	m := 0
	for j := 0; j < 100; j++ {
		if j%3 == 0 {
			if j > 0 {
				fmt.Println(j)
				m++

			}
		}
	}
	fmt.Println(m)

	//calculate the area
	fmt.Println(Areacalculator(21.0))
	studentsMarks := []float64{10.0, 10.0, 13.0, 10.0, 10.0}
	fmt.Println(MarkCalculater(studentsMarks))
	//variadic  function taking different number of arguments
	fmt.Println(add(1, 2, 3))
	fmt.Println(add(1, 2))
	//closure-creating functions inside functions
	recareaFinder := func(w, l int) (int, int) {
		return 0, w * l
	}
	fmt.Println(recareaFinder(3, 2))
	//function designed within another function-main() function
	k := 0
	multiplier := func() int {
		k++
		return k
	}
	fmt.Println(multiplier())
	salary := []float64{11.3, 12.5, 13.0, 14.5}
	fmt.Println(salary)
	for _, counter := range salary {
		fmt.Println(counter)
	}
	c := Circle{7, 8}
	fmt.Println(c.r, c.d)

	str21 := "hitman21m"
	fmt.Println(str21)
	carl := "hitman21m"
	fmt.Println(carl)
	//defer go has a special keyword "defer" which makes a certain fuction to be run first
	//example
	//defer printer2()
	//printer1()
	//defer func() {
	//	str := recover()
	//	fmt.Println(str)
	//}()
	//panic("PANIC")
	//sum is a function which takes a slice of numbers and adds them together. What would its function signature look like in Go?
	//Answer
	//signature is the return type,params and name of function
	//func sum(mySlice [])int) int {
	//Structs
	p := 15
	pm := &p
	fmt.Println(*pm)

}
