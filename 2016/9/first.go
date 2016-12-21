package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

func decompress(out io.Writer, b []byte) {
	o := 0
	m := -1
	for {
		c := b[o]
		if c == '(' {
			m = o
		} else if c == ')' {
			var cs, co int
			fmt.Sscanf(string(b[m:o+1]), "(%dx%d)", &cs, &co)
			off := o + 1
			for i := 0; i < co; i++ {
				out.Write(b[off : off+cs])
			}
			o = o + cs
			m = -1
		} else if m == -1 {
			out.Write([]byte{c})
		}
		o++
		if o >= len(b) {
			break
		}
	}
}

func main() {
	bin := bufio.NewReader(os.Stdin)
	bs := new(bytes.Buffer)
	for {
		line, _ := bin.ReadBytes('\n')
		if len(line) == 0 {
			break
		}
		line = bytes.TrimRight(line, "\n")
		bs.Reset()
		decompress(bs, line)
		fmt.Println(bs.Len())
	}
}
