package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const FILE_LOCATION = "./input.txt"
const DEBUG = false

func main() {
	packet := readFile()
	if DEBUG {
		fmt.Printf("Encoded packet : %s\n", packet)
	}

	_, result := decodePacket(packet)

	fmt.Println(result)
}

func readFile() string {
	mapped := map[rune]string{
		'0': "0000",
		'1': "0001",
		'2': "0010",
		'3': "0011",
		'4': "0100",
		'5': "0101",
		'6': "0110",
		'7': "0111",
		'8': "1000",
		'9': "1001",
		'A': "1010",
		'B': "1011",
		'C': "1100",
		'D': "1101",
		'E': "1110",
		'F': "1111",
	}

	packet := ""

	file, err := os.Open(FILE_LOCATION)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	text := scanner.Text()

	for _, c := range text {
		packet += mapped[c]
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return packet
}

func decodePacket(packet string) (string, int) {
	packetVersion := packet[0:3]
	decodedVersion, err := strconv.ParseInt(packetVersion, 2, 64)
	if err != nil {
		panic(err)
	}

	packetType := packet[3:6]
	decodedType, err := strconv.ParseInt(packetType, 2, 64)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Packet version %s (%d) and type %s (%d)\n", packetVersion, decodedVersion, packetType, decodedType)

	packetWithoutVersionAndType := packet[6:]

	if decodedType == 4 {
		fmt.Printf("Literal packet decoding %s\n", packetWithoutVersionAndType)
		number, remainingPacket := decodeLiteralPacket(packetWithoutVersionAndType)

		fmt.Printf("Decoded number %s : %d\n", packetWithoutVersionAndType, number)

		return remainingPacket, int(number)
	} else {
		fmt.Printf("Operator packet decoding %s\n", packetWithoutVersionAndType)

		return decodeOperatorPacket(packetWithoutVersionAndType, int(decodedType))
	}
}

func decodeLiteralPacket(packet string) (int64, string) {
	remainingPacket := packet
	numberBinary := ""
	ended := false

	for !ended {
		numberBinary += remainingPacket[1:5]

		if DEBUG {
			fmt.Printf("Found the number %s\n", numberBinary)
		}

		ended = remainingPacket[0] == '0'
		remainingPacket = remainingPacket[5:]
	}

	value, err := strconv.ParseInt(numberBinary, 2, 64)
	if err != nil {
		panic(err)
	}

	return value, remainingPacket
}

func decodeOperatorPacket(packet string, operator int) (string, int) {
	var numberToProcess []int
	remainingPacket := ""

	if packet[0] == '0' {
		lengthOfSubPacketsBinary := packet[1:16]
		lengthOfSubPackets, err := strconv.ParseInt(lengthOfSubPacketsBinary, 2, 64)
		if err != nil {
			panic(err)
		}

		if DEBUG {
			fmt.Printf("Found the length of subpacket is %d (%s)\n", lengthOfSubPackets, lengthOfSubPacketsBinary)
		}

		remainingPacket = packet[16+lengthOfSubPackets:]
		packetToDecode := packet[16 : 16+lengthOfSubPackets]
		for len(packetToDecode) > 0 {
			var number int
			packetToDecode, number = decodePacket(packetToDecode)
			numberToProcess = append(numberToProcess, number)
		}
	} else {
		numberOfSubPacketsBinary := packet[1:12]
		numberOfSubPackets, err := strconv.ParseInt(numberOfSubPacketsBinary, 2, 64)
		if err != nil {
			panic(err)
		}

		if DEBUG {
			fmt.Printf("Found %d (%s) subpacket\n", numberOfSubPackets, numberOfSubPacketsBinary)
		}

		remainingPacket = packet[12:]
		for i := 0; i < int(numberOfSubPackets); i++ {
			if DEBUG {
				fmt.Printf("Decoding the %d packet\n", i+1)
			}

			var number int
			remainingPacket, number = decodePacket(remainingPacket)
			numberToProcess = append(numberToProcess, number)
		}
	}

	if DEBUG {
		fmt.Printf("Number to proces for %d : %v\n", operator, numberToProcess)
	}

	var result int
	if operator == 0 {
		if DEBUG {
			fmt.Println("SUM")
		}

		result = 0

		for _, number := range numberToProcess {
			result += number
		}
	} else if operator == 1 {
		if DEBUG {
			fmt.Println("PRODUCT")
		}

		result = 1

		for _, number := range numberToProcess {
			result *= number
		}
	} else if operator == 2 {
		if DEBUG {
			fmt.Println("MIN")
		}

		result = numberToProcess[0]

		for _, number := range numberToProcess {
			if number < result {
				result = number
			}
		}
	} else if operator == 3 {
		if DEBUG {
			fmt.Println("MAX")
		}

		result = numberToProcess[0]

		for _, number := range numberToProcess {
			if number > result {
				result = number
			}
		}
	} else if operator == 5 {
		if DEBUG {
			fmt.Println("GREATER")
		}

		if numberToProcess[0] > numberToProcess[1] {
			result = 1
		} else {
			result = 0
		}
	} else if operator == 6 {
		if DEBUG {
			fmt.Println("LESS")
		}

		if numberToProcess[0] < numberToProcess[1] {
			result = 1
		} else {
			result = 0
		}
	} else if operator == 7 {
		if DEBUG {
			fmt.Println("EQUAL")
		}

		if numberToProcess[0] == numberToProcess[1] {
			result = 1
		} else {
			result = 0
		}
	} else {
		panic("Unsupported operation")
	}

	if DEBUG {
		fmt.Printf("The result is %d\n", result)
	}

	return remainingPacket, result
}
