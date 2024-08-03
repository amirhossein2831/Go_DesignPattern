package main

import (
	"designPattern/Bridge/pattern/computer"
	"designPattern/Bridge/pattern/monitor"
	"designPattern/Bridge/pattern/printer"
)

func main() {
	hpPrinter := &printer.Hp{}
	epsonPrinter := &printer.Epson{}

	lgMonitor := &monitor.LG{}
	samsungMonitor := &monitor.Samsung{}

	mac := computer.Mac{}
	windows := computer.Windows{}

	mac.SetPrinter(hpPrinter)
	mac.SetMonitor(lgMonitor)
	mac.Show()
	mac.Print()

	windows.SetPrinter(epsonPrinter)
	windows.SetMonitor(samsungMonitor)
	windows.Show()
	windows.Print()
}
