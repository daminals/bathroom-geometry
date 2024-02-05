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

type VoronoiPoint struct {
	point Point
	id int
}

//Priority Queue to be used in A* 
type PriorityQueue []Point  

func(pq PriorityQueue) Len()int {return len(pq)}
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].x < pq[j].x || (pq[i].x == pq[j].x && pq[i].y < pq[j].y)
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

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
	fScore := make(map[Point]int)
	closedSet := make(map[Point]bool)

	heap.Push(&openSet, start)
	gScore[start] = 0
	fScore[start] = distanceFormula(start, end)

	for len(openSet) > 0 {
		current := heap.Pop(&openSet).(Point)

		//Check if we have reached the end
		if current == end {
			path := reconstructPath(cameFrom, start, end)
			return path, gScore[end]
		}

		closedSet[current] = true

		neighbors := getNeighbors(current, matrix)
		for _, neighbor := range neighbors {
			if closedSet[neighbor] {
				continue
			}

			tentG := gScore[current] + 1

			if _, ok := gScore[neighbor]; !ok || tentG < gScore[neighbor] {
				gScore[neighbor] = tentG
				fScore[neighbor] = gScore[neighbor] + distanceFormula(neighbor, end)
				heap.Push(&openSet, neighbor)
				cameFrom[neighbor] = current
			}
		}
	}
	return nil, -1
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

	movements := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

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
	if len(voronoiPoints) > 5 {
		// return top 3
		return voronoiPoints[:5]
	}
	// else return all
	return voronoiPoints
}

// creates an array of voronoi points which are the closest voronois to this points neighbors
func friendlyVoronoiPoints(outputMatrix [][]int, voronoiTable map[int]Point, point Point) []Point {
	neighbors := getNeighbors(point, outputMatrix)
	// voronoi points
	neighborVoronoiPoints := make([]Point, 0)
	for _, neighbor := range neighbors {
		neighborVoronoiId := outputMatrix[neighbor.x][neighbor.y]
		if (neighborVoronoiId > 0) {
			// get voronoi point
			neighborVoronoiPoint := voronoiTable[neighborVoronoiId]
			neighborVoronoiPoints = append(neighborVoronoiPoints, neighborVoronoiPoint)
		}
	}
	return neighborVoronoiPoints
}

func checkWithinBounds(point Point, sizeX, sizeY int) bool {
	return point.x < 0 || point.x >= sizeX || point.y < 0 || point.y >= sizeY
}

// create some sample points
func createInitSamplePoints(voronoiPoints []Point, numSamplePoints, sizeX, sizeY int) []Point {
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
			if checkWithinBounds(samplePoint, sizeX, sizeY) {
				continue
			}


			
			// // check if the sample point is already in the list
			// breakFlag := false
			// for _, point := range samplePoints {
			// 	if samplePoint.x == point.x && samplePoint.y == point.y {
			// 		breakFlag = true
			// 	}
			// }

			// for _, point := range voronoiPoints {
			// 	if samplePoint.x == point.x && samplePoint.y == point.y {
			// 		breakFlag = true
			// 	}
			// }	

			if !(isInArray(samplePoints, samplePoint) || isInArray(voronoiPoints, samplePoint)){
				samplePoints = append(samplePoints, samplePoint)
			}
		}
	
  // generate some sample points which are generally far from voronoi points
	tries = 0
	maxTries = numSamplePoints;
	for maxTries > tries {
		tries += 1
		samplePoint := Point{rand.Intn(sizeX-1)+1, rand.Intn(sizeY-1)+1}

		// check if within bounds
		if checkWithinBounds(samplePoint, sizeX, sizeY) {
			continue
		}

		// check if the sample point is already in the list
		// breakFlag := false
		// for _, point := range samplePoints {
		// 	if samplePoint.x == point.x && samplePoint.y == point.y {
		// 		breakFlag = true
		// 	}
		// }

		// for _, point := range voronoiPoints {
		// 	if samplePoint.x == point.x && samplePoint.y == point.y {
		// 		breakFlag = true
		// 	}
		// }


		// if !breakFlag {
			if !(isInArray(samplePoints, samplePoint) || isInArray(voronoiPoints, samplePoint)){
			samplePoints = append(samplePoints, samplePoint)
		}
	}

	return samplePoints
}


// create a voronoi table data structure which maps an ID to a point value
func createVoronoiTable(voronoiPoints []VoronoiPoint) map[int]Point {
	voronoiTable := make(map[int]Point)
	for _, point := range voronoiPoints {
		voronoiTable[point.id] = point.point
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
	if maxCount >= len(neighbors) {
		if maxColor > 0 {
			return maxColor
		}
	}
	// otherwise, return -1
	return -1
}

func isInArray(points []Point, point Point) bool {
	for _, p := range points {
		if p.x == point.x && p.y == point.y {
			return true
		}
	}
	return false
}

func combinePointList(points1, points2 []Point) []Point {
	combinedPoints := make([]Point, len(points1))
	copy(combinedPoints, points1)
	for _, point := range points2 {
		if !(isInArray(combinedPoints, point)) {
			combinedPoints = append(combinedPoints, point)
		}
	}
	return combinedPoints
}

func calculateNearestVoronoiID(matrix, outputMatrix [][]int, voronoiPoints []Point, voronoiTable map[int]Point, point Point) int {
	minDistance := MaxInt
	voronoiId := -1

	// check matrix for if wall
	if matrix[point.x][point.y] == -1 {
		return -1
	}

	friendlyVoronoiPoints := friendlyVoronoiPoints(outputMatrix, voronoiTable, point)
	top3VoronoiPoints := top3Voronoi(voronoiPoints, point)
	// combine voronoi points into one list
	voronoiPointChecklist := combinePointList(friendlyVoronoiPoints, top3VoronoiPoints)

	// check if point is inside voronoi point list
	for _, voronoiPoint := range voronoiPointChecklist {
		// calculate distance from sample point to voronoi point
		distance := distance(matrix, point, voronoiPoint)
		// update matrix with voronoi point id
		if distance < minDistance {
			minDistance = distance
			voronoiId = getVoronoiId(voronoiTable, voronoiPoint)
		}
	}
	// fmt.Println("Voronoi ID: ", voronoiId, "Distance: ", minDistance)
	if minDistance == -1 {
		return 0
	}

	return voronoiId
}


func Voronoi(matrix [][]int, voronoiPointsWithIds []VoronoiPoint) [][]int {

	// get voronoi points
	voronoiPoints := make([]Point, len(voronoiPointsWithIds))
	// voronoiIds := make([]int, len(voronoiPointsWithIds))
	for i, voronoiPoint := range voronoiPointsWithIds {
		voronoiPoints[i] = voronoiPoint.point
	}

	// get sizeX
	sizeX := len(matrix)
	// get sizeY
	sizeY := len(matrix[0])
	
	// create output matrix
	outputMatrix := make([][]int, sizeX)
	for i := range outputMatrix {
		outputMatrix[i] = make([]int, sizeY)
	}
	
	// create initial sample points
	samplePoints := createInitSamplePoints(voronoiPoints, sizeX*15, sizeX, sizeY)
	// fmt.Println(samplePoints)
	filledPoints := 0
	filledPointList := make([]Point, sizeX*sizeY)

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

	voronoiTable := createVoronoiTable(voronoiPointsWithIds)
	// fmt.Println(voronoiTable)

	// calculate distance from each sample point to some voronoi points
	for _, point := range samplePoints {
		voronoiId := calculateNearestVoronoiID(matrix,outputMatrix, voronoiPoints, voronoiTable, point)
		// update output matrix with voronoi id 
		outputMatrix[point.x][point.y] = voronoiId
	}


	for x := 0; x < sizeX; x += 1 {
		for y := 0; y < sizeY; y += 1 {
		checkPoint := Point{x, y}
		// check if the sample point is already in the list
		if !(isInArray(filledPointList, checkPoint)) {
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
				voronoiId := calculateNearestVoronoiID(matrix,outputMatrix, voronoiPoints, voronoiTable, checkPoint)
				// update output matrix with voronoi id
				outputMatrix[checkPoint.x][checkPoint.y] = voronoiId
				if (filledPoints == sizeX*sizeY) {
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
	var wg sync.WaitGroup
	for _, voronoiPointWithId := range voronoiPointsWithIds {
		wg.Add(1)
		go func(point Point, id int) {
			defer wg.Done()
			outputMatrix[point.x][point.y] = id
		}(voronoiPointWithId.point, voronoiPointWithId.id)
	}
	wg.Wait()


	// print out filled points
	// fmt.Println(filledPointList) 
	ID := calculateNearestVoronoiID(matrix,outputMatrix, voronoiPoints, voronoiTable, Point{0,0}) 
	outputMatrix[0][0] = ID 
	return outputMatrix
}

func FindBathrooms(matrix [][]int) ([]VoronoiPoint, []Point) {
	// get size x
	sizeX := len(matrix)
	// get size y
	sizeY := len(matrix[0])

	bathrooms := make([]VoronoiPoint, 0)
	bathroomPoints := make([]Point, 0)
	for x := 0; x < sizeX; x += 1 {
		for y := 0; y < sizeY; y += 1 {
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
// 	// grid := [][]int{
// 	// 	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
// 	// 	{0, -1, -1, -1, -1, -1, -1, -1, -1, 0},
// 	// 	{0, -1, 0, 0, 0, 0, 0, 0, -1, 0},
// 	// 	{0, -1, 0, 0, 0, 0, 0, 0, -1, 0},
// 	// 	{0, -1, 0, 0, 0, 0, 0, 0, -1, 0},
// 	// 	{0, -1, 0, 0, 0, 0, 0, 45, -1, 0},
// 	// 	{0, -1, -1, -1, 0, -1, -1, -1, -1, 0},
// 	// 	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
// 	// 	{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
// 	// 	{0, 0, 0, 0, 0, 0, 0, 0, 34, 0},
// 	// }

// 	grid := [][]int{
//   {0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,},
//   {0,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,0,0,0,0,0,0,0,0,0,0,0,},
// 	{0,-1,0,0,0, 4,-1,0,0,-1,-1,0,0,0,0,0,0,0,0,0,0,0,},
// 	{0,-1,0,0,0,0,0,0,0,-1,-1,0,0,0,0,0,0,0,0,0,0,0,},
// 	{0,-1,0,0,0,5,-1,0,0,-1,-1,0,0,0,0,0,0,0,0,0,0,0,},
// 	{0,-1,-1,-1,-1,-1,-1,0,0,-1,-1,0,0,0,0,0,0,0,0,0,0,0,},
// 	{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,},
// 	{0,0,0,-1,-1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,},
// 	{0,0,-1,0,0,-1,0,0,0,0,0,-1,-1,-1,-1,-1,-1,-1,-1,0,0,0,},
// 	{0,0,-1,0,0,-1,0,-1,-1,0,-1,0,0,0,0,0,0,0,0,-1,0,0,},
// 	{0,0,-1,0,2,-1,0,0,-1,0,-1,0,0,0,0,0,0,0,0,-1,0,0,},
// 	{0,0,-1,0,-1,-1,0,-1,0,0,-1,-1,-1,-1,-1,-1,-1,0,10,-1,0,0,},
// 	{0,0,-1,0,0,0,0,-1,0,0,-1,6,0,0,0,0,0,0,9,-1,0,0,},
// 	{0,0,-1,0,1,-1,0,-1,0,0,-1,0,0,0,0,0,0,0,0,-1,0,0,},
// 	{0,0,-1,0,-1,-1,0,-1,0,0,-1,7,0,0,0,0,0,0,0,0,0,0,},
// 	{0, 0, -1, 0, 0, -1, 0, -1, 0, 0, -1, 0, 0, -1, -1, 0, 0, 0, 0, -1, 0, 0,},
// 	{0, 0, 0, -1, -1, -1, 0, -1, 0, 0, -1, 0, 0, 0, -1, 0, 0, 0, 0, -1, 0, 0,},
// 	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, -1, 0, 0, 0, -1, -1, -1, 0, 0, -1, 0, 0,},
// 	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, -1, 0, 0, 0, -1, 0, 8, 0, -1, 0, 0, 0,},
// 	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, -1, 0, -1, -1, -1, 0, 11, 0, -1, 0, 0, 0,},
// 	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, -1, -1, 0, 0, 0, 0,},
// }
	
// 	// {{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0},{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0},{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0},{0,-1,-1,-1,-1,-1,-1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,0,0,0,0,0,0,0,0,0,0,0},{0,-1,0,0,0,0,-1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,-1,0,0,0,0,0,0,0,0,0,-1,0,0,0,0,0,0,0,0,0,0,0},{0,-1,0,0,0,0,-1,-1,0,0,0,0,0,0,0,0,-1,-1,-1,-1,-1,-1,0,0,0,0,0,0,0,0,0,-1,-1,-1,0,0,0,0,0,0,0,0,0},{0,-1,0,0,0,0,0,-1,0,0,0,0,0,0,0,0,-1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,-1,0,0,0,0,0,0,0,0,0},{0,-1,0,0,0,0,0,-1,-1,0,-1,-1,0,0,0,0,-1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,2,0,-1,0,0,0,0,0,0,0,0,0},{0,-1,0,0,0,0,0,0,0,0,0,-1,0,0,0,0,-1,0,1,4,0,0,0,0,0,0,0,0,0,0,0,5,0,-1,-1,0,0,0,0,0,0,0,0},{0,-1,0,0,0,0,0,0,0,0,0,-1,0,0,0,0,-1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,-1,0,0,0,0,0,0,0,0},{-1,-1,0,0,0,0,0,0,0,0,0,-1,0,0,0,0,-1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,-1,0,0,0,0,0,0,0,0},{-1,0,0,0,0,0,0,0,0,0,0,-1,0,0,0,0,-1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,-1,0,0,0,0,0,0,0,0},{-1,0,0,0,0,0,0,0,0,0,0,-1,0,0,0,0,-1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,-1,0,0,0,0,0,0,0,0},{-1,0,0,0,0,0,0,0,0,0,0,-1,0,0,0,0,-1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,-1,0,0,0,0,0,0,0,0},{-1,0,0,0,0,0,0,0,0,0,0,-1,0,0,0,0,-1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,-1,0,0,0,0,0,0,0,0},{-1,0,0,-1,0,8,0,0,0,0,0,-1,0,0,0,0,-1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,-1,0,0,0,0,0,0,0,0},{-1,0,0,-1,0,0,0,0,0,0,-1,-1,0,0,0,0,-1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,-1,0,0,0,0,0,0,0,0},{-1,0,0,-1,0,0,0,0,0,0,-1,0,0,0,-1,-1,-1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,-1,0,0,0,0,0,0,0,0},{-1,0,0,-1,0,0,0,0,0,0,-1,0,0,0,-1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0},{-1,-1,-1,-1,0,0,0,0,0,0,-1,0,0,0,-1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,-1,0,0,0,0,0,0,0,0},{0,0,0,0,0,0,-1,-1,-1,-1,-1,0,0,0,-1,-1,-1,-1,-1,-1,-1,0,-1,0,0,0,0,0,0,0,0,0,0,0,-1,0,0,0,0,0,0,0,0},{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,-1,0,0,0,0,0,0,-1,0,0,0,0,0,0,0,0,0,0,-1,-1,0,0,0,0,0,0,0,0},{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,-1,0,0,0,0,0,0,-1,0,0,0,0,0,0,0,0,0,0,-1,0,0,0,0,0,0,0,0,0},{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,-1,0,0,0,0,0,0,-1,0,0,0,0,0,0,0,3,0,0,-1,0,0,0,0,0,0,0,0,0},{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,-1,0,0,0,0,0,0,-1,0,0,0,0,0,0,0,6,0,0,-1,0,0,0,0,0,0,0,0,0},{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,-1,0,0,0,0,0,0,-1,0,0,0,0,0,0,0,0,0,0,-1,0,0,0,0,0,0,0,0,0},{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,-1,0,0,0,0,0,0,-1,0,0,0,0,0,0,0,0,0,0,-1,0,0,0,0,0,0,0,0,0},{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,-1,-1,0,-1,-1,-1,0,-1,0,0,0,0,0,0,0,-1,-1,-1,-1,0,0,0,0,0,0,0,0,0},{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,-1,-1,-1,-1,-1,-1,-1,-1,0,-1,-1,0,0,0,0,0,0,0,0,0,0,0,0},{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0},{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}}

// 	bathrooms, _ := FindBathrooms(grid)
// 	// print out bathrooms
// 	for _, bathroom := range bathrooms {
// 		fmt.Println(bathroom.id, bathroom.point)
// 	}

// 	// voronoiPoints := Voronoi(grid, bathrooms)

// 	// // print out voronoi points matrix
// 	// for _, row := range voronoiPoints {
// 	// 	fmt.Println(row)
// 	// }

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
// 	start := Point{0, 5}
// 	end := Point{2, 5} // 4

// 	fmt.Println("Start:", start)
// 	fmt.Println("End:", end)
// 	path, cost := astar(grid, start, end)
// 	if path != nil {
// 		fmt.Println("Shortest Path:")
// 		for _, node := range path {
// 			fmt.Printf("(%d, %d) -> ", node.x, node.y)
// 		}
// 		fmt.Printf("\nTotal Cost: %d\n", cost)
// 	} else {
// 		fmt.Println("No path found.")
// 	}

// 	fmt.Println("Start:", start)
// 	fmt.Println("End:", end)
// 	end = Point{12, 11} // 6
// 	path, cost = astar(grid, start, end)
// 	if path != nil {
// 		fmt.Println("Shortest Path:")
// 		for _, node := range path {
// 			fmt.Printf("(%d, %d) -> ", node.x, node.y)
// 		}
// 		fmt.Printf("\nTotal Cost: %d\n", cost)
// 	} else {
// 		fmt.Println("No path found.")
// 	}

// 	// fmt.Println("Start:", start)
// 	// fmt.Println("End:", end)
// 	// end = Point{14, 11} // 7
// 	// path, cost = astar(grid, start, end)
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