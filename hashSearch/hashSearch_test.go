package hashsearch

import (
	"fmt"
	"testing"
)

func TestHashSearch(t *testing.T) {
	searchElements := []int{3, 2, 1, 5, 6, 7, 4}
	result := HashSearch(2, searchElements)
	fmt.Println("index is", result)
	if result == -1 {
		t.Errorf("Search(2, %v) returned -1, expected a valid index", searchElements)
	}
}
func TestHashSearchNotFound(t *testing.T) {
	searchElements := []int{3, 2, 1, 5, 6, 7, 4}
	result := HashSearch(200, searchElements)

	if result != -1 {
		t.Errorf("Search(2, %v) returned -1, expected a valid index", searchElements)
	}
}
