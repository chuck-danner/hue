package hue

import (
	"fmt"

	"github.com/hashicorp/mdns"
)

func discover() {

	entriesCh := make(chan *mdns.ServiceEntry, 4)
	go func() {
		for entry := range entriesCh {
			fmt.Printf("Got new entry: %v\n", entry)
			fmt.Print(entry.Addr)
		}
	}()

	// Start the lookup
	err := mdns.Lookup("_hue._tcp", entriesCh)
	if err != nil {
		fmt.Println(err)
	}
	close(entriesCh)
}
