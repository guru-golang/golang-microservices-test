package net_lib

import (
	"fmt"
	"net"
	"sync"
)

type ConnPool struct {
	conns   chan net.Conn
	factory func() (net.Conn, error)
	mu      sync.Mutex
}

func NewConnPool(factory func() (net.Conn, error), size int) (*ConnPool, error) {
	pool := &ConnPool{
		conns:   make(chan net.Conn, size),
		factory: factory,
	}

	for i := 0; i < size; i++ {
		conn, err := factory()
		if err != nil {
			return nil, err
		}
		pool.conns <- conn
	}

	return pool, nil
}

func (p *ConnPool) Get() (net.Conn, error) {
	select {
	case conn := <-p.conns:
		return conn, nil
	default:
		p.mu.Lock()
		defer p.mu.Unlock()
		conn, err := p.factory()
		if err != nil {
			return nil, err
		}
		return conn, nil
	}
}

func (p *ConnPool) Put(conn net.Conn) error {
	select {
	case p.conns <- conn:
		return nil
	default:
		conn.Close()
		return fmt.Errorf("connection pool is full")
	}
}
