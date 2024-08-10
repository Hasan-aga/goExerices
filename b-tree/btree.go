import ("errors")


const (
	degree      = 5
	maxChildren = 2 * degree       // 10
	maxItems    = maxChildren - 1  // 9
	minItems    = degree - 1       // 4
)

type Btree struct{
  root *node
}

func NewBtree() *Btree {
  return &Btree{}
}

func (t *Btree) find(key []byte) ([]byte, error){
  for next := t.root; next != nil; {
    pos, found := next.search(key)

    if found {
      return next.items[pos].val, nil
    }

    next = next.children[pos]
  }
  return nil, errors.New("key not found") 
}
