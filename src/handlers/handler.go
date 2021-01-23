package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	. "github.com/matthewojenkins/programservice/database"
	. "github.com/matthewojenkins/programservice/model"
)

func HandleProgramRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		fmt.Println("Handling program POST")
		HandleProgramPost(w, r)
	}
	if r.Method == http.MethodGet {
		fmt.Println("Handling program GET ")
		HandleProgramGet(w, r)
	}
}

func HandleSetRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		fmt.Println("Handling set POST - done a set - well done!")
		HandleSetPost(w, r)
	}
	if r.Method == http.MethodGet {
		fmt.Println("Handling program GET ")
		HandleSetGet(w, r)
	}
}

func HandleWorkoutRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		fmt.Println("Handling workout POST - doing a workout - nice.")
		HandleWorkoutPost(w, r)
	}
	if r.Method == http.MethodGet {
		fmt.Println("Handling workout GET ")
		HandleWorkoutGet(w, r)
	}
}

func HandleWorkoutGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	workoutId, _ := strconv.Atoi(strings.Join(r.URL.Query()["workoutid"], ""))

	// See if we want a list of sets for this workout
	if workoutId > 0 {
		workout := GetWorkout(workoutId)
		fmt.Println("Workout found : ", workoutId)

		err := json.NewEncoder(w).Encode(workout)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func HandleWorkoutPost(w http.ResponseWriter, r *http.Request) {
	var workout Workout

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&workout)
	if err != nil {
		panic(err)
	}

	workout = SaveWorkout(workout)

	w.WriteHeader(201)
	err = json.NewEncoder(w).Encode(workout)
	if err != nil {
		fmt.Println(err)
	}
}

func HandleProgramPost(w http.ResponseWriter, r *http.Request) {
	var program Program

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&program)
	if err != nil {
		panic(err)
	}

	err = SaveProgram(program.ProgramName, program.Description)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	w.WriteHeader(201)
}

func HandleSetPost(w http.ResponseWriter, r *http.Request) {
	var set Set

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&set)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	err = SaveSet(set)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	w.WriteHeader(201)
}

func HandleSetGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	workoutId, _ := strconv.Atoi(strings.Join(r.URL.Query()["workoutid"], ""))

	// See if we want a list of sets for this workout
	if workoutId > 0 {
		sets := GetWorkoutSets(workoutId)
		fmt.Println("Workout found : ", workoutId)

		err := json.NewEncoder(w).Encode(sets)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func HandleHomePage(w http.ResponseWriter, r *http.Request) {

}

func HandleStatusPage(w http.ResponseWriter, r *http.Request) {
	var dbCheck string
	if IsDatabaseUp() {
		dbCheck = "Database is alive"
	} else {
		dbCheck = "Database is down"
	}
	err := json.NewEncoder(w).Encode(dbCheck)
	if err != nil {
		fmt.Println(err)
	}
}

func HandleProgramGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(strings.Join(r.URL.Query()["ID"], ""))

	if id > 0 {
		program := GetProgram(id)
		fmt.Println("Program found : ", program)

		err := json.NewEncoder(w).Encode(program)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		programs := GetPrograms()
		err := json.NewEncoder(w).Encode(programs)
		if err != nil {
			fmt.Println(err)
		}
	}
}
