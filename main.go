package main

import (
	"fmt"
	"math"
)

type player struct {
	rotation int
	position [2]float32
}

func defaultGame() player {
	return player{position: [2]float32{0.0, 0.0}, rotation: 0}
}

const screenHeight = int32(1000)
const screenWidth = int32(1000)

func radToDeg(in float32) float32 {
	return in * (180.0 / math.Pi)
}

func degToRad(in float32) float32 {
	return in * (math.Pi / 180.0)
}

func ray(angle int, position [2]float32, mapIn *[][]bool) [2]float32 {

	angle = angle % 360
	for angle < 0 {
		angle += 360
	}
	if angle == 360 {
		angle = 0
	}

	angleRad := degToRad(float32(angle))

	var blockMap map[float32][2]float32 = make(map[float32][2]float32)

	//angleRad := float32(angle) / 3.141592653
	if angle > 0 && angle < 90 {
		//x -> inf
		//y -> inf

		//y := x*tan(angleRad)
		//x := y*tan(pi/2 - angleRad)
		for x := 0; x < 20.0; x++ {
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
				blockMap[float32(math.Sqrt(math.Pow(float64(tempX), 2)+math.Pow(float64(currentY), 2)))] = [2]float32{float32(tempX), float32(currentY)}
			}
			fmt.Print("x: ")
			fmt.Print(tempX)
			fmt.Print("\ty: ")
			fmt.Print(currentY)
			fmt.Println("")
		}

		fmt.Println("#####################")

		for y := 0; y < 20.0; y++ {
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
				blockMap[float32(math.Sqrt(math.Pow(float64(currentX), 2)+math.Pow(float64(tempY), 2)))] = [2]float32{float32(currentX), float32(tempY)}
			}
			fmt.Print("x: ")
			fmt.Print(currentX)
			fmt.Print("\ty: ")
			fmt.Print(tempY)
			fmt.Println("")
		}
	} else if angle > 90 && angle < 180 {
		//x ->  - inf
		//y -> inf

		for x := 0; x > -20.0; x-- {
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
				blockMap[float32(math.Sqrt(math.Pow(float64(tempX), 2)+math.Pow(float64(currentY), 2)))] = [2]float32{float32(tempX), float32(currentY)}
			}
			fmt.Print("x: ")
			fmt.Print(tempX)
			fmt.Print("\ty: ")
			fmt.Print(currentY)
			fmt.Println("")
		}

		fmt.Println("#####################")

		for y := 0; y < 20.0; y++ {
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
				blockMap[float32(math.Sqrt(math.Pow(float64(currentX), 2)+math.Pow(float64(tempY), 2)))] = [2]float32{float32(currentX), float32(tempY)}
			}
			fmt.Print("x: ")
			fmt.Print(currentX)
			fmt.Print("\ty: ")
			fmt.Print(tempY)
			fmt.Println("")
		}

		//y := x*tan(angleRad)
		//x := y*tan(pi/2 - angleRad)
	} else if angle > 180 && angle < 270 {
		//x ->  - inf
		//y ->  - inf

		for x := 0; x > -20.0; x-- {
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
				blockMap[float32(math.Sqrt(math.Pow(float64(tempX), 2)+math.Pow(float64(currentY), 2)))] = [2]float32{float32(tempX), float32(currentY)}
			}
			fmt.Print("x: ")
			fmt.Print(tempX)
			fmt.Print("\ty: ")
			fmt.Print(currentY)
			fmt.Println("")
		}

		fmt.Println("#####################")

		for y := 0; y > -20.0; y-- {
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
				blockMap[float32(math.Sqrt(math.Pow(float64(currentX), 2)+math.Pow(float64(tempY), 2)))] = [2]float32{float32(currentX), float32(tempY)}
			}
			fmt.Print("x: ")
			fmt.Print(currentX)
			fmt.Print("\ty: ")
			fmt.Print(tempY)
			fmt.Println("")
		}

		//y := x*tan(angleRad)
		//x := y*tan(pi/2 - angleRad)
	} else if angle > 270 && angle < 360 {
		//x ->  - inf
		//y ->  - inf

		for x := 0; x < 20.0; x++ {
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
				blockMap[float32(math.Sqrt(math.Pow(float64(tempX), 2)+math.Pow(float64(currentY), 2)))] = [2]float32{float32(tempX), float32(currentY)}
			}
			fmt.Print("x: ")
			fmt.Print(tempX)
			fmt.Print("\ty: ")
			fmt.Print(currentY)
			fmt.Println("")
		}

		fmt.Println("#####################")

		for y := 0; y > -20.0; y-- {
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
				blockMap[float32(math.Sqrt(math.Pow(float64(currentX), 2)+math.Pow(float64(tempY), 2)))] = [2]float32{float32(currentX), float32(tempY)}
			}
			fmt.Print("x: ")
			fmt.Print(currentX)
			fmt.Print("\ty: ")
			fmt.Print(tempY)
			fmt.Println("")
		}
	}

	smallest := float32(9999999.0)
	for iter := range blockMap {
		if iter < float32(smallest) {
			smallest = float32(iter)
		}
	}

	if smallest == float32(9999999.0) {
		return position
	}

	return blockMap[float32(smallest)]
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

	gameMap[50][50] = true

	fmt.Println(ray(225, [2]float32{60.0, 60.0}, &gameMap))

	//rl.InitWindow(screenWidth, screenHeight, "omwtfyb")
	//defer rl.CloseWindow()
	//rl.SetTargetFPS(60)
	//for !rl.WindowShouldClose() {
	//	rl.BeginDrawing()
	//	rl.EndDrawing()
	//}
}
