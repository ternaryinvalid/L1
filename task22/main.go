package main

import "fmt"

type newPrintS struct{}

type basePrintS struct{}

type adapter struct {
	adapted basePrintS
}

func (b *basePrintS) InitialPrint(str string) string {
	return "initial print " + str
}

func (a *adapter) AdaptedPrint() string {
	return a.adapted.InitialPrint("adapted hello")
}

func (s *newPrintS) NewPrint() string {
	return "new print"
}

func main() {
	str := &basePrintS{}
	newStr := &newPrintS{}
	adapt := &adapter{*str}
	fmt.Println(newStr.NewPrint())
	fmt.Println(adapt.AdaptedPrint())
}
