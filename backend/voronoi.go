package main

import (
	"fmt" 
	"container/heap"
	"sync"
	"sort"
	"math/rand"
	"math"
)

// MaxUint is the maximum value for uint
const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1) 

// point data structure
type Point struct {
	x int
	y int
}  
//Priority Queue to be used in A* 
type PriorityQueue []Point  

func(pq PriorityQueue) Len()int {return len(pq)}
func(pq PriorityQueue) Less(i, j int) bool {return false} 
func(pq PriorityQueue) Swap(i, j int) {pq[i], pq[j] = pq[j], pq[i]} 

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(Point)
	*pq = append(*pq, item)
} 

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}
 
func astar(matrix [][]int, start, end Point) ([]Point, int) { 
	//0, -1, and bathrooms greater than 1  
	//Distance to nearest bathroom 
	openSet := make(PriorityQueue, 0) 
	heap.Init(&openSet)  

	cameFrom := make(map[Point]Point) 
	gScore := make(map[Point]int) 

	heap.Push(&openSet, start) 
	gScore[start] = 0  

	for len(openSet) > 0 { 
		current := heap.Pop(&openSet).(Point) 

		if current == end { 
			path := reconstructPath(cameFrom, start, end) 
			return path, gScore[end]
		}

		neighbors := getNeighbors(current, matrix) 
		for _, neighbor := range neighbors { 
			tentG := gScore[current] + 1 

			if _, ok := gScore[neighbor]; !ok || tentG < gScore[neighbor] { 
				heap.Push(&openSet, neighbor) 
				cameFrom[neighbor] = current
			}
		}
	} 
	return nil, 0  
}

func reconstructPath(cameFrom map[Point]Point, start, current Point) []Point { 
	path := make([]Point, 0)
	for current != start {
		path = append(path, current)
		current = cameFrom[current]
	}
	path = append(path, start)
	reversePath(path)
	return path
}

func getNeighbors(point Point, matrix [][]int) []Point { 
	neighbors := make([]Point, 0) 

	movements := [][2]int {{-1, 0}, {}, {0, -1}, {0, 1}}

	for _, move := range movements {
		x, y := point.x+move[0], point.y+move[1]

		// Check if the neighbor is within the grid boundaries and is passable
		if x >= 0 && x < len(matrix) && y >= 0 && y < len(matrix[0]) && matrix[x][y] == 0 {
			neighbors = append(neighbors, Point{x, y})
		}
	}
	return neighbors
} 
func reversePath(path []Point) {
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
}
// distance based on astar formula
func distance(matrix [][]int, start, end Point) int {
  // utilize astar
	_, cost := astar(matrix, start, end)
	return cost
}

// distance based on pythagorean theorem
func distanceFormula(p1, p2 Point) int {
	return int(math.Sqrt(math.Pow(float64(p2.x-p1.x), 2) + math.Pow(float64(p2.y-p1.y), 2)))
}

func top3Voronoi(voronoiPoints []Point, point Point) []Point {
	// sort voronoi points using distance formula to point
	sort.Slice(voronoiPoints, func(i, j int) bool {
		return distanceFormula(voronoiPoints[i], point) < distanceFormula(voronoiPoints[j], point)
	})
	// return top 3
	return voronoiPoints[:3]
}

// create some sample points that are generally close to the voronoi points
func createInitSamplePoints(voronoiPoints []Point, numSamplePoints int) []Point {
	// initialize sample points
	samplePoints := make([]Point, numSamplePoints)

	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < numSamplePoints; i += 1 {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			// pick a random voronoi point
			voronoiPoint := voronoiPoints[rand.Intn(len(voronoiPoints))]
			// add some random noise to the point
			samplePoints[i] = Point{voronoiPoint.x + rand.Intn(3), voronoiPoint.y + rand.Intn(3)}

			// protect concurrent access to samplePoints
			mu.Lock()
			defer mu.Unlock()
		}(i)
	}

	wg.Wait()
	return samplePoints
}

// create a voronoi table data structure which maps an ID to a point value
func createVoronoiTable(voronoiPoints []Point) map[int]Point {
	voronoiTable := make(map[int]Point)
	for i, point := range voronoiPoints {
		voronoiTable[i] = point
	}
	return voronoiTable
}

// write a function which takes in voronoi points and returns the voronoi id from a table
func getVoronoiId(voronoiTable map[int]Point, point Point) int {
	// iterate through the voronoi table and return the id of the point
	for id, voronoiPoint := range voronoiTable {
		if voronoiPoint == point {
			return id
		}
	}
	return -1
}

// func medianNeighborVornoiID(matrix [][]int, point Point) int {
// 	// get the color of the points neighbors
// 	neighbors := getNeighbors(point, matrix)
// 	return 1
// }


func voronoi(matrix [][]int, voronoiPoints []Point, size int) {
	// create initial sample points
	samplePoints := createInitSamplePoints(voronoiPoints, size / 5)
	filled_points := 0

	// create output matrix
	outputMatrix := make([][]int, size)
	for i := range outputMatrix {
		outputMatrix[i] = make([]int, size)
	}

	voronoiTable := createVoronoiTable(voronoiPoints)

	// calculate distance from each sample point to each voronoi point
	for _, point := range samplePoints {
		minDistance := MaxInt
		voronoiId := -1
		for _, voronoiPoint := range top3Voronoi(voronoiPoints, point) {
			// calculate distance from sample point to voronoi point
			distance := distance(matrix, point, voronoiPoint)
			// update matrix with voronoi point id
			if distance < minDistance {
				minDistance = distance
				voronoiId = getVoronoiId(voronoiTable, voronoiPoint)
			}
		}
		// update output matrix with voronoi id
		outputMatrix[point.x][point.y] = voronoiId
		filled_points += 1
	}

	for filled_points < size*size { 


	}

}



// func jumpFlood(matrix [][]int, voronoiPoints [][]int, size int) {
// 	step = size / 2
// 	outputMatrix := make([][]int, size)
// 	for i := range outputMatrix {
// 		outputMatrix[i] = make([]int, size)
// 	}
// 	// initialize output matrix
// 	for x := range voronoiPoints {
// 		for y := range voronoiPoints[x] {
// 			outputMatrix[x][y] = 
// 		}
// 	}

// 	for x := 0; x < size; x += 1 {
// 		for y := 0; y < size; y += 1 {
// 			// For each neighbor  q at  ( x + i , y + j ) where  i , j ∈ { − k , 0 , k }
// 			for i := -step; i <= step; i += step {
// 				for j := -step; j <= step; j += step {
// 					// if q is -1, discard color
// 					if x+i < 0 || x+i >= size || y+j < 0 || y+j >= size {
// 						continue
// 					}
// 					if matrix[x+i][y+j] == -1 {
// 						continue
// 					}
// 					// if q is closer to p than the current value at p, update p
// 					if (dista)
// 				}
// 			}
// 		}
// 	}
// }


func main() {
	grid := [][]int{
		{0, 0, 0, 0, 0},
		{0, 1, 0, 1, 0},
		{0, 1, 0, 1, 0},
		{0, 0, 0, 0, 0},
	}

	start := Point{0, 0}
	end := Point{4, 4}

	path, cost := astar(grid, start, end)
	if path != nil {
		fmt.Println("Shortest Path:")
		for _, node := range path {
			fmt.Printf("(%d, %d) -> ", node.x, node.y)
		}
		fmt.Printf("\nTotal Cost: %d\n", cost)
	} else {
		fmt.Println("No path found.")
	}
}