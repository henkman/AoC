package main

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	var key string
	_, err := fmt.Fscanf(os.Stdin, "%s\n", &key)
	if err != nil {
		if err == io.EOF {
			return
		}
		log.Fatal(err)
	}
	pass := make([]byte, 8)
	h := md5.New()
	var i uint64
	for i = 0; ; i++ {
		s := fmt.Sprint(key, i)
		h.Write([]byte(s))
		m := h.Sum(nil)
		x := hex.EncodeToString(m)
		if strings.HasPrefix(x, "00000") {
			p := x[5] - '0'
			if p >= 0 && p <= 7 && pass[p] == 0 {
				pass[p] = x[6]
				if !bytes.ContainsRune(pass, rune(0)) {
					break
				}
			}
		}
		h.Reset()
	}
	fmt.Println(string(pass))
}
