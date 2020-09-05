package bst

type NodeFloat64 struct {
    Left  *NodeFloat64
    Val   float64
    Right *NodeFloat64
}

// Insert number into the Tree
func (t *NodeFloat64) Insert(n float64) error {
    // If the NodeFloat64 is Empty
    if t.Val == 0 {
        t.Val = n
        return nil
    }

    if t.Val == n {
        panic("This Value already exist")
    }

    // Left
    if t.Val > n {
        if t.Left == nil {
            t.Left = &NodeFloat64{Val: n}
            return nil
        }

        return t.Left.Insert(n)
    }
    // Right
    if t.Val < n {
        if t.Right == nil {
            t.Right = &NodeFloat64{Val: n}
            return nil
        }
        return t.Right.Insert(n)
    }
    return nil
}

// Returns an array with the sorted Tree
func (this *NodeFloat64) Sort() []float64 {
    var res []float64
    this.srt(&res)
    return res
}

func (this *NodeFloat64) srt(res *[]float64) {
    if this == nil {
        return
    }
    this.Left.srt(res)
    *res = append(*res, this.Val)
    this.Right.srt(res)
}
