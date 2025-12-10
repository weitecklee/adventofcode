package utils

type UnionFind struct {
	Parents []int
	Sizes   []int
}

func NewUnionFind(n int) *UnionFind {
	uf := UnionFind{
		make([]int, n),
		make([]int, n),
	}
	for i := range uf.Parents {
		uf.Parents[i] = i
		uf.Sizes[i] = 1
	}
	return &uf
}

func (uf *UnionFind) FindParent(n int) int {
	if uf.Parents[n] != n {
		uf.Parents[n] = uf.FindParent(uf.Parents[n])
	}
	return uf.Parents[n]
}

func (uf *UnionFind) Union(a, b int) {
	rootA := uf.FindParent(a)
	rootB := uf.FindParent(b)

	if rootA == rootB {
		return
	}

	if uf.Sizes[rootA] < uf.Sizes[rootB] {
		rootA, rootB = rootB, rootA
	}

	uf.Parents[rootB] = rootA
	uf.Sizes[rootA] += uf.Sizes[rootB]
}

func (uf UnionFind) Count() int {
	res := 0
	for i := range uf.Parents {
		if i == uf.FindParent(i) {
			res += 1
		}
	}
	return res
}
