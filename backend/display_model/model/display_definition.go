package model

import "time"

type DisplayDefinition struct {
	Lines int
	LineLength int
}

type Display interface {
	GetLedIndicesToLightUp(time time.Time) []int
	GetDisplayDefinition() DisplayDefinition
}