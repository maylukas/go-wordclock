package core

import "time"

type Display interface {
	GetLedIndicesToLightUp(time time.Time) []int
	GetDisplayDefinition() DisplayDefinition
}
