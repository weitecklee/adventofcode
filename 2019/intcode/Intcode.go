package intcode

import "fmt"

type IntcodeProgram struct {
	program      map[int]int
	programIndex int
	relativeBase int
	channel      chan int
}

func NewIntcodeProgram(prog []int, intcodeChan chan int) *IntcodeProgram {
	program := make(map[int]int)
	for i, n := range prog {
		program[i] = n
	}
	ic := &IntcodeProgram{program, 0, 0, intcodeChan}
	return ic
}

func (ic *IntcodeProgram) getParams(parameterModes []int, nParams int) []int {
	parameters := make([]int, nParams)
	for j := 0; j < nParams; j++ {
		if j >= len(parameterModes) || parameterModes[j] == 0 {
			// position mode
			parameters[j] = ic.program[ic.programIndex+j+1]
		} else if parameterModes[j] == 1 {
			// immediate mode
			parameters[j] = ic.programIndex + j + 1
		} else if parameterModes[j] == 2 {
			// relative move
			parameters[j] = ic.program[ic.programIndex+j+1] + ic.relativeBase
		}
	}
	return parameters

}

func (ic *IntcodeProgram) Run() {
	for ic.programIndex >= 0 {
		opcode := ic.program[ic.programIndex] % 100
		parameterModeNumber := ic.program[ic.programIndex] / 100
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
			tmp := ic.getParams(parameterModes, 2)
			params = make([]int, len(tmp))
			for i, n := range tmp {
				params[i] = ic.program[n]
			}
		case 1:
			fallthrough
		case 2:
			fallthrough
		case 7:
			fallthrough
		case 8:
			params = ic.getParams(parameterModes, 3)
		case 3:
			fallthrough
		case 4:
			fallthrough
		case 9:
			params = ic.getParams(parameterModes, 1)
		}

		switch opcode {
		case 1:
			// add
			ic.program[params[2]] = ic.program[params[0]] + ic.program[params[1]]
			ic.programIndex += 3
		case 2:
			// multiply
			ic.program[params[2]] = ic.program[params[0]] * ic.program[params[1]]
			ic.programIndex += 3
		case 3:
			// save input
			ic.program[params[0]] = <-ic.channel
			ic.programIndex++
		case 4:
			// output
			ic.channel <- ic.program[params[0]]
			ic.programIndex++
		case 5:
			// jump if true
			if params[0] != 0 {
				ic.programIndex = params[1]
				continue
			} else {
				ic.programIndex += 2
			}
		case 6:
			// jump if false
			if params[0] == 0 {
				ic.programIndex = params[1]
				continue
			} else {
				ic.programIndex += 2
			}
		case 7:
			// less than
			if ic.program[params[0]] < ic.program[params[1]] {
				ic.program[params[2]] = 1
			} else {
				ic.program[params[2]] = 0
			}
			ic.programIndex += 3
		case 8:
			// equal
			if ic.program[params[0]] == ic.program[params[1]] {
				ic.program[params[2]] = 1
			} else {
				ic.program[params[2]] = 0
			}
			ic.programIndex += 3
		case 9:
			// adjust relative base
			ic.relativeBase += ic.program[params[0]]
			ic.programIndex++
		case 99:
			// halt
			ic.programIndex = -99
		default:
			panic(fmt.Sprintf("Unknown opcode: %d", ic.program[ic.programIndex]))
		}
		ic.programIndex++
	}

	ic.channel <- ic.program[0]
}
