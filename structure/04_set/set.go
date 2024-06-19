package set

type Set interface {
	Add(item any)
	Delete(item any)
	Len() int
	GetItems() []any
	In(item any) bool
	IsSubsetOf(set2 Set) bool
	IsProperSubsetOf(set2 Set) bool
	IsSupersetOf(set2 Set) bool
	IsProperSupersetOf(set2 Set) bool
	Union(set2 Set) Set
	Intersection(set2 Set) Set
	Difference(set2 Set) Set
	SymmetricDifference(set2 Set) Set
}

type set struct {
	elements map[any]bool
}
