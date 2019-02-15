// +build js,wasm
package net

import (
	"errors"
	"fmt"
	"net"
	"sync"
	"syscall/js"
	"time"

	"github.com/maxmcd/wasabi/pkg/wasm"
)

type eventState struct {
	state int
	cond  *sync.Cond
}

func (e *eventState) readable() bool {
	return ((e.state >> 0) & 1) == 1
}
func (e *eventState) writeable() bool {
	return ((e.state >> 1) & 1) == 1
}
func (e *eventState) hup() bool {
	return ((e.state >> 2) & 1) == 1
}
func (e *eventState) error() bool {
	return ((e.state >> 3) & 1) == 1
}

func (e *eventState) dead() bool {
	return e.error() || e.hup()
}

func (e *eventState) writewait() error {
	if e.writeable() {
		return nil
	}
	e.cond.L.Lock()
	for {
		if e.dead() {
			return errors.New("wasabi: Connection closed")
		}
		if e.writeable() {
			e.state = e.state ^ (1 << 1)
			break
		}
		e.cond.Wait()
	}
	e.cond.L.Unlock()
	return nil
}

func (e *eventState) readwait() error {
	e.cond.L.Lock()
	for {
		if e.dead() {
			return errors.New("wasabi: Connection closed")
		}
		if e.readable() {
			e.state = e.state ^ (1 << 0)
			break
		}
		e.cond.Wait()
	}
	e.cond.L.Unlock()
	return nil
}

func newEventState() *eventState {
	return &eventState{
		state: 0,
		cond:  sync.NewCond(&sync.Mutex{}),
	}
}

var connections map[int32]*eventState

func init() {
	connections = make(map[int32]*eventState)
	callback := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		for i, _ := range args {
			if i%2 != 0 {
				continue
			}
			token := args[i].Int()
			es := connections[int32(token)]
			es.state = args[i+1].Int()
			es.cond.Broadcast()
		}
		return nil
	})
	nlr := js.Global().Get("net_listener").New()
	nlr.Call("register", callback)
}

type Resolver struct {
	net.Resolver
}

// func ListenTCP(network string, laddr *net.TCPAddr) (*TCPListener, error) {

// }

type Listener struct {
	token int32
}

func acceptTcp(id int32) (int32, bool)

func (l *Listener) Accept() (c Conn, err error) {
	for {
		token, ok := acceptTcp(l.token)
		if ok {
			connections[token] = newEventState()
			c.token = int32(token)
			return c, nil
		}
		bytes, _ := wasm.GetBytes(token)
		err = errors.New(string(bytes))
		if err.Error() != "Resource temporarily unavailable (os error 35)" {
			return
		}
		if err = connections[l.token].readwait(); err != nil {
			return c, err
		}
	}
	return
}

type Conn struct {
	token int32
}

func readConn(id int32, b []byte) (int32, bool)

func (c *Conn) Read(b []byte) (ln int, err error) {
	// TODO EOF error?
	for {
		length, ok := readConn(c.token, b)
		if ok {
			return int(length), nil
		}
		bytes, _ := wasm.GetBytes(length) // ln is ref if there's an error
		err = errors.New(string(bytes))
		if err.Error() != "Resource temporarily unavailable (os error 35)" {
			return 0, err
		}
		if err := connections[c.token].readwait(); err != nil {
			return 0, err
		}
	}
	return
}

func writeConn(id int32, b []byte) (int32, bool)

func (c *Conn) Write(b []byte) (ln int, err error) {
	for {
		length, ok := writeConn(c.token, b)
		if ok {
			return int(length), nil
		}
		bytes, _ := wasm.GetBytes(length) // ln is ref if there's an error
		err = errors.New(string(bytes))
		if err.Error() != "Resource temporarily unavailable (os error 35)" {
			return 0, err
		}
		if err := connections[c.token].writewait(); err != nil {
			return 0, err
		}
	}
}

func (c *Conn) Close() error {
	return nil
}

func (c *Conn) LocalAddr() net.Addr {
	// TODO: fix
	return &net.TCPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 8668,
		Zone: "",
	}
}
func (c *Conn) RemoteAddr() net.Addr {
	// TODO: fix
	return &net.TCPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 8668,
		Zone: "",
	}
}
func (c *Conn) SetDeadline(t time.Time) error {
	return nil
}
func (c *Conn) SetReadDeadline(t time.Time) error {
	return nil
}
func (c *Conn) SetWriteDeadline(t time.Time) error {
	return nil
}

func listenTcp(addr string) (int32, bool)

func ListenTcp(addr string) (Listener, error) {
	id, ok := listenTcp(addr)
	if ok {
		connections[id] = newEventState()
		fmt.Printf("%#v\n", connections)
		return Listener{token: id}, nil
	}
	bytes, _ := wasm.GetBytes(id) // id is ref if there's an error
	err := errors.New(string(bytes))
	// TODO: don't just pass a negative number
	return Listener{token: -1}, err
}

// func DialContext(ctx context.Context, network, addr string) (net.Conn, error) {
// 	var c net.Conn
// 	c = &Conn{token: -1}
// 	return c, nil
// }

func dialTcp(addr string) (int32, bool)

func Dial(network, addr string) (c net.Conn, err error) {
	if network != "tcp" {
		return c, errors.New("tcp is the only protocal supported")
	}
	ref, ok := dialTcp(addr)
	if ok {
		connections[ref] = newEventState()
		if err := connections[ref].writewait(); err != nil {
			return c, err
		}
		var c net.Conn
		c = &Conn{token: ref}
		return c, nil
	}

	bytes, _ := wasm.GetBytes(ref)
	return c, errors.New(string(bytes))
}
