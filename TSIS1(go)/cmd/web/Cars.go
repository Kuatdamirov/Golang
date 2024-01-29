package main
import ("encoding/json"
"fmt" 
"net/http"
"github.com/gorilla/mux"
)

type Response struct {
	Cars []Car `json:"cars"`
}

type Car struct {
	Mark  string    `json:"mark"`
	Model string `json:"model"`
	Year  int `json:"Year"`
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	//specify status code
	w.WriteHeader(http.StatusOK)
  
  //update response writer 
	fmt.Fprintf(w, "API is up and running")
}

func Cars(w http.ResponseWriter, r *http.Request) {
	//declare response variable
	var response Response
  
	//Retrieve person details
	marks := prepareResponse()
  
	//assign person details to response
	response.Cars = marks
  
	//update content type
	w.Header().Set("Content-Type", "application/json")
  
	//specify HTTP status code
	w.WriteHeader(http.StatusOK)
  
	//convert struct to JSON
	jsonResponse, err := json.Marshal(response)
	if err != nil {
	 return
	}
  
	//update response
	w.Write(jsonResponse)
  }

  func prepareResponse() []Car {
	var cars []Car
  
	var myCar Car
	myCar.Mark = "Hyundai"
	myCar.Model = "Sonata"
	myCar.Year = 2021
	cars = append(cars, myCar)
  
	myCar.Mark = "Toyota"
	myCar.Model = "Supra"
	myCar.Year = 2000
	cars = append(cars, myCar)
  
	myCar.Mark = "Cadillac"
	myCar.Model = "Deville"
	myCar.Year = 1975
	cars = append(cars, myCar)

	return cars
  }

  func main() { router := mux.NewRouter()
	router.HandleFunc("/health-check", HealthCheck).Methods("GET")
	router.HandleFunc("/cars", Cars).Methods("GET")
	http.Handle("/", router)

	//start and listen to requests
	http.ListenAndServe(":8080", router)}