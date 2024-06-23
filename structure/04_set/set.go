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

func New(items ...any) Set {
	st := set{
		elements: make(map[any]bool),
	}
	for _, item := range items {
		st.Add(item)
	}
	return &st
}

func (st *set) Add(value any) {
	st.elements[value] = true
}

func (st *set) Delete(value any) {
	delete(st.elements, value)
}

func (st *set) GetItems() []any {
	keys := make([]any, 0, len(st.elements))
	for k := range st.elements {
		keys = append(keys, k)
	}
	return keys
}

func (st *set) Len() int {
	return len(st.elements)
}

func (st *set) In(value any) bool {
	if _, in := st.elements[value]; in {
		return true
	}
	return false
}

func (st *set) IsSubsetOf(superSet Set) bool {
	if st.Len() > superSet.Len() {
		return false
	}

	for _, value := range superSet.GetItems() {
		if !st.In(value) {
			return false
		}
	}

	return true
}

func (st *set) IsProperSubsetOf(superSet Set) bool {
	if st.Len() == superSet.Len() {
		return false
	}
	return st.IsSubsetOf(superSet)
}

func (st *set) IsSupersetOf(subSet Set) bool {
	return subSet.IsSubsetOf(st)
}

func (st *set) IsProperSupersetOf(subSet Set) bool {
	if st.Len() == subSet.Len() {
		return false
	}
	return st.IsSupersetOf(subSet)
}

// 并集
func (st *set) Union(st2 Set) Set {
	unionSet := New()
	for _, value := range st.GetItems() {
		unionSet.Add(value)
	}
	for _, value := range st2.GetItems() {
		unionSet.Add(value)
	}
	return unionSet
}

// 交集
func (st *set) Intersection(st2 Set) Set {
	intersectionSet := New()
	var minSet, maxSet Set
	if st.Len() > st2.Len() {
		minSet = st2
		maxSet = st
	} else {
		minSet = st
		maxSet = st2
	}
	for _, item := range minSet.GetItems() {
		if maxSet.In(item) {
			intersectionSet.Add(item)
		}
	}
	return intersectionSet
}

// 差集
func (st *set) Difference(st2 Set) Set {
	differenceSet := New()
	for _, value := range st.GetItems() {
		if !st2.In(value) {
			differenceSet.Add(value)
		}
	}
	return differenceSet
}

func (st *set) SymmetricDifference(st2 Set) Set {
	symmetricDifferenceSet := New()
	dropSet := New()
	for _, item := range st.GetItems() {
		if st2.In(item) {
			dropSet.Add(item)
		} else {
			symmetricDifferenceSet.Add(item)
		}
	}
	for _, item := range st2.GetItems() {
		if !dropSet.In(item) {
			symmetricDifferenceSet.Add(item)
		}
	}
	return symmetricDifferenceSet
}
