package iteration

import "errors"

func Repeat(r string, num int) (string, error) {
	if num < 0 {
		return "", errors.New("cannot repeat negative times")
	}
	s := ""
	for i := 0; i < num; i++ {
		s += r
	}
	return s, nil
	// strings library already does this
	// strings.Repeat(r, num)

}
