package ui

import (
	"image/color"
	"log"

	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/image"
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

	// Add sliders for AI settings
	//explorationSlider := widget.NewSlider(
	//	widget.SliderOpts.Value(g.gameManager.ExplorationRate),
	//	widget.SliderOpts.Range(0, 1),
	//	widget.SliderOpts.OnChanged(func(args *widget.SliderChangedEventArgs) {
	//		g.gameManager.UpdateExplorationRate(args.Value)
	//	}),
	//)

	sliderContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewStackedLayout(
			widget.StackedLayoutOpts.Padding(widget.Insets{
				Top: 10, Bottom: 10, Left: 10, Right: 10,
			}),
		)),
	)

	trackGraphic := widget.NewGraphic(
		widget.GraphicOpts.Image(manager.SliderResources.TrackImage),
		widget.GraphicOpts.WidgetOpts(widget.WidgetOpts.Dropped(
			func(args *widget.DragAndDropDroppedEventArgs) {
				log.Println("Dropped")
				log.Println(args)
			},
		)),
	)

	handleGraphic := widget.NewGraphic(
		widget.GraphicOpts.Image(manager.SliderResources.HandleImage),
		widget.GraphicOpts.WidgetOpts(widget.WidgetOpts.Dropped(
			func(args *widget.DragAndDropDroppedEventArgs) {
				log.Println("Dropped")
				log.Println(args)
			},
		)),
	)

	explorationSlider := widget.NewSlider(
		widget.SliderOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.RowLayoutData{
			Position: widget.RowLayoutPositionCenter,
		}), widget.WidgetOpts.MinSize(200, 6)),
		widget.SliderOpts.MinMax(1, 10),
		widget.SliderOpts.Images(&widget.SliderTrackImage{}, &widget.ButtonImage{
			Idle:    image.NewNineSliceColor(color.RGBA{255, 255, 255, 255}),
			Pressed: image.NewNineSliceColor(color.RGBA{255, 255, 255, 255}),
		}),
		widget.SliderOpts.FixedHandleSize(10),
		widget.SliderOpts.TrackOffset(5),
		widget.SliderOpts.PageSizeFunc(func() int {
			return 1
		}),
		widget.SliderOpts.ChangedHandler(func(args *widget.SliderChangedEventArgs) {
			log.Println(args.Current)
		}),
	)

	sliderContainer.AddChild(handleGraphic)
	sliderContainer.AddChild(trackGraphic)
	sliderContainer.AddChild(explorationSlider)

	//riskSlider := widget.NewSlider(
	//	widget.SliderOpts.Value(g.gameManager.RiskAppetite),
	//	widget.SliderOpts.Range(0, 1),
	//	widget.SliderOpts.OnChanged(func(args *widget.SliderChangedEventArgs) {
	//		g.gameManager.UpdateRiskAppetite(args.Value)
	//	}),
	//)

	//learningSlider := widget.NewSlider(
	//	widget.SliderOpts.Value(g.gameManager.LearningRate),
	//	widget.SliderOpts.Range(0, 1),
	//	widget.SliderOpts.OnChanged(func(args *widget.SliderChangedEventArgs) {
	//		g.gameManager.UpdateLearningRate(args.Value)
	//	}),
	//)

	settingsPanel.AddChild(sliderContainer)
	//settingsPanel.AddChild(riskSlider)
	//settingsPanel.AddChild(learningSlider)

	rootContainer.AddChild(gameView)
	rootContainer.AddChild(settingsPanel)

	return &ebitenui.UI{Container: rootContainer}
}
