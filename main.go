package main

import (
	"encoding/json"
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Quotes Generator")

	quote := binding.NewString()
	quote.Set("Hi!")

	hello := widget.NewLabelWithData(quote)

	hello.Wrapping = fyne.TextWrapWord

	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Generate Quote", func() { RetrieveNewQuote(quote) }),
	))

	w.ShowAndRun()
}

type Quote struct {
	Id         string   `json:"_id"`
	Content    string   `json:"content"`
	Author     string   `json:"author"`
	Tags       []string `json:"tags"`
	AuthorSlug string   `json:"authorSlug"`
	Length     int      `json:"length"`
}

func RetrieveNewQuote(text binding.String) {

	text.Set("Retrieving new quote...")
	r, err := http.Get("https://api.quotable.io/random")

	if err != nil {
		text.Set(err.Error())
		return
	}

	defer r.Body.Close()

	quote := Quote{}
	err = json.NewDecoder(r.Body).Decode(&quote)

	if err != nil {
		text.Set(err.Error())
		return
	}

	text.Set(quote.Content + " --- " + quote.Author)
}
