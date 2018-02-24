package core

type Display interface {
	GetLedIndicesToLightUp(time time.Time) []int
	GetDisplayDefinition() DisplayDefinition
}
