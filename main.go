package main

import (
	btree "github.com/jamillepp/b-tree/btree"
)

func main() {
	t := btree.Initbtree(2)
	t.Insert(50)
	t.Insert(20)
	t.Insert(80)
	t.Insert(10)
	t.Insert(100)
	t.Insert(30)
	t.Insert(40)
	t.Insert(150)
	t.Insert(43)
	t.Insert(23)
	t.Insert(26)
	t.Delete(100)
	t.Insert(160)
	t.Insert(170)
	t.Delete(43)
	t.Print()
}
