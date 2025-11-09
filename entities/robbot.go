package entities

import (
	"github.com/hajimehoshi/ebiten/v2"
	"pybot-simulator/utils"
)

type Robot struct {
	Position      utils.Vector2D
	Velocity      utils.Vector2D
	CansCollected int
	Sprite        *ebiten.Image
	Sprites       map[string]*ebiten.Image
	minX, maxX    float64
	minY, maxY    float64
}

func NewRobot(x, y float64, sprite *ebiten.Image) *Robot {
	return &Robot{
		Position:      utils.Vector2D{X: x, Y: y},
		Velocity:      utils.Vector2D{X: 0, Y: 0},
		CansCollected: 0,
		Sprite:        sprite,
		Sprites:       make(map[string]*ebiten.Image),
	}
}

func (r *Robot) Update() {
	newX := r.Position.X + r.Velocity.X
	newY := r.Position.Y + r.Velocity.Y
	
	if newX >= r.minX && newX <= r.maxX {
		r.Position.X = newX
	}
	if newY >= r.minY && newY <= r.maxY {
		r.Position.Y = newY
	}
}

func (r *Robot) SetVelocity(vx, vy float64) {
	r.Velocity.X = vx
	r.Velocity.Y = vy
}

func (r *Robot) SetBounds(minX, maxX, minY, maxY float64) {
	r.minX = minX
	r.maxX = maxX
	r.minY = minY
	r.maxY = maxY
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