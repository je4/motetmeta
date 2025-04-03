package main

import (
	"database/sql"
	"emperror.dev/errors"
	"fmt"
)

func NewCycle(id int, title string) *Cycle {
	return &Cycle{
		ID:     id,
		Title:  title,
		Motets: []*Motet{},
	}
}

type Cycle struct {
	ID     int
	Title  string
	Motets []*Motet
}

func (c *Cycle) String() string {
	str := fmt.Sprintf("Cycle % 4d: %s", c.ID, c.Title)
	for _, motet := range c.Motets {
		str += "\n   " + motet.String()
	}
	return str
}

func (c *Cycle) loadMotets(db *sql.DB) error {
	sql := "SELECT m.nid, m.title FROM motet_cycle mc, motets m WHERE mc.motet_id=m.nid AND mc.cycle_id = ?"
	rows, err := db.Query(sql, c.ID)
	if err != nil {
		return errors.Wrapf(err, "cannot query motets for cycle %d", c.ID)
	}
	defer rows.Close()
	for rows.Next() {
		var nid int
		var title string
		if err := rows.Scan(&nid, &title); err != nil {
			return errors.Wrapf(err, "cannot scan motet for cycle %d", c.ID)
		}
		motet := NewMotet(nid, title)
		c.Motets = append(c.Motets, motet)
	}
	if err := rows.Err(); err != nil {
		return errors.Wrapf(err, "error iterating over motets for cycle %d", c.ID)
	}
	return nil
}
