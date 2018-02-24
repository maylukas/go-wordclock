package led_display

import (
	"time"
	"github.com/maylukas/go-wordclock/backend/display_model/model"
	"github.com/maylukas/go-wordclock/backend/repository/config_repository"
)

type LedDisplay interface {
	UpdateDisplay(t time.Time)
}

type ledDisplayImpl struct {
	model  model.Display
	config config_repository.Config
	//strip ws281x.Matrix
}

func NewDisplay(model model.Display /*, config *ws281x.HardwareConfig*/) LedDisplay {
	//def := model.GetDisplayDefinition()
	//s, err := ws281x.NewWS281x(def.LineLength * def.Lines, config)
	//s.Initialize()
	//if err != nil {
	//	log.Fatalf("Could not initialize LED strip: %v", err)
	//}
	c := config_repository.GetConfig()
	c.DisplayConfig.Color.RGBA()
	return &ledDisplayImpl{
		model:  model,
		config: c.DisplayConfig,
	}
}

func (l *ledDisplayImpl) UpdateDisplay(t time.Time) {
	//indices := l.model.GetLedIndicesToLightUp(t)

	// Calculate the led indices to light up
	// Will map the indices, if the flow should be alternated (See config)
	//calcIndices := l.calcIndices(indices)
	//for _, ledIdx := range calcIndices {
	//	l.strip.Set(ledIdx, l.config.Color)
	//}
	//l.strip.Render()
}

func (l *ledDisplayImpl) calcIndices(indices []int) []int {
	if l.config.AlternateFlow {
		mappedIndices := make([]int, len(indices))
		for idx, ledIdx := range indices {
			mappedIndices[idx] = l.mapIndex(ledIdx)
		}
		return mappedIndices
	}
	return indices
}

func (l *ledDisplayImpl) mapIndex(ledIdx int) int {
	dd := l.model.GetDisplayDefinition()

	l.model.GetDisplayDefinition()
	isEvenRow := ledIdx/dd.LineLength%2 == 1
	if l.config.AlternateEven == isEvenRow {
		return ledIdx/dd.LineLength*dd.LineLength + dd.LineLength - (ledIdx % dd.LineLength) - 1
	} else {
		return ledIdx
	}
}
