package multidatadogclient

import (
	"net"
	"testing"
)

func TestClient_Get(t *testing.T) {
	//spawn 3 udp server to mock datadog agent
	udpAddr1, _ := net.ResolveUDPAddr("udp", "localhost:10000")
	net.ListenUDP("udp", udpAddr1)
	udpAddr2, _ := net.ResolveUDPAddr("udp", "localhost:20000")
	net.ListenUDP("udp", udpAddr2)
	udpAddr3, _ := net.ResolveUDPAddr("udp", "localhost:30000")
	net.ListenUDP("udp", udpAddr3)

	c := New("localhost:10000", "localhost:20000", "localhost:30000")
	client1 := c.Get()
	client2 := c.Get()
	client3 := c.Get()
	client4 := c.Get()
	client5 := c.Get()
	client1.Namespace = "client1"
	client2.Namespace = "client2"
	client3.Namespace = "client3"

	if client1.Namespace == client2.Namespace {
		t.Error("client1 and client2 should be not equal", client1.Namespace, client2.Namespace)
	}
	if client2.Namespace == client3.Namespace {
		t.Error("client2 and client3 should be not equal", client2.Namespace, client3.Namespace)
	}
	if client1.Namespace != client4.Namespace {
		t.Error("client1 and client4 should be equal", client1.Namespace, client4.Namespace)
	}
	if client2.Namespace != client5.Namespace {
		t.Error("client2 and client5 should be equal", client2.Namespace, client5.Namespace)
	}
}
