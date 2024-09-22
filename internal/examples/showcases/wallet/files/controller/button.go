package controller

import (
	"github.com/erry-az/qt-go/core"

	_ "github.com/erry-az/qt-go/internal/examples/showcases/wallet/files/dialog/controller"
)

var ButtonController *buttonController

type buttonController struct {
	core.QObject

	_ func() `constructor:"init"`

	_ func(string) `signal:"clicked,->(controller.Controller.Show)"`
}

func (c *buttonController) init() {
	ButtonController = c
}
