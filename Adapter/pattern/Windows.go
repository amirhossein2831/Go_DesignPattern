package pattern

type Windows struct{}

type WindowsAdapter struct {
	Windows *Windows
}

func (w *Windows) InsertUsbPort() {
	println("Inserting usb Port...")
}

func (w *WindowsAdapter) InsertLightningPort() {
	println("Converting Usb to LightningPort...")

	w.Windows.InsertUsbPort()
}
