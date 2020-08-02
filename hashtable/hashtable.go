package hashtable

type implMap map[string]string

// Go's maps are HashTables, but to imbed the map[string]string it needed a custom type
type HashTable struct {
	implMap
}
