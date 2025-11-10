package game

import (
	"fmt"
	"log"
	"math/rand"
	"time"
	"image/color"
	
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"pybot-simulator/config"
	"pybot-simulator/entities"
)

type Game struct {
	width  int
	height int
	
	robot *entities.Robot
	cans  []*entities.Can
	
	spawnButton Button
	
	canSprite *ebiten.Image
	
	rng              *rand.Rand
	animationCounter int
}

type Button struct {
	X, Y, Width, Height float64
	Text                string
}

func NewGame(width, height int) *Game {
	g := &Game{
		width:            width,
		height:           height,
		cans:             make([]*entities.Can, 0),
		rng:              rand.New(rand.NewSource(time.Now().UnixNano())),
		animationCounter: 0,
	}
	
	// Crear robot en el centro (sin sprite todavía)
	g.robot = entities.NewRobot(
		float64(width)/2,
		float64(height)/2,
		nil, // El sprite se cargará después
	)
	
	// Cargar sprites
	g.LoadAssets()
	
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
	// Cargar sprites del robot
	g.loadRobotSprites()
	
	// Cargar sprite de la lata
	var err error
	g.canSprite, _, err = ebitenutil.NewImageFromFile("assets/can.png")
	if err != nil {
		log.Printf("No se pudo cargar can.png: %v", err)
		// Sprite temporal para la lata
		g.canSprite = ebiten.NewImage(config.CanSize, config.CanSize)
		g.canSprite.Fill(color.RGBA{255, 200, 50, 255})
	}
}

func (g *Game) loadRobotSprites() {
	sprites := make(map[string]*ebiten.Image)
	
	spriteFiles := map[string]string{
		"idle":  "assets/pybot-moves/pybot_idle.png",
		"up":    "assets/pybot-moves/pybot_walk_up.png",
		"left":  "assets/pybot-moves/pybot_walk_left.png",
		"right": "assets/pybot-moves/pybot_walk_right.png",
	}
	
	allLoaded := true
	for name, path := range spriteFiles {
		img, _, err := ebitenutil.NewImageFromFile(path)
		if err != nil {
			log.Printf("Error cargando sprite %s: %v", name, err)
			allLoaded = false
		} else {
			sprites[name] = img
		}
	}
	
	// Si se cargaron los sprites, asignarlos al robot
	if allLoaded && len(sprites) > 0 {
		g.robot.Sprites = sprites
	} else {
		// Crear sprite temporal si no se pudieron cargar
		log.Println("Usando sprites temporales para el robot")
		tempSprite := ebiten.NewImage(config.RobotSize, config.RobotSize)
		tempSprite.Fill(color.RGBA{100, 150, 255, 255})
		sprites["idle"] = tempSprite
		sprites["up"] = tempSprite
		sprites["left"] = tempSprite
		sprites["right"] = tempSprite
		g.robot.Sprites = sprites
	}
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
	g.animationCounter++
	g.HandleInput()
	g.robot.Update()
	g.CheckCollisions()
	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.width, g.height
}