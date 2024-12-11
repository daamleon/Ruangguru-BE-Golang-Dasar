package main

import "fmt"

func SlurredTalk(pwords *string) {
    result := ""
    for _, c := range *pwords {
        if c == 'S' || c == 'R' || c == 'Z' {
            result += "L"
        } else if c == 's' || c == 'r' || c == 'z' {
            result += "l"
        } else {
            result += string(c)
        }
    }
    fmt.Println(result)
    *pwords = result
}

func main() {
	// bisa dicoba untuk pengujian test case
	var words string = ""
	SlurredTalk(&words)
	fmt.Println(words)
}
