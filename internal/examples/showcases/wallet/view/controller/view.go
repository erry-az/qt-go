package controller

import "github.com/erry-az/qt-go/core"

var Controller *viewController

type viewController struct {
	core.QObject

	_ func() `constructor:"init"`

	_ func(bool) `signal:"blur"`
}

func (c *viewController) init() {
	Controller = c
}
