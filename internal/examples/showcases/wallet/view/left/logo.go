package left

import (
	"github.com/erry-az/qt-go/quick"

	_ "github.com/erry-az/qt-go/internal/examples/showcases/wallet/view/left/controller"
)

func init() { logoTemplate_QmlRegisterType2("LeftTemplate", 1, 0, "LogoTemplate") }

type logoTemplate struct {
	quick.QQuickItem

	_ func() `signal:"clicked,->(controller.NewLogoController(nil))"`
}
