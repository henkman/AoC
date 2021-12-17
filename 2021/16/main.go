package main

import (
	"sort"
	"io/ioutil"
	"os"
	"constraints"
	"encoding/hex"
	"fmt"
)

func main() {
	raw, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	data := make([]byte, hex.DecodedLen(len(raw)))
	_, err = hex.Decode(data, raw)
	if err != nil {
		panic(err)
	}

	pp := PacketParser{Data: data}
	p := pp.Parse()
	fmt.Println("first:", versionSum(p))
	fmt.Println("second:", p.Evaluate())
}

func versionSum(p Packet) int {
	s := int(p.Version)
	for _, sp := range p.SubPackets {
		s += versionSum(sp)
	}
	return s
}

type Packet struct {
	Version    byte
	Type       byte
	Value      uint64
	SubPackets []Packet
}

func (p *Packet) Evaluate() int {
	switch p.Type {
		case 0: {
			s := 0
			for _, sp := range p.SubPackets {
				s += sp.Evaluate()
			}
			return s
		}
		case 1: {
			s := 1
			for _, sp := range p.SubPackets {
				s *= sp.Evaluate()
			}
			return s
		}
		case 2: {
			vals := make([]int, 0, len(p.SubPackets))
			for _, sp := range p.SubPackets {
				vals = append(vals, sp.Evaluate())
			}
			sort.Ints(vals)
			return vals[0]
		}
		case 3:{
			vals := make([]int, 0, len(p.SubPackets))
			for _, sp := range p.SubPackets {
				vals = append(vals, sp.Evaluate())
			}
			sort.Ints(vals)
			return vals[len(vals)-1]
		}
		case 4:
		return int(p.Value)
		case 5: {
			first := p.SubPackets[0].Evaluate()
			second := p.SubPackets[1].Evaluate()
			if first > second {
				return 1
			} else {
				return 0
			}
		}
		case 6: {
			first := p.SubPackets[0].Evaluate()
			second := p.SubPackets[1].Evaluate()
			if first < second {
				return 1
			} else {
				return 0
			}
		}
		case 7: {
			first := p.SubPackets[0].Evaluate()
			second := p.SubPackets[1].Evaluate()
			if first == second {
				return 1
			} else {
				return 0
			}
		}
	}
	panic("unreachable")
}

type PacketParser struct {
	Data []byte
	Cur int
}

func (pp *PacketParser) Parse() Packet {
	p := Packet{}
	for {
		p =  pp.parsePacket()
		if pp.Cur >= (len(pp.Data)*8)-11 {
			break
		}
	}
	return p
}

func (pp *PacketParser) parsePacket() Packet {
	var p Packet
	p.Version = readBits[byte](pp.Data, &pp.Cur, 3)
	p.Type = readBits[byte](pp.Data, &pp.Cur, 3)
	if p.Type == 4 {
		for {
			end := !readBit(pp.Data, &pp.Cur)
			b := readBits[byte](pp.Data, &pp.Cur, 4)
			p.Value |= uint64(b)
			if end {
				break
			}
			p.Value <<= 4
		}
		return p
	}

	lengthType := readBit(pp.Data, &pp.Cur)
	if !lengthType {
		length := readBits[uint16](pp.Data, &pp.Cur, 15)
		spl := pp.Cur + int(length)
		for {
			p.SubPackets = append(p.SubPackets, pp.parsePacket())
			if pp.Cur >= spl {
				break
			}
		}
		return p
	}

	count := readBits[uint16](pp.Data, &pp.Cur, 11)
	for i := 0; i < int(count); i++ {
		p.SubPackets = append(p.SubPackets, pp.parsePacket())
	}
	return p
}

func readBits[T constraints.Unsigned] (data []byte, cur *int, count int) T {
	var d T
	for i := 0; i < count; i++ {
		if readBit(data, cur) {
			d |= 1 << (count - i - 1)
		}
	}
	return d
}

func readBit(data []byte, cur *int) bool {
	b, o := *cur / 8, *cur % 8
	m := data[b] & (1 << (7 - o))
	(*cur)++
	return m != 0
}
