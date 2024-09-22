//source: https://doc.qt.io/qt-5/qtsensors-accelbubble-example.html

package main

import (
	"os"

	"github.com/erry-az/qt-go/core"
	"github.com/erry-az/qt-go/gui"
	"github.com/erry-az/qt-go/qml"
	_ "github.com/erry-az/qt-go/sensors"
)

func main() {
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	gui.NewQGuiApplication(len(os.Args), os.Args)

	var app = qml.NewQQmlApplicationEngine(nil)
	app.Load(core.NewQUrl3("qrc:///accelbubble.qml", 0))

	gui.QGuiApplication_Exec()
}
