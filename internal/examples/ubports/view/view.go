package main

import (
	"os"

	"github.com/erry-az/qt-go/gui"
	"github.com/erry-az/qt-go/quick"
)

func main() {
	EnableHighDPI()

	gui.NewQGuiApplication(len(os.Args), os.Args)

	var view = quick.NewQQuickView(nil)
	view.SetSource(NewQUrl("qrc:/qml/view.qml"))
	view.SetResizeMode(quick.QQuickView__SizeRootObjectToView)
	view.Show()

	gui.QGuiApplication_Exec()
}
