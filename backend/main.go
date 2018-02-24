package main

import (
	"github.com/maylukas/go-wordclock/backend/led_display"
	"github.com/maylukas/go-wordclock/backend/display_model/api"
	"golang.org/x/text/language"
	"github.com/maylukas/go-wordclock/backend/http_api"
	"time"
)

func main() {
	displayModel := api.GetDisplayModelForLanguage(language.German)
	display := led_display.NewDisplay(displayModel)
	go updateDisplayPeriodically(display)
	http_api.Configure()
}

func updateDisplayPeriodically(display led_display.LedDisplay) {
	for {
		display.UpdateDisplay(time.Now())
		time.Sleep(500 * time.Millisecond)
	}
}
