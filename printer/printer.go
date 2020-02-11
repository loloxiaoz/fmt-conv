package printer

import (
	"github.com/loloxiaoz/fmt-conv/xlog"
)

// Printer is a type that can create a PDF file from a source.
// The source is defined in the underlying implementation.
type Printer interface {
	Print(destination string) error
}

func logOptions(logger xlog.Logger, opts interface{}) {
	const op string = "printer.logOptions"
	logger.DebugfOp(op, "options: %+v", opts)
}
