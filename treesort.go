/*
	@saulpanders
	treesort.go -- testing recursive data structs & tree review

*/

package main

import (
	//"flag"
	"fmt"
	"os"
	"strconv"
)

// basic tree struct
type tree struct {
	val         int
	left, right *tree
}

//to use, type ints into command line you want sorted (seperated by spaces)
//will return low to high
func main() {

	var values []int

	//if no input, print usage
	if len(os.Args) == 1 {
		fmt.Println("Usage: treesort <int_1> <int_2>.... <int_n>")

	} else {
		for i := 1; i < len(os.Args); i++ {
			//strconv library handles int-string stuff
			arg, err := strconv.Atoi(os.Args[i])
			if err != nil {
				//handle error
				fmt.Println(err)
				os.Exit(2)
			}
			//append is a built in modifier, sweet!
			values = append(values, arg)
		}

		//treesort our array
		Sort(values)
		fmt.Println(values)
		//sorts values in place
	}
}
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

func add(t *tree, value int) *tree {
	if t == nil {
		//equivalent to return &tree{value: value}
		t = new(tree)
		t.val = value
		return t
	}

	if value < t.val {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.val)
		values = appendValues(values, t.right)
	}
	return values
}
