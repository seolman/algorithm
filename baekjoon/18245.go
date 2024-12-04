package baekjoon

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Solution() {
    reader := bufio.NewReader(os.Stdin)
    result := []string{}

    for i := 2; ;i++ {
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(input)
        if input == "Was it a cat I saw?" {
            break
        }
        result = append(result, Decode(input, i))
    }

    for _, line := range result {
        fmt.Println(line)
    }

}

func Decode(s string, c int) string {
    var decoded []byte
    for i := 0; i < len(s); i += c {
        decoded = append(decoded, s[i])
    }
    return string(decoded)
}
