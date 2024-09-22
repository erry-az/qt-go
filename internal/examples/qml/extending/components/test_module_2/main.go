package main

import (
	"os"

	"github.com/erry-az/qt-go/core"
	"github.com/erry-az/qt-go/gui"
	"github.com/erry-az/qt-go/quick"

	_ "github.com/erry-az/qt-go/internal/examples/qml/extending/components/test_module_2/component"
)

func main() {
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	gui.NewQGuiApplication(len(os.Args), os.Args)

	view := quick.NewQQuickView(nil)
	view.Engine().AddImportPath("qrc:/qml")
	view.SetResizeMode(quick.QQuickView__SizeRootObjectToView)
	view.SetSource(core.NewQUrl3("qrc:/qml/app.qml", 0))
	view.Show()

	gui.QGuiApplication_Exec()
}
