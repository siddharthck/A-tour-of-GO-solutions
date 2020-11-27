package main

import ("golang.org/x/tour/tree"
"fmt")

// Walk walks the tree t sending all values
// from the tree to the channel ch.


/*
have to include another function which does real traversing
because calling a function recursively and closing 
the channel is not posiible in same function
*/
func WalkReal(t *tree.Tree, ch chan int){
// traverse a BST and put the elements in channel (inorder traversal)

	if t == nil {
		return
	}
	
	WalkReal(t.Left, ch)
	//fmt.Println(t.Value)
		ch <- t.Value
	 WalkReal(t.Right,ch)
	


}

// in this function we close the channel
func Walk(t *tree.Tree, ch chan int){

	WalkReal(t,ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.

func Same(t1, t2 *tree.Tree) bool{
ch1:= make(chan int) //to get values from t1
ch2:= make(chan int) // from t2
go Walk(t1,ch1)
go Walk(t2,ch2)

for {

	// if channel is closed,second assignment is set to false
	v1,ok1:= <-ch1
	v2,ok2:= <-ch2
	if (ok1 && ok2) {// if both channels are available
		
		if (v1!=v2){ 
		
			return false
			}
		
	 } else if (!ok1 && !ok2) {// if both are closed at same time (equal number of nodes)
		return true
	} else{// if both are not open or closed at same time( unequal number of nodes)
	
	return false
	}
	
	}






}

func main() {


//testing with each case

fmt.Println(Same(tree.New(10),tree.New(10)))
fmt.Println(Same(tree.New(2),tree.New(1)))
fmt.Println(Same(tree.New(10),tree.New(15)))



}
