package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	var ps []Passport
	{
		reKeyValue := regexp.MustCompile(`(\S+):(\S+)`)
		bin := bufio.NewReader(os.Stdin)
		for {
			line, err := bin.ReadString('\n')
			line = strings.TrimSpace(line)
			if len(line) > 0 {
				ms := reKeyValue.FindAllStringSubmatch(line, -1)
				p := Passport{}
				for _, m := range ms {
					p[m[1]] = m[2]
				}
				ps = append(ps, p)
			}
			if err != nil {
				if err == io.EOF {
					break
				}
				panic(err)
			}
		}
	}

	{
		first := 0
	next_simple:
		for _, p := range ps {
			for _, k := range []string{
				"byr", "iyr", "eyr", "hgt",
				"hcl", "ecl", "pid"} {
				if _, ok := p[k]; !ok {
					continue next_simple
				}
			}
			first++
		}
		fmt.Println("first:", first)
	}

	{
		reHeight := regexp.MustCompile(`^(\d+)(cm|in)$`)
		reHairColor := regexp.MustCompile(`^#[0-9a-f]{6}$`)
		rePassportId := regexp.MustCompile(`^[0-9]{9}$`)
		second := 0
	next_validation:
		for _, p := range ps {
			if s, ok := p["byr"]; ok {
				n, err := strconv.Atoi(s)
				if err != nil || n < 1920 || n > 2002 {
					continue next_validation
				}
			} else {
				continue next_validation
			}

			if s, ok := p["iyr"]; ok {
				n, err := strconv.Atoi(s)
				if err != nil || n < 2010 || n > 2020 {
					continue next_validation
				}
			} else {
				continue next_validation
			}

			if s, ok := p["eyr"]; ok {
				n, err := strconv.Atoi(s)
				if err != nil || n < 2020 || n > 2030 {
					continue next_validation
				}
			} else {
				continue next_validation
			}

			if s, ok := p["hgt"]; ok {
				mh := reHeight.FindStringSubmatch(s)
				if mh == nil {
					continue next_validation
				}
				n, err := strconv.Atoi(mh[1])
				if err != nil {
					continue next_validation
				}
				if mh[2] == "cm" {
					if n < 150 || n > 193 {
						continue next_validation
					}
				} else {
					if n < 59 || n > 76 {
						continue next_validation
					}
				}
			} else {
				continue next_validation
			}

			if s, ok := p["hcl"]; !ok || !reHairColor.MatchString(s) {
				continue next_validation
			}

			if s, ok := p["ecl"]; ok {
				valid := false
				for _, ec := range []string{
					"amb", "blu", "brn",
					"gry", "grn", "hzl", "oth"} {
					if s == ec {
						valid = true
						break
					}
				}
				if !valid {
					continue next_validation
				}
			} else {
				continue next_validation
			}

			if s, ok := p["pid"]; !ok || !rePassportId.MatchString(s) {
				continue next_validation
			}

			second++
		}
		fmt.Println("second:", second)
	}
}

type Passport map[string]string
