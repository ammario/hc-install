package releasesjson

import (
	"io"
	"sync"
)

var bufPool = &sync.Pool{
	New: func() interface{} {
		return make([]byte, 32*1024)
	},
}

// ioCopy is like io.Copy but uses a sync.Pool to improve performance, since
// copy buffers are large and can be reused without zeroing.
func ioCopy(dst io.Writer, src io.Reader) (written int64, err error) {
	buf := bufPool.Get().([]byte)
	//lint:ignore SA6002 slices are pointers
	defer bufPool.Put(buf)
	return io.CopyBuffer(dst, src, buf)
}
