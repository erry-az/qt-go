//go:build !qml
// +build !qml

package main

import (
	"os"

	"github.com/erry-az/qt-go/core"
	"github.com/erry-az/qt-go/widgets"

	"github.com/erry-az/qt-go/internal/examples/sql/masterdetail_qml/controller"

	"github.com/erry-az/qt-go/internal/examples/sql/masterdetail_qml/view"
)

func main() {
	qApp := widgets.NewQApplication(len(os.Args), os.Args)

	controller.NewController(nil).InitWith(core.NewQFile2(":/albumdetails.xml"), qApp)

	view.NewViewController(nil, 0).Show()

	qApp.Exec()
}
