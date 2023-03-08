package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/logrusorgru/aurora"
	"github.com/nipunsgolangrepo/nipunsgolangrepo/Datatypes"
	"github.com/nipunsgolangrepo/nipunsgolangrepo/Docker"
	"github.com/nipunsgolangrepo/nipunsgolangrepo/Multithreading"
	"github.com/nipunsgolangrepo/nipunsgolangrepo/MyRest"
	"log"
	"net/http"
	"reflect"
)

var Flagprogram string

func init() {
	flag.StringVar(&Flagprogram, "Flagprogram", "", "Program to Execute out of the following.")
}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

func main() {
	//To be changed - Implement a mec that based on the case
	// We execute a function and all of the functions are getting executed async
	fmt.Println(aurora.Yellow("Welcome to Nipun's main package"))
	flag.Parse()

	if !isFlagPassed("Flagprogram") {
		log.Fatal("Mandatory -Flagprogram is not provided")
	}

	fmt.Printf(`Running a function out of the follwing Depending on flag provided
	1. Maps_freq_calculator
	2. Mymatchingsubarray
	3. Getuniquesubstring
	4. GetTriplets
	5. Rundatatypes
	6. Middlelement
	7. Bubblesort
	8. Bubblesortwithrecursion
	9. Insertionsort
	10. Mergesort
	11. DirectRecursion
	12. Multithreading
	13. Basicchannel
	14. SortFrequency
	` + "\n")

	fmt.Printf("Algorithm to run is: %v and flag type is: %v\n\n", Flagprogram, reflect.TypeOf(Flagprogram))

	switch Flagprogram {

	case ("Maps_freq_calculator"):
		arr := []int{30, 40, 50, 60, 70, 40, 60, 90, 100, 30, 50}

		fmt.Println("Frequency map: ", Datatypes.Maps_freq_calculator(arr))

	case ("Mymatchingsubarray"):

		fmt.Println("Frequency map: ", Datatypes.Mymatchingsubarray())

	case ("Getuniquesubstring"):

		fmt.Println("Frequency map: ", Datatypes.Getuniquesubstring())

	case ("GetTriplets"):

		fmt.Println("Frequency map: ", Datatypes.GetTriplets())

	case ("Rundatatypes"):

		Datatypes.Rundatatypes()

	case ("Middlelement"):
		arr1 := []int{2, 4, 5, 1, 5, 3, 5, 1, 1}
		Datatypes.Middletype(arr1)

	case ("Bubblesort"):
		arr1 := []int{2, 4, 8, 9, 5, 3, 1, 7, 6}
		Datatypes.Bubblesort(arr1)

	case ("Bubblesortwithrecursion"):
		arr1 := []int{2, 4, 8, 9, 5, 3, 1, 7, 6}
		arrlength := len(arr1)
		count := 0
		Datatypes.Bubblesortwithrecursion(arr1, arrlength, count)

	case ("Insertionsort"):
		arr1 := []int{2, 4, 8, 9, 5, 3, 1, 7, 6}
		Datatypes.InsertionSort(arr1)

	case ("Selectionsort"):
		// arr1 := []int{2, 4, 8, 9, 5, 3, 1, 7, 6}
		arr1 := []int{2, 5, 2, 8, 5, 5, 6, 8, 8, 8}
		Datatypes.Selectionsort(arr1)

	case ("Mergesort"):
		arr1 := []int{2, 4, 8, 9, 5, 3, 1, 6}
		fmt.Println(Datatypes.Mergesort(arr1))

	case ("DirectRecursion"):
		num := 5
		fmt.Println(Datatypes.DirectRecursion(num))

	case ("Basicgoroutines"):
		Multithreading.Basicgoroutines()

	case ("Basicchannel"):
		Multithreading.Basicchannel()

	case ("Leader"):
		arr := []int{5, 18, 2, 10, 9, 3}
		Datatypes.Leader(arr)

	case ("SortFrequency"):
		arr := []int{2, 5, 2, 8, 5, 5, 6, 8, 8, 8}
		Datatypes.SortByFrequency(arr)

	case ("EquilibriumPoint"):
		arr := []int{1, 4, 2, 5}
		Datatypes.EquilibriumPoint(arr)

	case ("RevArrGrp"):
		//Shrinking array problem
		arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
		k := 4
		Datatypes.RevArrGrp(arr, k)

	case ("Setzigzagarray"):
		//Shrinking array problem
		arr := []int{4, 1, 3, 7, 9, 10, 2, 8, 6, 5}
		// arr := []int{4, 3, 7, 8, 6, 2, 1}
		Datatypes.Setzigzagarray(arr)

	case ("SushiRollsDetails"):
		//Shrinking array problem

		MyRest.Rolls = append(MyRest.Rolls, MyRest.Roll{ID: "1",
			ImageNumber: "8",
			Name:        "Spicy Tuna Roll",
			Ingredients: "Tuna, Chili sauce, Nori, Rice"})
		MyRest.Rolls = append(MyRest.Rolls, MyRest.Roll{ID: "2",
			ImageNumber: "9",
			Name:        "Mild Tuna Roll",
			Ingredients: "salt and pepper"})
		// arr := []int{4, 3, 7, 8, 6, 2, 1}
		router := mux.NewRouter()
		router.HandleFunc("/sushi", MyRest.GetSushiRollsDetails).Methods("GET")
		router.HandleFunc("/sushi/{id}", MyRest.GetSushiRollDetails).Methods("GET")
		router.HandleFunc("/sushi", MyRest.CreateSushiRoll).Methods("POST")
		log.Fatal(http.ListenAndServe(":5000", router))

	case ("CreateDocker"):
		dockerfile := "Dockerfiles/ubuntu_dockerfile.txt"
		tag := "ubuntu:nipun_java_1.0.0"
		err := Docker.CreateDocker(dockerfile, tag)
		if err != nil {
			log.Fatalf("Aborting docker creation with Error:%v", err)
		}
	case ("ChannelOp"):
		c := Multithreading.Secondgo()
		fmt.Printf("channel returned is: %v %T", <-c, c)

	case ("Rungomultiple"):
		c := Multithreading.MyMain()

	}
}
