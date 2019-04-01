package impl

import "fmt"

type Api12Impl struct {
	Name string
}

func (api12 *Api12Impl) Api1() {
	fmt.Println("api1", api12.Name)
}

func (api12 *Api12Impl) Api2() {
	fmt.Println("api2", api12.Name)
}
