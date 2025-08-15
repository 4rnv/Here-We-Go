package main

import (
	"encoding/json"
	"fmt"
	"image/color"
	"log"
	"net/url"
	"os"
	"time"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"

	hs "github.com/thecsw/haruhi"
)

func main() {
	//fmt.Println("Running")
	go func() {
		w := new(app.Window)
		err := run(w)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()

}

func callAPI(showOrCharacter string, count string) (string, error) {
	//fmt.Println("Function called")
	type StructJSON struct {
		Id        string `json:"_id"`
		Character string `json:"character"`
		Quote     string `json:"quote"`
		Show      string `json:"show"`
	}

	requrl := "https://yurippe.vercel.app/api/quotes"
	params := url.Values{
		"show":   {showOrCharacter},
		"random": {count},
	}
	var responseJSON []StructJSON

	er :=
		hs.
			URL(requrl).
			Params(params).
			ResponseJson(&responseJSON)
	if er != nil {
		fmt.Println(er)
	}

	jsonBytes, err := json.MarshalIndent(responseJSON, "", "  ")
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Println(string(jsonBytes))
	return string(jsonBytes), nil
}

func run(w *app.Window) error {
	w.Option(
		app.Title("Go App"),
		app.StatusColor(color.NRGBA{R: 200, G: 10, B: 50, A: 255}),
		app.NavigationColor(color.NRGBA{R: 200, G: 10, B: 50, A: 255}),
		//app.Decorated(false),
	)

	type (
		Context = layout.Context
		Dims    = layout.Dimensions
	)
	var button widget.Clickable
	theme := material.NewTheme()
	accent := color.NRGBA{R: 200, G: 10, B: 50, A: 255}
	white := color.NRGBA{R: 255, G: 255, B: 255, A: 255}
	//theme.Shaper = text.NewShaper(text.WithCollection(gofont.Collection()))
	theme.Fg = white
	theme.TextSize = 12
	var quoteText string
	var ops op.Ops
	var (
		progress float32
		loading  bool
	)
	var showEditor widget.Editor
	var countEditor widget.Editor
	showEditor.SetText("")
	showEditor.Alignment = text.Middle
	showEditor.SingleLine = true
	countEditor.SetText("1")
	countEditor.Alignment = text.Middle
	countEditor.SingleLine = true
	for {
		switch e := w.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)
			paint.Fill(gtx.Ops, color.NRGBA{R: 0, G: 0, B: 0, A: 255})
			if loading {
				layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					gtx.Constraints.Max.X = gtx.Dp(120)
					return material.ProgressBar(theme, progress).Layout(gtx)
				})
			}
			if button.Clicked(gtx) {
				userInputShow := showEditor.Text()
				count := countEditor.Text()
				loading = true
				quoteText = ""
				go func() {
					response, err := callAPI(userInputShow, count)
					progress = 0.0
					start := time.Now()
					for loading && progress < 1.0 {
						progress = float32(time.Since(start).Seconds() / 3.0)
						if progress > 1.0 {
							progress = 1.0
						}
						w.Invalidate()
						time.Sleep(50 * time.Millisecond)
					}
					if err != nil {
						log.Printf("API Err: %s", err)
						return
					}
					loading = false
					progress = 0.0
					quoteText = response
					w.Invalidate()
				}()
			}
			layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Flexed(1, func(gtx Context) Dims {
					return layout.Center.Layout(gtx, func(gtx Context) Dims {
						heading := material.Label(theme, unit.Sp(48), "Go App")
						heading.Font.Typeface = "Segoe UI"
						heading.Font.Weight = -200
						heading.Color = white
						return heading.Layout(gtx)
					})
				}),
				layout.Flexed(4, func(gtx Context) Dims {
					return layout.Center.Layout(gtx, func(gtx Context) Dims {
						quote := material.Label(theme, unit.Sp(12), quoteText)
						quote.Color = white
						return quote.Layout(gtx)
					})
				}),
				layout.Rigid(func(gtx Context) Dims {
					gtx.Constraints.Min.Y = gtx.Dp(100)
					return layout.Center.Layout(gtx, func(gtx Context) Dims {
						ed := material.Editor(theme, &showEditor, "Enter Show")
						border := widget.Border{
							Color:        color.NRGBA{255, 255, 255, 255},
							CornerRadius: unit.Dp(0),
							Width:        unit.Dp(1),
						}
						inset := layout.UniformInset(unit.Dp(40))
						return inset.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							return border.Layout(gtx, ed.Layout)
						})
					})
				}),
				layout.Rigid(layout.Spacer{Height: unit.Dp(20)}.Layout),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					gtx.Constraints.Min.Y = gtx.Dp(100)
					return layout.Center.Layout(gtx, func(gtx Context) Dims {
						ed := material.Editor(theme, &countEditor, "Number of Quotes (1-5)")
						border := widget.Border{
							Color:        color.NRGBA{255, 255, 255, 255},
							CornerRadius: unit.Dp(0),
							Width:        unit.Dp(1),
						}
						inset := layout.UniformInset(unit.Dp(20))
						return inset.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							return border.Layout(gtx, ed.Layout)
						})
					})
				}),
				layout.Flexed(1, func(gtx Context) Dims {
					return layout.Center.Layout(gtx, func(gtx Context) Dims {
						btn := material.Button(theme, &button, "New Quote")
						btn.Background = accent
						btn.CornerRadius = 0
						btn.TextSize = 24
						btn.Font.Weight = -200
						btn.Font.Typeface = "Segoe UI"
						return btn.Layout(gtx)
					})
				}),
			)
			e.Frame(gtx.Ops)
		}
	}
}
