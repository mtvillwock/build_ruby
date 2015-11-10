// Code generated by go-bindata.
// sources:
// data/Dockerfile-centos.template
// data/Dockerfile-lucid.template
// data/Dockerfile.template
// DO NOT EDIT!

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

var _dataDockerfileCentosTemplate = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\x52\x4d\x73\xd3\x30\x10\xbd\xe7\x57\xec\xf8\x5a\x54\x1f\x60\x3a\x5c\x72\xe8\x24\x84\xe9\x00\x09\x93\x12\x38\x50\x0e\xb2\xb5\x76\x45\xf5\xe1\x91\xe4\x90\xb4\xe4\xbf\xb3\x56\x64\x27\x2e\x6d\x66\x92\xec\x3e\xad\x76\xf5\xde\xbe\xc5\x7a\xf5\x05\x9e\x9e\x2e\xe7\xd2\x07\x67\x0f\x87\xc9\x7a\xb3\x84\x7d\xab\xa1\x6d\x04\x0f\x08\x6c\x3f\x20\xd2\xf8\xc0\x95\x22\x08\xea\xb2\xec\xbe\xac\xbc\xb8\x00\xcd\x1f\x10\x78\x1b\x6c\x0c\x1e\xd0\x19\x54\x4c\xe0\x16\x15\xb8\x46\xb3\xa2\x95\x4a\xc0\xdd\x04\xe8\x53\x2b\x59\x94\xe9\x8c\xc2\xaa\x92\x29\xa9\x45\xa1\x53\x68\xca\xd6\x79\xf4\x29\x3b\xde\x73\xc8\x85\x92\x06\x13\x68\x1b\x34\xde\xab\x53\xa3\x3d\xd7\x7d\xf6\x48\xe9\xf8\x6e\x5b\xec\xe3\x4f\xff\x26\x0a\x6b\xd4\x1e\x68\x8e\x82\xc0\x5d\xe4\xf7\x33\xcb\x5b\xef\xf2\x42\x9a\x9c\x0e\xb3\x37\x90\x25\xb6\x5d\x58\x35\x11\x61\x8c\x8e\x85\x74\xd3\xbe\xf4\x08\x1a\xcb\x9c\xb0\xe5\x29\x91\xd9\xaf\xd8\x33\x0e\xe8\xa4\xb5\x7f\x8c\xb2\x5c\x6c\x9c\x3a\x1c\xfe\xd2\x44\xb0\xbb\xc7\x19\xe4\x41\x37\x93\x1f\xab\xf5\xa7\xf9\xcd\x3a\x26\x79\x7c\x25\x5d\x58\xd3\xff\x77\x74\x5e\x5a\x93\x16\x32\x5b\x7c\xbe\xfe\x78\x3b\xcd\x98\xe6\xae\xbc\x9f\xee\xde\x5f\xb1\xab\x77\xc0\x56\x6f\x33\xb8\xcc\x4b\x6b\x2a\x59\xb7\x0e\x23\x63\xc6\x1a\x87\x95\xdc\x4d\x73\xdb\x84\xd8\xf2\x79\xc7\x54\x86\x86\x17\x0a\x99\xbf\xe7\x0e\x45\xc2\x84\xf4\x11\x4c\xe4\x19\xf1\x1a\x57\x77\x3c\x98\x43\xc5\x83\xdc\x62\x7c\x5a\xdc\x3a\xfb\x4d\x43\x96\xad\x9e\x7d\xdd\x50\xff\xde\x28\xf3\x0f\xb7\xdf\x88\xdc\x34\x92\x23\x11\x27\x27\xba\xf1\x2a\x41\x69\x49\xcc\x03\x09\xdb\x27\xa1\x33\x4e\x9f\x18\x78\x51\x96\xfe\x98\x77\x0a\x5f\x93\x28\x27\x68\x0b\xaf\x14\x13\x7c\x13\xd0\xd1\xdb\xa3\xb0\xb1\x58\x8c\x4c\x79\xd7\x83\x23\x77\x0e\xe8\x99\x4d\x07\xec\x25\xbf\x12\xfc\xcc\xb2\x03\x3e\xf6\xee\xf9\xbc\x33\x13\x0f\xf0\x7f\x6e\x66\x47\xdf\xe4\x67\xd2\x35\x90\x13\xb1\x85\x54\xb8\xe4\x1a\x07\xb2\xb4\xfd\xc9\xbf\x00\x00\x00\xff\xff\x1b\x1d\xd6\x78\xdd\x03\x00\x00")

func dataDockerfileCentosTemplateBytes() ([]byte, error) {
	return bindataRead(
		_dataDockerfileCentosTemplate,
		"data/Dockerfile-centos.template",
	)
}

func dataDockerfileCentosTemplate() (*asset, error) {
	bytes, err := dataDockerfileCentosTemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/Dockerfile-centos.template", size: 989, mode: os.FileMode(420), modTime: time.Unix(1446829732, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataDockerfileLucidTemplate = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\x53\x4f\x4f\x1b\x3f\x10\xbd\xe7\x53\x8c\x72\xe0\x66\x5b\xf0\xe3\x87\xda\x4a\x41\x42\x50\x2a\xd4\x16\xaa\x50\xda\x0b\x17\xaf\x3d\xbb\x71\xeb\xb5\x57\xfe\x43\x09\x34\xdf\xbd\xf6\x6e\xbc\xd9\x50\x9a\x4b\x9e\x9f\x67\xc6\xf3\xde\xce\x5c\x2e\x6f\x3e\xc3\xf3\x33\xbd\x50\x3e\x38\xbb\xd9\xcc\x96\x77\xd7\x80\x62\x65\x61\x2e\xb1\x82\x55\x08\xdd\x3b\xc6\x3c\x8a\xe8\x54\x58\xd3\x58\x45\x13\x22\x15\xb6\x65\x03\x04\x1d\x85\x92\xa4\x04\x40\xcb\x95\x99\xc3\xe9\x29\x30\x0c\x82\xf1\x2e\x30\x6f\xa3\x13\xe8\xa9\x4e\x2f\xf4\xd5\x13\x49\x1a\x0c\x10\x3b\xc9\x03\xee\x51\xca\xf8\xc0\xb5\x06\xb2\x06\x17\xab\xf5\x21\x7d\x4b\x0f\x49\x1d\x13\x53\x45\xa5\x25\x41\xef\xd1\x04\xc5\x35\xdc\xcf\x20\xfd\xb4\xaa\xc4\x09\x91\xf8\x90\x51\x5d\xab\x02\x1b\x59\xb5\x05\x9b\xd4\x99\x47\xff\x7f\x7f\x1e\xd3\x1c\x72\xa9\x95\xc1\x12\xe5\xbd\x2e\x70\xcd\xdb\x01\x3f\xa5\xc3\x61\xb3\x9f\x67\x3b\x34\x39\x76\x6c\x6f\xd2\x68\x0e\x4c\x8f\xe9\x5e\x52\x06\xc5\xbe\xce\x59\x19\x45\x50\xd6\x50\x51\xd3\x9c\xd0\x60\xeb\xa9\x75\x0d\x2b\x87\x11\x90\x23\x7a\x4c\x8f\x68\x68\x9e\xe0\x77\xe0\x0e\xec\xe3\xd3\x39\xb0\xd0\x76\x43\x55\xd9\xe3\x17\xd1\x70\x70\xb0\x6b\x03\x3c\x86\xd8\x51\x57\xf5\x09\x29\x68\x60\x8b\xb5\x75\xd7\x02\x21\x95\x32\x52\xb9\x05\x8b\xde\xb1\x84\x13\x63\x2c\x71\xd2\x8a\x2d\x52\x3b\x0d\x79\x38\xec\x2f\xa3\x2d\x97\x77\x4e\x6f\x36\x2f\xba\xfa\x7e\xb3\xfc\x78\x71\xb5\xdc\xb5\x45\x52\xc2\x32\xfd\x7f\x43\xe7\x93\xe2\xed\x48\x9d\x5f\x7e\x3a\xfb\x70\xbb\x98\x93\x96\x3b\xb1\x5a\x3c\xbe\x39\x21\x27\xc7\x40\x6e\xfe\x9b\x03\x65\xc2\x9a\x5a\x35\xd1\x61\xef\x33\x21\x9d\xc3\x5a\x3d\x2e\x98\x4d\xe3\x93\x4b\xbe\xac\xb8\x0d\x43\xc3\x2b\x8d\xc4\xaf\xb8\x43\xb9\xe5\xa4\xf2\x3d\xb9\x95\x4b\xb2\xa4\xbd\xe8\xac\x83\x38\xd4\x3c\xa8\x87\x61\xf8\x5a\xfe\x13\x81\xfc\x48\x8f\x5c\xc7\xf6\xfc\xcb\x5d\xaa\x5f\xcc\xba\x78\x7f\xfb\x35\x89\x5b\xf4\xe2\x92\x73\xb3\x9d\xdc\x3e\x35\x9b\x39\x8c\x06\xf1\x90\x0c\x2d\x87\x00\x79\x79\xb6\x07\x03\xaf\xda\x52\xae\x79\x76\xf8\x2c\x99\xb2\xa3\x1e\xe0\x1f\xc1\x89\xbe\x0a\xe8\x78\x18\x8c\xed\x83\xe5\x64\x0b\xee\x27\x54\x59\x87\x29\x37\xee\xc5\x94\x7c\x65\x41\x86\x8b\xbd\x1d\x99\x5e\x94\x65\x99\x72\xe3\xd6\x8c\xe4\x5f\xeb\x43\x86\x91\x61\x13\xd7\x3a\x60\x49\xd3\xa5\xd2\x78\xcd\x5b\x1c\x75\xa6\x0f\x3f\xfb\x13\x00\x00\xff\xff\x30\x8f\x0e\x3f\x9a\x04\x00\x00")

func dataDockerfileLucidTemplateBytes() ([]byte, error) {
	return bindataRead(
		_dataDockerfileLucidTemplate,
		"data/Dockerfile-lucid.template",
	)
}

func dataDockerfileLucidTemplate() (*asset, error) {
	bytes, err := dataDockerfileLucidTemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/Dockerfile-lucid.template", size: 1178, mode: os.FileMode(420), modTime: time.Unix(1446829732, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataDockerfileTemplate = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\x52\x4d\x73\xd3\x30\x10\xbd\xe7\x57\xec\xf8\xcc\xc6\xd3\x29\x64\xe0\xe0\x43\x27\x21\x4c\x07\x48\x18\x97\xc0\x81\x72\x90\xad\xb5\x2b\x90\x65\x8f\x24\x87\xa4\x25\xff\x1d\x59\xb6\x1c\x07\x4a\x2e\xd9\x7d\xfb\x61\xbd\xb7\x6f\x9d\x6e\x3f\xc2\xd3\xd3\x7c\x25\x8c\xd5\xf5\xe9\x34\x4b\x77\x1b\x60\x8d\xc5\x92\x2c\xb4\x0d\x67\x96\x2e\x20\xa1\x8c\x65\x52\x02\x1e\x41\xb7\xd9\xf1\x6a\xfe\x66\x7e\x0d\x59\x2b\x24\x47\x32\x86\x94\x15\x4c\xc2\xfd\x0c\xdc\x4f\x8a\x2c\x5f\x20\xa7\x7d\x17\x15\x85\x08\x61\xc9\xb3\x2a\xc4\x2a\x6f\xb5\x21\xf3\xca\xe7\xe3\x98\x26\xc6\xa5\x50\x14\xba\x8c\x91\x21\x3c\xb2\xaa\x8f\x1f\x5d\x72\x55\x4e\xe6\xdc\x26\xe9\x9f\xfa\x2d\x8a\x5b\xa3\xe3\x4c\xa8\xb8\xa4\x2a\x7a\x01\xd1\xf0\xe8\x2e\x2c\x1a\x8f\x20\xba\x32\x17\x3a\x09\xad\x3d\xa8\x6a\xd4\xbc\xce\xcf\x89\x88\xbe\xfb\x9d\xdd\x72\x2f\x53\xfd\x4b\xc9\x9a\xf1\x9d\x96\xa7\xd3\x6f\xcb\x34\xd4\x87\xc7\x25\xc4\xb6\x6a\x66\x5f\xb7\xe9\xfb\xd5\x6d\xea\x93\xb8\x13\x07\xdd\x40\xea\xfe\xbf\x90\x36\xa2\x56\x83\xb8\xcb\xf5\x87\x9b\x77\x77\x49\x84\x15\xd3\xf9\x43\x72\x78\xbd\xc0\xc5\x4b\xc0\xed\x75\x04\xf3\x38\xaf\x55\x21\xca\x56\x93\x27\x85\xd8\x68\x2a\xc4\x21\x89\xeb\xc6\xfa\x95\x7f\x6f\x1c\xda\x48\xb1\x4c\x12\x9a\x07\xa6\x89\x0f\x18\x17\xc6\x83\x03\x79\x74\xbc\x2e\xbb\x3b\x1e\xa8\x49\x32\x2b\xf6\xfd\x91\x2b\xf6\x93\x00\x7f\xb8\x8f\x6c\xda\x6a\xf9\x69\xe7\xf6\x87\x7b\xaf\xde\xde\x7d\x76\xe4\x12\x4f\xce\x89\x38\x3b\xd3\xf5\xa3\x0e\x1a\xee\x80\x06\x9c\xb0\x21\xb1\xc0\x29\x0b\x89\x82\x67\x65\x09\x65\xd6\x29\x7c\xe3\x44\x39\x43\x7b\xf8\x4f\xb3\x83\x6f\x2d\x69\xf7\x76\x2f\xac\x6f\xe6\x13\xcb\xdd\x4f\xa0\xe0\xbd\x29\x36\x9a\x70\x0a\x3e\xe3\xc6\xbe\x70\x61\xc8\x69\x21\x38\x73\x8a\x8d\x16\x1d\xc1\x7f\xbc\x8a\xbd\x65\xe2\x89\x6a\x0d\xc4\x8e\xd3\x5a\x48\xda\xb0\x8a\x46\x9e\xee\xf0\xb3\x3f\x01\x00\x00\xff\xff\x0a\xf1\x0c\xed\xa4\x03\x00\x00")

func dataDockerfileTemplateBytes() ([]byte, error) {
	return bindataRead(
		_dataDockerfileTemplate,
		"data/Dockerfile.template",
	)
}

func dataDockerfileTemplate() (*asset, error) {
	bytes, err := dataDockerfileTemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/Dockerfile.template", size: 932, mode: os.FileMode(420), modTime: time.Unix(1446829732, 0)}
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
	"data/Dockerfile-centos.template": dataDockerfileCentosTemplate,
	"data/Dockerfile-lucid.template": dataDockerfileLucidTemplate,
	"data/Dockerfile.template": dataDockerfileTemplate,
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
	"data": &bintree{nil, map[string]*bintree{
		"Dockerfile-centos.template": &bintree{dataDockerfileCentosTemplate, map[string]*bintree{}},
		"Dockerfile-lucid.template": &bintree{dataDockerfileLucidTemplate, map[string]*bintree{}},
		"Dockerfile.template": &bintree{dataDockerfileTemplate, map[string]*bintree{}},
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

