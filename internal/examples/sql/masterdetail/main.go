//source: https://doc.qt.io/qt-5/qtsql-masterdetail-example.html

package main

import (
	"os"

	"github.com/erry-az/qt-go/core"
	"github.com/erry-az/qt-go/widgets"
)

var qApp *widgets.QApplication

func main() {
	qApp = widgets.NewQApplication(len(os.Args), os.Args)

	if !createConnection() {
		return
	}

	albumDetails := core.NewQFile2("albumdetails.xml")
	window := NewMainWindow(nil, 0)
	window.initWith("artists", "albums", albumDetails, nil)
	window.Show()

	qApp.Exec()
}
