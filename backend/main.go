package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

const bathroomsDB = "bathroomsDB.json"

// VoronoiRequest represents the JSON input structure.
type VoronoiRequest struct {
	Matrix [][]int `json:"matrix"`
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

	// run some code to fish out the bathrooms (all points who are greater than 0)
	bathroomVoronoi, _ := FindBathrooms(voronoiReq.Matrix)

	// create the voronoi output array
	voronoiOutput := Voronoi(voronoiReq.Matrix, bathroomVoronoi)

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

// Bathroom represents the bathroom details.
type Bathroom struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	Gender           string `json:"gender"`
	Accessible       bool   `json:"accessible"`
	MenstrualProduct bool   `json:"menstrualProducts"`
}

// Coordinates represents the latitude and longitude of a location.
type Coordinates struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

// BathroomMap represents the main structure for unmarshaling the JSON.
type BathroomMap struct {
	Name        string        `json:"name"`
	Coordinates []Coordinates `json:"coordinates"`
	Grid        [][]int       `json:"grid"`
	Bathrooms   []Bathroom    `json:"bathrooms"`
}

type BathroomMapOutput struct {
	Name        string        `json:"name"`
	Coordinates []Coordinates `json:"coordinates"`
	Grid        [][]int       `json:"grid"`
	Bathrooms   []Bathroom    `json:"bathrooms"`
	ID          int           `json:"ID"`
	Time        time.Time     `json:"time"`
	Delete      bool          `json:"delete"`
}

func generateUniqueID() int {
	// Seed the random number generator with the current time
	rand.Seed(time.Now().UnixNano())

	// Generate a random 9-digit integer ID
	id := rand.Intn(900000000) + 100000000

	return id
}

func ConvertBathroomMapToOutput(bathroomMap BathroomMap) BathroomMapOutput {
	bathroomMapOutput := BathroomMapOutput{
		Name:        bathroomMap.Name,
		ID:          generateUniqueID(),
		Time:        time.Now(),
		Delete:      true,
		Coordinates: bathroomMap.Coordinates,
		Grid:        bathroomMap.Grid,
		Bathrooms:   bathroomMap.Bathrooms,
	}
	return bathroomMapOutput
}

// check if more than one hour ago
func isMoreThanOneHourAgo(t time.Time) bool {
	return time.Now().Sub(t) > time.Hour
}

// check if more than one minute ago
func isMoreThanOneMinuteAgo(t time.Time) bool {
	return time.Now().Sub(t) > time.Minute
}

func bathroomWriteHandler(w http.ResponseWriter, r *http.Request) {
	// Only allow POST requests
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decode JSON request
	var bathroomMap BathroomMap
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&bathroomMap); err != nil {
		http.Error(w, "Invalid JSON input", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	bathroomMapOutput := ConvertBathroomMapToOutput(bathroomMap)

	// Write the bathroomMap to the file
	if err := writeBathroomMapToFile(bathroomMapOutput); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// turn bathroomMap back into JSON
	jsonResponse, err := json.Marshal(bathroomMapOutput)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

// write a new bathroom map to the file
func writeBathroomMapToFile(bathroomMap BathroomMapOutput) error {
	// Read existing data from file
	file, err := os.ReadFile(bathroomsDB)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	// add time and delete to the bathroomMap
	bathroomMap.Time = time.Now()
	bathroomMap.Delete = true

	// Unmarshal the JSON data into a slice of BathroomMap objects
	var bathroomMaps []BathroomMapOutput
	err = json.Unmarshal(file, &bathroomMaps)
	if err != nil {
		fmt.Println("Error Unmarshal JSON Write to File:", err)
		return err
	}

	// Append the new bathroomMap to the JSON
	bathroomMaps = append(bathroomMaps, bathroomMap)

	// // Remove entries older than 5 minutes
	// currentTime := time.Now()
	// var updatedBathroomMaps []BathroomMapOutput
	// for _, bm := range bathroomMaps {
	// 	if currentTime.Sub(bm.CreatedAt) <= 5*time.Minute {
	// 		updatedBathroomMaps = append(updatedBathroomMaps, bm)
	// 	}
	// }

	// Convert the bathroomMaps to JSON
	jsonData, err := json.Marshal(bathroomMaps)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	// Print the parsed data
	fmt.Printf("Name: %s\n", bathroomMap.Name)
	fmt.Println("ID:", bathroomMap.ID)
	fmt.Println("Coordinates:")
	for _, coord := range bathroomMap.Coordinates {
		fmt.Printf("Lat: %f, Lng: %f\n", coord.Lat, coord.Lng)
	}

	fmt.Println("Grid:")
	for _, row := range bathroomMap.Grid {
		fmt.Println(row)
	}

	fmt.Println("Bathrooms:")
	for _, bath := range bathroomMap.Bathrooms {
		fmt.Printf("ID: %d, Name: %s, Gender: %s, Accessible: %t, MenstrualProduct: %t\n",
			bath.ID, bath.Name, bath.Gender, bath.Accessible, bath.MenstrualProduct)
	}

	// Write the new JSON to the file
	err = os.WriteFile(bathroomsDB, jsonData, 0644)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	return nil
}

type BathroomGet struct {
	Name string `json:"name"`
	ID   int    `json:"ID"`
}

// Converts BathroomMapOutput to BathroomGet
func ConvertOutputToGet(bathroomMap BathroomMapOutput) BathroomGet {
	bathroomGet := BathroomGet{
		Name: bathroomMap.Name,
		ID:   bathroomMap.ID,
	}
	return bathroomGet
}

// Get BathroomMaps from the file
func getBathroomMapsFromFile() ([]BathroomGet, error) {
	// Read existing data from file
	file, err := os.ReadFile(bathroomsDB)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	// Unmarshal the JSON data into a slice of BathroomMap objects
	var bathroomOutputs []BathroomMapOutput
	err = json.Unmarshal(file, &bathroomOutputs)
	if err != nil {
		fmt.Println("Error Unmarshal JSON get Map:", err)
		return nil, err
	}

	// transform the data into an array of BathroomGet Structs
	var bathroomGets []BathroomGet
	var updatedBathroomOutputs []BathroomMapOutput
	for _, maps := range bathroomOutputs {
		// fmt.Println(maps)
		// fmt.Println(maps.time)
		// fmt.Println(maps.delete)

		// check if the bathroom is more than one hour old and delete is true
		if isMoreThanOneHourAgo(maps.Time) && maps.Delete {
			continue
		}
		updatedBathroomOutputs = append(updatedBathroomOutputs, maps)
		bathroomGets = append(bathroomGets, ConvertOutputToGet(maps))
	}

	// update the file with the new data
	jsonData, err := json.MarshalIndent(updatedBathroomOutputs, "", " ")
	if err != nil {
		fmt.Println("Error Writing Back to JSON:", err)
		return nil, err
	}
	err = os.WriteFile(bathroomsDB, jsonData, 0644)

	return bathroomGets, err
}

// bathroom maps by both name and ID
func bathroomGetHandler(w http.ResponseWriter, r *http.Request) {
	//Allow only request
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	bathroomMaps, err := getBathroomMapsFromFile()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.Marshal(bathroomMaps)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

type BathroomID struct {
	ID int `json:"ID"`
}

// bathroom map by id handler
func bathroomGetByIDHandler(w http.ResponseWriter, r *http.Request) {
	// Decode JSON request
	var bathroomID BathroomID
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&bathroomID); err != nil {
		http.Error(w, "Invalid JSON input", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	bathroomMap, err := getBathroomMapsByID(bathroomID.ID)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// turn bathroomMap back into JSON
	jsonResponse, err := json.Marshal(bathroomMap)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

// get bathroom map by ID
func getBathroomMapsByID(id int) (BathroomMapOutput, error) {
	// Read existing data from file
	file, err := os.ReadFile(bathroomsDB)
	if err != nil {
		fmt.Println("Error:", err)
		return BathroomMapOutput{}, err
	}

	// Unmarshal the JSON data into a slice of BathroomMap objects
	var bathroomMaps []BathroomMapOutput
	err = json.Unmarshal(file, &bathroomMaps)
	if err != nil {
		fmt.Println("Error:", err)
		return BathroomMapOutput{}, err
	}

	// find the bathroom map with the given ID
	for _, bathroomMap := range bathroomMaps {
		if bathroomMap.ID == id {
			return bathroomMap, nil
		}
	}

	return BathroomMapOutput{}, errors.New("BathroomMap not found")
}

// enableCORS is a middleware function to enable CORS for all origins
func enableCORS(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Allow all origins
		w.Header().Set("Access-Control-Allow-Origin", "*")
		// Allow the necessary methods
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		// Allow the necessary headers
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
		// Allow credentials if needed
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		// Handle the actual request
		if r.Method == "OPTIONS" {
			// Preflight request, respond with 200 OK
			w.WriteHeader(http.StatusOK)
			return
		}

		// Call the original handler function
		handler(w, r)
	}
}

func main() {
	// Define the endpoint and handler function
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprint(w, "Welcome to the bathroom finder API!")
	// })
	http.HandleFunc("/api/voronoi", enableCORS(voronoiHandler))
	http.HandleFunc("/api/bathroom/write", enableCORS(bathroomWriteHandler))
	http.HandleFunc("/api/bathroom/maps/id", enableCORS(bathroomGetByIDHandler))
	http.HandleFunc("/api/bathroom/maps", enableCORS(bathroomGetHandler))

	// Specify the directory containing the files
	dir := "./images/"
	// Create a file server handler for the specified directory
	fileServer := http.FileServer(http.Dir(dir))
	// Create a handler function to serve files with modified URLs
	http.Handle("/images/", http.StripPrefix("/images", fileServer))

	log.Fatal(http.ListenAndServe(":8080", nil))
}