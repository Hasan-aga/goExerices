package binarysearch

import (
	"fmt"
	"slices"
)

// function gets list of elements and a search element [3,1,2], 1
// sort list ascending 3,1,2 -> 1,2,3
// check if mid element matches search element
// if so return index, else
// if smaller than search -> create a list from 0 up to index and repeat
// if larger than search -> create list from index to list length and repeat
func Search(element int, list []int)bool {
	slices.Sort(list)
	return search(element, list)
}

func search(element int , list []int ) bool  {
	if len(list) < 1{
		return false
	}
	middleIndex := len(list) / 2
	middleElement := list[middleIndex]
	fmt.Println(list, middleElement)
	if element == middleElement{
		return true
	} else if element > middleElement {
		return Search(element, list[middleIndex+1:])
	} else {
		return Search(element, list[0:middleIndex-1])
	}
}

