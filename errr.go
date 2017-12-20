package errr

import "log"

type Errr struct {
	ptr *error
}

func New(opt ...*error) Errr {
	if L := len(opt); L > 1 {
		log.Fatal("too many errr arguments")
	} else if L == 1 {
		return Errr{opt[0]}
	}
	return Errr{}
}

func (errr Errr) Set(err error) {
	if errr.ptr != nil {
		*errr.ptr = err
	} else {
		panic(err)
	}
}
