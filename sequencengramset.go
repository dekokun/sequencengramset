package sequencengramset

import (
	"errors"
)

func Sequencengramset(strings string, n int) ([]string, error) {
	if n <= 0 {
		return nil, errors.New("nは1以上で")
	}
	stringRune := []rune(strings)
	resultMap := make(map[string]struct{})
	for i := 0; i < len(stringRune)-n+1; i += 1 {
		resultMap[string(stringRune[i:i+n])] = struct{}{}
	}
	keys := make([]string, 0, len(resultMap))
	for k := range resultMap {
		keys = append(keys, k)
	}
	return keys, nil

}
