package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func decryptRoom(cipher []byte, sectorId int) []byte {
	alpha := []byte("abcdefghijklmnopqrstuvwxyz")
	room := make([]byte, len(cipher))
	for i, c := range cipher {
		if c == '-' {
			room[i] = ' '
			continue
		}
		room[i] = alpha[(int(c)-'a'+sectorId)%len(alpha)]
	}
	return room
}

func main() {
	bin := bufio.NewReader(os.Stdin)
	for {
		line, err := bin.ReadString('\n')
		if line == "" {
			break
		}
		m := reLine.FindStringSubmatch(line)
		cipher := []byte(m[1])
		if !isValidRoom(cipher, []byte(m[3])) {
			continue
		}
		id, err := strconv.Atoi(m[2])
		if err != nil {
			log.Fatal(err)
		}
		room := decryptRoom(cipher, id)
		if bytes.Equal(room, []byte("northpole object storage")) {
			fmt.Println(id)
			return
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
	}
}
