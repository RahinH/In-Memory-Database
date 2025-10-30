package main

import (
	"bufio"
	"io"
	"os"
	"sync"
	"time"
)

// Aof handles writing and reading commands to and from the AOF file
type Aof struct {
	file *os.File
	rd   *bufio.Reader
	mu   sync.Mutex
}

// NewAof opens or creates the AOF file and starts periodic syncing
func NewAof(path string) (*Aof, error) {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return nil, err
	}

	aof := &Aof{
		file: f,
		rd:   bufio.NewReader(f),
	}

	// Sync file to disk every 1 second for durability
	go func() {
		for {
			aof.mu.Lock()
			aof.file.Sync()
			aof.mu.Unlock()
			time.Sleep(time.Second)
		}
	}()

	return aof, nil
}

// Close safely closes the AOF file
func (aof *Aof) Close() error {
	aof.mu.Lock()
	defer aof.mu.Unlock()
	return aof.file.Close()
}

// Write appends a command (as RESP) to the AOF file
func (aof *Aof) Write(value Value) error {
	aof.mu.Lock()
	defer aof.mu.Unlock()

	_, err := aof.file.Write(value.Marshal())
	if err != nil {
		return err
	}
	return nil
}

// Read replays all commands from the AOF file and executes them via callback
func (aof *Aof) Read(callback func(value Value)) error {
	aof.mu.Lock()
	defer aof.mu.Unlock()

	resp := NewResp(aof.file)

	for {
		value, err := resp.Read()
		if err == nil {
			callback(value)
			continue
		}
		if err == io.EOF {
			break
		}
		return err
	}

	return nil
}
