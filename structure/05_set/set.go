package set

func New(items ...any) Set {
	st := set{
		elements: make(map[any]bool),
	}
	for _, item := range items {
		st.Add(item)
	}
	return &st
}

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

// 注释：判断是否是子集
func (st *set) IsSubsetOf(superSet Set) bool {
	if st.Len() > superSet.Len() {
		return false
	}

	for _, item := range st.GetItems() {
		if !superSet.In(item) {
			return false
		}
	}
	return true
}

// 注释：判断是否是真子集
func (st *set) IsProperSubsetOf(superSet Set) bool {
	if st.Len() == superSet.Len() {
		return false
	}
	return st.IsSubsetOf(superSet)
}

// 注释：判断是否是超集
func (st *set) IsSupersetOf(subSet Set) bool {
	return subSet.IsSubsetOf(st)
}

// 注释：判断是否是真超集
func (st *set) IsProperSupersetOf(subSet Set) bool {
	if st.Len() == subSet.Len() {
		return false
	}
	return st.IsSupersetOf(subSet)
}

// 注释：求并集
func (st *set) Union(st2 Set) Set {
	unionSet := New()
	for _, item := range st.GetItems() {
		unionSet.Add(item)
	}
	for _, item := range st2.GetItems() {
		unionSet.Add(item)
	}
	return unionSet
}

// 注释：求交集
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

// 注释：求差集
func (st *set) Difference(st2 Set) Set {
	differenceSet := New()
	for _, item := range st.GetItems() {
		if !st2.In(item) {
			differenceSet.Add(item)
		}
	}
	return differenceSet
}

// 注释：求对称差集
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
