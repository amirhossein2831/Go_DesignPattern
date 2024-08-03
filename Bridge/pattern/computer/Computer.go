package computer

import (
	"designPattern/Bridge/pattern/monitor"
	"designPattern/Bridge/pattern/printer"
)

type Computer interface {
	Print()
	Show()
	SetMonitor(monitor.Monitor)
	SetPrinter(printer.Printer)
}
