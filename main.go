package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type player struct {
	rotation float64
	position [2]float64
}

func defaultPlayer() player {
	return player{position: [2]float64{0.0, 0.0}, rotation: 0}
}

const screenHeight = int32(1000)
const screenWidth = int32(1000)
const fov = int(90)
const fovIter = int(1000)
const maxView = int(20)
const displayScale = int(20)

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

func disScale(dis float64) [2]float64 {
	size := 1.0 / (dis / disNormal([2]float64{0.0, 0.0}, [2]float64{float64(maxView), float64(maxView)}))
	size = size * float64(displayScale)
	return [2]float64{float64(screenHeight)/2.0 - size, float64(screenHeight)/2.0 + size}
}

func ray(angle float64, position [2]float64, mapIn *[][]bool) [2]float64 {

	for angle > 360.0 {
		angle -= 360.0
	}
	for angle < 0.0 {
		angle += 360.0
	}
	if angle == 360.0 {
		angle = 0.0
	}

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
				(*mapIn)[int(tempX)][currentY] {
				blockMap[disNormal(position, [2]float64{float64(currentY), float64(tempX)})] = [2]float64{float64(tempX), float64(currentY)}
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
				(*mapIn)[currentX][int(tempY)] {
				blockMap[disNormal(position, [2]float64{float64(currentX), float64(tempY)})] = [2]float64{float64(currentX), float64(tempY)}
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
				(*mapIn)[int(tempX)][currentY] {
				blockMap[disNormal(position, [2]float64{float64(currentY), float64(tempX)})] = [2]float64{float64(tempX), float64(currentY)}
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
				(*mapIn)[currentX][int(tempY)] {
				blockMap[disNormal(position, [2]float64{float64(currentX), float64(tempY)})] = [2]float64{float64(currentX), float64(tempY)}
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
				(*mapIn)[int(tempX)][currentY] {
				blockMap[disNormal(position, [2]float64{float64(currentY), float64(tempX)})] = [2]float64{float64(tempX), float64(currentY)}
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
				(*mapIn)[currentX][int(tempY)] {
				blockMap[disNormal(position, [2]float64{float64(currentX), float64(tempY)})] = [2]float64{float64(currentX), float64(tempY)}
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
				(*mapIn)[int(tempX)][currentY] {
				blockMap[disNormal(position, [2]float64{float64(currentY), float64(tempX)})] = [2]float64{float64(tempX), float64(currentY)}
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
				(*mapIn)[currentX][int(tempY)] {
				blockMap[disNormal(position, [2]float64{float64(currentX), float64(tempY)})] = [2]float64{float64(currentX), float64(tempY)}
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

	gameMap := [][]bool{}

	for x := 0; x < 200; x++ {
		temp := []bool{}
		for y := 0; y < 200; y++ {
			temp = append(temp, false)
		}
		gameMap = append(gameMap, temp)
	}

	//gameMap[42][42] = true

	//gameMap[49][49] = true
	gameMap[50][50] = true
	gameMap[44][48] = true
	//gameMap[51][51] = true

	mainPlayer := defaultPlayer()
	mainPlayer.position[0] = 40
	mainPlayer.position[1] = 40

	rl.InitWindow(screenWidth, screenHeight, "omwtfyb")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		for i := 0; i < fov; i++ {
			returnPos := ray(mainPlayer.rotation+float64(i), mainPlayer.position, &gameMap)
			diff := disNormal(mainPlayer.position, returnPos)
			if diff != 0 {
				points := disScale(diff)
				rl.DrawLine(int32(i), int32(points[0]), int32(i), int32(points[1]), rl.White)
			}

		}

		rl.EndDrawing()
	}
}

/*TODO
change iter --> get new angle
--> new screen position
*/
