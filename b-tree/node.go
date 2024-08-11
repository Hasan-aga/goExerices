package btree

import (
	"bytes"
)

type item struct {
	key []byte
	val []byte
}

type node struct {
	items       [maxItems]*item
	children    [maxChildren]*node
	numItems    int
	numChildren int
}

func (n *node) isLeaf() bool {
	return n.numChildren == 0
}

func (n *node) search(key []byte) (int, bool) {
	low, high := 0, n.numItems
	var mid int
	for low < high {
		mid = (high + low) / 2
		result := bytes.Compare(key, n.items[mid].key)
		switch {
		case result > 0:
			low = mid + 1
		case result < 0:
			high = mid
		case result == 0:
			return mid, true
		}
	}
	return low, false
}

func (n *node) insertItemAt(pos int, i *item) {
	if pos < n.numItems {
		// move items to make room when inserting anywhere other than the end
		copy(n.items[pos+1:n.numItems+1], n.items[pos:n.numItems])
	}
	n.items[pos] = i
	n.numItems++
}

func (existingNode *node) insertChildAt(pos int, newChild *node) {
	if pos < existingNode.numChildren {
		copy(existingNode.children[pos+1:existingNode.numChildren+1], existingNode.children[pos:])
	}

	existingNode.children[pos] = newChild
	existingNode.numChildren++
}

func (n *node) split() (*item, *node) {
	// move mid node to parent
	mid := minItems
	midItem := n.items[mid]
	// take second half of items and create a new node from it
	newNode := &node{}
	copy(newNode.items[:], n.items[mid+1:])
	// copy the children pointers to the new node
	if !n.isLeaf() {
		copy(newNode.children[:], n.children[mid+1:])
		newNode.numChildren = minItems + 1
	}
	// cleanup the existing node from child pointers and items of copied node
	for i, l := mid, n.numItems; i < l; i++ {
		if !n.isLeaf() {
			n.children[i+1] = nil
			n.numChildren--
		}
		n.items[i] = nil
		n.numItems--
	}

	// return the mid item and the new node so we can link them to parent
	return midItem, newNode
}

func (n *node) insert(item *item) bool {
	// search for item, if its key exists then simply update the item
	pos, found := n.search(item.key)
	if found {
		n.items[pos] = item
		return false
	}
	// if item does not exist and node is not full then insert item at position
	if n.isLeaf() {
		n.insertItemAt(pos, item)
		return true
	}
	// if item does not exist and node is full perform a split
	if n.children[pos].numItems >= maxItems {
		midItem, newNode := n.children[pos].split()
		// note that we are inserting the mid item here not the new item
		n.insertItemAt(pos, midItem)
		n.insertChildAt(pos+1, newNode)
		// We may need to change our direction after promoting the middle item to the parent, depending on its key.
		switch cmp := bytes.Compare(item.key, n.items[pos].key); {
		case cmp < 0:
			// The key we are looking for is still smaller than the key of the middle item that we took from the child,
			// so we can continue following the same direction.
		case cmp > 0:
			// The middle item that we took from the child has a key that is smaller than the one we are looking for,
			// so we need to change our direction.
			pos++
		default:
			// The middle item that we took from the child is the item we are searching for, so just update its value.
			n.items[pos] = item
			return true
		}
	}
	return n.children[pos].insert(item)
}
