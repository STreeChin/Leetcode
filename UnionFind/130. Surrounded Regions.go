/*
130. Surrounded Regions
Given a 2D board containing 'X' and 'O' (the letter O), capture all regions surrounded by 'X'.

A region is captured by flipping all 'O's into 'X's in that surrounded region.

Example:

X X X X
X O O X
X X O X
X O X X
After running your function, the board should be:

X X X X
X X X X
X X X X
X O X X
Explanation:

Surrounded regions shouldn’t be on the border, which means that any 'O' on the border of the board are not flipped to 'X'. Any 'O' that is not on the border and it is not connected to an 'O' on the border will be flipped to 'X'. Two cells are connected if they are adjacent cells connected horizontally or vertically.
*/
/*
1. Solution1, DFS
first set the O as a another flag
*/
/*
2. Solution2, UnionFind
set a dummyNode as the root
*/
package main

//DFS solve
func solveDFS(board [][]byte) {
	if board == nil || len(board) == 0 {
		return
	}
	m, n := len(board), len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			isEdge := i == 0 || j == 0 || i == m-1 || j == n-1
			if isEdge && board[i][j] == 'O' {
				dfs(board, i, j)
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
			if board[i][j] == '#' {
				board[i][j] = 'O'
			}
		}
	}
}

func dfs(board [][]byte, i, j int) {
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) || board[i][j] == 'X' || board[i][j] == '#' {
		return
	}
	board[i][j] = '#'
	dfs(board, i-1, j)
	dfs(board, i+1, j)
	dfs(board, i, j-1)
	dfs(board, i, j+1)
}

//UnionFind
type UnionFind struct {
	parents []int
}

func NewUnionFind(totalNodes int) *UnionFind {
	uf := &UnionFind{make([]int, totalNodes)}
	for i := 0; i < totalNodes; i++ {
		uf.parents[i] = i
	}
	return uf
}

func (uf *UnionFind) Union(node1, node2 int) {
	root1 := uf.Find(node1)
	root2 := uf.Find(node2)
	if root1 != root2 {
		uf.parents[root2] = root1
	}
}

func (uf *UnionFind) Find(node int) int {
	for uf.parents[node] != node {
		uf.parents[node] = uf.parents[uf.parents[node]]
		node = uf.parents[node]
	}
	return node
}

func (uf *UnionFind) IsConnect(node1, node2 int) bool {
	return uf.Find(node1) == uf.Find(node2)
}
func solveUnionFind(board [][]byte) {
	if board == nil || len(board) == 0 {
		return
	}
	rows, cols := len(board), len(board[0])
	uf := NewUnionFind(rows * cols+1)
	dummyNode := rows * cols

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == 'O' {
				if i == 0 || i == rows-1 || j == 0 || j == cols-1 {
					uf.Union(i*cols+j, dummyNode)
				} else {
					// union with the around
					if i > 0 && board[i-1][j] == 'O' {
						uf.Union(i*cols+j, (i-1)*cols+j)
					}
					if i < rows-1 && board[i+1][j] == 'O' {
						uf.Union(i*cols+j, (i+1)*cols+j)
					}
					if j > 0 && board[i][j-1] == 'O' {
						uf.Union(i*cols+j, i*cols+j-1)
					}
					if j < cols-1 && board[i][j+1] == 'O' {
						uf.Union(i*cols+j, i*cols+j+1)
					}
				}
			}
		}
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if uf.IsConnect(i*cols+j, dummyNode) {
				board[i][j] = 'O'
			} else {
				board[i][j] = 'X'
			}
		}
	}
}
