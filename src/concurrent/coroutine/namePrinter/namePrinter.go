package namePrinter

import (
	"fmt"
	"time"
)

type NamePrinter struct {
	Name string
}

func (np *NamePrinter) Task() {
	fmt.Println(np.Name)
	time.Sleep(time.Second)
}
