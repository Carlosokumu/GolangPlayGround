package main


import "fmt"

//Functions
//first function
func compare(m int, n int) {
	//closure:creating functons inside functions
	add := func(o, p int) int {
		return o + p
	}
	//recursion


	fmt.Println(add(4, 6))
	if m >= n {
		fmt.Println("first argument greater than second number")
	}
	if m < n {
		fmt.Println("first argument is smaller than second argument")
	}

}
func average(testScores []float64) float64 {
	total := 0.0
	for _, val := range testScores {
		total += val
	}
	return total / float64(len(testScores))
}

func main() {
	compare(6, 5)
	var a int = 4
	var ()
	var firstName string = "carlos"
	var secondName string = "Okumu"
	y := "carlos"
	//discovering the for loop trick
	var once string = "once"
	var timer string = "executed " + once
	fmt.Println(timer)
	for a <= 10 {
		var times string = "4 times"
		xs := []float64{98, 93, 77, 82, 83}
		fmt.Println(a)
		fmt.Println(average(xs))
		fmt.Println(firstName)
		fmt.Println(timer + times)
		fmt.Println(secondName)
		a++

	}

	fmt.Println(a)
	fmt.Println(y)
	var b int = 10
	fmt.Println(b)
	var smallerNumber int = 9
	var biggerNumber int = 9
	if biggerNumber == smallerNumber {
		fmt.Println("Equal numbers")
		if (biggerNumber / smallerNumber) == 1 {
			println("Am also executed")
		}
	}
	

}

