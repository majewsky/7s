// Code generated by go-bindata. DO NOT EDIT.
// sources:
// static/presenter.css
// static/presenter.js
// static/slides.css
// static/slides.js

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
	info  fileInfoEx
}

type fileInfoEx interface {
	os.FileInfo
	MD5Checksum() string
}

type bindataFileInfo struct {
	name        string
	size        int64
	mode        os.FileMode
	modTime     time.Time
	md5checksum string
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
func (fi bindataFileInfo) MD5Checksum() string {
	return fi.md5checksum
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _bindataStaticPresentercss = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x92\xd1\x6a\xc2\x30\x14\x86\xef\xf3\x14\x07\xbc\x13\x3a\x6d\xa1\xca" +
	"\x22\xbb\xf2\x49\x12\x13\xdb\x60\x9a\x13\xd2\x53\xad\x1b\x7d\xf7\x91\xb4\x4a\x8b\x63\x8c\x5d\x9e\xf4\xff\xf2\x9f" +
	"\x7c\xb4\xa6\xc6\xc2\x17\x03\x68\x44\xa8\x8c\xe3\xb0\x3d\x30\x00\x2f\x94\x32\xae\x9a\xa6\x33\x3a\xca\x5a\xf3\xa9" +
	"\x39\xe4\x3b\xdf\x1f\x60\xb3\x86\x0f\xc8\x83\x6e\x60\xbd\x61\x03\x93\xa8\xee\xbf\xdf\x31\x30\x66\xce\x41\x34\xfa" +
	"\x35\x26\x31\x28\x1d\x38\xe4\xbe\x87\x16\xad\x51\x20\xad\x38\x5d\x5e\xb7\xf0\xd8\x1a\x32\xe8\x38\x08\xd9\xa2\xed" +
	"\x48\xc7\x53\x42\xcf\xd3\x2a\xb3\x92\x95\x0f\xfa\x9a\x9a\xac\x3e\x13\x87\xfc\x7a\x8b\xd1\x9b\x51\x54\x73\x28\xb6" +
	"\xe3\x58\x6b\x53\xd5\x94\xe6\x3a\xc2\x13\x7b\xea\x42\x98\xb1\x45\xb1\x80\xcb\xdd\x12\x2e\x77\x0b\xd8\xe9\x9e\x12" +
	"\x1c\xc6\xcf\x7f\x6b\x66\xb2\x23\x42\x97\x40\x65\x5a\x6f\xc5\x9d\x83\xb4\x38\x59\xf8\xf1\xdd\x12\x89\xb0\xe1\x50" +
	"\xee\xe3\x1d\xff\x71\xfa\x56\x26\x6b\x00\x52\x9c\x2e\x55\xc0\xce\x29\x0e\xab\xe3\xf1\x18\xcf\xac\x71\x3a\x7b\xec" +
	"\x99\x3f\x82\xd8\xc7\xdf\x20\xe1\x63\x45\x26\xb1\x9f\xbd\x60\x25\xc9\x65\x4f\xf9\x93\x83\x22\x49\x18\xe6\x91\xa7" +
	"\xa6\xd1\xf1\xfe\x7d\x4c\x7c\x07\x00\x00\xff\xff\xdf\xef\xaf\x89\x8c\x02\x00\x00")

func bindataStaticPresentercssBytes() ([]byte, error) {
	return bindataRead(
		_bindataStaticPresentercss,
		"static/presenter.css",
	)
}



func bindataStaticPresentercss() (*asset, error) {
	bytes, err := bindataStaticPresentercssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "static/presenter.css",
		size: 652,
		md5checksum: "",
		mode: os.FileMode(420),
		modTime: time.Unix(1527197051, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataStaticPresenterjs = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x92\x41\x6b\xdb\x40\x10\x85\xef\xfa\x15\x53\xf5\xe0\x35\x4e\xe4\xf6" +
	"\x6c\x44\xa1\x4d\x68\x5d\xd2\xba\x54\x09\x14\x4a\x29\x8b\xf6\x29\x59\x90\x67\xe5\xdd\x91\xad\x52\xfc\xdf\xcb\x4a" +
	"\xb6\xdc\x83\x0d\x21\x27\x1d\x34\xf3\xbd\x37\xef\xad\xaa\x5a\x2e\xc5\x3a\x56\x53\xfa\x9b\x24\x44\x5b\xed\x29\x40" +
	"\x3e\xb4\xde\x83\xa5\xa8\xad\xc1\x92\x1f\x96\x94\xd3\x38\x69\x4d\x17\x87\x89\x88\x8c\x2b\xdb\x35\x58\xb2\x4d\x0b" +
	"\xff\xa7\x40\x8d\x52\x9c\x57\xa9\xad\xbc\x5e\xe3\x75\xe3\xb1\x4d\xa7\x59\xe9\x58\xc0\x72\x73\x1c\xae\x5d\xa9\x23" +
	"\x29\x7b\xf2\xa8\x28\xa7\xc9\x3c\x44\x9d\x30\x9f\xd0\x8c\x22\x9e\xae\xe9\xed\x74\xf1\x1c\x85\xb2\xf5\xfe\x45\x0a" +
	"\x33\x7a\xf3\x3c\x05\x46\x27\x2f\x54\x88\x37\xf4\x12\x3b\xcb\xc6\xed\xb2\xf2\xbf\x50\x29\x27\x6b\xba\xe8\x60\xdf" +
	"\x0f\x9d\xc9\x5c\x9d\x59\x1b\x88\x67\x4a\x5a\x71\x01\xbf\x85\xbf\x50\x54\xdc\xf0\xd8\x50\x4e\x8c\x1d\xfd\xf8\x72" +
	"\xf7\x49\xa4\xf9\x8e\x4d\x8b\x20\xea\x10\x84\xc7\x26\xd3\xc6\xdc\x6e\xc1\x72\x67\x83\x80\xe1\x55\x5a\x3b\x6d\xd2" +
	"\xab\x13\x14\x47\x24\x51\xe9\x38\xb8\x1a\x59\xed\x1e\x15\x32\xd1\xfe\x11\x92\x79\x84\xc6\x71\xc0\x3d\x3a\x39\x80" +
	"\x07\x79\xc6\x6e\x69\x3a\xca\xe9\x73\xb1\xfa\x9a\x35\xda\x07\x5c\xd8\xfa\x99\x1e\x4e\xfe\xdd\x67\x9a\xfe\x3a\x62" +
	"\x6c\x45\xea\x80\x79\x95\xe7\xd4\xb2\x41\x65\x19\xe6\x64\xe9\x7c\x90\xc3\xce\xe8\x66\xdf\x7f\xf7\xc7\x76\xe2\xdd" +
	"\xae\x01\xab\xf4\xdb\xaa\xb8\x4f\xaf\x4e\x2f\x62\x2c\xb9\xd1\xf2\xc4\x7a\x0d\x9a\xd1\xe4\x5d\x80\x5c\xf7\xc6\xf2" +
	"\xd8\xb5\x1d\xc1\x91\x13\xc0\x66\xc8\x73\xa8\x75\x3e\xbf\xb9\x7d\xff\xf0\x31\x19\xdf\xc0\xe5\xda\x2e\xfc\x59\x24" +
	"\xc9\x5e\x4d\xa7\x8b\xe4\x5f\x00\x00\x00\xff\xff\xa1\x3e\x52\x9b\xab\x03\x00\x00")

func bindataStaticPresenterjsBytes() ([]byte, error) {
	return bindataRead(
		_bindataStaticPresenterjs,
		"static/presenter.js",
	)
}



func bindataStaticPresenterjs() (*asset, error) {
	bytes, err := bindataStaticPresenterjsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "static/presenter.js",
		size: 939,
		md5checksum: "",
		mode: os.FileMode(420),
		modTime: time.Unix(1527196715, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataStaticSlidescss = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x92\x4d\x8e\x83\x30\x0c\x46\xf7\x39\x85\x0f\x10\xaa\xf2\xd3\x59\x04" +
	"\x69\xee\x92\x92\x80\xad\x86\x04\x85\xb4\x0c\x33\xe2\xee\x23\x52\x8a\x68\x61\x89\xe5\xf7\xec\x98\x0f\x43\x6b\xe0" +
	"\x8f\x01\xd4\xce\x86\xa4\xa7\x5f\x2d\x20\x7b\x0c\x25\x03\x68\xc9\x26\xa8\xa9\xc1\x20\x20\x3d\x9f\x1f\x38\x17\xdd" +
	"\x43\xfb\xda\xb8\x21\x19\x05\x20\x29\xa5\x6d\xc9\x26\xc6\x66\x0d\x87\xab\x53\x23\x07\x4c\x39\x60\xc6\x01\x73\x0e" +
	"\x58\x70\xc0\x0b\x07\xfc\xe2\xd0\xc5\x39\xad\xf4\x0d\x59\x01\xe7\xd9\xd6\x49\xa5\xc8\x36\xcb\x57\x5c\xa1\x96\x2d" +
	"\x99\x51\x40\x2f\x6d\x9f\xf4\xda\x53\x1d\x07\xcc\xea\xc8\xaf\x48\xea\x75\xbb\x71\x24\xc1\x75\xd1\x33\x6f\x73\xb8" +
	"\xc1\xfb\xf8\xd3\xc5\xeb\x76\xed\xff\x3c\x41\xbe\xc8\x0d\x59\xbd\x1e\xa1\x88\xc5\x89\x61\xb6\xbb\x58\xb4\xed\x80" +
	"\x7c\x29\x4f\x0c\xf3\x1d\x72\x08\xbc\xda\x8b\xcf\xf6\xf4\x78\x42\xb6\x4e\x60\x1d\x87\xbb\xe1\xe0\x0c\x07\x43\x3b" +
	"\xfc\x08\x4e\x37\x30\x7c\x03\xb5\x4d\xc4\x14\xf5\x9d\x91\xa3\x80\xab\x71\xd5\xad\xdc\x1e\x0d\xe4\x3d\xb8\xf5\x7f" +
	"\x9c\xac\x4b\x7a\x43\x4a\x47\x6c\x20\x15\xf0\x99\x94\x18\x9f\x5d\x74\xae\xb2\xba\x35\xde\xdd\xad\x9a\xd5\xf2\xa9" +
	"\xde\x26\x60\xfa\xb0\x3e\x03\xf3\xbe\xf2\x4b\x56\x39\xe3\xbc\x80\x01\x29\xe8\xf2\xfd\xad\xc5\xf2\xd6\xa0\x7f\x42" +
	"\x22\x0d\x35\x56\x40\xa5\x6d\xd0\x3e\x6e\xfe\x1f\x00\x00\xff\xff\x8f\x75\x65\xfa\xf2\x02\x00\x00")

func bindataStaticSlidescssBytes() ([]byte, error) {
	return bindataRead(
		_bindataStaticSlidescss,
		"static/slides.css",
	)
}



func bindataStaticSlidescss() (*asset, error) {
	bytes, err := bindataStaticSlidescssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "static/slides.css",
		size: 754,
		md5checksum: "",
		mode: os.FileMode(420),
		modTime: time.Unix(1527179541, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataStaticSlidesjs = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x56\x3d\x8f\xe3\x36\x10\xed\xf7\x57\x3c\xa8\xb2\xe1\x5d\xd9\x38\xc0" +
	"\x85\xe1\xd5\x1e\xae\x08\x70\x45\x90\xe6\x8a\xd4\xb4\x38\xb6\x88\xa5\x39\x0a\x49\x59\x51\x02\xff\xf7\x80\xfa\xb0" +
	"\x48\xef\x07\x72\x40\x90\xc6\x80\x38\xc3\x37\xc3\xf7\x1e\xc7\x5c\x1c\x1b\x53\x7a\xc5\x66\xb1\xc4\xdf\x0f\x0f\xc0" +
	"\x45\x58\x58\x72\xea\x2f\xfa\x2e\x8c\xd4\x64\x51\x20\xc9\x01\x86\xa4\xb2\xb1\x96\x8c\xff\x5d\x49\x5f\xa1\xc0\x76" +
	"\xb3\xbf\x85\x3c\xf3\x8f\xb3\xd0\x7a\x8a\x35\x46\xd2\x51\x19\x92\x49\xca\xaf\xc2\x9e\xe8\x93\x14\x49\x4e\x59\x92" +
	"\xdf\x49\x9d\x2a\x8f\x02\xad\x32\x92\xdb\x5c\x19\x43\x76\x58\x7c\x5b\xf1\x96\xfc\x49\xc9\xf7\x72\xfa\xa4\x92\x8d" +
	"\x63\x4d\x79\xa9\x49\xd8\xc5\x72\x9f\x2c\x6a\x3e\x2d\xb2\xb1\x25\x54\x13\x44\x86\x55\xda\xe7\xb8\xeb\xc8\x16\x8b" +
	"\x50\x54\x79\xb2\x22\x70\x87\x02\x9b\x7d\xf4\xf9\x1c\x18\x9b\xbf\x57\xab\x89\x5b\x60\xbd\x16\x75\xad\xbb\x84\xe1" +
	"\x31\xd4\x43\x9e\xc5\x89\x7e\x63\x49\x0e\x05\x24\x97\xcd\x99\x8c\xcf\xff\x68\xc8\x76\x3f\x48\x53\xe9\xd9\x7e\xd3" +
	"\x7a\x91\xd5\x78\x81\x3a\x9f\xb2\xb1\x25\xe0\x9b\xb5\xa2\xcb\x6b\xcb\x9e\x7d\x57\x53\x7e\x64\xfb\x8b\x28\xab\xbc" +
	"\x14\x5a\x2f\x66\xd4\xc7\x59\xef\xdb\xe2\xdc\x1c\xe6\xfa\xb9\xf3\x9d\xa6\xbc\x1d\x25\x4c\x0c\xb1\x42\x76\x69\xb3" +
	"\xa9\xf2\x75\x39\x52\x1c\x0e\xd7\x56\xc2\x4f\x0c\x4a\x26\x07\x1f\x16\x6a\xcb\xb2\x29\xe9\x6b\x74\xd0\x11\xf0\xa6" +
	"\xd7\xed\xac\x07\x96\x5d\xee\x4a\xcb\x93\xe0\x53\x1d\x75\xc4\xe2\x6e\x57\x71\xa7\x4f\x74\x90\x83\x25\xf1\xba\xc7" +
	"\x7a\x7d\x21\xdb\xa1\x31\x5a\xbd\x92\xee\xa6\x96\x3f\x84\x7c\xfe\x18\xf1\xde\xf8\x31\x25\xfb\x37\x59\xb7\x83\x25" +
	"\xf8\x73\x5e\x54\x7a\x00\x7c\x29\xb0\xdb\xc6\xf5\x02\x9d\x67\xb6\x14\x28\x34\x10\x07\x6e\x3c\x76\xdb\x4b\x35\xf0" +
	"\x6a\xd8\xe3\x2c\x5e\x09\x8e\x8c\xa3\x3d\x7c\x45\x83\x78\x68\xb9\xd1\x12\x7c\x21\x7b\xd4\xdc\x82\x4d\x1f\xb3\xa1" +
	"\x7c\x84\xfd\x1f\xf9\xe5\x33\xcf\x64\xbb\x6d\x64\x93\xc9\x2a\xf3\x57\x72\xfb\x0e\xdc\x18\x89\x43\x87\x7e\x77\x96" +
	"\x24\x0e\x5a\xde\x16\x26\xf5\xae\x20\xed\x28\x55\x28\x99\x3b\x1f\x2a\x94\x8e\x8a\x77\x15\xba\xce\x9e\x96\x8c\x96" +
	"\x50\x89\x0b\xc1\x33\xe8\xcf\x5a\x07\x55\x7a\x69\xa4\xb2\xd4\xb3\xe3\xbe\x46\x96\xba\xf7\x40\x11\x8d\xa3\x98\xbc" +
	"\x44\xff\xa7\x02\x5f\x36\x9f\xf8\xe3\x19\x9b\x7b\x77\x08\x57\x53\xe9\xd1\x0f\x19\x28\x17\x4e\x86\x96\x94\x95\x78" +
	"\x7a\xc1\x41\x28\x0d\x6e\xfe\x6f\xcd\xb7\x9b\x9f\xd0\x5c\x28\xad\xcc\x29\x74\xf9\x33\x82\x8f\x14\x27\x22\xfe\x1b" +
	"\x86\x57\x31\xc3\x6f\xcc\xb3\x5e\x4f\x22\x1f\xd8\x57\x10\x77\xf7\x5d\x18\x39\xac\xcd\x0e\x7b\x8c\xf6\x3a\xc6\x41" +
	"\xb9\xa0\x86\xe7\x56\x58\xe9\x20\x70\x62\x96\xb8\x08\xdd\xd0\xfb\xfd\x14\xb3\x53\xa6\xb9\x9a\xe0\x2f\xb1\xc6\x97" +
	"\x37\x7e\x4c\x18\x5c\x64\xc3\xa6\x04\x24\x0b\xf2\xc7\x81\x91\xa4\x15\xb2\x25\x9e\xfa\xc8\xfd\x38\xc7\x13\x26\xa8" +
	"\xe8\x06\xa5\x50\x31\xdf\x01\x2a\x8b\xe7\x3e\x9b\x92\x42\x6b\x17\xb2\x27\x92\x8f\x70\x9e\xeb\xe9\x0f\xd0\x9c\xde" +
	"\xb9\x1b\x43\x85\xf5\x5d\xc5\x17\x6c\xf2\xdd\x2e\x51\x30\x31\x4c\xf0\x45\x70\xcc\x81\x4a\xd1\xb8\xa8\x62\xec\x9f" +
	"\xc4\x3d\x83\x77\xc2\xef\xb5\x6f\x37\x79\xfe\x0c\xef\x80\xf1\xed\xc1\x66\x88\xa1\x48\x93\xf6\x0f\x0f\xd7\x65\xc8" +
	"\x7c\xf8\x27\x00\x00\xff\xff\xef\xfd\xb1\x0f\x50\x09\x00\x00")

func bindataStaticSlidesjsBytes() ([]byte, error) {
	return bindataRead(
		_bindataStaticSlidesjs,
		"static/slides.js",
	)
}



func bindataStaticSlidesjs() (*asset, error) {
	bytes, err := bindataStaticSlidesjsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "static/slides.js",
		size: 2384,
		md5checksum: "",
		mode: os.FileMode(420),
		modTime: time.Unix(1527179541, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}


//
// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
//
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

//
// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
// nolint: deadcode
//
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

//
// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or could not be loaded.
//
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

//
// AssetNames returns the names of the assets.
// nolint: deadcode
//
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

//
// _bindata is a table, holding each asset generator, mapped to its name.
//
var _bindata = map[string]func() (*asset, error){
	"static/presenter.css": bindataStaticPresentercss,
	"static/presenter.js":  bindataStaticPresenterjs,
	"static/slides.css":    bindataStaticSlidescss,
	"static/slides.js":     bindataStaticSlidesjs,
}

//
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
//
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, &os.PathError{
					Op: "open",
					Path: name,
					Err: os.ErrNotExist,
				}
			}
		}
	}
	if node.Func != nil {
		return nil, &os.PathError{
			Op: "open",
			Path: name,
			Err: os.ErrNotExist,
		}
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

var _bintree = &bintree{Func: nil, Children: map[string]*bintree{
	"static": {Func: nil, Children: map[string]*bintree{
		"presenter.css": {Func: bindataStaticPresentercss, Children: map[string]*bintree{}},
		"presenter.js": {Func: bindataStaticPresenterjs, Children: map[string]*bintree{}},
		"slides.css": {Func: bindataStaticSlidescss, Children: map[string]*bintree{}},
		"slides.js": {Func: bindataStaticSlidesjs, Children: map[string]*bintree{}},
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
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
