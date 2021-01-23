package database

import (
	. "github.com/matthewojenkins/programservice/model"
	"log"
)

func SaveProgram(program string, description string) error {
	db := GetDB()
	_, err := db.Exec("INSERT INTO program (program_name, description) VALUES($1, $2)", program, description)
	return err
}

func SaveWorkout(workout Workout) Workout {
	db := GetDB()
	result, err := db.Exec("INSERT INTO workout (program_id, cycle, cycle_name, phase, phase_name, workout_date, comment) "+
		" VALUES($1, $2, $3, $4, $5, $6, $7) returning workout_id", workout.ProgramID, workout.Cycle, workout.CycleName, workout.Phase, workout.PhaseName, workout.WorkoutDate, workout.Comment)
	if err != nil {
		log.Fatal(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Println("Yeah, messy - no new id returned for workout")
	}
	workout.WorkoutID = int(id)
	return workout
}

func GetWorkout(id int) Workout {
	db := GetDB()

	sqlStatement := "SELECT workout_id, program_id, COALESCE(cycle,0), coalesce(cycle_name, ''), coalesce(phase,0), coalesce(phase_name,''), workout_date, comment  FROM workout " +
		"WHERE workout_id=$1;"
	var workout Workout

	row := db.QueryRow(sqlStatement, id)
	err := row.Scan(&workout.WorkoutID, &workout.ProgramID, &workout.Cycle, &workout.CycleName, &workout.Phase, &workout.PhaseName, &workout.WorkoutDate, &workout.Comment)
	if err != nil {
		log.Fatal(err)
	}

	return workout
}

func IsDatabaseUp() bool {
	db := GetDB()
	_, err := db.Exec("select 1")
	if err != nil {
		return false
	} else {
		return true
	}
}

func GetProgram(id int) Program {
	db := GetDB()

	sqlStatement := "SELECT program_id, program_name, description FROM program WHERE program_id=$1;"
	var programName, description string
	var programID int

	row := db.QueryRow(sqlStatement, id)
	err := row.Scan(&programID, &programName, &description)
	if err != nil {
		log.Fatal(err)
	}

	return Program{programID, programName, description}
}

func GetPrograms() []Program {
	db := GetDB()

	sqlStatement := "SELECT program_id, program_name, description FROM program order by program_name;"
	var programName string
	var programID int
	var description string
	var program Program
	var programs []Program

	rows, err := db.Query(sqlStatement)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&programID, &programName, &description)
		if err != nil {
			log.Fatal(err)
		}
		program = Program{programID, programName, description}
		programs = append(programs, program)
	}
	return programs
}

func GetNextWorkout(lastWorkout Workout) []Set {
	var sets []Set

	return sets
}

func SaveSet(set Set) error {
	db := GetDB()
	_, err := db.Exec("INSERT INTO set (workout_id, exercise, weight, weight_unit, reps, hold_time, set_order, start_date, end_date) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)", set.WorkoutID, set.Exercise, set.Weight, set.WeightUnit, set.Reps, set.HoldTime, set.SetOrder, set.SetStart, set.SetEnd)
	return err
}

func GetWorkoutSets(id int) []Set {
	db := GetDB()
	sqlStatement := "SELECT set_id, exercise, start_date, end_date, coalesce(weight,0), coalesce(weight_unit,''), reps, coalesce(hold_time,0), set_order FROM set where workout_id = $1 order by set_order;"
	var sets []Set

	rows, err := db.Query(sqlStatement, id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var set Set
		err := rows.Scan(&set.SetID, &set.Exercise, &set.SetStart, &set.SetEnd, &set.Weight, &set.WeightUnit, &set.Reps, &set.HoldTime, &set.SetOrder)
		if err != nil {
			log.Fatal(err)
		}
		sets = append(sets, set)
	}
	return sets
}
