// Code generated by go-bindata.
// sources:
// template/bcp47.tmpl
// template/dive.tmpl
// template/gt.tmpl
// template/gte.tmpl
// template/hex.tmpl
// template/len.tmpl
// template/lt.tmpl
// template/lte.tmpl
// template/main.tmpl
// template/max.tmpl
// template/min.tmpl
// template/notnil.tmpl
// template/required.tmpl
// template/uuid.tmpl
// DO NOT EDIT!

package generator

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

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _templateBcp47Tmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\x92\x41\x6b\xdc\x30\x10\x85\xef\xfb\x2b\xa6\x02\x37\xf6\xb2\xb1\x2f\x81\x40\x4a\x2e\x2d\x2d\xf4\x52\x72\x28\xbd\x2c\x0b\x51\xac\x91\x2d\x2a\x4b\xae\x24\xd2\x06\xa1\xff\xde\x91\x9c\x6d\xbd\x65\xcb\x9e\x3c\x8c\xa4\xf7\xbe\x37\x9e\x18\xbb\xed\x06\xe0\xfd\x87\x87\x9b\xdb\x6f\x5c\x2b\xc1\x83\x75\x30\xa0\x41\xc7\x03\x7a\xe8\xad\x40\x08\x23\x0f\xf0\x53\x69\x0d\xcf\xe8\x94\x7c\x01\x25\x81\x83\x54\xa8\x05\x28\x4f\x65\x79\x4e\x77\xa7\x99\x07\xf5\xa4\x11\x7c\x70\xca\x0c\x24\x3c\x86\x30\xfb\xbb\xae\x0b\xd6\x6a\xdf\x2a\x0c\xb2\xb5\x6e\xe8\xc6\x30\xe9\xee\xa9\x9f\x6f\x6e\x37\xdb\x2e\xa5\x4d\x8c\x02\xa5\x32\x08\xac\x34\x59\x69\x65\x1b\xfc\x01\xed\xa7\x6c\xf4\xf5\x65\xa6\xd3\x45\x97\x01\xdb\x1e\x2b\xba\x99\xaf\x39\x07\x77\xf7\x99\xbb\x1f\xb1\xff\xde\x7e\xf6\x05\xa9\x8e\x91\x0e\x8d\x0d\x50\x2b\xff\x10\x1c\xb4\xd0\xa4\xf4\x36\x46\x34\x22\x25\xdf\xc6\xb8\x88\x7f\xe1\x13\xa6\xd4\xbc\x2b\x3a\x6f\xee\xc1\x28\x0d\x91\xe8\x89\x81\x0b\xf1\xd1\x39\x9b\x9f\x32\x3a\x2d\x86\x05\x0e\xb5\xc7\x33\x84\xfb\xc3\x1f\xc6\xfd\xe1\x48\x49\x6f\x24\x49\x28\x23\xf0\xd7\x0e\x3c\xb5\x34\x9e\x7a\x67\x7a\xc7\xcd\x40\xa3\xfb\x87\xaa\x80\x5c\x8a\xf8\x3f\x88\x55\xd8\x33\xa6\x97\x02\xd7\x33\x89\x04\x09\x8f\x72\x0a\x6d\x69\xca\x9a\x55\x95\xdf\x57\x95\x38\xc0\x35\x50\xc9\x76\xa7\xa9\x2a\xbf\xcb\x9a\xcd\x23\xfc\x35\x6a\xf2\xcc\x00\xd6\x73\x5b\xfe\xef\xeb\x9a\x29\x6b\x16\xc7\xa3\xdf\xb2\x04\x79\xb5\xf2\xaf\x7b\xce\x6b\x09\xd6\xbc\x2e\xdc\x55\xe5\xa1\xf2\x57\x6c\x65\xb0\xca\xde\x2c\xd2\x25\xf2\xf2\x85\x6b\xaa\x7e\x07\x00\x00\xff\xff\x2e\xc7\xf3\xb6\xe7\x02\x00\x00")

func templateBcp47TmplBytes() ([]byte, error) {
	return bindataRead(
		_templateBcp47Tmpl,
		"template/bcp47.tmpl",
	)
}

func templateBcp47Tmpl() (*asset, error) {
	bytes, err := templateBcp47TmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/bcp47.tmpl", size: 743, mode: os.FileMode(420), modTime: time.Unix(1478296916, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateDiveTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x64\x90\xcb\x4a\x34\x31\x10\x85\xf7\xfd\x14\xe7\x0f\x0c\xd3\x3d\x30\x3d\xfb\x5f\x66\xa7\x2e\x45\x50\xdc\x87\x4e\xb5\x16\xc6\xf4\x90\xc4\x11\x29\xf2\xee\x56\x77\xe6\x02\xba\x49\xea\x72\xaa\xce\x47\x89\xec\x36\x0d\xe0\xf8\x48\xf8\x62\xef\x31\x58\x7d\x5e\x29\x0c\x6f\x34\xbc\xe3\x68\x3d\x3b\x9b\x09\x53\x80\x45\xca\xf1\x73\xc8\x2a\xdf\xec\x4a\x69\x44\x1c\x8d\x1c\x08\x66\x9e\x36\x4b\x05\x3c\x62\x8a\x68\x39\x3d\x2d\x5a\xf4\xdd\x35\x79\xcc\x51\xf3\x8b\x8e\xd3\x52\x40\x29\x9a\xa4\x5e\xa4\xbf\x67\xf2\xee\xc1\x7e\x50\x29\xf8\xb7\x47\x60\x0f\x81\xc8\x16\x14\x9c\x8e\xa9\x8c\x62\xc4\xff\xfd\x85\xaf\x7f\x39\xf1\xb5\xbf\xe7\xbb\x9b\x45\x7b\xde\xa2\xcc\x6a\x6a\x9d\xbb\x8b\x71\x9a\x4d\x8d\x76\x8d\x5a\x37\x7f\x69\x8a\x48\xf5\xd3\x06\xf9\x44\xa8\xa1\x7a\x52\xb4\x99\xa7\x50\x77\xb4\x87\xc8\x21\x8f\x30\xb7\xf3\xed\x38\x21\x4c\xb9\xde\x6b\x3e\xd6\x38\xa3\x60\xbd\x4a\x58\xa5\xb5\xc1\x15\xed\x14\x3e\x7f\x1f\xa8\xab\x9b\xcf\x66\xfa\x63\xab\xd1\x4f\x00\x00\x00\xff\xff\xde\x64\x88\x1d\x93\x01\x00\x00")

func templateDiveTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateDiveTmpl,
		"template/dive.tmpl",
	)
}

func templateDiveTmpl() (*asset, error) {
	bytes, err := templateDiveTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/dive.tmpl", size: 403, mode: os.FileMode(420), modTime: time.Unix(1478296916, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateGtTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x92\xcf\x6a\xc3\x30\x0c\xc6\xef\x79\x0a\x55\xac\x60\x97\xd5\x0f\x30\xe8\xa1\x8c\xed\x58\x7a\xe8\x6e\x85\x92\x35\xf2\x30\xc4\x49\x67\x1b\xba\x61\xfc\xee\x93\x9d\xfe\x4b\x61\xb0\xcb\x2e\x45\x95\xf5\xe9\xd3\x4f\x51\x8c\xd0\x90\x36\x1d\x01\x7e\x04\x4c\xa9\xe2\x84\xd1\x40\x9f\x20\x82\x33\x76\xed\xf8\xf1\x0b\x70\x86\xa0\x5e\x0d\xb5\xcd\xe6\xfb\x40\x12\x30\x18\x4b\x6a\xc3\x3f\x08\xac\x01\x60\x55\x20\x7b\x68\xeb\x50\x1a\xed\x86\x27\x05\x43\x43\x6a\x3d\x9d\xc2\x6b\x95\x35\xdd\x4d\x45\xd7\x0c\x51\x0e\xe6\x1c\x55\xd5\x68\xb2\xdd\xc5\x2b\xc6\x87\x7a\xbf\x27\xef\x7b\x07\x4f\x0b\x10\x07\x67\xba\xa0\x01\xbd\x9a\xfa\xf3\x90\xab\xda\x92\xbc\xb0\x18\xbf\x0e\xae\x38\xf1\x9f\x1b\x79\x4a\x30\x59\x40\x67\x5a\x88\x9c\x9e\xe7\x19\x8a\x66\x7e\x5a\x00\x72\xbb\x75\xed\x6a\x9b\x6d\x43\x8c\xd7\xde\xac\x64\xeb\xb2\x83\x55\x7f\x14\x52\xbd\x6d\x9e\x85\xac\x58\x36\x19\x59\xa6\x24\x66\x23\x43\xc9\x80\xbc\x8b\x94\x46\xd9\x42\x9d\x92\x5a\xea\x40\x4e\xdc\x39\x49\x88\xc3\x7e\xeb\xa6\x79\x71\xae\xcf\x24\x17\x68\xca\x09\xaf\x56\x74\x14\x5b\x34\x1e\xde\x49\xf7\x8e\x60\xea\xb7\x28\x11\xb0\xeb\x8f\x28\xf3\xf4\x65\xb3\xc5\x98\x7b\xfd\x01\x45\x2d\x9b\x46\x70\x59\xc1\xe7\x19\x58\xf5\x8f\x70\xbf\xe2\x69\x1b\x54\x49\xea\x31\x5e\xe6\x7b\x84\x30\xf5\x72\xf4\xc5\x87\x53\xbc\xbd\xa2\xfb\x03\x38\x4f\x73\xba\xb8\x5c\xf2\x13\x00\x00\xff\xff\xc6\xd3\x10\x05\x01\x03\x00\x00")

func templateGtTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateGtTmpl,
		"template/gt.tmpl",
	)
}

func templateGtTmpl() (*asset, error) {
	bytes, err := templateGtTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/gt.tmpl", size: 769, mode: os.FileMode(420), modTime: time.Unix(1478895574, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateGteTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x52\xcd\x4a\x03\x31\x10\xbe\xef\x53\x8c\x83\x85\xa4\xd8\x3c\x80\xd0\x83\x8a\x1e\x4b\x0f\xf5\x56\x90\xb5\x99\x95\xc0\x66\xb7\x26\x81\x2a\x4b\xde\xdd\x49\xd2\x9f\xdd\x82\xe0\xc5\x4b\x99\x4e\xe6\xfb\x9b\x9d\x61\x00\x4d\x8d\xe9\x08\xf0\x23\x10\xc6\x58\x71\xc7\x34\x40\x9f\x20\x82\x33\x76\xed\xf8\xf5\x0b\x70\x8e\xa0\x5e\x0c\xb5\x7a\xf3\xbd\x27\x09\x18\x8c\x25\xb5\xe1\x1f\x04\xc6\x00\x30\x2a\x90\xdd\xb7\x75\x28\x4c\x6f\xe5\x4d\x41\x61\xa4\xd6\xd3\xb1\xbc\x8c\x59\xd3\x8d\x26\x3a\x5d\xaa\x54\x2c\xb8\xaa\xaa\xa9\xb7\xc2\x98\x47\x6e\xeb\xdd\x8e\xbc\xef\x1d\xdc\x2f\x41\xec\x9d\xe9\x42\x03\xe8\xd5\xcc\x9f\x5c\xae\x6a\x4b\xf2\x1c\xc6\xf8\x75\x70\x59\x89\xff\x8c\xe0\x31\xc2\xcd\x12\x3a\xd3\xc2\xc0\xed\x45\xf2\x90\x31\x8b\xe3\x06\x90\xe9\xd6\xb5\xab\x6d\x72\x16\x86\xe1\xc2\xcd\x48\x96\xce\x4b\x58\xf5\x07\x21\xd5\xeb\xe6\x49\xc8\x2a\xd3\x8f\x14\x63\x14\xf3\x89\x9e\xe4\x7c\xbc\x8a\x18\x27\xdd\x1c\x3a\x46\xf5\x48\x4d\xef\x48\x5c\x29\x49\x18\xca\x82\x6b\xad\x9f\x9d\xeb\x53\x92\x73\x68\x4a\x0d\xaf\x56\x74\x10\x5b\x34\x1e\xde\x33\x05\xcc\xfc\x16\x25\x02\x76\xfd\x01\x65\x72\x9f\x37\x9b\x95\x99\xeb\x0f\x51\xd4\x83\xd6\x82\xc7\x72\x7c\xf6\xc0\xa8\x7f\x0c\xf7\x6b\xbc\xc6\x06\x95\x9b\xcd\x34\x5e\xca\x77\x07\x61\xe6\xe5\xe4\x8b\x97\x5b\x1c\x5f\xd1\xf5\x01\x9c\xec\x8c\x46\x7e\x02\x00\x00\xff\xff\xe3\x75\x4e\x13\x03\x03\x00\x00")

func templateGteTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateGteTmpl,
		"template/gte.tmpl",
	)
}

func templateGteTmpl() (*asset, error) {
	bytes, err := templateGteTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/gte.tmpl", size: 771, mode: os.FileMode(420), modTime: time.Unix(1478895574, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateHexTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\x91\x41\x4b\x3b\x31\x10\xc5\xef\xfd\x14\xef\x1f\xc8\xbf\xbb\xd2\xee\x07\x50\x7a\x54\xf4\x22\x1e\xbc\x95\x42\x97\x66\xd2\x06\xb7\xd9\x9a\x04\xa9\x84\x7c\x77\x27\x59\xaa\x5b\xa9\x7a\xca\x30\x93\xbc\xf7\x7b\x99\x49\x8c\x8a\xb4\xb1\x04\xb1\xa3\xa3\x48\x89\x1b\x30\x1a\xf4\x8a\xe6\xce\x50\xa7\x9e\xdf\x0f\x3c\xf3\xc1\x19\xbb\x15\x10\x57\xa7\x8a\x6f\xe6\x6b\xce\xe1\x7a\x81\x2d\xd9\xcd\x8e\x36\x2f\xcd\x83\xbf\xa7\x63\x15\x23\x8f\x6c\x1f\x50\x19\xff\x14\x1c\x1a\xd4\x29\xfd\x8f\x91\xac\x4a\xc9\x37\x31\x0e\xd2\x8f\xed\x9e\x52\xaa\x6f\x8a\xca\xbf\x05\xac\xe9\x10\x27\x00\x13\xb4\x4a\xdd\x3a\xd7\xe7\xa7\x82\xa7\xc5\xae\xa0\x51\xe7\xe9\x02\xdf\x72\xf5\x49\xb8\x5c\x9d\x18\xf9\x8d\x66\x09\x63\x15\x1d\x67\xf0\xdc\xea\xe8\xdc\x3b\xb3\xbb\xd6\x6e\x09\xdf\xa9\x0a\xc8\xef\x01\x7f\x42\x18\x45\xbd\x60\xf9\x57\xdc\xea\xc0\x22\x41\x63\xad\xf7\xa1\x29\x4d\x5d\x09\x29\xfd\x52\x4a\xb5\xc2\x1c\x5c\x8a\xd9\x79\x26\xe9\x67\x59\xb3\x5e\xe3\xcb\xa8\xce\x3f\x06\x8c\x7f\x6d\xd8\x2d\x27\x21\xd7\x06\xd3\xdb\xc1\xf1\xe4\x97\xd7\x0f\xe3\xcb\xda\xde\xda\xce\x28\xf4\x16\x3a\xcb\x61\x2a\x3d\xa4\x9f\x8a\x91\xfc\x28\x79\x3d\x08\x97\xc0\xc3\x89\x39\x57\x1f\x01\x00\x00\xff\xff\x1b\x7e\x02\xe1\x59\x02\x00\x00")

func templateHexTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateHexTmpl,
		"template/hex.tmpl",
	)
}

func templateHexTmpl() (*asset, error) {
	bytes, err := templateHexTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/hex.tmpl", size: 601, mode: os.FileMode(420), modTime: time.Unix(1478296916, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateLenTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xbc\x53\xbd\x6e\xdc\x3c\x10\xec\xbf\xa7\x98\x8f\x80\x61\x9d\x01\xeb\x60\x27\x30\xdc\xb8\x70\x91\x74\x31\x5c\x24\x9d\x1b\x1e\xb9\x92\x88\x48\xa4\xc2\x9f\xd8\x82\x70\xef\x9e\x25\x79\x02\x0e\x46\x90\x22\x45\x9a\xdb\x11\x77\x6e\x34\x3b\x5c\xfd\xb7\xae\x9a\x3a\x63\x09\x62\x24\x2b\x8e\x47\x3e\xd8\x5f\xe1\x5b\x20\xc4\x81\x70\x48\x66\x8c\x30\x16\xdc\x44\x97\xac\x8a\xc6\x31\x70\xbe\x74\xe3\x32\x53\x60\x24\x23\x42\x9a\x67\xe7\x99\x1a\x5b\x5c\xed\x8b\x0c\x4c\x07\x26\x36\x30\xe1\x8b\x9c\xd1\x62\x57\xf0\xa3\xf7\x72\xa9\x4f\xf4\x03\xed\x67\x43\xa3\xfe\xca\x42\x10\x21\x7a\x63\x7b\xc1\x8d\x41\x86\x67\xcf\xb6\xde\x20\xd4\x20\xad\x38\xa3\xed\x58\x1b\x59\xfa\xff\x86\x3d\x35\xa1\x5d\xd7\xda\x7c\x92\x13\x1d\x8f\x3b\x3c\x3c\x80\x8f\x9e\xa5\x97\x53\x7e\x5c\x99\x0d\x3e\x81\xd4\xfa\x93\xf7\x6c\xa8\x85\xa0\x0c\x42\xfb\x44\xaf\xcd\x4b\x9e\xbb\x8f\x03\x26\x13\x26\x19\xd5\xf0\x22\x76\x02\xe5\x25\xa7\x2c\x1e\xb5\x66\x5b\x50\x03\xa9\xef\xa1\xcc\x6e\x6c\xa4\x9e\x3c\x7e\xca\x31\xe5\x00\x1c\x94\x33\x56\x19\x4d\x78\x35\x2c\x95\xb3\x99\x3d\x29\xd2\x64\x39\x1a\x8a\x38\x2c\xe8\xdd\xf5\x3c\xca\xa5\xf7\x2e\x59\xbd\xe7\xbf\x1a\x2d\x23\x8b\x6d\x69\xd1\xc8\x99\xf3\x5c\xef\x42\xe1\x77\x89\xf2\x7b\x5f\xcb\xcd\x5d\xad\x1f\x6e\x6b\xbd\xfb\xc8\x35\x55\x56\x3a\xd1\xd2\xc6\x4b\x1b\x31\x6d\xcc\xc3\x12\x89\x8b\x4f\x96\xf2\x65\x97\x1c\xdf\x67\xf8\xbb\x08\xff\x22\xc0\x3f\x4d\xd5\x8d\x4e\x66\xc7\xa5\x16\x87\x05\x15\x8b\xca\x4d\xf3\x48\x6f\xe7\xf8\xe6\xf6\xbe\x48\xfe\x1b\xbb\xf5\x3e\x7a\xb2\xe4\x65\xde\xf7\x2a\xd3\xcc\xbc\x9d\xb1\x2b\xdf\x09\xaf\x31\xac\x8b\x28\xb7\x88\xfc\x45\x64\x3b\xb8\xbc\x08\xb8\x08\x97\xdb\xba\x66\x7b\xe7\x9b\x7b\x12\x26\xab\x2b\xca\xe0\x9a\xd1\xaf\x00\x00\x00\xff\xff\x4e\x8b\xbc\x81\x84\x03\x00\x00")

func templateLenTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateLenTmpl,
		"template/len.tmpl",
	)
}

func templateLenTmpl() (*asset, error) {
	bytes, err := templateLenTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/len.tmpl", size: 900, mode: os.FileMode(420), modTime: time.Unix(1478296916, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateLtTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x92\xcd\x6a\x2b\x31\x0c\x85\xf7\xf3\x14\x8a\xb9\x01\x3b\xdc\xf8\x01\x0a\x59\xb4\xa5\x5d\x86\x2c\xd2\x5d\x20\x98\x58\x06\xc3\x78\x26\xb5\x0d\x49\x31\x7e\xf7\xca\x9e\xfc\xcc\x04\x0a\xdd\x74\x13\x14\x59\x47\x47\x9f\x46\x29\x81\x46\x63\x3b\x04\xd6\x46\x96\x73\x43\x09\x6b\x00\x3f\x81\x47\x6f\xdd\xc6\xd3\xe3\x19\xd8\x82\x81\x7c\xb7\xd8\xea\xed\xd7\x11\x05\xb0\x68\x1d\xca\x2d\xfd\x30\x20\x0d\x00\xa9\x22\xba\x63\xab\x62\x6d\xb4\x1f\x9e\x24\x0c\x0d\xb1\x0d\x78\x09\xef\x55\x4e\x9d\x47\x15\x9d\x1e\xa2\x12\x2c\x29\x6a\x9a\xc9\x64\xfb\x9b\x57\x4a\xff\xd4\xe1\x80\x21\xf4\x1e\x9e\x56\xc0\x8f\xde\x76\xd1\x00\x0b\x72\x1e\xae\x43\xae\x95\x43\x71\x63\xb1\x61\x13\x7d\x75\xa2\x3f\x23\x79\xce\x30\x5b\x41\x67\x5b\x48\x94\x5e\x96\x19\xaa\x66\x79\x59\x00\xa3\x76\x1b\xe5\x95\x2b\xb6\x31\xa5\x7b\x6f\x52\x92\x75\xdd\xc1\xba\x3f\x71\x21\x3f\xb6\xaf\x5c\x34\x24\x9b\x4d\x2c\x73\xe6\x8b\x89\xa1\x20\x40\xda\x45\xce\x93\x6c\xa5\xce\x59\xbe\xa0\xe9\x3d\xf2\x07\x2b\x01\x69\x58\xb0\xd2\xfa\xcd\xfb\xbe\xa0\xdc\xa8\xb1\x24\x82\x5c\xe3\x89\xef\x98\x0d\xa0\x4c\x44\x0f\xf3\xb0\x63\x82\x01\xeb\xfa\x13\x13\x65\xfa\xba\xd9\x6a\x4c\xad\x7e\x81\x22\x9f\xb5\xe6\x54\x56\xf1\x69\x04\x52\xfd\x25\xdc\x8f\x78\xc6\x45\x59\x93\x66\x82\x57\xf8\xfe\x43\x9c\x07\x31\xf9\xe2\xc3\x29\x8e\xaf\xe8\xf1\x00\xae\xd3\x5c\x2e\xae\x94\x7c\x07\x00\x00\xff\xff\x1e\x40\xd9\xc0\x01\x03\x00\x00")

func templateLtTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateLtTmpl,
		"template/lt.tmpl",
	)
}

func templateLtTmpl() (*asset, error) {
	bytes, err := templateLtTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/lt.tmpl", size: 769, mode: os.FileMode(420), modTime: time.Unix(1478895574, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateLteTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x92\xcf\x6a\xc3\x30\x0c\xc6\xef\x79\x0a\x4d\xac\x60\x97\xd5\x0f\x30\xe8\xa1\x8c\xed\x58\x7a\xe8\x6e\x85\x61\x6a\x05\x0c\x71\xd2\xd9\x86\x76\x18\xbf\xfb\x14\xa7\x7f\x92\x8e\xc1\x2e\xbb\xb4\x8e\x2c\x7d\xfa\x7e\xb2\x52\x02\x43\xb5\x6d\x09\xb0\x89\x84\x39\x57\x1c\xb1\x35\xd0\x27\x88\xe8\xad\xdb\x78\xbe\x3d\x01\xce\x11\xd4\x9b\xa5\xc6\x6c\xbf\x0e\x24\x01\xa3\x75\xa4\xb6\xfc\x83\xc0\x35\x00\x29\x2d\x20\x92\x3b\x34\x3a\x0e\x52\x1f\xc3\xa5\x82\x41\x92\x9a\x40\xe7\xe3\x2d\xcd\xe9\xd3\x35\x83\x5a\x73\xf9\x87\x05\x9f\xaa\x6a\x6a\x6d\xd0\x2b\x29\x8f\x7a\xbf\xa7\x10\x3a\x0f\xcf\x4b\x10\x07\x6f\xdb\x58\x03\x06\x35\x0b\x17\x93\x6b\xed\x48\x5e\x59\x6c\xd8\x44\x5f\xfa\xf0\xc7\xa8\x3c\x67\x78\x58\x42\x6b\x1b\x48\xc5\xff\xc5\xc2\xe2\x3c\x00\x64\xb9\x8d\xf6\xda\xf5\x0e\x63\x4a\x37\x6d\xae\xe4\xd6\x65\x06\xeb\xee\x28\xa4\x7a\xdf\xbe\x08\x59\x15\xf9\x51\xc7\x9c\xc5\x7c\xd2\x4f\x32\x1f\x0f\x22\xe7\x49\xf4\x0c\xaf\x56\x75\x24\x2f\xee\x1a\x49\x48\x65\xbc\xa0\x8d\x79\xf5\xbe\xeb\x41\xae\xcc\xd4\x07\x82\x5a\xd3\x51\xec\xd0\x06\xd0\xbd\x02\xcc\xc2\x0e\x25\x02\xb6\xdd\x11\x65\xef\xbd\xcc\xb5\xf4\x65\xa9\x3f\x80\xa8\x95\x31\x82\xd3\x0a\x3c\x5b\xe0\xaa\xff\x43\xfb\x15\xae\x76\x51\x95\x60\x3d\x81\xeb\xe9\x9e\x20\xce\x82\x9c\xbc\xf6\xb0\x86\xe3\x0d\xba\x7f\xfc\xfc\x73\xc9\xbe\x03\x00\x00\xff\xff\xf4\x50\xe7\xe6\xfe\x02\x00\x00")

func templateLteTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateLteTmpl,
		"template/lte.tmpl",
	)
}

func templateLteTmpl() (*asset, error) {
	bytes, err := templateLteTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/lte.tmpl", size: 766, mode: os.FileMode(420), modTime: time.Unix(1478895574, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateMainTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\x53\xcd\x6e\xdb\x3c\x10\xbc\xeb\x29\x06\x86\x81\xcf\x0e\x62\x29\xe7\x00\x5f\x0f\x6d\x93\x22\x87\xa4\x40\x5b\xe4\x4e\x4b\x2b\x89\x88\x44\x0a\x24\xe5\x34\x10\xf4\xee\x5d\xfe\xd8\x32\x8c\xf8\x22\x50\xdc\xdd\xd9\xd9\xd9\xe1\x34\xed\x50\x51\x2d\x15\x61\xd5\x92\xa8\xc8\xac\xe6\x39\x1b\x44\xf9\x26\x1a\xc2\x34\xe5\xe9\xc8\x97\x99\xec\x07\x6d\x1c\x36\xd9\xaa\x91\xae\x1d\xf7\x79\xa9\xfb\x42\xec\x65\x49\x45\x43\xaa\x6c\xa9\x7c\x5b\x65\xdb\x6c\x9a\x48\x55\xd8\xf9\x8a\x73\x74\xeb\xcc\x58\x3a\x8f\x3e\x4d\x6b\x33\x76\xf4\x2c\x06\xdc\xff\x8f\xdc\x9f\x6d\xc8\x2f\x0a\xbc\x8a\x4e\x56\xc2\x11\xa4\x85\x50\x10\xa3\xd3\xbd\x70\xb2\x14\x5d\xf7\x01\xee\x42\x86\x83\x15\x0e\x31\x4d\x6a\x85\x9e\x5c\xab\x2b\x0c\x46\x1f\x64\xc5\xa1\xfd\x87\xc7\x39\x12\xca\xfd\xcf\x6f\x22\xb4\xce\x0d\xf6\xbe\x28\xae\x52\x47\xad\x0d\x7a\x6d\x88\x19\x3b\x21\x3b\x9b\x67\xf5\xa8\x4a\x6c\x2c\xcb\x20\x6b\xe4\x83\x33\xcf\xb1\xd7\x3c\xdf\x84\x21\xe7\x99\x05\x52\xa2\x67\x75\xb6\x0b\xf3\xcd\x16\x1b\x32\x46\x9b\x2d\xa6\x0c\x5c\x5c\xdc\xe0\xa7\x62\xfa\xbd\x78\x23\xb8\x96\x10\xa2\x96\xff\x07\x30\xf0\x3b\xfd\xc7\x4d\x1b\x2d\x55\x03\xa7\x31\x5a\x1e\xde\xe5\xb8\x29\x58\x11\x5f\xbe\xf3\x49\x0d\xeb\xde\x91\x4a\x6a\x6d\x71\x17\x04\x43\xa2\xd6\x74\x7a\x2f\xba\x47\x66\xfd\x28\xac\x0b\x81\x83\x30\x38\x3c\xc4\x46\x27\x31\x5e\x4f\xb2\xc5\x48\x00\xa0\xce\x52\x2c\x49\xe9\xbc\x14\x4f\x75\x73\xb5\xec\x16\x77\xb7\x5c\xb8\xf0\xe1\xf9\x23\x94\xd7\x24\x9c\xe0\x3d\x90\xce\x3c\xff\x93\x0b\x8b\xe3\xe9\x8d\x1e\x9b\x36\xa8\xe0\x2b\x83\x06\x3c\x74\x69\x28\x86\xe3\xb5\x5d\xa6\x37\x42\xb1\x11\xd7\x52\x55\xf4\xf7\x16\xeb\x5a\x52\x57\x5d\xd8\x06\x08\x1d\xd7\x7e\x13\x3e\x14\x93\xf2\x97\xb0\x98\x53\x34\xe6\x2f\xe1\x5f\x97\xf5\x67\x2a\xaf\x2f\x54\x06\xd8\x45\x5f\x1f\x7e\x3c\xbd\x2c\x8d\xe6\x19\x8b\x2e\x36\x81\xec\x90\xf8\x9a\xa7\x44\x38\x4c\xe9\xbb\xc6\xfe\x09\x2e\x00\x1e\x59\x05\xa6\xa7\x4e\x11\xe6\x1b\x1b\xfe\x0f\xf5\x43\xe7\x65\x89\x20\xf1\x7b\xe2\xbb\xc3\x51\xed\x00\xf6\xf0\xf2\xfd\x2a\xb7\xe3\x84\x67\x3b\x59\xaa\x17\x7f\xc6\x67\xf0\x99\x41\xd1\x8a\x43\xda\xcc\x99\x31\x3f\xf7\xe5\x05\xa8\xa8\xaa\x04\x7c\xb2\xba\xd2\x0e\xa2\x7b\x17\x1f\x16\x35\x5b\xd6\xfb\xbe\x66\xdb\x9e\x21\x73\xaa\x4f\xfa\xcc\xd7\x1c\xe2\x7e\x9b\xe4\xd5\x2d\xbe\xdc\x85\x57\x06\x18\x72\xa3\x51\x47\x13\xf3\xd5\x7c\x61\xc9\x65\xe4\x94\xaa\x64\x97\xcd\xd9\x31\xe5\x5f\x00\x00\x00\xff\xff\xc9\x8f\x9e\x25\x0a\x05\x00\x00")

func templateMainTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateMainTmpl,
		"template/main.tmpl",
	)
}

func templateMainTmpl() (*asset, error) {
	bytes, err := templateMainTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/main.tmpl", size: 1290, mode: os.FileMode(420), modTime: time.Unix(1478537994, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateMaxTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xdc\x53\x4d\x6f\xd3\x40\x10\xbd\xf7\x57\x3c\x56\x4a\x9b\x44\xc2\x55\x0b\xaa\x7a\x09\x12\x42\x54\xe2\x40\x95\x03\xdc\x2a\xa1\x6d\x3c\x4e\x56\xd8\x6b\xb3\x1f\x90\x68\xb5\xff\x9d\xd9\x5d\x1b\xa5\x85\x03\xb9\xf6\xe2\x99\xac\x9f\xdf\xbc\x79\xfb\x12\x42\x4d\x8d\xd2\x04\xd1\xc9\xbd\x88\xf1\x2c\x84\xcb\x25\xbe\x5a\x82\xdb\x11\x1e\xbd\x6a\x1d\x94\x46\x4b\x1a\x8d\xd7\x1b\xa7\x7a\x6e\x7a\x93\xdf\xba\xc3\x40\x96\x3b\xe9\x60\xfd\x30\xf4\x86\xa1\xae\xc2\xf2\x32\xd3\x40\x35\x60\xe0\x1c\xca\x7e\x96\x03\x2a\x2c\x72\xff\xde\x18\x79\x28\xbf\xe8\x07\xaa\x3b\x45\x6d\xfd\x85\x89\x20\xac\x33\x4a\x6f\x05\xbf\xd8\x49\xbb\x36\x2c\x6b\x0f\xb1\xd9\x49\x2d\x8e\x60\x0b\xe6\x46\xa2\x66\x45\x73\x5b\x85\x50\x5e\xdd\xcb\x8e\x62\x5c\xe0\x5d\x99\x9b\x98\xd3\x11\x44\xeb\x04\x62\x5c\x85\x40\xba\x8e\x11\x8c\x5f\x4b\x23\xbb\xd4\x32\x0f\xf8\x00\xb2\xae\x3f\x1a\xc3\x52\x2b\xcc\x07\x96\xe0\x1a\x08\x4a\x07\xb6\xba\xa7\x5f\xf3\x07\xc1\xa3\xb6\x6e\x87\x46\xaa\x96\x6a\x6c\x76\xb4\xf9\x9e\x3d\x98\xd9\xd5\xcc\x3e\x88\x85\x18\x87\x15\xea\x05\xb2\xc4\xd1\xc9\x0f\x09\x6d\x33\x9c\x99\x69\x4b\x06\x3f\x65\xeb\xd9\xb7\xc9\x26\x6a\xd9\xec\xa2\x79\xce\x0e\x74\xd3\xe6\xcb\x27\x6b\x43\xf0\xe7\x22\x3f\x6f\x4b\xb9\xba\x29\xf5\xcd\x75\xa9\x37\x6f\xb9\xfa\x82\xf2\x23\xcc\x4f\x38\x3f\x01\xfd\x84\x7c\x3c\x38\xe2\x62\xbc\x26\x91\x05\xb3\x14\x47\xdd\xd0\x4a\x57\xc2\xf0\xed\x53\xa2\xaa\x70\x9a\xca\xa6\xed\x65\x52\x90\x6b\x9e\x98\xbb\x3c\x72\xd3\x33\x3d\xed\x8f\xfb\xab\xeb\x5b\x81\x7f\x8f\xbf\x2b\x4c\x4f\x04\x4c\xc8\x2d\x69\x32\x32\x85\xb1\xdc\xdc\x9f\x7b\xe3\x0f\x39\x63\xd0\xbd\x4b\x36\xab\x1a\x29\xae\x49\x1e\x2e\x66\x96\x2f\xec\x62\x92\x5b\x2e\xec\x48\xf9\x38\x45\xd7\xa5\x4b\xcd\x6b\xee\xce\xd2\xe9\xd1\x5f\x24\xbb\xc2\xe7\x6c\x45\x08\xfc\x50\x76\xed\x38\x39\x31\x3e\xcf\x22\x5e\xad\xa0\x55\x8b\xf3\x73\x2c\xc7\xf4\xfd\x05\x39\x31\xad\xff\x93\xd5\x53\x42\x3a\xee\x99\xb6\x7c\xb6\x64\xf1\xfe\xc5\xad\xf9\x3b\x00\x00\xff\xff\xd1\x19\xf5\xe0\xed\x04\x00\x00")

func templateMaxTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateMaxTmpl,
		"template/max.tmpl",
	)
}

func templateMaxTmpl() (*asset, error) {
	bytes, err := templateMaxTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/max.tmpl", size: 1261, mode: os.FileMode(420), modTime: time.Unix(1478895574, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateMinTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xdc\x53\x3d\x6f\xdb\x30\x10\xdd\xfb\x2b\x5e\x09\x24\xb1\x0d\x44\x46\xd2\x22\xc8\x50\x0f\x19\x1a\xa0\x43\x03\x0f\xed\x16\xa0\xa0\xc5\x93\x4c\x54\xa6\x54\x7e\x34\x11\x04\xfd\xf7\x1e\x49\xa9\x70\xda\x0c\xf5\xda\x45\x77\x3e\x1e\xdf\x3b\xbe\x7b\x7e\x33\x0c\x8a\x2a\x6d\x08\xe2\xa0\x8d\x18\x47\x2e\xac\x57\xf8\xea\x08\x7e\x4f\xd8\x05\xdd\x78\x68\x83\x86\x0c\xaa\x60\x4a\xaf\x5b\x4e\x5a\x9b\x4e\x7d\xdf\x91\xe3\x4c\x7a\xb8\xd0\x75\xad\xe5\x56\x5f\x60\xb5\x4e\x30\xd0\x15\xb8\x71\x01\xed\x3e\xcb\x0e\x05\x96\x29\xbf\xb3\x56\xf6\xf9\x17\xfd\x40\x71\xaf\xa9\x51\x5f\x18\x08\xc2\x79\xab\x4d\x2d\xf8\x60\x2f\xdd\xd6\xf2\x58\xcf\x10\xe5\x5e\x1a\x71\xd4\xb6\x64\x6c\x44\x68\x9e\x68\xe1\x8a\x61\xc8\x47\x0f\xf2\x40\xe3\xb8\xc4\x87\xcc\x1b\x91\x63\x09\xa2\xf6\x02\xe3\xb8\x19\x06\x32\x6a\x1c\xc1\xfd\x5b\x69\xe5\x21\xa6\x8c\x03\x2e\x40\x2a\xf5\xd1\x5a\x1e\xb5\xc0\xa2\xe3\x11\x7c\x05\x41\xb1\xe0\x8a\x07\x7a\x5a\x3c\x0a\xa6\xaa\xfd\x1e\x95\xd4\x0d\x29\x94\x7b\x2a\xbf\x27\x0d\xce\xdc\xe6\xcc\x3d\x8a\xa5\x98\xc8\x32\xf4\x12\x69\xc4\x49\xc9\x3b\xa5\xf8\x51\xf9\x92\x4b\xb7\x98\x80\x6a\xb2\xf8\x29\x9b\x10\xe5\x6b\x51\xb6\xda\x94\x5a\x11\x9e\x34\xd3\x44\x65\x3b\x4b\x25\x29\x32\x2c\x2c\x79\xec\x7a\xd4\xed\x65\xd7\xc8\xbe\xb6\x6d\x30\x6a\xcd\x57\xb5\x92\x9e\xc1\x66\xad\xa9\xe1\x8d\xe5\x87\x2f\x58\xc6\xc3\x2c\xdf\xea\x85\x76\x10\x4c\x2e\xd2\xf7\x36\x87\xab\x9b\x1c\xdf\x5d\xe7\x78\xf3\x9e\x63\xc8\x5d\x61\x6a\x0b\x73\x5f\x98\x1b\xc3\xdc\xb9\xeb\x3d\x71\xb0\xc1\x50\xf6\x0e\x3c\x1d\x78\x50\x9f\xfd\xf4\xed\x53\x04\x2a\x70\xda\x8c\x55\xd3\xca\xc8\x9f\x62\xe2\x4b\x59\x22\x2c\x5b\x86\xa7\xe7\xe3\xfc\xea\xfa\x56\xe0\x35\xf2\xfb\x8c\xf3\x82\x3e\xa7\x35\x19\xb2\x32\x7a\x39\x2f\xfe\xf7\xda\xf9\x1a\x5b\x14\xa6\xf5\x48\x1a\x23\xba\x3d\x8e\x86\x8b\x33\xc7\xfb\xbe\x98\x47\xcd\xfb\x3e\x9a\x7a\xe2\x30\x2a\x67\x31\xb9\xe4\x2c\x16\x8f\xfe\x60\x49\x10\x2e\xb3\x0a\xc3\xc0\x1f\xed\xb6\x9e\x7d\x37\x8e\x7f\x3a\x19\x6f\x37\x30\xba\xc1\xf9\x39\x56\x93\x77\xff\x6a\x39\xd1\xeb\xff\xe2\xf4\x53\x2c\x3e\x3d\xf3\x95\x47\x66\xe1\xff\xbb\x67\xfe\x0a\x00\x00\xff\xff\x54\x3f\x19\xe1\x2c\x05\x00\x00")

func templateMinTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateMinTmpl,
		"template/min.tmpl",
	)
}

func templateMinTmpl() (*asset, error) {
	bytes, err := templateMinTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/min.tmpl", size: 1324, mode: os.FileMode(420), modTime: time.Unix(1478895574, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateNotnilTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x44\x8e\xb1\x6a\x43\x31\x0c\x45\xf7\x7e\xc5\xc5\x10\x92\x0c\x7d\x7f\x90\xb1\x1d\x3d\x75\xcc\xe2\x62\xb9\x08\x54\x39\xd8\xaf\x2d\x45\xf8\xdf\x2b\x27\x79\xcd\x60\x38\x96\xb8\xba\xc7\x2c\x53\x61\x25\x04\xad\xab\xb2\x84\x31\x9e\x00\x33\x70\x01\xf7\xf8\x25\x92\xde\x85\xb0\xe0\x3a\xf7\x61\x5f\xcc\x96\x57\x26\xc9\x31\x7d\xd2\x18\x38\x9d\xe0\x39\x98\xaf\xaf\xc1\x94\xf3\x4b\x6b\xb5\x79\x26\xd0\x84\xbe\x44\xfa\x39\x9c\x03\x77\x44\x96\x73\x38\x86\xdb\xb1\x7b\x11\x49\x27\x6c\xad\x1f\xa4\xd4\xd2\xca\x55\x6f\x37\x0e\x97\xc6\xba\x96\xcd\xce\x95\xe0\x84\xef\x24\x9c\x51\xd5\x3f\xfe\x36\xc9\x32\xad\xb0\xdf\x75\xec\xfa\x3e\xe0\x61\x79\xc7\xb7\xdf\x0b\x1d\xff\xab\x48\xf3\x64\xb3\x09\xcf\x4e\x7f\x01\x00\x00\xff\xff\x8d\xc2\xfb\xfb\x0b\x01\x00\x00")

func templateNotnilTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateNotnilTmpl,
		"template/notnil.tmpl",
	)
}

func templateNotnilTmpl() (*asset, error) {
	bytes, err := templateNotnilTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/notnil.tmpl", size: 267, mode: os.FileMode(420), modTime: time.Unix(1478296916, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateRequiredTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xcc\x92\x4f\x4f\x02\x31\x10\xc5\xef\x7e\x8a\x49\x4f\x7a\x70\x13\xd0\x10\x2e\x1c\xf5\xc8\x41\x3d\x72\x29\x74\x20\x4d\x4a\xbb\x4c\x3b\x28\x36\xfb\xdd\xed\x9f\x5d\x62\x8c\x22\xd1\x98\x78\xe1\x4d\xcb\x9b\x79\xbf\x74\x27\x46\x85\x6b\x6d\x11\x04\xe1\x8e\x35\xa1\x12\x5d\x77\x01\x10\x23\xe8\x35\x68\x3f\x67\x63\xe4\xd2\x20\x34\x50\xee\xd3\xa5\x6f\x62\x6c\xee\x35\x1a\x35\x97\x5b\xec\x3a\x98\xcd\xc0\x6a\x03\x31\xfd\x5d\x1a\xa5\x52\x77\x44\x8e\x52\x8f\xc0\x5c\xf8\x66\x8e\xcf\x97\x0b\xa1\x3d\x0c\x29\x0b\x71\x25\xea\xc4\x3e\x0d\x8d\xc7\x3c\x1d\x77\x50\x87\x3f\x1d\xda\x44\xe5\x03\x69\xbb\x11\x27\xc3\x85\xf8\x9b\x6c\x6d\x83\x28\xbf\xd3\x2a\xa3\x49\xd5\x9b\x71\xd5\xc9\x6d\x52\xae\x2e\xee\x6d\x3c\xf8\x78\x30\xf2\xe0\x5c\x1e\x02\x26\x21\xb6\x78\x7c\xe3\x0d\x5a\x24\x19\xb4\xb3\x95\x5a\x3c\xf4\x8c\xe0\x2c\xa4\x46\xdc\x20\xc1\x5e\x1a\x46\x9f\x3e\x06\x58\x17\xc0\x73\xdb\x3a\x0a\xa8\x9a\xe3\x94\x2f\xf8\xd7\xc6\xc9\xcc\x56\xb4\xb0\x94\xaa\xc0\xac\xdc\xb6\x35\xf8\xf2\xbe\x1e\x8d\xa7\xfd\xb3\x7c\x07\x66\x79\x8b\xa4\x57\xd2\xfc\x18\x6d\xe9\x9c\x39\x33\x2c\x5b\x51\xda\xf3\xa3\xb4\x7f\x0c\xc4\xab\x30\xac\xec\x5e\x12\xbc\x22\xb9\x0f\x8b\x33\x1c\x33\xd0\xa9\xed\xfa\xac\xf5\xf7\xdb\xf6\x3f\xc9\xac\xca\xe7\x18\x73\x71\x9d\xaa\xb7\x00\x00\x00\xff\xff\x08\xdc\x4c\x18\x1e\x04\x00\x00")

func templateRequiredTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateRequiredTmpl,
		"template/required.tmpl",
	)
}

func templateRequiredTmpl() (*asset, error) {
	bytes, err := templateRequiredTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/required.tmpl", size: 1054, mode: os.FileMode(420), modTime: time.Unix(1478296916, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateUuidTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\x92\x51\x4f\xc2\x30\x14\x85\xdf\xf9\x15\xc7\x26\x95\xcd\xc0\x7e\x80\x86\x37\x35\xe1\xc5\xf8\x20\x4f\x84\x84\x85\xde\x62\xe3\xe8\xb0\x1d\x46\xd3\xf4\xbf\x7b\xdb\x05\x1d\x06\xc3\xd3\xee\xee\x6d\xcf\xf9\x4e\xdb\x51\x08\x8a\xb4\xb1\x04\x71\x38\x18\x25\x62\xe4\x0e\x8c\x06\xbd\xa3\x7a\x34\xd4\xa8\x97\xaf\x3d\x0f\x7d\xe7\x8c\xdd\x0a\x88\x9b\x63\xc5\x2b\xd3\x32\xe7\x70\x3b\xc3\x96\xec\xe6\x95\x36\x6f\xd5\xdc\x2f\x16\xf3\xfb\x22\x04\x9e\xd9\xb6\x43\x61\xfc\x73\xe7\x50\xa1\x8c\xf1\x3a\x04\xb2\x2a\x46\x5f\x85\xd0\x6b\x3f\xd5\x3b\x8a\xb1\xbc\xcb\x32\x57\x33\x58\xd3\x20\x8c\x00\x46\xa8\x95\x7a\x70\xae\x4d\x5b\x05\x4f\xb3\x5f\x66\xa3\xc6\xd3\x19\xc0\xe5\xea\x07\x71\xb9\x3a\x42\xf2\x1e\xcd\x12\xc6\x2a\xfa\x9c\xc0\x73\xab\xa1\x53\xef\x04\xef\x6a\xbb\x25\xfc\xa5\xca\x20\x17\x12\xfe\xc7\x30\xc8\x7a\xc6\xf3\x52\xde\x62\xcf\x22\x9d\xc6\x5a\xef\xba\x2a\x37\x75\x21\xa4\xf4\x4b\x29\xd5\x0a\x53\x70\x29\x26\xa7\xa1\xa4\x9f\x24\xcd\x72\x8d\x5f\xa3\x32\x1d\x19\x30\x3c\xb6\xfe\x76\x39\x0a\xb9\xba\x33\xad\xed\x1d\x8f\x7e\xf9\x05\xc0\xf8\x7c\x71\x1f\x75\xc3\x3f\xad\x85\x4e\x7a\x18\x4b\x0f\xe9\xc7\x62\xa0\x3f\x88\x5e\xf6\xca\x39\x71\xff\xc5\x94\xab\xef\x00\x00\x00\xff\xff\xb3\x04\xf3\xf9\x5d\x02\x00\x00")

func templateUuidTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateUuidTmpl,
		"template/uuid.tmpl",
	)
}

func templateUuidTmpl() (*asset, error) {
	bytes, err := templateUuidTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/uuid.tmpl", size: 605, mode: os.FileMode(420), modTime: time.Unix(1478296916, 0)}
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
	"template/bcp47.tmpl":    templateBcp47Tmpl,
	"template/dive.tmpl":     templateDiveTmpl,
	"template/gt.tmpl":       templateGtTmpl,
	"template/gte.tmpl":      templateGteTmpl,
	"template/hex.tmpl":      templateHexTmpl,
	"template/len.tmpl":      templateLenTmpl,
	"template/lt.tmpl":       templateLtTmpl,
	"template/lte.tmpl":      templateLteTmpl,
	"template/main.tmpl":     templateMainTmpl,
	"template/max.tmpl":      templateMaxTmpl,
	"template/min.tmpl":      templateMinTmpl,
	"template/notnil.tmpl":   templateNotnilTmpl,
	"template/required.tmpl": templateRequiredTmpl,
	"template/uuid.tmpl":     templateUuidTmpl,
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
	"template": {nil, map[string]*bintree{
		"bcp47.tmpl":    {templateBcp47Tmpl, map[string]*bintree{}},
		"dive.tmpl":     {templateDiveTmpl, map[string]*bintree{}},
		"gt.tmpl":       {templateGtTmpl, map[string]*bintree{}},
		"gte.tmpl":      {templateGteTmpl, map[string]*bintree{}},
		"hex.tmpl":      {templateHexTmpl, map[string]*bintree{}},
		"len.tmpl":      {templateLenTmpl, map[string]*bintree{}},
		"lt.tmpl":       {templateLtTmpl, map[string]*bintree{}},
		"lte.tmpl":      {templateLteTmpl, map[string]*bintree{}},
		"main.tmpl":     {templateMainTmpl, map[string]*bintree{}},
		"max.tmpl":      {templateMaxTmpl, map[string]*bintree{}},
		"min.tmpl":      {templateMinTmpl, map[string]*bintree{}},
		"notnil.tmpl":   {templateNotnilTmpl, map[string]*bintree{}},
		"required.tmpl": {templateRequiredTmpl, map[string]*bintree{}},
		"uuid.tmpl":     {templateUuidTmpl, map[string]*bintree{}},
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
