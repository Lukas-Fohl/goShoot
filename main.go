package main

import (
	"fmt"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type player struct {
	rotation float64
	position [2]float64
}

func defaultTextureDesc() textureDesc {
	return textureDesc{isSet: false, texture: ""}
}

type textureDesc struct {
	isSet   bool
	texture string
}

func defaultBlock() block {
	return block{collision: false, floorTexture: defaultTextureDesc(), wallTexture: defaultTextureDesc(), ceilingTexture: defaultTextureDesc(), interaction: normal}
}

type blockType int

const (
	normal blockType = iota
)

type block struct {
	collision      bool
	floorTexture   textureDesc
	wallTexture    textureDesc
	ceilingTexture textureDesc
	interaction    blockType
}

func defaultPlayer() player {
	return player{position: [2]float64{0.0, 0.0}, rotation: 0}
}

const screenHeight = int32(1000)
const screenWidth = int32(1000)
const fov = int(90)
const fovIter = int(1000)
const maxView = int(200)
const displayScale = int(15)

func radToDeg(in float64) float64 {
	return in * (180.0 / math.Pi)
}

func degToRad(in float64) float64 {
	return in * (math.Pi / 180.0)
}

func disNormal(in1 [2]float64, in2 [2]float64) float64 {
	diffX := math.Abs(float64(in1[0] - in2[0]))
	diffY := math.Abs(float64(in1[1] - in2[1]))
	return float64(math.Sqrt(float64(diffX*diffX) + float64(diffY*diffY)))
}

func disCor(in1 [2]float64, in2 [2]float64, proc float64) float64 {
	return float64(disNormal(in1, in2) * math.Cos(degToRad(math.Sin(proc)*float64(fov))-math.Pi/4)) // * math.Cos(degToRad(proc*float64(fov))-math.Pi/4)
}

func disScale(dis float64) [2]float64 {
	size := 1.0 / (dis / disNormal([2]float64{0.0, 0.0}, [2]float64{float64(maxView), float64(maxView)}))
	size = size * float64(displayScale)
	return [2]float64{float64(screenHeight)/2.0 - size, float64(screenHeight)/2.0 + size}
}

func playerMove(playerIn *player, mapIn *[][]block) {
	currentRotation := playerIn.rotation
	move := false
	switch true {
	case rl.IsKeyDown(rl.KeyS):
		move = true
		currentRotation += 180
		break
	case rl.IsKeyDown(rl.KeyW):
		move = true
		break
	case rl.IsKeyDown(rl.KeyD):
		move = true
		currentRotation -= 90
		break
	case rl.IsKeyDown(rl.KeyA):
		move = true
		currentRotation += 90
		break
	case rl.IsKeyDown(rl.KeyLeft):
		playerIn.rotation -= 1
		break
	case rl.IsKeyDown(rl.KeyRight):
		playerIn.rotation += 1
		break
	default:
		move = false
		break
	}
	currentRotation = normRotationDeg(currentRotation)

	tempPositionX := 0.0
	tempPositionY := 0.0

	if move {
		tempPositionX += math.Sin(currentRotation * math.Pi / 180.0)
		tempPositionY += math.Cos(currentRotation * math.Pi / 180.0)
	}
	if (*mapIn)[int(tempPositionX)][int(tempPositionY)].collision {
		return
	}
	playerIn.position[0] = tempPositionX
	playerIn.position[1] = tempPositionY

	//check for keys --> change rotation + check for collision and move
}

func normRotationDeg(angle float64) float64 {
	for angle > 360.0 {
		angle -= 360.0
	}
	for angle < 0.0 {
		angle += 360.0
	}
	if angle == 360.0 {
		angle = 0.0
	}
	return angle
}

func ray(angle float64, position [2]float64, mapIn *[][]block) [2]float64 {

	angle = normRotationDeg(angle)

	angleRad := degToRad(float64(angle))

	var blockMap map[float64][2]float64 = make(map[float64][2]float64)

	//angleRad := float64(angle) / 3.141592653
	if angle > 0 && angle < 90 {
		//x -> inf
		//y -> inf

		//y := x*tan(angleRad)
		//x := y*tan(pi/2 - angleRad)
		for x := 0; x < maxView; x++ {
			currentYFloat := float64(x) * math.Tan(float64(angleRad))
			if currentYFloat > 0 {
				currentYFloat += 0.01
			} else if currentYFloat < 0 {
				currentYFloat -= 0.01
			}
			currentY := int(currentYFloat)
			currentY += int(position[1])
			tempX := x + int(position[0])
			if int(tempX) < len((*mapIn)[0]) &&
				int(currentY) < len(*mapIn) &&
				int(tempX) > 0 &&
				int(currentY) > 0 &&
				(*mapIn)[int(tempX)][currentY].wallTexture.isSet {
				blockMap[disNormal(position, [2]float64{float64(currentY), float64(tempX)})] = [2]float64{float64(tempX), float64(currentYFloat + position[1])}
			}
			//fmt.Print("x: ")
			//fmt.Print(tempX)
			//fmt.Print("\ty: ")
			//fmt.Print(currentY)
			//fmt.Println("")
		}

		//fmt.Println("#####################")

		for y := 0; y < maxView; y++ {
			currentXFloat := float64(y) * math.Tan((math.Pi/2.0)-float64(angleRad))
			if currentXFloat > 0 {
				currentXFloat += 0.01
			} else if currentXFloat < 0 {
				currentXFloat -= 0.01
			}
			currentX := int(currentXFloat)
			currentX += int(position[0])
			tempY := y + int(position[1])
			if int(currentX) < len((*mapIn)[0]) &&
				int(tempY) < len(*mapIn) &&
				int(currentX) > 0 &&
				int(tempY) > 0 &&
				(*mapIn)[currentX][int(tempY)].wallTexture.isSet {
				blockMap[disNormal(position, [2]float64{float64(currentX), float64(tempY)})] = [2]float64{float64(currentXFloat + position[0]), float64(tempY)}
			}
			//fmt.Print("x: ")
			//fmt.Print(currentX)
			//fmt.Print("\ty: ")
			//fmt.Print(tempY)
			//fmt.Println("")
		}
	} else if angle > 90 && angle < 180 {
		//x ->  - inf
		//y -> inf

		for x := 0; x > -maxView; x-- {
			currentYFloat := float64(x) * math.Tan(float64(angleRad))
			if currentYFloat > 0 {
				currentYFloat += 0.01
			} else if currentYFloat < 0 {
				currentYFloat -= 0.01
			}
			currentY := int(currentYFloat)
			currentY += int(position[1])
			tempX := x + int(position[0])
			if int(tempX) < len((*mapIn)[0]) &&
				int(currentY) < len(*mapIn) &&
				int(tempX) > 0 &&
				int(currentY) > 0 &&
				(*mapIn)[int(tempX)][currentY].wallTexture.isSet {
				blockMap[disNormal(position, [2]float64{float64(currentY), float64(tempX)})] = [2]float64{float64(tempX), float64(currentYFloat + position[1])}
			}
			//fmt.Print("x: ")
			//fmt.Print(tempX)
			//fmt.Print("\ty: ")
			//fmt.Print(currentY)
			//fmt.Println("")
		}

		//fmt.Println("#####################")

		for y := 0; y < maxView; y++ {
			currentXFloat := float64(y) * math.Tan((math.Pi/2.0)-float64(angleRad))
			if currentXFloat > 0 {
				currentXFloat += 0.01
			} else if currentXFloat < 0 {
				currentXFloat -= 0.01
			}
			currentX := int(currentXFloat)
			currentX += int(position[0])
			tempY := y + int(position[1])
			if int(currentX) < len((*mapIn)[0]) &&
				int(tempY) < len(*mapIn) &&
				int(currentX) > 0 &&
				int(tempY) > 0 &&
				(*mapIn)[currentX][int(tempY)].wallTexture.isSet {
				blockMap[disNormal(position, [2]float64{float64(currentX), float64(tempY)})] = [2]float64{float64(currentXFloat + position[0]), float64(tempY)}
			}
			//fmt.Print("x: ")
			//fmt.Print(currentX)
			//fmt.Print("\ty: ")
			//fmt.Print(tempY)
			//fmt.Println("")
		}

		//y := x*tan(angleRad)
		//x := y*tan(pi/2 - angleRad)
	} else if angle > 180 && angle < 270 {
		//x ->  - inf
		//y ->  - inf

		for x := 0; x > -maxView; x-- {
			currentYFloat := float64(x) * math.Tan(float64(angleRad))
			if currentYFloat > 0 {
				currentYFloat += 0.01
			} else if currentYFloat < 0 {
				currentYFloat -= 0.01
			}
			currentY := int(currentYFloat)
			currentY += int(position[1])
			tempX := x + int(position[0])
			if int(tempX) < len((*mapIn)[0]) &&
				int(currentY) < len(*mapIn) &&
				int(tempX) > 0 &&
				int(currentY) > 0 &&
				(*mapIn)[int(tempX)][currentY].wallTexture.isSet {
				blockMap[disNormal(position, [2]float64{float64(currentY), float64(tempX)})] = [2]float64{float64(tempX), float64(currentYFloat + position[1])}
			}
			//fmt.Print("x: ")
			//fmt.Print(tempX)
			//fmt.Print("\ty: ")
			//fmt.Print(currentY)
			//fmt.Println("")
		}

		//fmt.Println("#####################")

		for y := 0; y > -maxView; y-- {
			currentXFloat := float64(y) * math.Tan((math.Pi/2.0)-float64(angleRad))
			if currentXFloat > 0 {
				currentXFloat += 0.01
			} else if currentXFloat < 0 {
				currentXFloat -= 0.01
			}
			currentX := int(currentXFloat)
			currentX += int(position[0])
			tempY := y + int(position[1])
			if int(currentX) < len((*mapIn)[0]) &&
				int(tempY) < len(*mapIn) &&
				int(currentX) > 0 &&
				int(tempY) > 0 &&
				(*mapIn)[currentX][int(tempY)].wallTexture.isSet {
				blockMap[disNormal(position, [2]float64{float64(currentX), float64(tempY)})] = [2]float64{float64(currentXFloat + position[0]), float64(tempY)}
			}
			//fmt.Print("x: ")
			//fmt.Print(currentX)
			//fmt.Print("\ty: ")
			//fmt.Print(tempY)
			//fmt.Println("")
		}

		//y := x*tan(angleRad)
		//x := y*tan(pi/2 - angleRad)
	} else if angle > 270 && angle < 360 {
		//x ->  - inf
		//y ->  - inf

		for x := 0; x < maxView; x++ {
			currentYFloat := float64(x) * math.Tan(float64(angleRad))
			if currentYFloat > 0 {
				currentYFloat += 0.01
			} else if currentYFloat < 0 {
				currentYFloat -= 0.01
			}
			currentY := int(currentYFloat)
			currentY += int(position[1])
			tempX := x + int(position[0])
			if int(tempX) < len((*mapIn)[0]) &&
				int(currentY) < len(*mapIn) &&
				int(tempX) > 0 &&
				int(currentY) > 0 &&
				(*mapIn)[int(tempX)][currentY].wallTexture.isSet {
				blockMap[disNormal(position, [2]float64{float64(currentY), float64(tempX)})] = [2]float64{float64(tempX), float64(currentYFloat + position[1])}
			}
			//fmt.Print("x: ")
			//fmt.Print(tempX)
			//fmt.Print("\ty: ")
			//fmt.Print(currentY)
			//fmt.Println("")
		}

		//fmt.Println("#####################")

		for y := 0; y > -maxView; y-- {
			currentXFloat := float64(y) * math.Tan((math.Pi/2.0)-float64(angleRad))
			if currentXFloat > 0 {
				currentXFloat += 0.01
			} else if currentXFloat < 0 {
				currentXFloat -= 0.01
			}
			currentX := int(currentXFloat)
			currentX += int(position[0])
			tempY := y + int(position[1])
			if int(currentX) < len((*mapIn)[0]) &&
				int(tempY) < len(*mapIn) &&
				int(currentX) > 0 &&
				int(tempY) > 0 &&
				(*mapIn)[currentX][int(tempY)].wallTexture.isSet {
				blockMap[disNormal(position, [2]float64{float64(currentX), float64(tempY)})] = [2]float64{float64(currentXFloat + position[0]), float64(tempY)}
			}
			//fmt.Print("x: ")
			//fmt.Print(currentX)
			//fmt.Print("\ty: ")
			//fmt.Print(tempY)
			//fmt.Println("")
		}
	}

	smallest := float64(9999999.0)
	for iter := range blockMap {
		if iter < float64(smallest) {
			smallest = float64(iter)
		}
	}

	if smallest == float64(9999999.0) {
		return position
	}

	return blockMap[float64(smallest)]
}

func main() {

	gameMap := [][]block{}

	for x := 0; x < 2000; x++ {
		temp := []block{}
		for y := 0; y < 2000; y++ {
			temp = append(temp, defaultBlock())
		}
		gameMap = append(gameMap, temp)
	}

	//gameMap[42][42] = true

	//gameMap[49][49] = true
	//gameMap[50][50] = true
	//gameMap[49][50] = true
	//gameMap[50][49] = true
	for i := 0; i < 2000; i++ {
		gameMap[i][2000-1-i].wallTexture.isSet = true
		gameMap[i][i].wallTexture.isSet = true
	}

	mainPlayer := defaultPlayer()
	mainPlayer.position[0] = 970
	mainPlayer.position[1] = 960

	//for i := 0; i < 200; i++ {
	//	mainPlayer.position[1] += 0.1
	//	fmt.Print(mainPlayer.position)
	//	fmt.Print("\t")
	//	returnPos := ray(40, mainPlayer.position, &gameMap)
	//	fmt.Print(returnPos)
	//	fmt.Print("\t")
	//	diff := disNormal(mainPlayer.position, returnPos)
	//	fmt.Print(diff)
	//	fmt.Print("\n")
	//}

	fmt.Println(rl.White)
	rl.InitWindow(screenWidth, screenHeight, "omwtfyb")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)
	for !rl.WindowShouldClose() {
		//blocksSeen := []block{}

		//mainPlayer.position[1] += 1
		//mainPlayer.position[0] += 1
		playerMove(&mainPlayer, &gameMap)
		//mainPlayer.rotation += 2

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		for i := 0; i < rl.GetScreenWidth(); i++ {
			currentAngle := mainPlayer.rotation + (float64(i)/float64(rl.GetScreenWidth()))*(float64(fov))/2.0
			returnPos := ray(currentAngle, mainPlayer.position, &gameMap)
			diff := disCor(mainPlayer.position, returnPos, float64(i)/float64(rl.GetScreenWidth()))
			points := disScale(diff)
			rl.DrawLine(int32(i), int32(points[0]), int32(i), int32(points[1]), rl.White)
		}

		rl.EndDrawing()
	}
}

/*TODO
fix movement
	into function
	check for rotation
	check rotation
save seen block in on frame (in ray)
	--> floor and seeling
	block that got hit
	--> coords for entry and exit
	--> calc dis then height for them
	--> map texture for space
	--> draw
block from bool to struct
	col?
	wall - texture?
	floor - texture?
	seeling - texture?
	interaction?
texture for block
	--> draw as pixels
*/
