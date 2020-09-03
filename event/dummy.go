// This file is part of a workaround for `go mod vendor` which won't
// vendor C files if there are no Go files in the same directory.
// This prevents the C header files in event/ from being vendored.
//
// These two file make it so `go mod vendor` behaves correctly.
//
// See this issue for reference: https://github.com/golang/go/issues/26366

package event
