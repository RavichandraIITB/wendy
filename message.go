package pastry

import (
	"encoding/json"
	"net"
	"time"
)

// Message represents the messages that are sent through the cluster of Nodes
type Message struct {
	Purpose byte
	Hops    []NodeID  // A list of the nodes a message went through
	Key     NodeID    // The message's ID
	Value   string    // The message being passed
	Sent    time.Time // The time the message was initially sent
}

const NODE_JOIN = byte(0x01) // Used when a Node joins the cluster
const NODE_EXIT = byte(0x02) // Used when a Node leaves the cluster
const NODE_TEST = byte(0x03) // Used when a Node is being tested
const NODE_STAT = byte(0x04) // Used when a Node broadcasts state info

// String returns a string representation of a message.
func (m *Message) String() string {
	return m.Key.String() + ": " + m.Value
}

// send sends a message to the specified IP address.
func (m *Message) send(ip string) error {
	conn, err := net.Dial("tcp", ip)
	if err != nil {
		return err
	}
	encoder := json.NewEncoder(conn)
	return encoder.Encode(m)
}
