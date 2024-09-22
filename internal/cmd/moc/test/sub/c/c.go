package c

import (
	"github.com/erry-az/qt-go/core"

	_ "github.com/erry-az/qt-go/internal/cmd/moc/test/sub/conf"
)

type StructSubGoC struct{}
type StructSubMocC struct{ core.QObject }
