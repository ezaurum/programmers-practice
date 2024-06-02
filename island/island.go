package missle

import (
	"sort"
)

// Edge represents an edge between two nodes with a cost
type Edge struct {
	u, v, cost int
}

// DisjointSet represents a disjoint set data structure
type DisjointSet struct {
	parent, rank []int
}

// NewDisjointSet initializes a disjoint set
func NewDisjointSet(n int) *DisjointSet {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	return &DisjointSet{parent, rank}
}

// Find finds the root of the set containing x
func (ds *DisjointSet) Find(x int) int {
	if ds.parent[x] != x {
		ds.parent[x] = ds.Find(ds.parent[x])
	}
	return ds.parent[x]
}

// Union unites the sets containing x and y
func (ds *DisjointSet) Union(x, y int) {
	rootX := ds.Find(x)
	rootY := ds.Find(y)
	if rootX != rootY {
		if ds.rank[rootX] > ds.rank[rootY] {
			ds.parent[rootY] = rootX
		} else if ds.rank[rootX] < ds.rank[rootY] {
			ds.parent[rootX] = rootY
		} else {
			ds.parent[rootY] = rootX
			ds.rank[rootX]++
		}
	}
}

// MinimumCostToConnectIslands finds the minimum cost to connect all islands
func MinimumCostToConnectIslands(n int, edges []Edge) int {
	// Sort edges by cost
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].cost < edges[j].cost
	})

	ds := NewDisjointSet(n)
	totalCost := 0
	edgesUsed := 0

	for _, edge := range edges {
		if ds.Find(edge.u) != ds.Find(edge.v) {
			ds.Union(edge.u, edge.v)
			totalCost += edge.cost
			edgesUsed++
			if edgesUsed == n-1 {
				break
			}
		}
	}

	return totalCost
}

// solution https://school.programmers.co.kr/learn/courses/30/lessons/42861
func solution(n int, costs [][]int) int {
	var edges []Edge
	for _, cost := range costs {
		edges = append(edges, Edge{cost[0], cost[1], cost[2]})
	}
	return MinimumCostToConnectIslands(n, edges)
}
