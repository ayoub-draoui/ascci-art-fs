package ascciartfs

import "fmt"

func FindAndPrint(checkcharacter []string, readfile map[rune][]string) {
	is_Printed := false
	for idx, word := range checkcharacter {
		if word != "" {
			i := 0
			for i < 8 {
				line := ""
				for _, char := range word {
					line = readfile[char][i]
					fmt.Print(line)
				}
				fmt.Println()
				i++
				is_Printed = true
			}
		}else{
			if idx == len(checkcharacter)-1 && !is_Printed{
				continue
			} else{
				fmt.Println()
			}

		}
	}
}
