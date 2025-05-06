package conversion

import "strconv"

func StringsToFloat64(input []string) ([]float64, error) {
	var result []float64
	for _, str := range input {
		num, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return nil, err
		}
		result = append(result, num)
	}
	return result, nil
}
