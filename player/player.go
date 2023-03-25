package player

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	One string
	Two string
)

func Create(playerNo int) bool {
	if playerNo != 1 && playerNo != 2 {
		return false
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Enter player%v name: ", playerNo)
	name, err := reader.ReadString('\n')
	if err != nil {
		return false
	}

	name = strings.TrimSpace(name)

	if name == "" {
		fmt.Println("Please enter a valid name")
		return false
	}

	switch playerNo {
	case 1:
		One = name
		return true
	default:
		Two = name
		return true
	}
}
