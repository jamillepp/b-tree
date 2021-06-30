package btree

import (
	"fmt"
	"sort"
)

type btree struct {
	ord   int
	typ   int // -1: root, 0: inter, 1: leaf
	m     int
	page  []int
	child []*btree
	anc   *btree
}

func Initbtree(o int) btree {
	return btree{
		ord: o,
		typ: 1,
		m:   0,
	}
}

func (b *btree) Insert(k int) {
	if b.typ == 1 {
		if b.m < b.ord*2 {
			b.page = append(b.page, k)
			b.m += 1
			sort.Ints(b.page)
		} else if b.m == b.ord*2 {
			b.page = append(b.page, k)
			b.split()
		}
	} else {
		if k < b.page[0] {
			b.child[0].Insert(k)
		}
		if k > b.page[len(b.page)-1] {
			b.child[len(b.child)-1].Insert(k)
		}
		// for i := 0; i < len(b.page); i++ {
		// 	if i == len(b.page)-1 {
		// 		if
		// 	}
		// }
	}
}

func (b *btree) split() {
	if b.anc == nil {
		cp := make([]int, len(b.page))
		copy(cp, b.page)
		pivotid := int(len(cp) / 2)
		pivot := cp[pivotid]
		f := cp[:pivotid]
		s := cp[pivotid+1:]
		b.page = []int{pivot}
		b.typ = -1
		b.m = 1
		c0 := btree{
			ord:  b.ord,
			typ:  1,
			m:    2,
			page: f,
			anc:  b,
		}
		c1 := btree{
			ord:  b.ord,
			typ:  1,
			m:    2,
			page: s,
			anc:  b,
		}
		b.child = append(b.child, &c0)
		b.child = append(b.child, &c1)
	} else {
		pivotid := int(len(b.page) / 2)
		pivot := b.page[pivotid]
		b.anc.page = append(b.anc.page, pivot)
		b.page = append(b.page[:pivotid], b.page[pivotid+1:]...)
	}
}

func (b *btree) Print() {
	fmt.Println(b.typ, b.page)
	for _, p := range b.child {
		if p != nil {
			p.Print()
		}
	}
}
