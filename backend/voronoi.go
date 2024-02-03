package main

import (
	"fmt"
	"math"
)

// VoronoiCell represents a Voronoi cell
type VoronoiCell struct {
	Center Point
}

// Point represents a point in 2D space with integer coordinates
type Point struct {
	X, Y int
}

// JumpFlood performs the Jump Flood algorithm to compute Voronoi diagram
func JumpFlood(matrix [][]int, size int) [][]VoronoiCell {
	grid := make([][]VoronoiCell, size)
	for i := range grid {
		grid[i] = make([]VoronoiCell, size)
	}

	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			if matrix[x][y] > 0 {
				grid[x][y] = VoronoiCell{Center: Point{X: x, Y: y}}
			}
		}
	}

	step := size / 2
	for step > 0 {
		for x := 0; x < size; x += step * 2 {
			for y := 0; y < size; y += step * 2 {
				c1 := grid[(x+step)%size][(y+step)%size].Center
				c2 := grid[(x+step)%size][(y-step+size)%size].Center
				c3 := grid[(x-step+size)%size][(y+step)%size].Center
				c4 := grid[(x-step+size)%size][(y-step+size)%size].Center

				average := Point{
					X: (c1.X + c2.X + c3.X + c4.X) / 4,
					Y: (c1.Y + c2.Y + c3.Y + c4.Y) / 4,
				}

				grid[x][y] = VoronoiCell{Center: average}
			}
		}

		for x := 0; x < size; x += step {
			for y := 0; y < size; y += step {
				for i := -step; i <= step; i += step {
					for j := -step; j <= step; j += step {
						if i == 0 && j == 0 {
							continue
						}

						newX, newY := (x+i+size)%size, (y+j+size)%size
						grid[newX][newY] = grid[x][y]
					}
				}
			}
		}

		step /= 2
	}

	return grid
}

func main() {
	// Example usage
	matrix := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	size := len(matrix)
	voronoiGrid := JumpFlood(matrix, size)

	// Display the Voronoi cells
	for _, row := range voronoiGrid {
		for _, cell := range row {
			if cell.Center != (Point{}) {
				fmt.Printf("(%d, %d) ", cell.Center.X, cell.Center.Y)
			} else {
				fmt.Print("( - , - ) ")
			}
		}
		fmt.Println()
	}
}
