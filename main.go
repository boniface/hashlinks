package main

import (
	"fmt"
	"github.com/jasonlvhit/gocron"
	"hashlinks/services"
)

func processLinks() {
	fmt.Println("Processing Links .")
	services.ProcessLinks()

}
func main() {
	fmt.Println(" The Hashlinks Server has Started!!!")
	s := gocron.NewScheduler()
	s.Every(5).Minutes().Do(processLinks)
	<-s.Start()
}
