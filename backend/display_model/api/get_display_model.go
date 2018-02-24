package api

import (
	"github.com/maylukas/go-wordclock/backend/display_model/german"
	"golang.org/x/text/language"
	"github.com/maylukas/go-wordclock/backend/display_model/model"
)


func GetDisplayModelForLanguage(l language.Tag) model.Display {
	if l == language.German {
		return &german.LedDisplayModel{}
	}
	return nil
}