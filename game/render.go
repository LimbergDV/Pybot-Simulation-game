package game

import (
	"fmt"
	"image"
	"image/color"

	"pybot-simulator/config"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
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
	vel := g.robot.Velocity
	var spriteName string
	
	if vel.X == 0 && vel.Y == 0 {
		spriteName = "idle"
	} else {
		// Priorizar dirección horizontal sobre vertical
		if vel.X < -0.1 {
			spriteName = "left"
		} else if vel.X > 0.1 {
			spriteName = "right"
		} else if vel.Y < 0 {
			spriteName = "up"
		} else {
			spriteName = "idle"
		}
	}
		
	currentSprite := g.robot.Sprites[spriteName]
	
	// Si el sprite no existe, usar idle como fallback
	if currentSprite == nil {
		fmt.Printf("Sprite %s no encontrado, usando idle\n", spriteName)
		currentSprite = g.robot.Sprites["idle"]
	}
	
	// Si aún no hay sprite, no dibujar nada
	if currentSprite == nil {
		fmt.Println("No hay sprite idle, no se puede dibujar")
		return
	}
	
	// Calcular el frame actual de la animación (4 frames por sprite)
	frameWidth := 75.0
	frameHeight := 75.0
	
	// Animar solo cuando se está moviendo
	frameIndex := 0
	if vel.X != 0 || vel.Y != 0 {
		// Ciclar entre frames 0, 1, 2, 3 basado en el tiempo
		frameIndex = (g.animationCounter / 8) % 4 // Cambia de frame cada 8 ticks
	}
	
	// Crear subimagen del frame actual
	sx := float64(frameIndex) * frameWidth
	frameRect := image.Rect(int(sx), 0, int(sx+frameWidth), int(frameHeight))
	frameImg := currentSprite.SubImage(frameRect).(*ebiten.Image)
	
	// Dibujar el frame centrado en la posición del robot
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-frameWidth/2, -frameHeight/2)
	op.GeoM.Translate(pos.X, pos.Y)
	
	screen.DrawImage(frameImg, op)
	ebitenutil.DrawCircle(screen, pos.X, pos.Y, 2, color.RGBA{255, 0, 0, 200})
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