package ui

import (
	"image/color"
	"log"

	"github.com/ebitenui/ebitenui"
	eimage "github.com/ebitenui/ebitenui/image"
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
			widget.RowLayoutOpts.Spacing(20),
			widget.RowLayoutOpts.Padding(widget.Insets{
				Top: 10, Bottom: 10, Left: 10, Right: 10,
			}),
		)),
		widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.GridLayoutData{
			HorizontalPosition: widget.GridLayoutPositionEnd,
			MaxWidth:           300,
		})),
	)

	sliderContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewStackedLayout(
			widget.StackedLayoutOpts.Padding(widget.Insets{
				Top: 10, Bottom: 10, Left: 10, Right: 10,
			}),
		)),
	)

	trackGraphic := widget.NewGraphic(
		widget.GraphicOpts.Image(manager.SliderResources.TrackImage),
	)

	handleGraphic := widget.NewGraphic(
		widget.GraphicOpts.Image(manager.SliderResources.HandleImage),
	)

	explorationSlider := widget.NewSlider(
		widget.SliderOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.RowLayoutData{
			Position: widget.RowLayoutPositionCenter,
		}), widget.WidgetOpts.MinSize(200, 6)),
		widget.SliderOpts.MinMax(1, 10),
		widget.SliderOpts.Images(&widget.SliderTrackImage{}, &widget.ButtonImage{
			Idle:    eimage.NewNineSliceColor(color.RGBA{255, 255, 255, 255}),
			Pressed: eimage.NewNineSliceColor(color.RGBA{255, 255, 255, 255}),
		}),
		widget.SliderOpts.FixedHandleSize(10),
		widget.SliderOpts.TrackOffset(5),
		widget.SliderOpts.PageSizeFunc(func() int {
			return 1
		}),
		widget.SliderOpts.ChangedHandler(func(args *widget.SliderChangedEventArgs) {
			log.Println(args.Current)
			gld := settingsPanel.GetWidget().LayoutData.(widget.GridLayoutData)
			gld.HorizontalPosition = widget.GridLayoutPositionStart
			settingsPanel.GetWidget().LayoutData = gld
			settingsPanel.RequestRelayout()
		}),
	)

	sliderContainer.AddChild(handleGraphic)
	sliderContainer.AddChild(trackGraphic)
	sliderContainer.AddChild(explorationSlider)

	settingsPanel.AddChild(sliderContainer)

	rootContainer.AddChild(gameView)
	rootContainer.AddChild(settingsPanel)

	return &ebitenui.UI{Container: rootContainer}
}
