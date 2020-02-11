package printer

import (
	"fmt"

	"github.com/loloxiaoz/fmt-conv/xlog"
)

// NewHTMLPrinter returns a Printer which
// is able to convert an HTML file to PDF.
func NewHTMLPrinter(logger xlog.Logger, fpath string, opts ChromePrinterOptions) Printer {
	URL := fmt.Sprintf("file://%s", fpath)
	return chromePrinter{
		logger: logger,
		url:    URL,
		opts:   opts,
	}
}
