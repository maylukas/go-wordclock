package leddisplay

import "time"

type Model interface {
 GetLedIndicesToLightUp(time time.Time) []int
}

type Language int
const (
	GERMAN Language = iota
)

func GetDisplayForLanguage(l Language) Model {
	if l == GERMAN {
		return &GermanLedDisplayModel{}
	}
	return nil
}