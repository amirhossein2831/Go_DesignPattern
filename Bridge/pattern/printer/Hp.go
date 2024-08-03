package printer

type Hp struct{}

func (h *Hp) PrintFile(fileName string) {
	println("Printing file with HP printer..., file name:", fileName)
}
