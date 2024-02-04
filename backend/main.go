package main 

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http" 
	"os"
	"math/rand"
	"time"
	"errors"
	// "voronoi"
)

const bathroomsDB = "bathroomsDB.json"
const bathroomObjectDB = "bathroomObjectDB.json"

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
	Name        string       `json:"name"`
	Coordinates []Coordinates `json:"coordinates"`
	Grid        [][]int       `json:"grid"`
	Bathrooms   []Bathroom    `json:"bathrooms"` 
}

type BathroomMapOutput struct {
	Name        string       `json:"name"`
	Coordinates []Coordinates `json:"coordinates"`
	Grid        [][]int       `json:"grid"`
	Bathrooms   []Bathroom    `json:"bathrooms"` 
	ID 					int 					`json:"ID"`
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
		Name: bathroomMap.Name,
		ID: generateUniqueID(),
		Coordinates: bathroomMap.Coordinates,
		Grid: bathroomMap.Grid,
		Bathrooms: bathroomMap.Bathrooms,
	}
	return bathroomMapOutput
}

func ConvertBathroomObjectToOutput(bathroomObject bathroomObjectNoID) Bathroom {
	bathroomObjectOutput := Bathroom{
		// ID               int    `json:"id"`
		// Name             string `json:"name"`
		// Gender           string `json:"gender"`
		// Accessible       bool   `json:"accessible"`
		// MenstrualProduct bool   `json:"menstrualProducts"`
		ID: generateUniqueID(),
		Name: bathroomObject.Name,
		Gender: bathroomObject.Gender,
		Accessible: bathroomObject.Accessible,
		MenstrualProduct: bathroomObject.MenstrualProduct,
	}
	return bathroomObjectOutput
}

type bathroomObjectNoID struct {
	Name             string `json:"name"`
	Gender           string `json:"gender"`
	Accessible       bool   `json:"accessible"`
	MenstrualProduct bool   `json:"menstrualProducts"`
}

func bathroomObjectWriteHandler(w http.ResponseWriter, r *http.Request) {
	// Only allow POST requests
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decode JSON request
	var bathroomObject bathroomObjectNoID
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&bathroomObject); err != nil {
		http.Error(w, "Invalid JSON input", http.StatusBadRequest)
		fmt.Println("Error:", err)
		return
	}
	defer r.Body.Close()

	bathroomObjectOutput := ConvertBathroomObjectToOutput(bathroomObject)

	// Write the bathroomObject to the file
	if err := writeBathroomToFile(bathroomObjectOutput); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// turn bathroomObject back into JSON
	jsonResponse, err := json.Marshal(bathroomObjectOutput)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return 
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func writeBathroomToFile(bathroomObject Bathroom) error {
	// Read existing data from file
	file, err := os.ReadFile(bathroomObjectDB)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	// Unmarshal the JSON data into a slice of BathroomMap objects
	var bathrooms []Bathroom
	err = json.Unmarshal(file, &bathrooms)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	// Append the new bathroomMap to the JSON
	bathrooms = append(bathrooms, bathroomObject)

	// Convert the bathrooms to JSON
	jsonData, err := json.Marshal(bathrooms)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}	

	// Write the new JSON to the file
	err = os.WriteFile(bathroomObjectDB, jsonData, 0644)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	return nil
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
	// Unmarshal the JSON data into a slice of BathroomMap objects
	var bathroomMaps []BathroomMapOutput
	err = json.Unmarshal(file, &bathroomMaps)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	// Append the new bathroomMap to the JSON
	bathroomMaps = append(bathroomMaps, bathroomMap)

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
	ID int `json:"ID"`
}
//Converts BathroomMapOutput to BathroomGet
func ConvertOutputToGet(bathroomMap BathroomMapOutput) BathroomGet {
    bathroomGet := BathroomGet{
        Name: bathroomMap.Name,
        ID: bathroomMap.ID,
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
        fmt.Println("Error:", err)
        return nil, err
    }
    // transform the data into an array of BathroomGet Structs   
    var bathroomGets [] BathroomGet 
    for _, maps := range bathroomOutputs{ 
        bathroomGets = append(bathroomGets, ConvertOutputToGet(maps))
    } 
    return bathroomGets, err; 
}
//bathroom maps by both name and ID 
func bathroomGetHandler(w http.ResponseWriter, r *http.Request){ 
	//Allow only request 
	if r.Method != http.MethodGet { 
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	} 

	bathroomMaps, err := getBathroomMapsFromFile()
	if err!= nil { 
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
	// Only allow get requests
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
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

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
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
	http.HandleFunc("/", enableCORS(handler))
	// Define the endpoint and handler function
	http.HandleFunc("/api/voronoi", enableCORS(voronoiHandler))
	http.HandleFunc("/api/bathroom/write", enableCORS(bathroomWriteHandler))
	http.HandleFunc("/api/bathroom/object/write", enableCORS(bathroomObjectWriteHandler))
	http.HandleFunc("/api/bathroom/get/id", enableCORS(bathroomGetByIDHandler))
	http.HandleFunc("/api/bathroom/get/maps", enableCORS(bathroomGetHandler))

	// Specify the directory containing the files
	dir := "./images/"
	// Create a file server handler for the specified directory
	fileServer := http.FileServer(http.Dir(dir))
	// Create a handler function to serve files with modified URLs
	http.Handle("/images/", http.StripPrefix("/images", fileServer))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
