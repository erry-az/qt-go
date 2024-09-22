package dialog

import (
	"github.com/erry-az/qt-go/widgets"

	_ "github.com/erry-az/qt-go/internal/examples/showcases/wallet/files/dialog/controller"
)

type filesUploadTemplate struct {
	filesDialogTemplate

	_ func([]string) `slot:"uploadFiles,->(controller.Controller)"`
}

func (t *filesUploadTemplate) show(cident string) {
	if cident == "files" {
		t.Blur(true)
		t.UploadFiles(widgets.QFileDialog_GetOpenFileNames(nil, "Select one or more files to upload", "", "", "", 0))
		t.Blur(false)
	}
}
