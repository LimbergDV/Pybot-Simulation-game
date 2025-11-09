package game

import (
	"fmt"
	"image/color"
	"math/rand"
	"time"
	
	"github.com/hajimehoshi/ebiten/v2"
	"pybot-simulator/config"
	"pybot-simulator/entities"
)

type Game struct {
	width  int
	height int
	
	robot *entities.Robot
	cans  []*entities.Can
	
	spawnButton Button
	
	robotSprite *ebiten.Image
	canSprite   *ebiten.Image
	
	rng *rand.Rand
}

type Button struct {
	X, Y, Width, Height float64
	Text                string
}

func NewGame(width, height int) *Game {
	g := &Game{
		width:  width,
		height: height,
		cans:   make([]*entities.Can, 0),
		rng:    rand.New(rand.NewSource(time.Now().UnixNano())),
	}
	
	// Cargar sprites
	g.LoadAssets()
	
	// Crear robot en el centro
	g.robot = entities.NewRobot(
		float64(width)/2,
		float64(height)/2,
		g.robotSprite,
	)
	
	// Crear botón
	g.spawnButton = Button{
		X:      20,
		Y:      float64(height - 60),
		Width:  160,
		Height: 40,
		Text:   "Spawn Latas (S)",
	}
	
	// Spawn inicial
	g.SpawnCans(5)
	
	return g
}

func (g *Game) LoadAssets() {
	// implementación de carga de imágenes desde archivos, no sé si esté bien, falta checar eso:
	
	// var err error
	// g.robotSprite, _, err = ebitenutil.NewImageFromFile("assets/robot.png")
	// if err != nil {
	//     log.Printf("No se pudo cargar robot.png: %v", err)
	// }
	// g.canSprite, _, err = ebitenutil.NewImageFromFile("assets/can.png")
	// if err != nil {
	//     log.Printf("No se pudo cargar can.png: %v", err)
	// }
	
	// Sprites temporales
	g.robotSprite = ebiten.NewImage(config.RobotSize, config.RobotSize)
	g.robotSprite.Fill(color.RGBA{100, 150, 255, 255})
	
	g.canSprite = ebiten.NewImage(config.CanSize, config.CanSize)
	g.canSprite.Fill(color.RGBA{255, 200, 50, 255})
}

func (g *Game) SpawnCans(count int) {
	margin := float64(config.GridMargin)
	
	for i := 0; i < count; i++ {
		x := margin + g.rng.Float64()*(float64(g.width)-2*margin)
		y := margin + g.rng.Float64()*(float64(g.height)-2*margin)
		
		can := entities.NewCan(x, y, g.canSprite)
		g.cans = append(g.cans, can)
	}
}

func (g *Game) CheckCollisions() {
	robotPos := g.robot.Position
	collectRadius := float64(config.RobotSize/2 + config.CanSize/2)
	
	for _, can := range g.cans {
		if !can.Active {
			continue
		}
		
		distance := robotPos.Distance(can.Position)
		
		if distance < collectRadius {
			can.Deactivate()
			g.robot.CollectCan()
			fmt.Printf("¡Lata recogida! Total: %d\n", g.robot.CansCollected)
		}
	}
}

func (g *Game) GetActiveCansCount() int {
	count := 0
	for _, can := range g.cans {
		if can.Active {
			count++
		}
	}
	return count
}

func (g *Game) Update() error {
	g.HandleInput()
	g.CheckCollisions()
	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.width, g.height
}