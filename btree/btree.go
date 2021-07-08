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

	fmt.Printf("\nInserting %v", k)

	if b.typ == 1 { // Se for folha

		if b.m < b.ord*2 {
			fmt.Printf("\n%v inserted\n", k)
			b.page = append(b.page, &key{value: k})
			b.m = len(b.page)
			sortpage(b.page)
			return
		} else if b.m == b.ord*2 {
			fmt.Printf("\n%v inserted\n", k)
			b.page = append(b.page, &key{value: k})
			sortpage(b.page)
			b.split()
			return
		}
	} else {
		fmt.Println("\nNão é folha")
		if k < b.page[0].value { // Entra na p0
			fmt.Println("Entra na p0")
			b.page[0].c0.Insert(k)
			return
		} else if k > b.page[len(b.page)-1].value { // Entra na pm
			fmt.Println("Entra na pn")
			b.page[len(b.page)-1].c1.Insert(k)
			return
		} else {
			fmt.Println("Entra na pi")
			for i := 0; i < len(b.page)-1; i++ {
				if k > b.page[i].value && k < b.page[i+1].value {
					b.page[i+1].c0.Insert(k)
					break
				}
			}
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
	fmt.Println("Split")
	if b.anc == nil {
		cp := make([]*key, len(b.page))
		copy(cp, b.page)
		indpivot := int(len(cp) / 2)
		pivot := cp[indpivot]
		pagec0 := []*key{}
		pagec1 := []*key{}
		pagec0 = append(pagec0, cp[:indpivot]...)
		pagec1 = append(pagec1, cp[indpivot+1:]...)
		ch0 := btree{
			ord:  b.ord,
			typ:  1,
			m:    len(pagec0),
			page: pagec0,
			anc:  b,
		}
		ch1 := btree{
			ord:  b.ord,
			typ:  1,
			m:    len(pagec1),
			page: pagec1,
			anc:  b,
		}
		b.typ = -1
		b.page = []*key{
			{
				value: pivot.value,
				c0:    &ch0,
				c1:    &ch1,
			},
		}
		b.m = 1
	} else {
		indpivot := int(len(b.page) / 2)
		pivot := b.page[indpivot]
		pagec0 := []*key{}
		pagec1 := []*key{}
		pagec0 = append(pagec0, b.page[:indpivot]...)
		pagec1 = append(pagec1, b.page[indpivot+1:]...)
		c0 := btree{
			ord:  b.ord,
			typ:  1,
			m:    b.ord,
			page: pagec0,
			anc:  b.anc,
		}
		c1 := btree{
			ord:  b.ord,
			typ:  1,
			m:    b.ord,
			page: pagec1,
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

func (b *btree) Delete(k int) {
	for i, v := range b.page {
		if v.value == k {
			if b.typ != 1 {
				v.value = v.c1.findSucessor()
				return
			} else if b.m > b.ord {
				b.page = append(b.page[:i], b.page[i+1:]...)
				b.m -= 1
				return
			} else {
				b.page = append(b.page[:i], b.page[i+1:]...)
				b.m -= 1
				for i, w := range b.anc.page {
					if w.value < k {

						if i == len(b.anc.page)-1 {
							if w.c0.m <= b.ord {
								b.anc.concat(w)
							} else {
								b.anc.redistribute(w)
							}
							return
						} else if b.anc.page[i+1].value > k {
							if b.anc.page[i+1].c1.m <= b.ord {
								b.anc.concat(b.anc.page[i+1])
								return
							} else {
								b.anc.redistribute(b.anc.page[i+1])
								return
							}
						}
					}
				}
			}
		}
	}

	if k < b.page[0].value { // Entra na p0
		b.page[0].c0.Delete(k)
		return
	} else if k > b.page[len(b.page)-1].value { // Entra na pm
		b.page[len(b.page)-1].c1.Delete(k)
		return
	} else {
		for i := 0; i < len(b.page)-1; i++ {
			if k > b.page[i].value && k < b.page[i+1].value {
				b.page[i+1].c0.Delete(k)
			}
		}
	}
}

func (b *btree) findSucessor() (k int) {
	if b.page[0].c0 != nil {
		return b.page[0].c0.findSucessor()
	}
	sucessor := b.page[0].value
	b.page = b.page[1:]
	b.m -= 1
	if b.m < b.ord {
		fmt.Println("Menor")
	}
	return sucessor
}

func (b *btree) concat(w *key) {
	fmt.Println("Concat")
	newChild := btree{
		ord: b.ord,
		typ: 1,
		anc: b,
	}
	for i, k := range b.page {
		if k.value == w.value {
			newChild.page = append(newChild.page, k)
			newChild.page = append(newChild.page, k.c0.page...)
			newChild.page = append(newChild.page, k.c1.page...)
			sortpage(newChild.page)
			k.c0 = nil
			k.c1 = nil
			b.page = append(b.page[:i], b.page[i+1:]...)
			if len(b.page) == 0 {
				b.page = newChild.page
			} else {
				if i != 0 {
					b.page[i-1].c1 = &newChild
				}
				if i != len(b.page) {
					b.page[i+1].c0 = &newChild
				}
			}
		}
	}
}

func (b *btree) redistribute(w *key) {
	fmt.Println("Redistribute")
	temp := []*key{}

	temp = append(temp, &key{value: w.value})
	temp = append(temp, w.c0.page...)
	temp = append(temp, w.c1.page...)
	sortpage(temp)

	pivotind := int(len(temp) / 2)
	pivot := temp[pivotind]
	pagec0 := []*key{}
	pagec1 := []*key{}
	pagec0 = append(pagec0, temp[:pivotind]...)
	pagec1 = append(pagec1, temp[pivotind+1:]...)
	w.value = pivot.value
	w.c0 = &btree{
		ord:  b.ord,
		typ:  1,
		m:    len(pagec0),
		page: pagec0,
		anc:  b,
	}
	w.c1 = &btree{
		ord:  b.ord,
		typ:  1,
		m:    len(pagec1),
		page: pagec1,
		anc:  b,
	}
	for i, v := range b.page {
		if v.value > w.value || i == len(b.page)-1 {
			b.page[i-1].c1 = w.c0
		}
	}
}

func (b *btree) Print() {
	fmt.Println("")
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
