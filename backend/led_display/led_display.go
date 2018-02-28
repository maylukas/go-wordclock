package led_display

import (
	"time"
	"github.com/maylukas/go-wordclock/backend/display_model/core"
	"github.com/maylukas/go-wordclock/backend/repository/config_repository"
	led_display_core "github.com/maylukas/go-wordclock/backend/led_display/core"
	"fmt"
	"github.com/jgarff/rpi_ws281x/golang/ws2811"
)

type ledDisplayImpl struct {
	model  core.Display
	config config_repository.Config
}

func NewDisplay(model core.Display) led_display_core.LedDisplay {
	fmt.Println("Initializing strip...")
	c := config_repository.GetConfig()
	disp := &ledDisplayImpl{
		model:  model,
		config: c.DisplayConfig,
	}

	var color uint32 = 0
	color += 0 << 24
	color += 0 << 16
	color += 255 << 8
	color += 255

	for i := 0; i < 110; i++ {
		disp.ResetStrip()
		ws2811.SetLed(disp.mapIndex(i), color)
		ws2811.Render()
		time.Sleep(100 * time.Millisecond)
	}
	return disp
}

func (l *ledDisplayImpl) ResetStrip() {
	for i := 0; i < 110; i++ {
		var color uint32 = 0
		color += 0 << 24
		color += 0 << 16
		color += 0 << 8
		color += 0
		ws2811.SetLed(l.mapIndex(i), color)
	}
}

func (l *ledDisplayImpl) UpdateDisplay(t time.Time) {
	indices := l.model.GetLedIndicesToLightUp(t)

	// Calculate the led indices to light up
	// Will map the indices, if the flow should be alternated (See config)
	calcIndices := l.calcIndices(indices)
	fmt.Printf("Lighting up indices %v", calcIndices)
	l.ResetStrip()
	for _, ledIdx := range calcIndices {
		r, g, b, a := l.config.Color.RGBA()
		var color uint32 = 0
		color += r << 24
		color += g << 16
		color += b << 8
		color += a

		ws2811.SetLed(ledIdx,color)
	}
	ws2811.Render()
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
