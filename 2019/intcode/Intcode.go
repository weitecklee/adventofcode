package intcode

import (
	"fmt"
	"math"
)

type IntcodeProgram struct {
	Program      map[int]int
	programIndex int
	relativeBase int
	ch           chan int
}

const (
	REQUESTSIGNAL = math.MaxInt
	ENDSIGNAL     = math.MinInt
	STOPSIGNAL    = math.MinInt
)

func NewIntcodeProgram(prog []int, ch chan int) *IntcodeProgram {
	program := make(map[int]int)
	for i, n := range prog {
		program[i] = n
	}
	ic := &IntcodeProgram{program, 0, 0, ch}
	return ic
}

func (ic *IntcodeProgram) getParams(parameterModes []int, nParams int) []int {
	parameters := make([]int, nParams)
	for j := range nParams {
		if j >= len(parameterModes) || parameterModes[j] == 0 {
			// position mode
			parameters[j] = ic.Program[ic.programIndex+j+1]
		} else if parameterModes[j] == 1 {
			// immediate mode
			parameters[j] = ic.programIndex + j + 1
		} else if parameterModes[j] == 2 {
			// relative move
			parameters[j] = ic.Program[ic.programIndex+j+1] + ic.relativeBase
		}
	}
	return parameters

}

func (ic *IntcodeProgram) Run() {
	defer close(ic.ch)
	for ic.programIndex >= 0 {
		opcode := ic.Program[ic.programIndex] % 100
		parameterModeNumber := ic.Program[ic.programIndex] / 100
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
				params[i] = ic.Program[n]
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
			ic.Program[params[2]] = ic.Program[params[0]] + ic.Program[params[1]]
			ic.programIndex += 3
		case 2:
			// multiply
			ic.Program[params[2]] = ic.Program[params[0]] * ic.Program[params[1]]
			ic.programIndex += 3
		case 3:
			// save input
			ic.ch <- REQUESTSIGNAL
			received := <-ic.ch
			if received == STOPSIGNAL {
				return
			}
			ic.Program[params[0]] = received
			ic.programIndex++
		case 4:
			// output
			ic.ch <- ic.Program[params[0]]
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
			if ic.Program[params[0]] < ic.Program[params[1]] {
				ic.Program[params[2]] = 1
			} else {
				ic.Program[params[2]] = 0
			}
			ic.programIndex += 3
		case 8:
			// equal
			if ic.Program[params[0]] == ic.Program[params[1]] {
				ic.Program[params[2]] = 1
			} else {
				ic.Program[params[2]] = 0
			}
			ic.programIndex += 3
		case 9:
			// adjust relative base
			ic.relativeBase += ic.Program[params[0]]
			ic.programIndex++
		case 99:
			// halt
			ic.ch <- ENDSIGNAL
			return
		default:
			panic(fmt.Sprintf("Unknown opcode: %d", ic.Program[ic.programIndex]))
		}
		ic.programIndex++
	}

}
