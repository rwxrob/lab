package main

import (
	"io"
	"log"
	"os"

	"github.com/mikefarah/yq/v4/pkg/yqlib"
	logging "gopkg.in/op/go-logging.v1"
)

// creating a printer is a pain in the ass, but flexible

func NewYamlPrinter(
	w io.Writer,
	f yqlib.PrinterOutputFormat, // i mean, what the fuck
	unwrap bool,
	colors bool,
	indent int,
	sep bool,
) yqlib.Printer {
	enc := yqlib.NewYamlEncoder(indent, colors, sep, unwrap)
	pwr := yqlib.NewSinglePrinterWriter(w)
	return yqlib.NewPrinter(enc, pwr)
}

func setupLogging() {
	format := logging.MustStringFormatter(
		`%{color}%{time:15:04:05} %{shortfunc} [%{level:.4s}]%{color:reset} %{message}`,
	)
	backend := logging.AddModuleLevel(
		logging.NewBackendFormatter(logging.NewLogBackend(os.Stderr, "", 0), format))

	backend.SetLevel(logging.ERROR, "")
	logging.SetBackend(backend)
}

func main() {
	setupLogging()
	ev := yqlib.NewAllAtOnceEvaluator()
	pr := NewYamlPrinter(os.Stdout, yqlib.YamlOutputFormat, true, false, 2, true)
	dc := yqlib.NewYamlDecoder()
	err := ev.EvaluateFiles(os.Args[1], []string{"sample.yaml"}, pr, true, dc)
	if err != nil {
		log.Print(err)
	}

}
