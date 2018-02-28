package core

import "time"

type LedDisplay interface {
	UpdateDisplay(t time.Time)
}
