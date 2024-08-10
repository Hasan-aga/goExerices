import {
  "bytes"
}

type item struct {
  key []byte
  val []byte
}

type node struct {
  items [maxItems]*item
  children [maxChildren]*node
  numItems int
  numChildren int
}

func (n *node) isLeaf() bool {
  return n.numChildren == 0
}

func (n *node) search(key []byte) (int, bool){
  low, high := 0, n.numItems
  var mid int
  for low < high {
    mid = (high + low) / 2
    result := bytes.Compare(key, n.item[mid].key)
    switch  {
    case result > 0:
      low = mid + 1
    case result < 0:
      high = mid - 1
    case result == 0:
      return mid, true
    }
  }
  return low, false
}
