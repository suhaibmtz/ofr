package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	to_f := ""
	lower := false
	ShowLine := false
	ShowErr := false
	TotalOnly := false
	for i, a := range os.Args {
		switch a {
		case "-i":
			lower = true
		case "-v", "--version":
			fmt.Println("1.4")
			return
		case "-n":
			ShowLine = true
		case "-e":
			ShowErr = true
		case "-t":
			TotalOnly = true
		default:
			if i != 0 {
				to_f += " " + a
			}
		}
	}
	to_f = strings.TrimSpace(to_f)
	if lower {
		to_f = strings.ToLower(to_f)
	}
	Total := 0
	for {
		var s string
		fmt.Scan(&s)
		if s == "" {
			break
		}
		f, e := os.Open(s)
		if e == nil {
			to_P := "--------------------\n"
			to_P += s + "\n"
			v, err := io.ReadAll(f)
			if err != nil && ShowErr {
				fmt.Println(err.Error())
			}
			if err == nil {
				content := string(v)
				Pr := false
				i := 0
				l := ""
				for i, l = range strings.Split(content, "\n") {
					if !TotalOnly {
						lo := l
						if lower {
							lo = strings.ToLower(l)
						}
						if strings.Contains(lo, to_f) || to_f == "" {
							l = strings.ReplaceAll(l, to_f, "\u001b[31m"+to_f+"\u001b[0m")
							Pr = true
							if ShowLine {
								to_P = fmt.Sprint(i+1) + " "
							}
							to_P += l + "\n"
						}
					}
				}
				if TotalOnly {
					fmt.Println(i, ":", s)
				}
				Total += i
				if Pr {
					fmt.Print(to_P)
				}
			}
		}
	}
	fmt.Println(Total)
}
