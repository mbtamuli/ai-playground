package system

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"

	"github.com/mbtamuli/ai-playground/snake/src/game/component"
)

type Controls struct {
	query *query.Query
}

func NewControls() *Controls {
	return &Controls{
		query: query.NewQuery(
			filter.Contains(
				component.Position,
				component.Input,
				component.Velocity,
				component.Sprite,
			),
		),
	}
}
func (i *Controls) Update(w donburi.World) {
	i.query.EachEntity(w, func(entry *donburi.Entry) {
		input := component.GetInput(entry)
		velocity := component.GetVelocity(entry)

		velocity.X = 0
		velocity.Y = 0

		switch {
		case ebiten.IsKeyPressed(input.MoveUpKey):
			velocity.Y = -input.MoveSpeed
		case ebiten.IsKeyPressed(input.MoveDownKey):
			velocity.Y = input.MoveSpeed
		case ebiten.IsKeyPressed(input.MoveRightKey):
			velocity.X = input.MoveSpeed
		case ebiten.IsKeyPressed(input.MoveLeftKey):
			velocity.X = -input.MoveSpeed
		}
	})
}
