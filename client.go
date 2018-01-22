package multidatadogclient

import (
	"log"
	"sync"

	"github.com/DataDog/datadog-go/statsd"
)

// A Client is a handle for sending messages to dogstatsd.  It is safe to
type Client struct {
	sync.Mutex
	clients []*statsd.Client
	index   int
}

//New returns pointer to new multi datadog client
func New(hosts ...string) (mstatsd *Client) {
	mstatsd = &Client{}

	for _, host := range hosts {
		statsdClient, err := statsd.New(host)
		if err != nil {
			log.Println(err)
		}
		mstatsd.clients = append(mstatsd.clients, statsdClient)
	}
	return
}

//Get statsd client from pool
//implement simple round robin
func (c *Client) Get() (client *statsd.Client) {
	c.Lock()
	defer c.Unlock()

	for i := 0; i < len(c.clients); i++ {
		candidate := c.clients[c.index]
		c.index = (c.index + 1) % len(c.clients)
		client = candidate

		if client == nil {
			continue
		}
		break
	}

	if client == nil {
		log.Println("Couldn't find eligible client from pool")
	}

	return
}
