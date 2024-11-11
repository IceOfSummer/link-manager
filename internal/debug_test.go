package internal

import (
	"fmt"
	"github.com/symbolic-link-manager/internal/localizer"
	"testing"
)

func Obj() localizer.LocalizedError {
	obj := localizer.LocalizedError{}

	p := &obj
	fmt.Println(p)
	return obj
}

func TestReturn(t *testing.T) {
	o := Obj()
	p := &o
	fmt.Println(p)
}
