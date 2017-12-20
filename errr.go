package errr

import "log"

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

func (o Orrr) Err(or error) bool {
	if o.ptr == nil {
		panic(or)
	} else {
		*o.ptr = or
		return or != nil
	}
}
