package dialog

import (
	"github.com/erry-az/qt-go/core"

	_ "github.com/erry-az/qt-go/internal/examples/showcases/wallet/wallet/dialog/controller"
)

func init() { sendTemplate_QmlRegisterType2("DialogTemplate", 1, 0, "SendTemplate") }

type sendTemplate struct {
	dialogTemplate

	_ func(string, string) *core.QVariant `slot:"send,->(controller.Controller)"`
}
