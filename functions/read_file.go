package linear

import (
	"log"
	"os"
	"strings"
)

func ReadFile(s string) []string {
	content, err := os.ReadFile(s)
	if err != nil {
		log.Fatal(err)
	}
	data := strings.Fields(string(content))
	return data
}
