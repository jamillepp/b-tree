package btree

import (
	"fmt"
	"math"
)

type btree struct {
	ord  int
	typ  int // -1: root, 0: inter, 1: leaf
	m    int
	page []*key
	anc  *btree
}

type key struct {
	value int
	c0    *btree
	c1    *btree
}

func Initbtree(o int) btree {
	return btree{
		ord: o,
		typ: 1,
		m:   0,
	}
}

func (b *btree) Insert(k int) {

	newkey := &key{
		value: k,
	}

	if b.typ == 1 { // Se for folha

		if b.m < b.ord*2 {
			b.page = append(b.page, newkey)
			b.m += 1
			sortpage(b.page)
			return
		} else if b.m == b.ord*2 {
			b.page = append(b.page, newkey)
			sortpage(b.page)
			b.split()
			return
		}
	} else {
		if newkey.value < b.page[0].value { // Entra na p0
			b.page[0].c0.Insert(k)
		} else if newkey.value > b.page[len(b.page)-1].value { // Entra na pm
			b.page[len(b.page)-1].c1.Insert(k)
		} else {
			fmt.Println("pi")
		}
	}
}

func sortpage(p []*key) {
	for k := 0; k < int(math.Pow(float64((len(p))), 2)); k++ {
		for i := 0; i < len(p); i++ {
			if i != len(p)-1 && p[i].value > p[i+1].value {
				t := p[i]
				p[i] = p[i+1]
				p[i+1] = t
			}
		}
	}
}

func (b *btree) split() {
	if b.anc == nil {
		cp := make([]*key, len(b.page))
		copy(cp, b.page)
		indpivot := int(len(cp) / 2)
		pivot := cp[indpivot]
		ch0 := &btree{
			ord:  b.ord,
			typ:  1,
			m:    b.ord,
			page: cp[:indpivot],
			anc:  b,
		}
		ch1 := &btree{
			ord:  b.ord,
			typ:  1,
			m:    b.ord,
			page: cp[indpivot+1:],
			anc:  b,
		}
		b.typ = -1
		b.page = []*key{
			{
				value: pivot.value,
				c0:    ch0,
				c1:    ch1,
			},
		}
		b.m = 1
	} else {
		indpivot := int(len(b.page) / 2)
		pivot := b.page[indpivot]
		c0 := btree{
			ord:  b.ord,
			typ:  1,
			m:    b.ord,
			page: b.page[:indpivot],
			anc:  b.anc,
		}
		c1 := btree{
			ord:  b.ord,
			typ:  1,
			m:    b.ord,
			page: b.page[indpivot+1:],
			anc:  b.anc,
		}
		newkey := key{
			value: pivot.value,
			c0:    &c0,
			c1:    &c1,
		}
		b.anc.page = append(b.anc.page, &newkey)
		b.anc.m += 1
		sortpage(b.anc.page)
		for i, k := range b.anc.page {
			if k.value == pivot.value {
				if i-1 > 0 {
					b.anc.page[i-1].c1 = k.c0
				}
				if i+1 < len(b.anc.page) {
					b.anc.page[i+1].c0 = k.c1
				}
			}
		}
		if b.anc.m > b.ord*2 {
			b.anc.split()
		}
	}
}

func (b *btree) Print() {
	b.print([]*btree{})
}

func (b *btree) print(queue []*btree) {
	fmt.Printf("%v [", b.typ)
	for i, k := range b.page {
		fmt.Printf("%v", k.value)
		if k.c0 != nil {
			queue = append(queue, k.c0)
		}
		if k.c1 != nil {
			queue = append(queue, k.c1)
		}
		if i == len(b.page)-1 {
			fmt.Printf("]\n")
		} else {
			fmt.Printf(" ")
		}
	}

	for _, k := range queue {
		newQueue := []*btree{}
		k.print(newQueue)
	}
}
