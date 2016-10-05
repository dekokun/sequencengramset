package sequencengramset

import (
	"errors"
)

func Sequencengramset(strings string, n int) (map[string]struct{}, error) {
	if n <= 0 {
		return nil, errors.New("nは1以上で")
	}
	stringRune := []rune(strings)
	result := make(map[string]struct{})
	for i := 0; i < len(stringRune)-n+1; i += 1 {
		result[string(stringRune[i:i+n])] = struct{}{}
	}
	return result, nil

}
