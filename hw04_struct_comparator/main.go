package main

import (
	"fmt"

	"github.com/lisinearbird/otus-hw/hw04_struct_comparator/bookstruct"
)

type VsMode int

const (
	VsByYear VsMode = iota
	VsBySize
	VsByRate
)

type BookVs struct {
	mode VsMode
}

func NewBookVs(mode VsMode) *BookVs {
	return &BookVs{mode: mode}
}

func (bVs *BookVs) Versus(book1, book2 *bookstruct.Book) bool {
	switch bVs.mode {
	case VsByYear:
		return book1.Year() > book2.Year()
	case VsBySize:
		return book1.Size() > book2.Size()
	case VsByRate:
		return book1.Rate() > book2.Rate()
	default:
		return false
	}
}

func main() {
	book1 := bookstruct.NewBook(11, "Aaaa", "Bbbb", 2020, 200, 1.1)
	book2 := bookstruct.NewBook(12, "Cccc", "Dddd", 2022, 150, 0.7)

	vs := NewBookVs(VsByYear)
	fmt.Println("Сравнение по году:", vs.Versus(book1, book2))

	vs = NewBookVs(VsBySize)
	fmt.Println("Сравнение по размеру:", vs.Versus(book1, book2))

	vs = NewBookVs(VsByRate)
	fmt.Println("Сравнение по рейтингу:", vs.Versus(book1, book2))
}
