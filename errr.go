package errr

import (
	"errors"
	"fmt"
	"log"
)

// Orr is errr's suffix, acting as an ambiguous conditional and/or setter.
func Orr(err []*error, v ...interface{}) bool {
	return new(err).err(v)
}

type orr struct {
	ptr *error
}

func new(err []*error) orr {
	if L := len(err); L > 1 {
		log.Fatal("too many errr arguments")
	} else if L == 1 {
		return orr{err[0]}
	}
	return orr{}
}

func (o orr) err(v interface{}) bool {
	if v == nil {
		return *o.ptr != nil
	}
	switch err := v.(type) {
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
