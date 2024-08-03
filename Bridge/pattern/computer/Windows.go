package computer

import (
	"designPattern/Bridge/pattern/monitor"
	"designPattern/Bridge/pattern/printer"
)

type Windows struct {
	printer printer.Printer
	monitor monitor.Monitor
}

func (w *Windows) SetPrinter(printer printer.Printer) {
	w.printer = printer
}

func (w *Windows) Print() {
	w.printer.PrintFile("Windows file")
}

func (w *Windows) SetMonitor(monitor monitor.Monitor) {
	w.monitor = monitor
}

func (w *Windows) Show() {
	w.monitor.Show()
}
