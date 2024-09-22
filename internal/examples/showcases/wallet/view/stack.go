package view

import (
	"github.com/erry-az/qt-go/quick"

	_ "github.com/erry-az/qt-go/internal/examples/showcases/wallet/view/controller"
)

func init() { stackTemplate_QmlRegisterType2("ViewTemplate", 1, 0, "StackTemplate") }

type stackTemplate struct {
	quick.QQuickItem

	_ func(code string) `signal:"Clicked,<-(controller.NewStackController(nil))"`
}
