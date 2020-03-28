package main

import ( "fmt"
		// "errors"
		//  "math"
		 )

// Print

// func main(){
// 	x := 5;
// 	y := 6;
// 	sum := x + y;
// 	fmt.Println(sum);
// }


// IF else


// func main(){
// 	x := 5;

// 	if x > 6 {
// 		fmt.Println("More")
// 	} else {
// 		fmt.Println("les")
// 	}
// }


// Arrays


// func main() {
// 	// var a[5] int

// 	a :=[5]int{1,2,5,9,8}
// 	fmt.Println(a)
// }



// Objects


// func main(){
// 	vertices := make(map[string]int)

// 	vertices["triangle"] =2
// 	vertices["square"] = 1
// 	vertices["dodecagon"] = 12

// 	fmt.Println(vertices);
// }



// Loop


// func main(){
// 	for i:= 0; i < 5; i++ {
// 		fmt.Println(i)
// 	} 
// }


// loop through array


// func main() {
// 	arr := []string{"a", "b", "c"}

// 	for index, value := range arr {
// 		fmt.Println("index", index, "value", value)
// 	}
// }


// loop through object


// func main() {
// 	m := make(map[string]string)
// 	m["a"] = "alpha"
// 	m["b"] = "beta"

// 	for key, value := range m {
// 		fmt.Println("key", key, "value", value)
// 	}
// }


// Functions

// func main() {
// 	result := sum(2,3)
// 	fmt.Println(result)
// }

// func sum(x int, y int) int {
// 	return x + y;
// }


// Multipe return



// func main(){
// 	result, err := sqrt(-16)
	
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println(result)
// 	}
// }

// func sqrt(x float64) (float64, error) {
// 	if x < 0 {
// 		return 0, errors.New("Negative number")
// 	}
// 	return math.Sqrt(x), nil
// }



// Struct

// type person struct {
// 	name string
// 	age int
// }

// func main() {
// 	p := person{name: "Jake", age: 12}
// 	fmt.Println(p)
// }


// Pointer

func main() {
	i := 7
	fmt.Println(&i)
}