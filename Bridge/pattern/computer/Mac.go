package computer

import (
	"designPattern/Bridge/pattern/monitor"
	"designPattern/Bridge/pattern/printer"
)

type Mac struct {
	printer printer.Printer
	monitor monitor.Monitor
}

func (m *Mac) SetPrinter(printer printer.Printer) {
	m.printer = printer
}

func (m *Mac) Print() {
	m.printer.PrintFile("Windows file")
}

func (m *Mac) SetMonitor(monitor monitor.Monitor) {
	m.monitor = monitor
}

func (m *Mac) Show() {
	m.monitor.Show()
}
