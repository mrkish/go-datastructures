package hashtable

import "go-datastructures/model"

type implMap map[string]model.Object

// Go's maps are HashTables, but to embed the map[string]model.Object it needed a custom type.
type HashTable struct {
	implMap
}
