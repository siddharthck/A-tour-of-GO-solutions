package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	index:=0
	first := 0
	second := 1
	return func() int{
	switch(index){
		case 0:
			index++
			
			return 0
		case 1:
			index++
			return 1
		default:
			index++
			valueToReturn := first + second
			first = second
			second = valueToReturn
			return valueToReturn
			
		}
	}
	
	
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
