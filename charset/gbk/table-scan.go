package gbk

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var TransformUnicodeTable [][2]uint16

func ScanCp936Table() {
	// https://www.unicode.org/Public/MAPPINGS/VENDORS/MICSFT/WINDOWS/CP936.TXT
	fd, err := os.Open("./charset/gbk/CP936.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(fd)
	var start uint16 = 0x8140
	for scanner.Scan() {
		s := strings.TrimSpace(scanner.Text())
		if s == "" || s[0] == '#' {
			continue
		}
		x, y := uint16(0), uint16(0)
		if _, err := fmt.Sscanf(s, "0x%x\t0x%x", &x, &y); err != nil {
			if _, err = fmt.Sscanf(s, "0x%x\t", &x); err != nil {
				panic(err)
			}
		}
		if x >= 0x8140 && x <= 0xFE7E {
			for x > start+1 {
				var temp = [2]uint16{start, 0x0000}
				TransformUnicodeTable = append(TransformUnicodeTable, temp)
				start++
			}
			var temp = [2]uint16{x, y}
			TransformUnicodeTable = append(TransformUnicodeTable, temp)
			start++
		}
	}
}
