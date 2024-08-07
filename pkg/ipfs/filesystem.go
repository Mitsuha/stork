package ipfs

import (
	"context"
	"errors"
	"fmt"
	"github.com/ipfs/boxo/files"
	"github.com/ipfs/boxo/path"
	"github.com/ipfs/go-cid"
	iface "github.com/ipfs/kubo/core/coreiface"
	"io/fs"
	"net/http"
	"strings"
	"time"
)

// Filesystem implements the http.FileSystem interface for IPFS.
type Filesystem struct {
	IPFS iface.CoreAPI
	Ctx  context.Context
}

func NewFilesystem(ipfs iface.CoreAPI, ctx context.Context) *Filesystem {
	return &Filesystem{IPFS: ipfs, Ctx: ctx}
}

func (f *Filesystem) Open(name string) (http.File, error) {
	name = strings.Trim(name, "/")

	c, err := cid.Decode(name)
	if err != nil {
		return nil, fmt.Errorf("invalid cid: %s", name)
	}
	fn, err := f.IPFS.Unixfs().Get(f.Ctx, path.FromCid(c))
	if err != nil {
		return nil, fmt.Errorf("could not fetch cid: %s", err)
	}
	file, ok := fn.(files.File)
	if !ok {
		return nil, fmt.Errorf("node is not a file")
	}

	return NewFile(name, file), nil
}

// WithOption applies the given options to the filesystem.
func (f *Filesystem) WithOption(options ...func(filesystem *Filesystem)) {
	for _, option := range options {
		option(f)
	}
}

// File implements the http.File interface for IPFS.
type File struct {
	files.File
	Name string
}

func NewFile(name string, file files.File) *File {
	return &File{File: file, Name: name}
}

func (f *File) Readdir(_ int) ([]fs.FileInfo, error) {
	return nil, errors.New("not a directory")
}

func (f *File) Stat() (fs.FileInfo, error) {
	return NewFileInfo(f), nil
}

// FileInfo implements the fs.FileInfo interface for IPFS.
type FileInfo struct {
	file *File
}

func NewFileInfo(file *File) *FileInfo {
	return &FileInfo{file: file}
}

func (f *FileInfo) Name() string {
	return f.file.Name
}

func (f *FileInfo) Size() int64 {
	size, err := f.file.Size()
	if err != nil {
		return 0
	}
	return size
}

func (f *FileInfo) Mode() fs.FileMode {
	return fs.ModePerm
}

func (f *FileInfo) ModTime() time.Time {
	return time.Time{}
}

func (f *FileInfo) IsDir() bool {
	return false
}

func (f *FileInfo) Sys() any {
	return f.file.File
}
