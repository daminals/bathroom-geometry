package main

import (
	"fmt"
	"math"
)

// Point represents a point in 2D space
type Point struct {
	X, Y float64
}

// VoronoiCell represents a Voronoi cell
type VoronoiCell struct {
	Center Point
}


// JumpFlood performs the Jump Flood algorithm to compute Voronoi diagram
func JumpFlood(points []Point, size int) [][]VoronoiCell {
	grid := make([][]VoronoiCell, size)
	for i := range grid {
		grid[i] = make([]VoronoiCell, size)
	}

	for _, point := range points {
		x, y := int(point.X), int(point.Y)
		grid[x][y] = VoronoiCell{Center: point}
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
	points := []Point{
		{2, 2},
		{8, 8},
		{5, 5},
	}

	size := 16
	voronoiGrid := JumpFlood(points, size)

	// Display the Voronoi cells
	for _, row := range voronoiGrid {
		for _, cell := range row {
			if cell.Center != (Point{}) {
				fmt.Printf("(%0.2f, %0.2f) ", cell.Center.X, cell.Center.Y)
			} else {
				fmt.Print("( - , - ) ")
			}
		}
		fmt.Println()
	}
}