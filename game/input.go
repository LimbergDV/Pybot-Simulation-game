package game

import (
	"math"
	
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"pybot-simulator/config"
)

func (g *Game) HandleInput() {
	// Spawn de latas con tecla S
	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		g.SpawnCans(3)
	}
	
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		mx, my := ebiten.CursorPosition()
		if g.IsPointInButton(float64(mx), float64(my)) {
			g.SpawnCans(3)
		}
	}
	
	g.HandleRobotMovement()
}

func (g *Game) HandleRobotMovement() {
	dx, dy := 0.0, 0.0
	
	// Movimiento con flechas o WASD
	if ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.IsKeyPressed(ebiten.KeyW) {
		dy = -1
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		dy = 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyA) {
		dx = -1
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.IsKeyPressed(ebiten.KeyD) {
		dx = 1
	}
	
	// Normalizar el vector de dirección para movimiento diagonal
	if dx != 0 || dy != 0 {
		length := math.Sqrt(dx*dx + dy*dy)
		dx /= length
		dy /= length
	}
	
	// Establecer velocidad
	g.robot.SetVelocity(dx*config.RobotSpeed, dy*config.RobotSpeed)
	
	// Establecer límites del área de juego
	margin := float64(config.GridMargin)
	g.robot.SetBounds(
		margin,
		float64(g.width)-margin,
		margin,
		float64(g.height)-margin,
	)
}

func (g *Game) IsPointInButton(x, y float64) bool {
	btn := g.spawnButton
	return x >= btn.X && x <= btn.X+btn.Width &&
		y >= btn.Y && y <= btn.Y+btn.Height
}