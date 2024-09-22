package main

import (
	"os"
	"strings"

	"github.com/erry-az/qt-go/core"
	"github.com/erry-az/qt-go/felgo"
	"github.com/erry-az/qt-go/qml"
	"github.com/erry-az/qt-go/widgets"

	"github.com/erry-az/qt-go/internal/examples/3rdparty/qzxing"

	"github.com/erry-az/qt-go/internal/examples/felgo/appdemos/qtws/cpp"
)

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)
	felgoApp := felgo.NewFelgoApplication(nil)

	// Use platform-specific fonts instead of Felgo's default font
	felgoApp.SetPreservePlatformFonts(true)

	// QQmlApplicationEngine is the preferred way to start qml projects since Qt 5.2
	// if you have older projects using Qt App wizards from previous QtCreator versions than 3.1, please change them to QQmlApplicationEngine
	engine := qml.NewQQmlApplicationEngine(nil)
	felgoApp.Initialize(engine)

	//

	// 10MB cache for network data (chris bartsch style)
	engine.SetNetworkAccessManagerFactory(cpp.NewDiskCacheFactoryWithCacheSize(1024 * 1024 * 10))

	// register QZXing qml types for barcode scannning
	qzxing.RegisterQMLTypes()

	//

	// use this during development
	// for PUBLISHING, use the entry point below
	mainQmlFile := "QtWSMain.qml"
	if strings.Contains(os.Args[0], "/deploy/") {
		felgoApp.SetMainQmlFileName(strings.Split(os.Args[0], "/deploy/")[0] + "/qml/" + mainQmlFile) //to make qtdeploy -fast work
	} else if core.QSysInfo_ProductType() == "ios" || core.QSysInfo_ProductType() == "android" {
		felgoApp.SetMainQmlFileName("qml/" + mainQmlFile) //to make qtdeploy work
	} else {
		pwd, _ := os.Getwd()
		felgoApp.SetMainQmlFileName(pwd + "/qml/" + mainQmlFile) //to make go run/build work
	}

	// use this instead of the above call to avoid deployment of the qml files and compile them into the binary with qt's resource system qrc
	// this is the preferred deployment option for publishing games to the app stores, because then your qml files and js files are protected
	// felgoApp.SetMainQmlFileName("qrc:/qml/" + mainQmlFile)

	engine.Load(core.NewQUrl3(felgoApp.MainQmlFileName(), 0))

	// to start your project as Live Client, comment (remove) the lines "felgoApp.SetMainQmlFileName ..." & "engine.Load ...",
	// and uncomment the line below
	//felgo.NewFelgoLiveClient(engine, nil)

	widgets.QApplication_Exec()
}
