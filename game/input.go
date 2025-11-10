package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (g *Game) HandleInput() {
	// Spawn de latas con tecla S
	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		g.SpawnCans(3)
	}
	
	// Recargar batería con tecla R
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		g.robot.Recharge()
	}
	
	// Click en botones
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		mx, my := ebiten.CursorPosition()
		fx, fy := float64(mx), float64(my)
		
		if g.IsPointInButton(g.spawnButton, fx, fy) {
			g.SpawnCans(3)
		}
		
		if g.IsPointInButton(g.rechargeButton, fx, fy) {
			g.robot.Recharge()
		}
	}
}