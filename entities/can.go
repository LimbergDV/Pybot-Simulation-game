package entities

import (
	"github.com/hajimehoshi/ebiten/v2"
	"pybot-simulator/utils"
)

type Can struct {
	Position utils.Vector2D
	Active   bool
	Sprite   *ebiten.Image
}

func NewCan(x, y float64, sprite *ebiten.Image) *Can {
	return &Can{
		Position: utils.Vector2D{X: x, Y: y},
		Active:   true,
		Sprite:   sprite,
	}
}

func (c *Can) GetPosition() utils.Vector2D {
	return c.Position
}

func (c *Can) IsActive() bool {
	return c.Active
}

func (c *Can) Deactivate() {
	c.Active = false
}