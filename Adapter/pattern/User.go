package pattern

import "fmt"

type User struct{}

// InsertLightningPortToComputer now this can work with both mac and Windows because we use an adapter
func (u *User) InsertLightningPortToComputer(c Computer) {
	fmt.Println("User inserts Lightning connector into computer.")
	c.InsertLightningPort()
}
