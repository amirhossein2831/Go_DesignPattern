package pattern

// so our computer only support LightningPort for mac is ok but how to use a Windows as a computer we need an adaptor
//to wrap the Windows inside it and convert usb to lightning

type Computer interface {
	InsertLightningPort()
}
