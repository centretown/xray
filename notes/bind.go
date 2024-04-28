package notes

import (
	"fmt"
)

func bind[T any](bound *T, value any) (err error) {
	bind, ok := value.(*T)
	if !ok {
		err = fmt.Errorf("wrong type")
		return
	}
	*bound = *bind
	return
}
