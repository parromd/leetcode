package main

import "fmt"

func pacificAtlantic(heights [][]int) [][]int {
    visited := map[int]int {}
   
	var elper func(int, int)
    elper = func(x int, y int) {
        coord := x*len(heights[0]) + y
        if x < 0 || y < 0 {
            visited[coord] += 1
            return
        }
        
        if x >= len(height) || y >= len(height[0]) {
            visited[coord] += 2
            return
        }
        
        if heights()
    }
    
    
}


func main() {

}
