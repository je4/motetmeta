package main

import "fmt"

func NewMotet(id int, title string) *Motet {
	return &Motet{
		ID:    id,
		Title: title,
	}
}

type Motet struct {
	ID    int
	Title string
}

func (m *Motet) String() string {
	return fmt.Sprintf("Motet % 4d: %s", m.ID, m.Title)
}
