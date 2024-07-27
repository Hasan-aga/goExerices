// we get an element to search for and a list to search in
// we create a map of every element in the list map of element and index in list
// then we lookup the map for that element
package hashsearch

func HashSearch(searchElement int, list []int) int {
	if len(list) < 1 {
		return -1
	}
	elementMap := make(map[int]int)
	for index, element := range list {
		elementMap[element] = index
	}

	if index, found := elementMap[searchElement]; found {
		return index
	}
	return -1
}
