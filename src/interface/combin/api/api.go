package api

type Api1 interface {
	Api1()
}

type Api2 interface {
	Api2()
}

type Api12 interface {
	Api1
	Api2
}
