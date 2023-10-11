package main

import "patterns/internal/facade"

func main() {
	computer := facade.NewComputerFacade()
	computer.Start()
}
