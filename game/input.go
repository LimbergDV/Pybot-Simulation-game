package game

import (
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
	
	// Movimiento del robot
	g.HandleRobotMovement()
}

func (g *Game) HandleRobotMovement() {
	robot := g.robot
	
	// Movimiento con flechas o WASD
	if ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.IsKeyPressed(ebiten.KeyW) {
		robot.Position.Y -= config.RobotSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		robot.Position.Y += config.RobotSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyA) {
		robot.Position.X -= config.RobotSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.IsKeyPressed(ebiten.KeyD) {
		robot.Position.X += config.RobotSpeed
	}
	
	g.ClampRobotPosition()
}

func (g *Game) ClampRobotPosition() {
	margin := float64(config.GridMargin)
	
	if g.robot.Position.X < margin {
		g.robot.Position.X = margin
	}
	if g.robot.Position.X > float64(g.width)-margin {
		g.robot.Position.X = float64(g.width) - margin
	}
	if g.robot.Position.Y < margin {
		g.robot.Position.Y = margin
	}
	if g.robot.Position.Y > float64(g.height)-margin {
		g.robot.Position.Y = float64(g.height) - margin
	}
}

func (g *Game) IsPointInButton(x, y float64) bool {
	btn := g.spawnButton
	return x >= btn.X && x <= btn.X+btn.Width &&
		y >= btn.Y && y <= btn.Y+btn.Height
}