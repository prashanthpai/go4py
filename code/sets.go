// Example from:
// https://github.com/gluster/glusterd2/blob/master/servers/sunrpc/server.go

var clientsList = struct {
	sync.RWMutex
	c map[net.Conn]bool
}{
	// This map is used as a set. Values are not consumed.
	c: make(map[net.Conn]bool),
}
