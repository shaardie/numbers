package numbers

import (
	"regexp"

	"github.com/shaardie/set"
)

var NatualNumbers = set.CreateFromFunc(func(x interface{}) (bool, error) {

	str, ok := x.(string)
	if !ok {
		return false, nil
	}
	return regexp.MatchString("0|[1-9][0-9]*", str)
})
