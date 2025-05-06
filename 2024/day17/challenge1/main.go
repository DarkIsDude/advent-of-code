package main

import (
	"fmt"
	"math"
	"strings"
)

type Program struct {
	RegisterA int
	RegisterB int
	RegisterC int

	Output []int

	InstructionPointer int
	Instructions       []int
}

// var currentProgram = Program{
// RegisterA: 729,
// RegisterB: 0,
// RegisterC: 0,

// InstructionPointer: 0,
// Instructions: []int{
// 	0, 1, 5, 4, 3, 0,
// },
// }

var currentProgram = Program{
	RegisterA: 66171486,
	RegisterB: 0,
	RegisterC: 0,

	InstructionPointer: 0,
	Instructions: []int{
		2, 4, 1, 6, 7, 5, 4, 6, 1, 4, 5, 5, 0, 3, 3, 0,
	},
}

func main() {
	for true {
		if currentProgram.InstructionPointer >= len(currentProgram.Instructions)-1 {
			fmt.Println("End of program")
			break
		}

		currentProgram.ExecuteInstruction()
	}

	fmt.Println("Final state of registers:")
	fmt.Println("Register A:", currentProgram.RegisterA)
	fmt.Println("Register B:", currentProgram.RegisterB)
	fmt.Println("Register C:", currentProgram.RegisterC)

	outputAsString := []string{}
	for _, output := range currentProgram.Output {
		outputAsString = append(outputAsString, fmt.Sprintf("%d", output))
	}

	fmt.Println("Output:", strings.Join(outputAsString, ","))
}

func (p *Program) ExecuteInstruction() {
	shouldIncrement := true
	opcode := p.Instructions[p.InstructionPointer]
	literalOperand := p.Instructions[p.InstructionPointer+1]
	comboOperand := p.ReadComboOperand()

	fmt.Println("Current instruction pointer:", p.InstructionPointer, "opcode:", opcode, "literal operand:", literalOperand, "combo operand:", comboOperand)

	switch opcode {
	case 0:
		p.RegisterA = p.RegisterA / int(math.Pow(2, float64(comboOperand)))
	case 1:
		p.RegisterB = p.RegisterB ^ literalOperand
	case 2:
		p.RegisterB = comboOperand % 8
	case 3:
		if p.RegisterA != 0 {
			p.InstructionPointer = literalOperand
			shouldIncrement = false
		}
	case 4:
		p.RegisterB = p.RegisterB ^ p.RegisterC
	case 5:
		computed := comboOperand % 8
		fmt.Println("Output:", computed)
		p.Output = append(p.Output, computed)
	case 6:
		p.RegisterB = p.RegisterA / int(math.Pow(2, float64(comboOperand)))
	case 7:
		p.RegisterC = p.RegisterA / int(math.Pow(2, float64(comboOperand)))
	default:
		panic("Unknown opcode")
	}

	if shouldIncrement {
		p.InstructionPointer += 2
	}
}

func (p *Program) ReadComboOperand() int {
	literalOperand := p.Instructions[p.InstructionPointer+1]

	switch literalOperand {
	case 0:
	case 1:
	case 2:
	case 3:
		return literalOperand
	case 4:
		return p.RegisterA
	case 5:
		return p.RegisterB
	case 6:
		return p.RegisterC
	case 7:
		panic("Reserved operand")
	default:
		panic("Unknown operand")
	}

	return literalOperand
}
