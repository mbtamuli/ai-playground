package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"

	"github.com/mbtamuli/ai-playground/snake/assets"
	"github.com/mbtamuli/ai-playground/snake/src/game/component"
	"github.com/mbtamuli/ai-playground/snake/src/game/system"
)

const (
	snakeHead = "PNG/Blue/Default/arrow_decorative_n.png"
)

type System interface {
	Update(w donburi.World)
}
type Drawable interface {
	Draw(w donburi.World, screen *ebiten.Image)
}

type Game struct {
	world     donburi.World
	systems   []System
	drawables []Drawable
}

func createWorld(snakeImage *ebiten.Image) donburi.World {
	world := donburi.NewWorld()

	snakeEntity := world.Create(
		component.Position,
		component.Velocity,
		component.Sprite,
		component.Input)

	snake := world.Entry(snakeEntity)
	donburi.SetValue(snake, component.Position, component.PositionData{X: 100, Y: 100})
	donburi.SetValue(snake, component.Velocity, component.VelocityData{X: 1, Y: 1})
	donburi.SetValue(snake, component.Sprite, component.SpriteData{Image: snakeImage})
	donburi.SetValue(snake, component.Input, component.InputData{
		MoveUpKey:    ebiten.KeyW,
		MoveRightKey: ebiten.KeyD,
		MoveDownKey:  ebiten.KeyS,
		MoveLeftKey:  ebiten.KeyA,
		MoveSpeed:    1.0,
	})

	return world
}

// NewGame creates and returns a new Game instance
func NewGame() (*Game, error) {
	snakeImage, err := assets.LoadEbitenImageFromAsset(snakeHead)
	if err != nil {
		return nil, err
	}

	g := &Game{
		world: createWorld(snakeImage),
	}
	g.systems = []System{
		system.NewVelocity(),
		system.NewControls(),
	}
	g.drawables = []Drawable{
		system.NewRenderer(),
	}

	return g, nil
}

// Update handles the game logic updates for each frame
// Returns an error if any game update operations fail
func (g *Game) Update() error {
	for _, s := range g.systems {
		s.Update(g.world)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Clear()
	for _, s := range g.drawables {
		s.Draw(g.world, screen)
	}
}
