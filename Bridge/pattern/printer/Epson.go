package printer

type Epson struct{}

func (e *Epson) PrintFile(fileName string) {
	println("Printing file with Epson printer..., file name:", fileName)
}
