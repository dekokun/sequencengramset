package sequencengramset

import (
	"errors"
)

func Sequencengramset(strings string, n int) (map[string]struct{}, error) {
	if n <= 0 {
		return nil, errors.New("nは1以上で")
	}
	result := make(map[string]struct{})
	for i := 0; i < len(strings)-n+1; i += 1 {
		result[strings[i:i+n]] = struct{}{}
	}
	return result, nil

}
