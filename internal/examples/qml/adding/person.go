package main

import (
	"github.com/erry-az/qt-go/core"
)

type Person struct {
	core.QObject

	_ string `property:"name"`
	_ int    `property:"shoeSize"`
}
