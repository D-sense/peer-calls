package wsadapter

import "github.com/peer-calls/peer-calls/server/ws/wsmessage"

type Client interface {
	ID() string
	Write(message wsmessage.Message) error
	Metadata() string
	SetMetadata(metadata string)
}

type Adapter interface {
	Add(client Client) error
	Remove(clientID string) error
	Broadcast(msg wsmessage.Message) error
	Metadata(clientID string) (string, bool)
	SetMetadata(clientID string, metadata string) bool
	Emit(clientID string, msg wsmessage.Message) error
	Clients() (map[string]string, error)
	Size() (int, error)
	Close() error
}