package ui

import (
	"image/color"
	"log"

	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"
)

func createGameScreen(manager *UIManager) *ebitenui.UI {
	rootContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewGridLayout(
			widget.GridLayoutOpts.Columns(2),
			widget.GridLayoutOpts.Stretch([]bool{true, true}, []bool{true}),
			widget.GridLayoutOpts.Spacing(1, 1),
		)),
	)

	// Game view container (left side - 70%)
	gameView := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewRowLayout()),
		widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.GridLayoutData{
			HorizontalPosition: widget.GridLayoutPositionStart,
			MaxWidth:           700,
		})),
	)

	// Settings panel (right side - 30%)
	settingsPanel := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionVertical),
		)),
		widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.GridLayoutData{
			HorizontalPosition: widget.GridLayoutPositionEnd,
			MaxWidth:           300,
		})),
		widget.ContainerOpts.BackgroundImage(manager.BackgroundResources.Image),
	)

	// Exploration rate slider: 1-10
	explorationSlider := createCustomLabeledSlider(manager, "Exploration Rate", 1, 10, func(args *widget.SliderChangedEventArgs) {
		log.Printf("Exploration rate: %d\n", args.Current)
	})

	// Discount Factor slider: 1-5
	discountSlider := createCustomLabeledSlider(manager, "Discount Factor", 1, 5, func(args *widget.SliderChangedEventArgs) {
		log.Printf("Discount Factor: %d\n", args.Current)
	})

	// Learning Rate slider: 1-10
	learningSlider := createCustomLabeledSlider(manager, "Learning Rate", 1, 10, func(args *widget.SliderChangedEventArgs) {
		log.Printf("Learning Rate: %d\n", args.Current)
	})

	settingsPanel.AddChild(explorationSlider)
	settingsPanel.AddChild(discountSlider)
	settingsPanel.AddChild(learningSlider)

	rootContainer.AddChild(gameView)
	rootContainer.AddChild(settingsPanel)

	return &ebitenui.UI{Container: rootContainer}
}

func createCustomLabeledSlider(manager *UIManager, label string, min, max int, changedHandler func(args *widget.SliderChangedEventArgs)) *widget.Container {
	rootContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionHorizontal),
		)),
		widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.RowLayoutData{
			Position: widget.RowLayoutPositionStart,
			Stretch:  true,
		}), widget.WidgetOpts.MinSize(200, 6)),
	)

	slider := widget.NewSlider(
		widget.SliderOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.RowLayoutData{
			Position: widget.RowLayoutPositionEnd,
			Stretch:  true,
		}), widget.WidgetOpts.MinSize(100, 6)),
		widget.SliderOpts.MinMax(min, max),
		widget.SliderOpts.Images(manager.SliderResources.TrackImage, manager.SliderResources.HandleImage),
		widget.SliderOpts.FixedHandleSize(10),
		widget.SliderOpts.TrackOffset(1),
		widget.SliderOpts.PageSizeFunc(func() int {
			return 1
		}),
		widget.SliderOpts.ChangedHandler(changedHandler),
	)

	text := widget.NewText(
		widget.TextOpts.Text(label, manager.Fonts.Face, color.Black),
		widget.TextOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.RowLayoutData{
			Position: widget.RowLayoutPositionStart,
		})),
	)

	rootContainer.AddChild(text)
	rootContainer.AddChild(slider)

	return rootContainer
}
