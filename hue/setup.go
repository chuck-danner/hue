package hue

import (
	"fmt"

	"github.com/hashicorp/mdns"
)

func discoverMdns() {

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

type discover struct {
	Id         string `json:"id"`
	InternalIp string `json:"internalipaddress"`
}

func discoverEndpoint() string {
	var discover []discover

	//http.DefaultClient.Get("https://discovery.meethue.com", &discover)
	if len(discover) == 0 {
		return ""
	}
	return discover[0].InternalIp
}
