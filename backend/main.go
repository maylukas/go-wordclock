package main

import (
	"github.com/maylukas/go-wordclock/backend/led_display"
	"github.com/maylukas/go-wordclock/backend/display_model/api"
	"github.com/maylukas/go-wordclock/backend/http_api"
	"time"
	"golang.org/x/text/language"
	"github.com/maylukas/go-wordclock/backend/led_display/core"
	"fmt"
	"github.com/jgarff/rpi_ws281x/golang/ws2811"
	"os"
)

func main() {
	defer ws2811.Fini()

	fmt.Println("Wordclock started")
	pin := 18
	count := 150
	brightness := 255
	err := ws2811.Init(pin, count, brightness)
	if err != nil {
		fmt.Printf("Error when initilaizing strip: %v", err)
		os.Exit(-1)
	}

	displayModel := api.GetDisplayModelForLanguage(language.German)
	display := led_display.NewDisplay(displayModel)
	fmt.Println("Scheduling periodic update...")
	go updateDisplayPeriodically(display)
	fmt.Println("Setting up http server...")
	http_api.Configure()
}

func updateDisplayPeriodically(display core.LedDisplay) {
	for {
		display.UpdateDisplay(time.Now())
		time.Sleep(500 * time.Millisecond)
	}
}
