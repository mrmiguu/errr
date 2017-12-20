package errr

import (
	"errors"
	"fmt"
	"log"
)

type Orrr struct {
	ptr *error
}

func New(err ...*error) Orrr {
	if L := len(err); L > 1 {
		log.Fatal("too many errr arguments")
	} else if L == 1 {
		return Orrr{err[0]}
	}
	return Orrr{}
}

func Is(err ...*error) bool {
	return New(err...).Err(*err[0])
}

func (o Orrr) Err(or interface{}) bool {
	switch err := or.(type) {
	case error:
		if o.ptr == nil {
			panic(err)
		} else {
			*o.ptr = err
			return err != nil
		}
	default:
		if s := fmt.Sprint(err); o.ptr == nil {
			panic(s)
		} else {
			*o.ptr = errors.New(s)
			return len(s) > 0
		}
	}
}
