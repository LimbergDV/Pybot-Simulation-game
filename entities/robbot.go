package entities

import (
	"github.com/hajimehoshi/ebiten/v2"
	"pybot-simulator/utils"
)

type Robot struct {
	Position      utils.Vector2D
	CansCollected int
	Sprite        *ebiten.Image
}

func NewRobot(x, y float64, sprite *ebiten.Image) *Robot {
	return &Robot{
		Position: utils.Vector2D{X: x, Y: y},
		Sprite:   sprite,
	}
}

func (r *Robot) GetPosition() utils.Vector2D {
	return r.Position
}

func (r *Robot) SetPosition(x, y float64) {
	r.Position.X = x
	r.Position.Y = y
}

func (r *Robot) CollectCan() {
	r.CansCollected++
}

func (r *Robot) GetCansCollected() int {
	return r.CansCollected
}