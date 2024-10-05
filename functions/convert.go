package linear

import (
	"strconv"
)

func ConvertData(data []string) []float64 {
	result := []float64{}
	for i := 0; i < len(data); i++ {
		number, err := strconv.ParseFloat(data[i], 64)
		if err != nil {
			continue
		}
		result = append(result, number)
	}
	return result
}
