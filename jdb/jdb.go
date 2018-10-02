package jdb

import (
	"fmt"

	scribble "github.com/nanobox-io/golang-scribble"
)

func OpenDB() (jdb *scribble.Driver, err error) {
	jdb, err = scribble.New("./JDB", nil)
	if err != nil {
		fmt.Println("ErrorDB", err)
	}
	return
}
