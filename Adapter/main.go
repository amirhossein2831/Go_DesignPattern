package main

import "designPattern/Adapter/pattern"

func main() {

	// despite the windows does not have lightning we use an adapter(tabdil) to use it as our computer which has Lightning
	user := &pattern.User{}

	mac := &pattern.Mac{}
	windows := &pattern.Windows{}
	windowsAdapter := &pattern.WindowsAdapter{
		Windows: windows,
	}

	user.InsertLightningPortToComputer(mac)
	user.InsertLightningPortToComputer(windowsAdapter)
}
