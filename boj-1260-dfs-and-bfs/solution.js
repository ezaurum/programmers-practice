const input0 = `4 5 1
1 2
1 3
1 4
2 4
3 4`
const input1 = `5 5 3
5 4
5 2
1 2
3 4
3 1`

const input2 = `1000 1 1000
999 1000`

const input = `11 10 1

1 2

1 3

1 4

1 5

1 6

1 7

1 8

1 9

1 10

1 11`

const lines = input.split('\n')

const [n,m,v] = lines[0].split(' ').map(i=>parseInt(i))

const edges = lines.slice(1).map(l=>(l.split(' ').map(i=>parseInt(i))))
const graph = edges.reduce((g, [start, end])=>{
	const a = (g[start] || [])
	a.push(end)
	a.sort((a,b)=>a-b)
	g[start] = a

	const b = (g[end] || [])
	b.push(start)
	a.sort((a,b)=>a-b)
	g[end] = b

	return g
}, new Array(n+1))

// dfs
const dfsVisited = {}
const dfsResult = []
for (const stack = [v];stack.length > 0;) {
	const current = stack.pop()
	if (dfsVisited[current]) continue
	dfsVisited[current] = true
	dfsResult.push(current)
	const nodes = graph[current]
	if (!nodes) continue
	// 작은 숫자부터 방문하라고 했는데, 스택에서 꺼내는 순서는 반대니까
	// stack에 넣을 때는 뒤집어 넣어야 한다.
	for (let i = nodes.length - 1; i >=0 ; i--) {
		const end = nodes[i]
		if (!dfsVisited[end]) {
			stack.push(end)
		}
	}

}


// bfs
const bfsVisited = {}
const bfsResult = []
for (const queue = [v];queue.length > 0;) {
	const [current] = queue
	queue.splice(0,1)
	if (bfsVisited[current]) continue
	bfsVisited[current] = true
	bfsResult.push(current)
	const nodes = graph[current]
	if (!nodes) continue

	for (let i = 0; i < nodes.length; i++) {
		const next = nodes[i]
		if (!bfsVisited[next]) {
			queue.push(next)
		}
	}
}

console.log(dfsResult.join(" "))
console.log(bfsResult.join(" "))

