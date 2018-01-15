package test

import "syscall"

type NumberError struct {
	Func string
	Err error
}

func (e *NumberError) Error() string {
	return e.Func + " " + ":" + e.Err.Error()
}
func Foo(param int) (n int, err error) {
	var stat syscall.Stat_t

	err = syscall.Stat("/home/leare", &stat)

	if param <= 5 {
		return 0, &NumberError{"Foo", err}
	}
	return param * param, nil
}