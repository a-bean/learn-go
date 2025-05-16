package segmenttree

import "learn-go/structure/constraints"

func minInt[T constraints.Integer](values ...T) T {
	min := values[0]
	for _, value := range values {
		if value < min {
			min = value
		}
	}
	return min
}

func maxInt[T constraints.Integer](values ...T) T {
	max := values[0]
	for _, value := range values {
		if value > max {
			max = value
		}
	}
	return max
}

const emptyLazyNode = 0

type SegmentTree struct {
	Array       []int
	SegmentTree []int
	LazyTree    []int
}

func (s *SegmentTree) Propagate(node int, leftNode int, rightNode int) {
	if s.LazyTree[node] != emptyLazyNode {
		s.SegmentTree[node] += (rightNode - leftNode + 1) * s.LazyTree[node]

		if leftNode == rightNode {
			s.Array[leftNode] += s.LazyTree[node]
		} else {

			s.LazyTree[2*node] += s.LazyTree[node]
			s.LazyTree[2*node+1] += s.LazyTree[node]
		}

		s.LazyTree[node] = emptyLazyNode
	}
}

func (s *SegmentTree) Query(node int, leftNode int, rightNode int, firstIndex int, lastIndex int) int {
	if (firstIndex > lastIndex) || (leftNode > rightNode) {
		return 0
	}

	s.Propagate(node, leftNode, rightNode)

	if (leftNode >= firstIndex) && (rightNode <= lastIndex) {
		return s.SegmentTree[node]
	}

	mid := (leftNode + rightNode) / 2

	leftNodeSum := s.Query(2*node, leftNode, mid, firstIndex, minInt(mid, lastIndex))
	rightNodeSum := s.Query(2*node+1, mid+1, rightNode, maxInt(firstIndex, mid+1), lastIndex)

	return leftNodeSum + rightNodeSum
}

func (s *SegmentTree) Update(node int, leftNode int, rightNode int, firstIndex int, lastIndex int, value int) {
	s.Propagate(node, leftNode, rightNode)

	if (firstIndex > lastIndex) || (leftNode > rightNode) {
		return
	}

	if (leftNode >= firstIndex) && (rightNode <= lastIndex) {

		s.LazyTree[node] += value
		s.Propagate(node, leftNode, rightNode)
	} else {
		mid := (leftNode + rightNode) / 2

		s.Update(2*node, leftNode, mid, firstIndex, minInt(mid, lastIndex), value)
		s.Update(2*node+1, mid+1, rightNode, maxInt(firstIndex, mid+1), lastIndex, value)

		s.SegmentTree[node] = s.SegmentTree[2*node] + s.SegmentTree[2*node+1]
	}
}

func (s *SegmentTree) Build(node int, left int, right int) {
	if left == right {
		s.SegmentTree[node] = s.Array[left]
	} else {
		mid := (left + right) / 2

		s.Build(2*node, left, mid)
		s.Build(2*node+1, mid+1, right)

		s.SegmentTree[node] = s.SegmentTree[2*node] + s.SegmentTree[2*node+1]
	}
}

func NewSegmentTree(Array []int) *SegmentTree {
	if len(Array) == 0 {
		return nil
	}

	segTree := SegmentTree{
		Array:       Array,
		SegmentTree: make([]int, 4*len(Array)),
		LazyTree:    make([]int, 4*len(Array)),
	}

	for i := range segTree.LazyTree {
		segTree.LazyTree[i] = emptyLazyNode
	}

	segTree.Build(1, 0, len(Array)-1)

	return &segTree
}
