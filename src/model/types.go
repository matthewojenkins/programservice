package model

import (
	"time"
)

type Program struct {
	ProgramID   int
	ProgramName string
	Description string
}

type Set struct {
	SetID      int
	WorkoutID  int
	SetStart   time.Time
	SetEnd     time.Time
	Exercise   string
	Weight     int
	WeightUnit string
	Reps       int
	HoldTime   int
	SetOrder   int
}

type Workout struct {
	WorkoutID   int
	WorkoutDate time.Time
	ProgramID   int
	Cycle       int
	CycleName   string
	Phase       int
	PhaseName   string
	Comment     string
}

type Progression struct {
	ProgramID                int
	RepIncrease              int
	RepIncreaseStart         int
	RepIncreaseMax           int
	WeightIncrease           int
	WeightIncreasePercentage int
	PhaseCountBeforeIncrease int
}
