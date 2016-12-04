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
	h := md5.New()
	for i := 0; ; i++ {
		s := fmt.Sprint(key, i)
		h.Write([]byte(s))
		m := h.Sum(nil)
		if strings.HasPrefix(hex.EncodeToString(m), "00000") {
			fmt.Println(i)
			break
		}
		h.Reset()
	}
}
