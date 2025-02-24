package intcode

import "fmt"

func getParams(program map[int]int, parameterModes []int, nParams, programIndex, relativeBase int) []int {
	parameters := make([]int, nParams)
	for j := 0; j < nParams; j++ {
		if j >= len(parameterModes) || parameterModes[j] == 0 {
			// position mode
			parameters[j] = program[programIndex+j+1]
		} else if parameterModes[j] == 1 {
			// immediate mode
			parameters[j] = programIndex + j + 1
		} else if parameterModes[j] == 2 {
			// relative move
			parameters[j] = program[programIndex+j+1] + relativeBase
		}
	}
	return parameters
}

func IntcodeProgram(prog []int, intcodeChan chan int) {
	program := make(map[int]int)
	for i, n := range prog {
		program[i] = n
	}
	programIndex := 0
	relativeBase := 0

	for programIndex >= 0 {
		opcode := program[programIndex] % 100
		parameterModeNumber := program[programIndex] / 100
		parameterModes := make([]int, 0, 4)
		for parameterModeNumber > 0 {
			parameterModes = append(parameterModes, parameterModeNumber%10)
			parameterModeNumber /= 10
		}

		var params []int
		switch opcode {
		case 5:
			fallthrough
		case 6:
			tmp := getParams(program, parameterModes, 2, programIndex, relativeBase)
			params = make([]int, len(tmp))
			for i, n := range tmp {
				params[i] = program[n]
			}
		case 1:
			fallthrough
		case 2:
			fallthrough
		case 7:
			fallthrough
		case 8:
			params = getParams(program, parameterModes, 3, programIndex, relativeBase)
		case 3:
			fallthrough
		case 4:
			fallthrough
		case 9:
			params = getParams(program, parameterModes, 1, programIndex, relativeBase)
		}

		switch opcode {
		case 1:
			// add
			program[params[2]] = program[params[0]] + program[params[1]]
			programIndex += 3
		case 2:
			// multiply
			program[params[2]] = program[params[0]] * program[params[1]]
			programIndex += 3
		case 3:
			// save input
			program[params[0]] = <-intcodeChan
			programIndex++
		case 4:
			// output
			intcodeChan <- program[params[0]]
			programIndex++
		case 5:
			// jump if true
			if params[0] != 0 {
				programIndex = params[1]
				continue
			} else {
				programIndex += 2
			}
		case 6:
			// jump if false
			if params[0] == 0 {
				programIndex = params[1]
				continue
			} else {
				programIndex += 2
			}
		case 7:
			// less than
			if program[params[0]] < program[params[1]] {
				program[params[2]] = 1
			} else {
				program[params[2]] = 0
			}
			programIndex += 3
		case 8:
			// equal
			if program[params[0]] == program[params[1]] {
				program[params[2]] = 1
			} else {
				program[params[2]] = 0
			}
			programIndex += 3
		case 9:
			// adjust relative base
			relativeBase += program[params[0]]
			programIndex++
		case 99:
			// halt
			programIndex = -99
		default:
			panic(fmt.Sprintf("Unknown opcode: %d", program[programIndex]))
		}
		programIndex++
	}

	intcodeChan <- program[0]
}
