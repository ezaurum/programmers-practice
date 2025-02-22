const input = `00110
00011
11111
00000`

const lines = input.split('\n')

const map = lines.map(s=>s.split(""))
const visited = {} 
const maxX = map[0].length -1
const maxY = map.length -1

const chunck = []
const stack = []

const check = (x,y) => {
	if (x < 0 || x > maxX || y < 0 || y > maxY) {
		return false
	}
	const v = map[y][x]
	if (v !=='0') {
		return false
	}
	const key = `${x}:${y}`
	if (visited[key]) {
		return false
	}

	visited[key] = true
	// 인접 노드 찾기
	check(x-1, y)
	check(x+1, y)
	check(x, y-1)
	check(x, y+1)

	return true
}

// dfs
for ( let i =0; i <= maxY; i++) {
	for (let j=0; j <= maxX; j++) {
		stack.push([j, i])
	}
}

let count = 0
for (;stack.length > 0;) {
	const [x,y] = stack.pop()
	const v = check(x,y)
	if (v) {
		count++
	}
}

console.log(count)
