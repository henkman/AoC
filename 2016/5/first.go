package main

import (
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
	p := make([]byte, 0, 8)
	h := md5.New()
	var i uint64
	for i = 0; ; i++ {
		s := fmt.Sprint(key, i)
		h.Write([]byte(s))
		m := h.Sum(nil)
		x := hex.EncodeToString(m)
		if strings.HasPrefix(x, "00000") {
			p = append(p, x[5])
			if len(p) == cap(p) {
				break
			}
		}
		h.Reset()
	}
	fmt.Println(string(p))
}
