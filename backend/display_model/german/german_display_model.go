package german

import (
	"time"
	"log"
	"github.com/maylukas/go-wordclock/backend/display_model/utils"
	"github.com/maylukas/go-wordclock/backend/display_model/core"
)

type LedDisplayModel struct {
	nonHourWords map[string][]int
	hourWords    map[string][]int
}

func (disp *LedDisplayModel) GetDisplayDefinition() core.DisplayDefinition {
	return core.DisplayDefinition{
		Lines: 10,
		LineLength: 11,
	}
}

func (disp *LedDisplayModel) getDispCharacts() string {
	return "ESKISTAFÜNF" +
		"ZEHNZWANZIG" +
		"DREIVIERTEL" +
		"VORFUNKNACH" +
		"HALBAELFÜNF" +
		"EINSXAMZWEI" +
		"DREIPMJVIER" +
		"SECHSNLACHT" +
		"SIEBENZWÖLF" +
		"ZEHNEUNKUHR"
}

func (disp *LedDisplayModel) getNonHourWord(word string) []int {
	return utils.GetDisplayIndices(disp.nonHourWords, word, disp.getDispCharacts(), false)
}

func (disp *LedDisplayModel) getWordForHour(hour int) string {
	switch hour {
	case 1:
		return "EINS"
	case 2:
		return "ZWEI"
	case 3:
		return "DREI"
	case 4:
		return "VIER"
	case 5:
		return "FÜNF"
	case 6:
		return "SECHS"
	case 7:
		return "SIEBEN"
	case 8:
		return "ACHT"
	case 9:
		return "NEUN"
	case 10:
		return "ZEHN"
	case 11:
		return "ELF"
	case 12:
		return "ZWÖLF"
	default:
		log.Fatalf("Unsupported hour number: %d", hour)
	}
	return ""
}

func (disp *LedDisplayModel) getIndicesForHour(hourWord string) []int {
	return utils.GetDisplayIndices(disp.hourWords, hourWord, disp.getDispCharacts(), true)
}

func (disp *LedDisplayModel) GetLedIndicesToLightUp(t time.Time) []int {
	controlWords, hourWords := disp.translateTimeToWords(t)
	indices := make([]int, 0)

	for _, controlWord := range controlWords {
		indices = append(indices, disp.getNonHourWord(controlWord)...)
	}

	for _, hourWord := range hourWords {
		indices = append(indices, disp.getIndicesForHour(hourWord)...)
	}
	return indices
}

func (disp *LedDisplayModel) translateTimeToWords(t time.Time) (controlWords []string, hourWords []string ) {
	hour := t.Hour()
	minute := t.Minute()
	if hour == 0 && minute < 25 {
		hour = 12
	}
	if ( hour == 12 && minute >= 25) || hour > 12 {
		hour -= 12
	}
	// Use 5 min intervals as this is the clock's display precision
	minDivFive := minute / 5
	controlWords = []string{"ES", "IST"}
	hourWords = make([]string, 0)
	addControlWord := func(word string) {
		controlWords = append(controlWords, word)
	}
	addHour := func(hour int) {
		hourWords = append(hourWords, disp.getWordForHour(hour))
	}
	if minDivFive < 5 || minDivFive > 7 {
		fiveMinDiffToHour := minDivFive
		if minDivFive > 7 {
			fiveMinDiffToHour = 12 - minDivFive
			addControlWord("VOR")
			addHour(hour + 1)
		} else {
			addHour(hour)
			if minDivFive != 0 {
				addControlWord("NACH")
			}
		}
		switch fiveMinDiffToHour {
		case 0:
			addControlWord("UHR")
		case 1:
			addControlWord("FÜNF")
		case 2:
			addControlWord("ZEHN")
		case 3:
			addControlWord("VIERTEL")
		case 4:
			addControlWord("ZWANZIG")
		}
	} else {
		addHour(hour + 1)
		addControlWord("HALB")
		if minDivFive == 5 {
			addControlWord("FÜNF")
			addControlWord("VOR")
		}

		if minDivFive == 7 {
			addControlWord("FÜNF")
			addControlWord("NACH")
		}
	}
	return controlWords, hourWords
}
