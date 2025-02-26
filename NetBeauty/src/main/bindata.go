// Code generated for package main by go-bindata DO NOT EDIT. (@generated)
// sources:
// nbloader/nbloader.dll
package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _nbloaderNbloaderDll = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x58\x6d\x6c\x1c\xc5\xf9\xff\xcd\xee\xde\xde\xdd\xda\xde\x78\xef\x12\xc7\x49\x6c\xe7\x82\x63\xfe\xfe\x63\xb2\xbe\x10\x92\xe0\x40\x9a\xd8\x3e\x27\xb9\xc6\x6f\xf8\x9c\xf0\xa2\x52\x67\xef\x6e\x72\x59\xd8\xdb\x3d\x76\xf7\x42\x0e\x55\xc6\x88\x14\x45\x54\xb4\xcd\x87\x52\x91\x22\x84\x22\xa8\x8a\xfa\x22\x50\x11\x44\xa8\x54\x50\x45\xaa\x2a\x94\x0a\x24\x2a\xa4\xa2\x56\x2d\xaa\x54\x15\x55\x05\x55\x2d\x42\xa8\xc2\xd5\xcc\xee\xf9\x2e\xbe\x34\xc0\x27\x3e\xb4\x63\x79\x66\x9e\xdf\x3c\x6f\xf3\xcc\x3c\xcf\xde\xee\xd4\x9d\xdf\x86\x08\x40\x02\xb0\xbc\x0c\x5c\x40\xd0\xf6\xe3\xd3\xdb\x12\x00\x75\xf3\xcb\x2a\x5e\x88\x5f\xda\x72\x81\x4c\x5e\xda\x32\x7f\xc2\xf4\x52\x15\xd7\x29\xb9\x46\x39\x55\x30\x6c\xdb\xf1\x53\x79\x9a\x72\xab\x76\xca\xb4\x53\x99\x99\x5c\xaa\xec\x14\xa9\xde\xd1\xa1\x6c\x0d\x75\xcc\x4e\x00\x93\x44\xc4\xe1\x57\xd6\x3d\x52\xd7\xfb\x07\x5c\x93\x6a\x23\x69\xa0\x13\x80\x1c\x60\x7f\x1b\x06\x90\x5a\x71\xac\x93\xcf\x85\xc0\x6f\xa0\x31\x72\xa7\x84\x60\x2a\xe2\xd8\xd7\xb9\x92\xce\x40\x24\x18\x57\x06\xde\x9e\x18\x06\x66\xc2\x0d\xf7\x88\x57\xd8\xe4\x31\xa0\x9d\xa9\x1d\x06\x7a\x3e\x43\x4c\x56\x5a\x0a\x88\x35\x91\xb1\x14\x70\xa8\x89\xd6\x7d\x7a\xca\x07\xf0\x8f\x35\xe1\xbe\x3a\x1b\x7e\x37\xa9\x38\xa6\xbb\x9e\x5b\x40\xe8\xdb\xfe\x70\xa3\x89\xcb\xf9\xf6\x03\xfb\x75\x97\x5a\x4e\x21\xf0\x95\xf9\xcc\x75\x75\xb5\xf0\x8d\xad\x76\xf3\x8d\xe1\x60\x3c\xc4\x45\x22\xe8\xdc\x0a\x54\xda\x00\xf2\x79\xf6\xda\xd4\x5e\x1c\x6c\x03\x94\xb5\x9f\xc8\x12\x20\x7b\xed\x80\xe2\x74\x00\x4a\x1d\x8d\x30\x54\x65\xe8\x1a\x40\xb9\xae\x4f\x18\xec\x64\x23\x90\x4c\x4b\x30\xc1\xed\x6a\x8b\x32\x20\x6d\xfb\x8a\xa7\x01\xca\x12\x9b\xbb\x04\xa8\x0c\x26\x00\x65\x60\x9b\x74\xed\xda\xa1\x88\x93\xe4\x73\xf9\x5a\x77\x2f\x50\xd9\xfc\xb5\x2e\x67\x2d\xa0\x28\x5d\x6d\x43\x29\x39\x7a\x6e\xf3\x9e\x80\x6e\xe7\x8a\x62\x5d\xe7\x62\xeb\xcf\x39\xeb\x00\x25\xba\xfe\xf6\xb6\xa8\xfc\x4d\xf3\x86\xdf\xf2\x95\xeb\x10\xd8\xbd\x9b\xf0\x70\x69\x8b\x51\x40\x52\xba\xda\x6e\x9a\x23\x80\x1c\x3d\xd7\x1e\xe3\xea\x07\xbb\x00\x65\xcf\x1e\x02\xc4\x06\xd7\x03\xca\xb6\xc4\x22\x61\x4e\xed\x03\x2a\xb1\xc1\x6e\x40\x19\x22\xb1\x0e\xd1\xd9\xd0\xec\x53\x52\xd2\xa4\x64\x44\x8b\x04\x1a\x36\x02\xca\xf5\xaa\x3b\x0e\x54\xb4\xc8\xe0\x26\x40\x49\x46\x44\xa7\x07\x50\xdc\x23\x6c\x9d\x23\x72\x20\x18\x5d\x8c\x00\xd2\xc8\x19\x00\x9a\x9c\x8c\x69\x52\x93\x8a\x4e\x2d\x30\xab\x05\x76\x93\xb1\x41\x11\x90\xb5\xd8\x20\x01\x86\x92\x71\x2d\xbe\x2d\x1a\x28\x89\x6f\x8e\x3f\x9a\x04\xc8\x40\x57\xfc\xfc\xc0\x7a\x2e\x73\x7e\xa0\x5b\x8b\x9c\x1f\xd8\x10\x12\x1b\x35\xf9\xfc\xc0\xa6\x90\xe8\xd1\xe2\xe7\x07\x7a\x43\xa2\x4f\x93\xcf\x0f\xf6\xb1\x43\xdb\xcc\x6c\x44\x87\xd2\x9b\x3e\x83\xae\x66\x09\x2d\x3a\x98\x62\xee\xc6\x05\x2d\xea\x6c\x61\x07\x5c\x8f\xfc\xbe\x67\x3e\x59\x5e\x5e\x1b\x9e\xf7\x6b\x3c\x47\x9b\xe2\xfe\x0c\x5a\xe3\xfe\x28\x3e\x25\xee\xc9\x20\x40\xc9\x08\x8f\xdb\xf5\x3b\x2f\x73\x76\x21\x70\x76\x51\x02\xa4\x66\x7f\xa5\x66\x7f\x23\x43\x6b\xe3\x41\x58\x25\xae\x36\x44\xb5\x48\xb0\x0b\x55\x8b\x24\xe4\xc1\x6b\xd8\x54\xd4\xe4\xc6\x56\xe6\x96\x97\x97\x17\xfb\x83\xeb\x8b\x64\x5a\xc4\xcb\x61\x0e\x0c\x6e\x6d\xbe\x09\x4b\xdc\xe9\xda\x55\x6f\xf0\x12\xab\x64\xee\xf7\xae\xce\x23\x32\x9e\x37\xaf\xce\xc3\xf6\xb9\xc8\x19\x1b\x97\x46\x0b\x80\xf7\xeb\xc0\x10\xe9\x5a\x62\xa1\x5a\x64\x46\x37\xdf\x1c\xe4\xca\x52\x34\x48\x88\xb1\xdc\x97\xc7\x48\x98\xfd\xac\x96\x9c\xbc\x51\x4f\xeb\x3b\xd2\x3b\xb6\x8f\x30\x24\x02\x0b\xc0\x69\x11\xe8\x5f\x04\x3e\x10\x81\x57\x25\xa0\x3f\xe7\xbb\xa6\x5d\xf2\x18\xc7\x73\x31\xe0\x23\x00\xfd\x47\x72\xf8\x41\x3c\x28\x6b\xfd\x07\x8f\x64\x33\x00\x2e\xc4\x81\x0b\x04\xe8\x1f\xb3\x9c\x7c\x18\x2b\x01\x20\xb7\xad\x3b\x1f\x8f\xb3\x62\xf9\x31\xd9\xc1\x6a\x16\xb3\xbe\x31\xb8\x1b\x88\x86\x8f\x00\x56\xf6\xd8\x33\x43\x09\x71\xd2\xf4\x5f\xa7\x65\xd4\xeb\xd6\x53\x42\x30\xca\x78\x92\x9c\x16\x65\x7c\xc0\xfb\x8f\x70\x58\x5c\x83\xa7\x59\xca\xa0\x9f\x6c\x10\x64\xbc\x4d\x5e\x17\x64\x3c\xc0\xfb\x5b\x78\x7f\x07\xef\x7f\xce\xfb\x76\xde\x77\x49\x17\x05\x05\x0a\x28\x64\x9c\x90\x5e\x17\x54\x8c\x4a\xef\x0a\x32\x7e\x09\xb6\xba\x1e\x17\x05\x19\x47\x24\xd6\x67\x04\xd6\xff\x53\xbc\x28\x24\x50\x15\xc6\xa1\xe0\x92\xc8\xa4\x08\x28\xba\xf0\x2c\xc6\xb1\x11\xf7\x4b\x54\x94\x91\x11\x2f\xf2\x5a\xdf\xc7\xbd\x0c\xe2\xdd\x89\xfb\x05\x60\x3b\xa7\x96\x48\x27\x3e\x12\x5e\x90\x18\x25\x62\x17\xf6\xe0\x4f\xd8\x85\x29\xde\x7f\x97\xf7\xff\xc7\xfb\xef\xe0\xcf\xd0\xb0\x4b\x78\x0f\xbb\xf0\x55\xf1\xef\x98\x65\x4f\x12\x3c\x06\x55\xf8\x10\x04\x8b\x9c\x7a\xb8\x7b\xb7\xc8\x7c\x38\x13\xac\xc5\x76\x08\x1f\x33\x7b\x5b\x02\xce\xcd\xa2\x48\x08\x4e\x5d\x13\x50\x8a\xd8\x4e\x44\x1c\xea\x67\xd4\xd9\xee\xbd\xe2\x87\x88\x70\xff\x8e\x49\xec\xa4\xde\xe2\xf3\x7b\xf9\xfc\x8f\x88\x63\xb7\x48\xa0\x81\x69\xdf\x80\xdd\xa2\x82\xff\xc7\x6e\xb1\x13\xdb\x79\x3f\xc2\xfb\x51\xde\x67\x79\x7f\x2b\xef\xef\xe0\x7d\x0d\xfd\xd2\x3a\xb4\x63\xb7\x38\x80\x1a\x0e\x0b\x43\x58\x1b\xce\x1f\x16\xc6\x60\x70\x9d\x3d\xbc\x3f\x8b\xbb\x70\x17\x0c\x8c\x09\x79\x3c\x8e\x5e\xe9\x38\x7a\xf0\x63\xc1\xc6\xe3\x38\x27\x9d\xc6\xb3\x38\x83\x47\xf0\x38\xd6\x48\xdf\xc2\x83\x78\x09\x8c\xe7\xfb\xd2\x69\x8e\x3c\x86\x07\xf1\x13\xe4\xf1\x3c\x86\xa5\x27\x38\xf2\x3c\x9e\xc5\x1d\xc2\x4b\x78\x0d\x31\xe9\x11\xd4\x40\x85\x9f\xe1\x57\x58\xc2\x25\xbc\x89\xdf\x09\x6f\xe1\x2c\xce\x48\x6f\x43\x47\x1b\x7a\x88\x8e\x24\x06\x88\x8e\x4d\xc8\x10\x1d\xfd\x98\x22\x3a\x86\xf0\x10\xd1\xb1\x03\x67\x89\x8e\x9b\xf1\x24\xd1\x31\x8e\x1f\x12\x1d\x87\x39\x9e\xc3\x43\x64\x1a\x27\xf1\x0b\x88\xfc\x3c\x77\x0b\x49\x22\x40\x84\x88\x8d\x18\xc1\x51\x48\x4b\xf5\x93\xae\x37\xc2\xf2\x12\xc2\x0a\xf6\x0e\xfe\xc2\x47\xa9\x89\xef\x1d\xbc\x2b\xd6\x31\x34\x30\x21\xc0\x9a\x65\x7f\x2a\xb6\xca\xbe\x78\x05\x3e\xca\xf9\x6e\xc6\x8f\x02\x20\x3b\x47\x8d\xe2\x8c\x6d\xd5\x32\x66\xc1\x37\x1d\xdb\x70\x6b\xc7\x6e\xc0\x81\xaa\x5d\x38\xb6\x03\xb7\x4c\x39\xc5\xaa\x45\xbf\x84\xdc\x09\xc3\xa5\xc5\xb9\xaa\xed\x9b\x65\x3a\x5a\xa9\x64\x33\x18\x9d\x9d\x5d\x18\x1b\xcd\x4d\x20\x57\xf3\x7c\x5a\xd6\xb3\x33\x98\xcc\x8e\x2d\x64\xb2\x73\x13\xe3\xf3\x33\x73\xd9\x89\x1c\x0e\x52\x3f\x63\xf8\x46\x9d\x63\xdc\xb1\x2c\xca\x8d\x78\xfa\x41\x6a\x53\xd7\x2c\x60\xde\xad\x4d\x3a\x46\x11\x59\x6f\xd6\xf0\x4f\xcc\x39\x8e\x4f\x19\x71\x99\xbd\x29\xa7\x48\x71\xc0\xb4\x28\x4a\xd4\x5f\x98\x36\xca\xc1\x64\xbc\x6a\xf9\x55\x97\x72\x7a\xd4\xf3\x68\x39\x6f\xd5\x38\x61\x34\x13\x45\xcb\xb2\xd9\x18\xfa\x10\xaa\x44\x86\xe6\xab\xa5\x92\x91\xb7\xe8\xa8\xef\xbb\x66\xbe\xea\x37\x94\xcc\x9b\x7e\x33\x3c\x6f\xb8\x25\xea\x1f\x70\x8d\x32\xbd\xcf\x71\xef\x69\xe5\x67\xae\x1d\xa5\xae\x67\x3a\x76\xeb\x62\xd6\x3e\xee\xb8\x65\x83\x87\xd6\xfa\x8f\x5c\xe3\x8e\x7d\xdc\x2c\x55\x5d\xce\xd7\x58\x1e\x77\xca\x15\xd3\xe2\xe0\x1c\xb5\x8c\x53\x7c\xe6\xb5\x8a\xcf\xba\x4e\xb1\x5a\xf0\xaf\xa4\xb7\x5c\x31\xec\x5a\x63\x21\xdc\x3f\xc7\x7d\x33\x6f\x5a\xa6\xdf\xb4\x9a\xb5\x4d\xdf\x34\x2c\xf3\xfe\xd5\xf1\xd2\x43\xcf\x4d\xbb\xc4\x63\xef\xb9\xe5\x29\xa3\x52\x61\xe4\xbc\x13\x94\x7f\x18\xc5\xe2\xc2\x1c\xf5\x1c\xeb\x24\xa3\x0e\x52\xff\x40\xd5\xb2\xd8\xa9\x82\x9d\xf0\x01\xd7\x29\xaf\xb8\xcb\xc0\x9c\x6f\xb8\x7e\xb5\x72\xc8\x71\xee\xb9\x5c\xf4\x88\x5d\x36\x6c\xa3\x44\x8b\x19\xcb\x82\x9d\xb7\x1c\xa3\x48\x5d\xbd\x68\x59\xf0\xa8\xbf\x90\xf5\x69\xb9\x71\xe5\x74\x16\xfb\x80\x5a\xf1\x98\x1e\x0f\x6f\x19\xee\xa4\xae\xb3\x7a\x23\x93\x5c\x1d\xa6\xc7\xc2\x49\xdd\x00\x9a\xcd\x06\xbe\x50\x17\x53\x01\x54\x77\x7c\x05\xd7\x0b\xbe\xc3\x7a\x3e\x64\x6d\x7f\xd6\x77\xeb\x86\x32\xa6\x51\xb2\x1d\xcf\x37\x0b\x1e\x2a\xae\x93\xa7\xde\x6a\x17\xb2\xb6\x4f\x5d\xa7\x92\xa3\xee\x49\xb3\xd0\xba\x1c\x1c\x3a\x75\x57\xd6\x83\xbb\x6a\xda\x25\x96\x08\x2d\xec\x13\xa7\x7c\x6a\xb3\xb3\xf1\xd0\x48\xb0\x26\xb0\x35\xf9\x10\x1c\x58\xae\x62\x99\xfe\x4c\x25\x80\x26\x4e\x99\x9e\xcf\x34\xd8\x05\xc3\xc7\x4c\xfe\x6e\x5a\xf0\xc1\x39\xf8\x79\x67\xe8\x71\xa3\x6a\xf9\xec\x54\x8f\x1a\x56\x95\xce\xb8\x75\xa4\x1e\x1a\x16\xce\x71\xc7\xe6\xaf\x35\xa3\x95\x4a\x7d\x5a\xa8\x43\x21\xdb\x4a\x72\x62\xda\xf0\xcd\x93\x74\xd2\xcc\xbb\x86\x5b\xe3\x46\xc6\x0c\x8f\x66\x4c\x97\xb2\xa0\xd6\xe0\x54\x16\x26\xee\xad\x1a\xec\x82\xb2\x79\xd6\xa6\x75\x6a\x9a\xfa\x63\xd4\xa8\xfa\x35\x5e\xc8\xa6\x41\xe1\x63\x0c\x14\x06\xaa\xf0\x51\x43\x0e\x27\x60\xc0\x05\x45\x11\x73\xa8\xc2\x86\x0f\x13\x65\x50\x4c\xc1\x40\x05\x15\x98\xb0\x51\x62\xe5\x58\x1c\x06\x34\x0b\x0e\x0a\x30\x60\x81\xc2\xc3\x30\x10\xd7\x51\x84\xc5\x7e\xeb\x6c\x18\x86\x07\x17\x65\x2c\xc0\x86\xc1\xf5\x9c\x04\x65\x3c\x5b\x5a\xed\x4e\xc2\x44\x1e\x1e\x32\x30\xe1\x02\x3b\x3f\x9f\x67\x0e\x8a\xa0\x40\x6f\xab\xd4\x28\xf7\x39\x8b\x0c\x10\xb1\xe1\xf0\xb2\xfd\x5e\xff\x83\x33\xbf\x7e\x6a\x62\xfa\xe9\x11\xe9\xaf\x13\xcb\x13\xdf\x80\x94\x22\x24\x26\xa6\x40\x22\x29\x42\x34\x8d\x91\xaa\x04\x24\xf6\x2a\xeb\x12\x13\x62\x62\x6f\x62\x34\x31\x12\x49\x09\xa4\xa7\xbb\x23\x45\x48\x03\x8b\xf1\xe9\x88\xda\xdd\x56\x87\x47\xd4\xee\x68\x54\xec\x55\x63\xbd\xaa\xbc\x2e\xb1\x53\x50\x55\x09\xa4\x47\x15\x53\x50\xa3\x29\xa1\x57\x15\xb5\x5b\xa3\x29\x81\x24\x91\x24\x1d\x51\xa5\x57\x8d\xa9\x61\x8b\x40\x08\x98\x05\x55\x86\x18\x22\xaa\xaa\xae\xe9\x14\x84\x3e\xb2\x2e\x71\x97\xd0\x87\x3e\xd2\x07\x49\x11\xd8\x12\x51\x7b\x19\xb3\xaa\x46\x52\x24\x31\xa2\x2a\xd1\x68\xa8\xab\x5b\x66\x8a\x3a\xbb\x05\xb9\x5b\x04\xd4\xd8\x73\xfb\x16\x1e\xd0\x7e\xa3\xec\x11\x64\x55\x90\x85\x68\xe8\x94\x28\xf7\xaa\x22\x40\x62\x40\x00\xc4\x20\x24\x46\xd8\xa6\x64\x08\xdd\x89\x11\x35\x16\xab\x2f\x90\xf0\x55\xbc\x8f\x3d\x1b\xe7\x85\xae\xdb\x5c\xa3\x32\xcd\x92\xa4\x40\xf9\xf5\x9f\x3f\xe1\x3a\xf7\x79\x24\x46\xc2\x37\xf0\x9d\x04\xdd\xfa\xf4\xc4\xfc\xb8\xe3\xb2\x87\xdd\xf5\x61\xf1\xdb\x7b\x72\x87\x9e\x26\x98\x57\xd7\xae\x3c\x0c\x32\xa6\x57\xb1\x8c\xe0\x49\xd3\x41\x10\x5b\xa9\x26\x68\x27\x88\xce\x51\x8b\x1a\x1e\x0d\x88\xed\xfa\x76\x3d\xad\xa7\x01\x85\x20\xc2\x89\xd5\xef\xce\x9d\xab\xe8\x8b\xc3\x8d\xf9\xef\xeb\xdf\x3f\xae\xd0\xde\x18\x6e\xa6\x16\xc6\x1d\x37\x63\x59\x53\x86\x69\xa3\xec\x15\x1c\x97\x52\x5e\x3c\x59\x5b\x1e\x40\xaa\xc5\x4c\xf3\x2f\x12\xbe\xd8\x1d\x7c\x45\xb9\x0c\x67\xbf\x1f\xd2\x57\xc0\x11\x7e\x3b\xb8\x7d\x3f\xf0\x6a\xd3\xf7\x8b\x57\x85\x1b\x01\x1c\x45\x0e\x0b\x38\x8a\x09\xcc\x21\x87\x2c\x66\x30\x8d\x05\x64\x31\x8d\x03\xc1\x57\x17\xbc\x22\xbd\xff\x49\xa0\x9f\x34\x59\x02\xf6\x85\x7a\xa4\xd6\xcf\x22\x2c\x0f\x40\x70\x94\x67\xd2\x01\x98\x3c\x6f\xb3\xb0\x71\x3c\xcc\x8d\xad\x5c\x6a\x1e\x2e\x0c\xd8\xf0\x60\x85\x99\xeb\xc0\x0e\x35\x3c\x27\xa5\xd8\x8b\x07\x72\xf0\xe1\x86\xd5\xa0\x55\xd3\xbf\x08\xe3\x49\xaf\xfc\xdd\x88\x3c\x8b\x01\x6e\x40\x1c\x04\xe3\x70\x50\x46\x85\xdb\xa8\x61\x1a\x06\xcf\x63\xd6\x6c\xe4\xc1\x6a\x8a\xc1\x33\xda\xe5\xd8\x1e\x2e\x53\xb7\x91\xe1\x95\xa6\xc0\x6d\x57\x2e\xf3\xed\x4a\xb2\x69\xc4\x9a\x64\x8f\x72\xdc\x6b\x92\xd9\x0e\x9d\xff\xa7\xf9\x3f\xb3\xd5\x01\xc2\xf7\xe1\x73\x5e\x9b\xd7\xb6\x86\x87\xab\x6d\xac\x54\x3b\xf6\xee\xca\x7f\x2b\x4e\x82\xa2\xc4\xa5\xd8\x2e\x2b\xa8\x71\x4f\x4b\x38\x01\x3f\xbc\x86\x63\xdc\xc6\x4c\x88\x9b\xa1\x8d\xba\x8f\xf6\x67\xb2\x15\xc4\x71\x16\x2e\xaf\x7d\x55\x14\xe0\x7f\x6a\x1c\xd3\xfc\x7d\xea\x72\x99\xd5\x11\x69\xc4\x03\xb8\x89\xc7\x6e\x14\x1e\x3c\x50\x94\xb9\xc6\x1a\x52\x57\x91\x09\xe4\xfe\x6b\x5a\x2a\x78\xcf\xff\x60\xdf\x17\xed\xc8\xff\xda\x17\xd1\xfe\x1d\x00\x00\xff\xff\xda\x58\xd4\x71\x00\x18\x00\x00")

func nbloaderNbloaderDllBytes() ([]byte, error) {
	return bindataRead(
		_nbloaderNbloaderDll,
		"nbloader/nbloader.dll",
	)
}

func nbloaderNbloaderDll() (*asset, error) {
	bytes, err := nbloaderNbloaderDllBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "nbloader/nbloader.dll", size: 6144, mode: os.FileMode(438), modTime: time.Unix(1652960154, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"nbloader/nbloader.dll": nbloaderNbloaderDll,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"nbloader": &bintree{nil, map[string]*bintree{
		"nbloader.dll": &bintree{nbloaderNbloaderDll, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
