package main 

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	// "voronoi"
)

// VoronoiRequest represents the JSON input structure.
type VoronoiRequest struct {
	Matrix [][]int `json:"matrix"`
	Size   int     `json:"size"`
}

func voronoiHandler(w http.ResponseWriter, r *http.Request) {
	// Only allow POST requests
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decode JSON request
	var voronoiReq VoronoiRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&voronoiReq); err != nil {
		http.Error(w, "Invalid JSON input", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Process the VoronoiRequest (replace this with your actual Voronoi algorithm implementation)
	// Here, we simply print the received data for demonstration purposes.
	fmt.Println("Received matrix:", voronoiReq.Matrix)
	fmt.Println("Received size:", voronoiReq.Size)

	// run some code to fish out the bathrooms (all points who are greater than 0)
	bathroomVoronoi, bathroomPoints := FindBathrooms(voronoiReq.Matrix, voronoiReq.Size)

	// create the voronoi output array
	voronoiOutput := Voronoi(voronoiReq.Matrix, bathroomPoints, voronoiReq.Size)

	for _, voronoiPoint := range bathroomVoronoi {
		voronoiOutput[voronoiPoint.point.x][voronoiPoint.point.y] = voronoiPoint.id
	}

	// create the response
	jsonResponse, err := json.Marshal(voronoiOutput)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

type BathroomMap struct{ 
	ID int64 `json:"id omitempty"`  
	Name string `json:"name"`
}


func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	// Define the endpoint and handler function
	http.HandleFunc("/api/voronoi", voronoiHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
