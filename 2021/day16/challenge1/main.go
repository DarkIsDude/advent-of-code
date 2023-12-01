package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const FILE_LOCATION = "./input.txt"
const DEBUG = false

var sumVersion int = 0

func main() {
	packet := readFile()
	if DEBUG {
		fmt.Printf("Encoded packet : %s\n", packet)
	}

	decodePacket(packet)

	fmt.Println(sumVersion)
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

func decodePacket(packet string) string {
	packetVersion := packet[0:3]
	decodedVersion, err := strconv.ParseInt(packetVersion, 2, 64)
	if err != nil {
		panic(err)
	}

	sumVersion += int(decodedVersion)

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

		return remainingPacket
	} else {
		fmt.Printf("Operator packet decoding %s\n", packetWithoutVersionAndType)

		return decodeOperatorPacket(packetWithoutVersionAndType)
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

func decodeOperatorPacket(packet string) string {
	if packet[0] == '0' {
		lengthOfSubPacketsBinary := packet[1:16]
		lengthOfSubPackets, err := strconv.ParseInt(lengthOfSubPacketsBinary, 2, 64)
		if err != nil {
			panic(err)
		}

		if DEBUG {
			fmt.Printf("Found the length of subpacket is %d (%s)\n", lengthOfSubPackets, lengthOfSubPacketsBinary)
		}

		remainingPacket := packet[16+lengthOfSubPackets:]
		packetToDecode := packet[16 : 16+lengthOfSubPackets]
		for len(packetToDecode) > 0 {
			packetToDecode = decodePacket(packetToDecode)
		}

		return remainingPacket
	} else {
		numberOfSubPacketsBinary := packet[1:12]
		numberOfSubPackets, err := strconv.ParseInt(numberOfSubPacketsBinary, 2, 64)
		if err != nil {
			panic(err)
		}

		if DEBUG {
			fmt.Printf("Found %d (%s) subpacket\n", numberOfSubPackets, numberOfSubPacketsBinary)
		}

		remainingPacket := packet[12:]
		for i := 0; i < int(numberOfSubPackets); i++ {
			if DEBUG {
				fmt.Printf("Decoding the %d packet\n", i+1)
			}

			remainingPacket = decodePacket(remainingPacket)
		}

		return remainingPacket
	}
}
