package model

// Just a dummy object that'd ideally be a generic type, whenever Go gets those implemented.
// Using a string as the value is a simple way to do comparisons rather than having interface{}.
// This may get refactored since it's kinda cumbersome, even if it's "more representative" of use cases.
type Object struct {
	Value string
}
