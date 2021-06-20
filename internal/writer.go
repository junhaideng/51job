package internal

import (
	"io"
	"sync"
	"sync/atomic"
)

// Writer 可以安全的被多线程调用
type Write struct {
	f       io.WriterAt
	woffset int64
	m sync.Mutex
}

func (w *Write) Write(data []byte) (int, error) {
	var offset int64 
	for {
		offset = atomic.LoadInt64(&w.woffset)
		if atomic.CompareAndSwapInt64(&w.woffset, offset, (offset + int64(len(data)))) {
			break
		}
	}
	w.m.Lock()
	n, err := w.f.WriteAt(data, offset)
	if err != nil {
		w.m.Unlock()
		return 0, err
	}
	w.m.Unlock()
	return n, err
}

func NewWriter(f io.WriterAt) *Write{
	return &Write{
		f: f,
		woffset: 0,
	}
}