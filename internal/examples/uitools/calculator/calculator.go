package main

import (
	"os"

	"github.com/erry-az/qt-go/widgets"

	"github.com/erry-az/qt-go/internal/examples/uitools/calculator/ui"
)

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)

	ui.NewCalculatorForm(nil).Show()

	widgets.QApplication_Exec()
}
