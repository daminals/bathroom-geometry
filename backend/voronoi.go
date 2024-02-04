package main

import (
	"fmt" 
	"container/heap"
	// "sync"
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

type VoronoiPoint struct {
	point Point
	id int
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

		//Check if we have reached the end
		if current == end {  
			// fmt.Println("Reached?")
			path := reconstructPath(cameFrom, start, end) 
			return path, gScore[end]
		}

		neighbors := getNeighbors(current, matrix) 
		for _, neighbor := range neighbors {  
			//fmt.Println("Reached?")
			tentG := gScore[current] + 1 

			if _, ok := gScore[neighbor]; !ok || tentG < gScore[neighbor] { 
				//fmt.Println("Reached? 2")  
				gScore[neighbor] = tentG
				heap.Push(&openSet, neighbor) 
				cameFrom[neighbor] = current
			}
		}
		// fmt.Println("(",current.x,",", current.y,")")
		// fmt.Println(len(openSet))
	} 
	return nil, 0  
}
//Recreates Path when end of algorithm is reached
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
//Get's Neighboring Points in order to see what a good move is 
func getNeighbors(point Point, matrix [][]int) []Point { 
	neighbors := make([]Point, 0) 

	movements := [][2]int {{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for _, move := range movements {
		x, y := point.x+move[0], point.y+move[1]

		// Check if the neighbor is within the grid boundaries and is passable
		if x >= 0 && x < len(matrix) && y >= 0 && y < len(matrix[0]) && matrix[x][y] != -1 {
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
	if len(voronoiPoints) > 3 {
		// return top 3
		return voronoiPoints[:3]
	}
	// else return all
	return voronoiPoints
}

// create some sample points
func createInitSamplePoints(voronoiPoints []Point, numSamplePoints int, size int) []Point {
	// initialize sample points
	samplePoints := make([]Point, 0)
	maxTries := numSamplePoints*2;
	tries := 0

	// this loop will create a number of sample points which are close to the voronoi points
	for maxTries > tries {
			tries += 1

			// pick a random voronoi point
			voronoiPoint := voronoiPoints[rand.Intn(len(voronoiPoints))]
			// add some random noise to the point
			samplePoint := Point{voronoiPoint.x + rand.Intn(3)+1, voronoiPoint.y + rand.Intn(3)+1}

			// check if within bounds
			if samplePoint.x < 0 || samplePoint.x >= size || samplePoint.y < 0 || samplePoint.y >= size {
				continue
			}

			// check if the sample point is already in the list
			breakFlag := false
			for _, point := range samplePoints {
				if samplePoint.x == point.x && samplePoint.y == point.y {
					breakFlag = true
				}
			}

			for _, point := range voronoiPoints {
				if samplePoint.x == point.x && samplePoint.y == point.y {
					breakFlag = true
				}
			}	

			if !breakFlag {
				samplePoints = append(samplePoints, samplePoint)
			}
		}
	
  // generate some sample points which are generally far from voronoi points
	tries = 0
	maxTries = numSamplePoints;
	for maxTries > tries {
		tries += 1
		samplePoint := Point{rand.Intn(size-1)+1, rand.Intn(size-1)+1}

		// check if within bounds
		if samplePoint.x < 0 || samplePoint.x >= size || samplePoint.y < 0 || samplePoint.y >= size {
			continue
		}

		// check if the sample point is already in the list
		breakFlag := false
		for _, point := range samplePoints {
			if samplePoint.x == point.x && samplePoint.y == point.y {
				breakFlag = true
			}
		}

		for _, point := range voronoiPoints {
			if samplePoint.x == point.x && samplePoint.y == point.y {
				breakFlag = true
			}
		}


		if !breakFlag {
			samplePoints = append(samplePoints, samplePoint)
		}
	}

	return samplePoints
}


// create a voronoi table data structure which maps an ID to a point value
func createVoronoiTable(voronoiPoints []Point) map[int]Point {
	voronoiTable := make(map[int]Point)
	for i, point := range voronoiPoints {
		voronoiTable[i+1] = point
	}
	// print out voronoi table
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

func medianNeighborVornoiID(matrix [][]int, colorMatrix [][]int, point Point) int {
	// check if point is wall
	if matrix[point.x][point.y] == -1 {
		return -1
	}	
	// get the color of the points neighbors
	neighbors := getNeighbors(point, matrix)
	// make a map of the colors of the neighbors with their frequency
	neighborColors := make(map[int]int)
	for _, neighbor := range neighbors {
		neighborColor := colorMatrix[neighbor.x][neighbor.y]
		neighborColors[neighborColor] += 1
	}
	// check how many of the neighbors have the same color
	// if the majority of the neighbors have the same color, return that color
	// otherwise, return -1
	if len(neighborColors) == 0 {
		return -1
	}

	maxColor := -1
	maxCount := 0
	// loop through the neighbor colors and find the most frequent color
	for color, count := range neighborColors {
		if count > maxCount {
			maxCount = count
			maxColor = color
		}
	}
	// if the most frequent color is the majority, return that color
	if maxCount > (len(neighbors) / 2 )+1 {	
		if maxColor > 0 {
			return maxColor
		}
	}
	// otherwise, return -1
	return -1
}

func calculateNearestVoronoiID(matrix [][]int, voronoiPoints []Point, voronoiTable map[int]Point, point Point) int {
	minDistance := MaxInt
	voronoiId := -1

	// check matrix for if wall
	if matrix[point.x][point.y] == -1 {
		return -1
	}

	// check if point is inside voronoi point list
	for _, voronoiPoint := range top3Voronoi(voronoiPoints, point) {
		// calculate distance from sample point to voronoi point
		distance := distance(matrix, point, voronoiPoint)
		// update matrix with voronoi point id
		if distance < minDistance {
			minDistance = distance
			voronoiId = getVoronoiId(voronoiTable, voronoiPoint)
		}
	}
	return voronoiId
}


func Voronoi(matrix [][]int, voronoiPoints []Point, size int) [][]int {
	
	// create output matrix
	outputMatrix := make([][]int, size)
	for i := range outputMatrix {
		outputMatrix[i] = make([]int, size)
	}
	
	// create initial sample points
	samplePoints := createInitSamplePoints(voronoiPoints, size*5, size)
	// fmt.Println(samplePoints)
	filledPoints := 0
	filledPointList := make([]Point, size*size)

	// add voronoi points to filledPointList
	for _, point := range voronoiPoints {
		filledPointList[filledPoints] = point
		filledPoints += 1
		outputMatrix[point.x][point.y] = 0
	}

	// add sample points to filledPointList
	for _, point := range samplePoints {
		filledPointList[filledPoints] = point
		filledPoints += 1
	}

	voronoiTable := createVoronoiTable(voronoiPoints)
	// fmt.Println(voronoiTable)

	// calculate distance from each sample point to some voronoi points
	for _, point := range samplePoints {
		voronoiId := calculateNearestVoronoiID(matrix, voronoiPoints, voronoiTable, point)
		// update output matrix with voronoi id
		outputMatrix[point.x][point.y] = voronoiId
	}


	for x := 0; x < size; x += 1 {
		for y := 0; y < size; y += 1 {
		checkPoint := Point{x, y}
		// check if the sample point is already in the list
		breakFlag := false
		for _, point := range filledPointList {
			if checkPoint.x == point.x && checkPoint.y == point.y {
				breakFlag = true
			}
		}
		if !breakFlag {
			// print the sample point
			// fmt.Println(samplePoint)

			// calculate median neighbor voronoi id
			medianNeighborVornoiID := medianNeighborVornoiID(matrix, outputMatrix, checkPoint)
			if medianNeighborVornoiID != -1 {
				// update output matrix with voronoi id
				outputMatrix[checkPoint.x][checkPoint.y] = medianNeighborVornoiID
				filledPointList[filledPoints] = checkPoint
				filledPoints += 1
			} else {
				// calculate distance to near voronoi point
				voronoiId := calculateNearestVoronoiID(matrix, voronoiPoints, voronoiTable, checkPoint)
				// update output matrix with voronoi id
				outputMatrix[checkPoint.x][checkPoint.y] = voronoiId
				if (filledPoints == size*size) {
					fmt.Println("Filled all points")
					fmt.Println(filledPointList)
					break;
				}

				filledPointList[filledPoints] = checkPoint
				filledPoints += 1
			}
		}
	}
}
	// loop through voronoi points and add in the actual voronoi id from the table
	for _, point := range voronoiPoints {
		outputMatrix[point.x][point.y] = 0
	}

	// print out filled points
	// fmt.Println(filledPointList) 
	ID := calculateNearestVoronoiID(matrix, voronoiPoints, voronoiTable, Point{0,0}) 
	outputMatrix[0][0] = ID 
	return outputMatrix
}

func FindBathrooms(matrix [][]int, size int) ([]VoronoiPoint, []Point) {
	bathrooms := make([]VoronoiPoint, 0)
	bathroomPoints := make([]Point, 0)
	for x := 0; x < size; x += 1 {
		for y := 0; y < size; y += 1 {
			if matrix[x][y] > 0 {
				bathroomPoint := Point{x, y}
				bathroomPoints = append(bathroomPoints, bathroomPoint)
				bathrooms = append(bathrooms, VoronoiPoint{bathroomPoint, matrix[x][y]})
			}
		}
	}
	return bathrooms, bathroomPoints
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


// func main() {
// 	grid := [][]int{
// 		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
// 		{0, -1, -1, -1, -1, -1, -1, -1, -1, 0},
// 		{0, -1, 0, 0, 0, 0, 0, 0, -1, 0},
// 		{0, -1, 0, 0, 0, 0, 0, 0, -1, 0},
// 		{0, -1, 0, 0, 0, 0, 0, 0, -1, 0},
// 		{0, -1, 0, 0, 0, 0, 0, 45, -1, 0},
// 		{0, -1, -1, -1, 0, -1, -1, -1, -1, 0},
// 		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
// 		{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
// 		{0, 0, 0, 0, 0, 0, 0, 0, 34, 0},
// 	}

// 	bathrooms := []Point{
// 		{9,8},
// 		{5,7},
// 	}

// 	voronoiPoints := voronoi(grid, bathrooms, 10)

// 	// print out voronoi points matrix
// 	for _, row := range voronoiPoints {
// 		fmt.Println(row)
// 	}



// 	// grid := [][]int{
// 	// 	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
// 	// 	{0, -1, -1, -1, -1, -1, -1, -1, -1, 0},
// 	// 	{0, -1, 0, 0, 0, 0, 0, 0, -1, 0},
// 	// 	{0, -1, 0, 0, 0, 0, 0, 0, -1, 0},
// 	// 	{0, -1, 0, 0, 0, 0, 0, 0, -1, 0},
// 	// 	{0, -1, 0, 0, 0, 0, 0, 0, -1, 0},
// 	// 	{0, -1, -1, -1, 0, -1, -1, -1, -1, 0},
// 	// 	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
// 	// 	{0, -1, -1, -1, -1, -1, -1, -1, -1, -1},
// 	// 	{0, 0, 0, 0, 0, 0, 0, 0, 32, 0},
// 	// }
// 	// start := Point{0, 0}
// 	// end := Point{7, 5}

// 	// fmt.Println(grid[9][8])

// 	// path, cost := astar(grid, start, end)
// 	// if path != nil {
// 	// 	fmt.Println("Shortest Path:")
// 	// 	for _, node := range path {
// 	// 		fmt.Printf("(%d, %d) -> ", node.x, node.y)
// 	// 	}
// 	// 	fmt.Printf("\nTotal Cost: %d\n", cost)
// 	// } else {
// 	// 	fmt.Println("No path found.")
// 	// }
// }