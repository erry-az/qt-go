package dialog

import (
	"github.com/erry-az/qt-go/core"

	_ "github.com/erry-az/qt-go/internal/examples/showcases/wallet/wallet/dialog/controller"
)

func init() { unlockTemplate_QmlRegisterType2("DialogTemplate", 1, 0, "UnlockTemplate") }

type unlockTemplate struct {
	dialogTemplate

	_ func(string) *core.QVariant `slot:"unlock,->(controller.Controller)"`
}
