package main

import (
	// "fmt"
)


func numIslands(grid [][]byte) int {
	var max_island int = 0

	for i := 0; i<len(grid); i++ {
		for j:=0; j< len(grid[0]); j++ {
			if isIsland(i, j, grid) {
				max_island++
			}
		}
	}
	return max_island
}

func isIsland(i, j int, grid [][]byte) bool{
	if i < 0 || j < 0 || i > len(grid)-1 || j > len(grid[0])-1 || grid[i][j] == '0' {
		return false
	}
	grid[i][j] = '0'
	isIsland(i+1, j,grid)
	isIsland(i, j+1,grid)
	isIsland(i-1, j,grid)
	isIsland(i, j-1,grid)
	return true

}