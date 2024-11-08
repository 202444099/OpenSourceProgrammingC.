package keybord

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func getInteger() int {
	in := bufio.NewReader(os.Stdin)
	a, err := in.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	a = strings.TrimSpace(a)
	n, _ := strconv.Atoi(a)
	if err != nil {
		log.Fatal(err)
	}
	return n
}
