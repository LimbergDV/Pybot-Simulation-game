package game

import (
	"fmt"
	"image/color"
	
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"pybot-simulator/config"
)

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{40, 40, 50, 255})
	g.DrawPlayArea(screen)
	g.DrawCans(screen)
	g.DrawRobot(screen)
	g.DrawButton(screen)
	g.DrawInfo(screen)
}

func (g *Game) DrawPlayArea(screen *ebiten.Image) {
	margin := float64(config.GridMargin)
	areaColor := color.RGBA{80, 80, 100, 255}
	
	x := margin
	y := margin
	w := float64(g.width) - 2*margin
	h := float64(g.height) - 2*margin
	
	ebitenutil.DrawRect(screen, x, y, w, 2, areaColor)
	ebitenutil.DrawRect(screen, x, y+h-2, w, 2, areaColor)
	ebitenutil.DrawRect(screen, x, y, 2, h, areaColor)
	ebitenutil.DrawRect(screen, x+w-2, y, 2, h, areaColor)
}

func (g *Game) DrawRobot(screen *ebiten.Image) {
	pos := g.robot.Position
	
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-config.RobotSize/2, -config.RobotSize/2)
	op.GeoM.Translate(pos.X, pos.Y)
	
	if g.robot.Sprite != nil {
		screen.DrawImage(g.robot.Sprite, op)
	}
	
	ebitenutil.DrawCircle(screen, pos.X, pos.Y, 4, color.RGBA{255, 255, 255, 255})
}

func (g *Game) DrawCans(screen *ebiten.Image) {
	for _, can := range g.cans {
		if !can.Active {
			continue
		}
		
		pos := can.Position
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(-config.CanSize/2, -config.CanSize/2)
		op.GeoM.Translate(pos.X, pos.Y)
		
		if can.Sprite != nil {
			screen.DrawImage(can.Sprite, op)
		}
	}
}

func (g *Game) DrawButton(screen *ebiten.Image) {
	btn := g.spawnButton
	
	ebitenutil.DrawRect(screen, btn.X, btn.Y, btn.Width, btn.Height,
		color.RGBA{70, 130, 180, 255})
	
	ebitenutil.DrawRect(screen, btn.X, btn.Y, btn.Width, 2, color.RGBA{200, 200, 200, 255})
	ebitenutil.DrawRect(screen, btn.X, btn.Y+btn.Height-2, btn.Width, 2, color.RGBA{200, 200, 200, 255})
	ebitenutil.DrawRect(screen, btn.X, btn.Y, 2, btn.Height, color.RGBA{200, 200, 200, 255})
	ebitenutil.DrawRect(screen, btn.X+btn.Width-2, btn.Y, 2, btn.Height, color.RGBA{200, 200, 200, 255})
	
	ebitenutil.DebugPrintAt(screen, btn.Text, int(btn.X+15), int(btn.Y+13))
}

func (g *Game) DrawInfo(screen *ebiten.Image) {
	activeCans := g.GetActiveCansCount()
	
	info := fmt.Sprintf("Recolectadas: %d | Activas: %d | FPS: %.1f",
		g.robot.CansCollected, activeCans, ebiten.ActualTPS())
	ebitenutil.DebugPrintAt(screen, info, 10, 10)
	
	controls := "Controles: Flechas o WASD para mover | S para spawn latas"
	ebitenutil.DebugPrintAt(screen, controls, 10, 30)
}