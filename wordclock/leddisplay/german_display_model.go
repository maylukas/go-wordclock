package leddisplay

import (
	"time"
	"log"
)

type GermanLedDisplayModel struct {
	nonHourWords    map[string][]int
	hourWords       map[string][]int
}

func (disp *GermanLedDisplayModel) getDispCharacts() string {
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

func (disp *GermanLedDisplayModel) getNonHourWord(word string) []int {
	return getDisplayIndices(disp.nonHourWords, word, disp.getDispCharacts(), false)
}

func (disp *GermanLedDisplayModel) getWordForHour(hour int) string {
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

func (disp *GermanLedDisplayModel) getIndicesForHour(hour int) []int {
	return getDisplayIndices(disp.hourWords, disp.getWordForHour(hour), disp.getDispCharacts(), true)
}

func (disp *GermanLedDisplayModel) GetLedIndicesToLightUp(t time.Time) []int {
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
	indices := append(disp.getNonHourWord("ES"), disp.getNonHourWord("IST")...)

	addControlWord := func(word string) {
		indices = append(indices, disp.getNonHourWord(word)...)
	}

	addHour := func(hour int) {
		indices = append(indices, disp.getIndicesForHour(hour)...)
	}

	if minDivFive < 5 || minDivFive > 7 {
		fiveMinDiffToHour := minDivFive
		if minDivFive > 7 {
			fiveMinDiffToHour = 12 - minDivFive
			addControlWord("VOR")
			addHour(hour+1)
		} else {
			addHour(hour)
			if minDivFive != 0 {
				addControlWord("NACH")
			}
		}
		switch fiveMinDiffToHour {
		case 1: addControlWord("FÜNF")
		case 2: addControlWord("ZEHN")
		case 3: addControlWord("VIERTEL")
		case 4: addControlWord("ZWANZIG")
		}
	} else {
		addHour(hour +1)
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

 	return indices

}
