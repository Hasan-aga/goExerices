package binarysearch

import (
	"testing"
)

func TestSearch(t *testing.T) {
	searchElements := []int {3,2,1,5,6,7,4}
	result := Search(2, searchElements)

	if result == false {
        t.Errorf("Search(2, %v) returned -1, expected a valid index", searchElements)
    }

}
func TestSearchNotFound(t *testing.T) {
	searchElements := []int {3,2,1,5,6,7,4}
	result := Search(200, searchElements)

	if result != false {
        t.Errorf("Search(2, %v) returned an index other than -1", searchElements)
    }

}