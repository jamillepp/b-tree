package main

import (
	btree "github.com/jamillepp/b-tree/btree"
)

func main() {
	t := btree.Initbtree(2)
	t.Insert(10)
	t.Insert(9)
	t.Insert(11)
	t.Insert(15)
	t.Insert(12)
	t.Insert(8)
	t.Insert(16)
	t.Insert(13)
	t.Insert(7)
	t.Insert(6)
	t.Insert(17)
	t.Insert(18)
	t.Insert(19)
	t.Insert(20)
	t.Insert(21)
	t.Insert(22)
	t.Insert(23)
	t.Print()
}
