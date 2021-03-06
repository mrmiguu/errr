package errr

import (
	"errors"
	"fmt"
	"log"
)

// Is is errr's suffix, acting as an ambiguous conditional and/or setter.
func Is(e []*error, v ...interface{}) bool {
	if L := len(v); L > 1 {
		log.Fatal("too many errr arguments")
	} else if L == 1 {
		return new(e).err(v[0])
	}
	return new(e).err(nil)
}

type orr struct {
	ptr *error
}

func new(e []*error) orr {
	if L := len(e); L > 1 {
		log.Fatal("too many errr arguments")
	} else if L == 1 {
		return orr{e[0]}
	}
	return orr{}
}

func (o orr) err(v interface{}) bool {
	if v == nil {
		return o.ptr != nil && *o.ptr != nil
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
