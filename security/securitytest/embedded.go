// Code generated by go-bindata.
// sources:
// ../../resource/test_certs/ca.crt
// ../../resource/test_certs/ca.key
// ../../resource/test_certs/node.client.crt
// ../../resource/test_certs/node.client.key
// ../../resource/test_certs/node.server.crt
// ../../resource/test_certs/node.server.key
// ../../resource/test_certs/root.client.crt
// ../../resource/test_certs/root.client.key
// DO NOT EDIT!

package securitytest

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
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

var _test_certsCaCrt = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x64\x92\xcd\x72\xa2\x40\x14\x85\xf7\x3c\xc5\xec\x53\x53\x21\x0d\x29\x75\x79\x9b\xbe\x68\x23\xdd\x4e\x23\x3f\x81\x9d\x68\x00\x11\x21\x24\x4a\x43\x9e\x7e\xd4\xc5\xd4\x4c\x4d\x2f\xbf\xc5\xe9\xaf\xce\x3d\x3f\xef\x8f\xe2\x92\xcb\x1f\x0e\x06\x21\x77\xb9\x03\x21\x3e\xa8\x21\x38\xa7\xef\xa1\xe3\xc0\x36\x29\x41\x73\x0a\x25\x0f\x60\x7d\xf0\x1a\x7b\x5f\x91\x5d\x66\x39\xde\xa6\xc8\xb7\xaf\x38\x64\x5a\x3b\x3a\xf5\xd6\x5d\xc6\xab\x61\x2f\x41\xa1\x2f\x9c\xaf\xd1\xc0\x1a\x14\x2d\x65\x4c\xa1\x0b\x9d\x48\x0e\x29\xb9\x4c\x39\x71\xeb\x1d\xc3\x58\x50\xb1\x84\x97\x08\x61\x14\x42\x91\x45\xbd\xb3\xbc\x21\x4d\x64\xc7\x51\x52\x41\xed\x37\x16\xe2\x8b\x21\x58\x3a\xca\x1a\x27\x19\x8a\x17\xd1\x74\x77\x48\x1e\x2c\xfc\xc3\xf4\xfa\x1b\xb7\x82\xc2\x23\xcc\xa9\x84\xf7\x77\x98\x21\x82\x48\xa3\x4e\x59\xac\x14\x43\x3d\xb2\x9c\xc8\xcf\xfd\x79\x51\xa5\xa4\x2c\x95\x89\xfa\x8d\x81\xa4\xe5\xa9\xaf\x4e\xc7\xe5\x42\x9b\xf4\xa6\xee\x02\x48\x1f\x18\xe5\x70\x42\x30\x5a\x3a\xaf\xeb\x84\x90\x59\x2e\x16\xae\xdd\xbb\x4d\xd4\x9f\xfa\xe3\xf6\x30\xfe\x62\x22\xd5\x4e\x50\x78\x7e\x73\x2a\x2a\xbf\x79\xea\xf3\x59\x57\x4d\x41\xbb\x8d\x2f\xaf\xee\xb5\xb5\xac\xf4\x83\xf0\xdc\x48\xe0\xf9\xea\x91\x99\x6f\xc7\x89\xa5\x58\x98\x9d\x93\xcb\x18\x69\xce\x40\x01\xed\xa6\x54\x7b\x0c\x36\xf7\x8e\x56\x6a\x4e\xa1\x98\x23\x05\xe1\xc0\x46\x69\x2c\x6f\xd6\x81\x19\x1a\xa0\x56\xcf\xb7\xea\x35\x2d\x91\x3e\xdf\xaf\x10\x82\xff\xaf\xf2\x17\x53\x0a\x6e\xdf\x7c\xbe\x17\x7d\xd4\x56\x66\x53\xc8\x21\xb2\x76\x41\x32\xad\x26\x2c\x8d\xd9\xfa\x72\xcd\x06\xd7\x74\xab\xd6\x3c\xfa\x24\x3b\xd8\xe7\x38\xb4\x5f\xdb\xac\xfc\xf0\xaf\xc7\xb4\x7a\x2a\x4d\x89\xea\x3b\x4b\xf7\x56\xd4\x64\x95\x59\x1f\xfb\xd9\x79\xd2\x44\x8d\x93\x67\x3c\x96\x80\x92\xfd\xbf\x8e\xdf\x01\x00\x00\xff\xff\xd6\xf2\x00\x70\x3a\x02\x00\x00")

func test_certsCaCrtBytes() ([]byte, error) {
	return bindataRead(
		_test_certsCaCrt,
		"test_certs/ca.crt",
	)
}

func test_certsCaCrt() (*asset, error) {
	bytes, err := test_certsCaCrtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "test_certs/ca.crt", size: 570, mode: os.FileMode(420), modTime: time.Unix(1400000000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _test_certsCaKey = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\xd1\xbb\x92\xb2\x4a\x18\x85\xe1\x9c\xab\x30\xb7\x76\x81\xa2\x02\xc1\x04\x1f\xd8\x40\x83\x20\x0d\xd8\x03\x64\xd3\x20\xc8\x19\x15\x44\xb9\xfa\x5d\x33\xf1\xbf\xd2\x55\xf5\x26\xcf\x7f\xbf\x53\x91\x81\xdd\x95\x1f\xc0\xca\xf3\x31\x85\x10\xad\x6c\x14\xff\x3d\x9c\x83\xb1\x7a\x9e\xb1\x0a\x60\xa9\x60\xcd\x39\xde\x6d\xc6\x71\xda\x2e\x9e\x7f\xb5\x6f\x09\xb5\x07\xbb\xaf\xbb\x34\x9c\x17\xc3\x81\x3a\xfa\x04\xfb\xc4\xdc\x05\xfb\xfc\xde\xae\x0f\x38\xad\x13\x81\x6b\xd8\xd5\x67\x83\xa2\x6c\xed\xac\x34\x36\x85\x27\x95\xd9\xf4\x59\xeb\x74\x14\x40\xd8\x26\xcd\x23\xa5\x8e\x06\x33\x02\x20\x16\xa0\x45\xea\x32\x35\xa4\xdf\x1f\x5c\xc4\x63\x6b\x8d\x5a\x79\xe0\xae\xd9\x83\x7f\xf2\x83\xe8\x4b\x28\xf2\x6a\xe7\xd6\xc8\xc5\x11\x25\xb3\xa6\xa6\x28\xa3\x6e\xf4\x7e\x03\x4a\x3b\x69\x84\x0b\xe8\xa1\x62\x91\x7c\xc0\x47\xa5\xed\x8c\x9f\xe3\xa9\xb2\xfa\x27\x09\xb8\x92\xe0\x1b\x38\xe7\x7c\x51\xd8\xf5\x7e\xd0\x63\x5c\x81\x67\x74\x2f\xcd\xed\x4e\xaa\x6a\xf2\x8f\x39\x74\x8d\xdd\xf8\xa6\xe6\x25\xcf\x99\x1d\x98\x50\x22\x58\xd0\x4e\x11\x7c\x31\xb9\x0f\x0e\xcf\x9d\x94\xef\x42\x64\xfd\x39\x08\xa4\x92\xff\xa9\xab\x63\x9f\x35\x32\xe5\x1d\x7e\xd7\xd0\x40\xb9\x68\x98\x68\xa1\x6c\x45\x15\x5b\xc4\x7e\xb0\x8a\xd6\xd5\xa2\x44\x2f\xc2\x93\x27\xeb\x6a\x62\xc8\x94\x4b\x2b\x72\x88\xcd\x57\x7b\x71\xd8\xba\x23\xb8\x50\xcb\x2d\xdf\xb8\xcd\x4c\xd2\xe4\xea\xb9\x31\x33\x69\x6b\x61\x5f\x17\xa5\x35\xbf\x54\x11\x9b\x04\xdd\x5b\xe2\x40\xc6\x68\x53\xff\xc6\x75\xee\x19\x13\xc7\x16\x5e\x9b\x78\x7f\xe8\x4c\x19\x98\xff\x1a\xb7\x0c\x62\xd8\x97\x13\x4f\x7a\x71\x26\x92\x33\x3d\xb2\x75\x6d\xce\x5f\x5f\xdc\x1f\x1c\x72\x8f\xff\x06\xfd\x3f\x00\x00\xff\xff\x0c\xf3\x27\x95\xf1\x01\x00\x00")

func test_certsCaKeyBytes() ([]byte, error) {
	return bindataRead(
		_test_certsCaKey,
		"test_certs/ca.key",
	)
}

func test_certsCaKey() (*asset, error) {
	bytes, err := test_certsCaKeyBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "test_certs/ca.key", size: 497, mode: os.FileMode(420), modTime: time.Unix(1400000000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _test_certsNodeClientCrt = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x64\x92\x4b\x93\x9a\x40\x14\x85\xf7\xfc\x8a\xec\x53\x29\x05\x5f\xe3\x22\x8b\xdb\x0f\xb0\x91\xdb\xda\xbc\x75\x07\x18\x50\x06\x45\x07\xb5\x75\x7e\x7d\xd4\x4a\xa5\x92\x9a\xbb\xfc\x16\xa7\xbe\x7b\xea\xfc\x78\x1e\xe1\x8e\x90\xdf\x28\xf7\x43\x61\x0b\x0a\x21\x7f\x51\x03\x85\x20\x45\x4d\x29\xf8\xe3\x0a\xb4\x20\x50\x09\x1f\x3c\x73\x3d\xd0\x98\xc9\xa3\xee\xe6\x69\xa0\xaf\xfd\xb2\x28\x36\x6f\x9a\xea\x95\x3b\x6f\xd7\x62\x7b\x2d\x24\x28\xee\x21\xed\x6e\x06\xaf\x41\x91\x4a\xc6\x04\xda\x90\x46\xf2\xba\xb2\xce\xf7\xdc\xb2\xeb\x8c\xf1\x18\x09\x3a\x60\x46\x1c\x6e\x88\xca\x9a\xd6\xd9\xc0\xbd\xae\x12\xd9\x0a\x2e\x09\x92\x61\xca\x42\x6e\x1a\xc8\x56\x37\x59\xf3\xbb\x0c\xd1\x94\x71\xfb\x84\xd6\x8b\x85\x7f\x99\x16\x9f\x3c\x40\x02\xaf\x30\xba\x45\xf7\xdf\x30\x03\x55\xff\xa9\xc6\x62\xa5\x18\xd7\xfe\x25\xb7\xfc\x06\x6d\xad\x99\xfa\x4f\x97\x10\x05\x2c\xd0\xa0\x03\x70\x09\x2c\xf3\xa8\xf3\xed\x6e\xeb\x37\x0b\x23\xf4\xf6\x53\x87\xfb\xe7\xdb\x70\xd1\x98\xef\x0d\x29\x13\xbe\x2b\x82\x4e\x05\x87\x20\x5a\xc0\xec\x94\xef\x7a\xce\x10\x0e\x40\x7a\x61\x79\x9d\xec\x86\x38\x5f\xdf\x3e\xd6\xf7\x4d\x72\x0e\xf3\xe4\x34\x49\x8d\xe0\xd8\xda\x4b\x7b\x94\x3e\xfe\xa4\xa0\x39\x40\x86\x07\xa4\x91\x66\xd5\xc3\xca\xef\x2f\x41\xcd\x7a\x04\x14\x83\x8a\x56\x7f\x3a\xd9\xb8\x4a\x21\x42\xeb\x50\xda\x39\x06\xa8\xc8\x26\x1a\x29\x42\xe7\xd0\xe0\xe4\x04\x22\x1f\xb0\x87\x32\xd5\x92\x80\x9d\xaa\xfa\xb8\x1d\xae\x97\xdf\x31\x71\x16\xee\x64\x7c\x2f\xad\x8c\xed\x7b\x54\x3b\x17\x36\x02\x0b\x27\x33\x63\xac\xd3\xa2\x6f\x97\x56\xe1\x1d\xfb\x0b\x6f\x57\x75\xcd\x2e\xde\xcf\xa7\x34\xae\xca\x4b\x52\x84\xef\xbf\x46\xf6\xca\x05\xb3\xa3\x5e\x59\x64\x9f\xdb\xcd\x47\xf3\xf6\xd3\x78\x2d\x80\x4b\xf6\x75\x15\xbf\x03\x00\x00\xff\xff\x9a\xe1\x5d\xb2\x32\x02\x00\x00")

func test_certsNodeClientCrtBytes() ([]byte, error) {
	return bindataRead(
		_test_certsNodeClientCrt,
		"test_certs/node.client.crt",
	)
}

func test_certsNodeClientCrt() (*asset, error) {
	bytes, err := test_certsNodeClientCrtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "test_certs/node.client.crt", size: 562, mode: os.FileMode(420), modTime: time.Unix(1400000000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _test_certsNodeClientKey = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\xd1\x3b\x97\xa2\x30\x00\x05\xe0\x9e\x5f\x31\xbd\x67\x0f\x4f\x15\x8b\x29\x12\x0c\x21\x28\xf2\x08\x88\xd8\x01\x1a\x40\x50\x5e\x4a\xc4\x5f\xbf\xb3\x53\xef\x6d\x6f\x71\xef\x39\xdf\x9f\x7f\x81\x08\x93\xc3\x57\x40\xc1\x97\x17\x90\x23\x08\xd1\xd7\x0e\x25\xbf\x8d\xe0\x10\x02\x5d\x4e\x20\x00\x36\x04\x5e\x16\x8d\x81\x39\x96\x41\xe3\x86\xfb\xfb\x06\xa3\xe0\xf9\xd6\xdc\x46\xae\x1b\xc8\x62\x54\xe5\x74\xf4\xe9\x83\x46\x2e\xb0\xfa\xac\x12\xb1\x06\x1e\x40\x80\x62\xc8\xa6\x75\xa5\x39\xbb\xf3\x7b\x38\xcf\x97\xf8\x19\x66\x71\xbf\x3e\xd1\xae\x35\x3d\x73\x79\x02\x6f\xc7\x00\x1c\x01\xe0\xff\x0c\xec\x86\x3d\x4e\x86\x44\xb7\x20\xf6\xf3\x9b\x84\x94\xfe\x2e\xe0\xf5\xa4\xe5\xc4\x4e\x50\xab\xa8\x2b\x79\x8d\x74\x0e\x98\xaa\x9a\xf4\x76\x8d\xd8\xc6\x28\xa5\x53\xc4\x01\xac\x49\xe6\x0d\xe2\xd4\xa9\x25\xed\x97\x2a\x7f\xa0\x83\x43\xd2\x62\x71\x50\x53\x2a\xec\x28\x32\x88\xbf\x55\x16\xde\x10\x6f\xca\x69\x24\x53\x9d\xaa\xaf\x75\xa2\x15\x86\xaa\x4d\x53\x23\x1e\xf5\xda\x57\x31\x8e\x36\xda\x91\x22\xf9\x66\xf9\xa4\x04\x9e\x98\x26\x7d\x57\x17\x70\x08\x05\x98\x47\xf3\xf5\x42\xd9\xe1\xf9\xd9\x73\x74\xcd\xa4\x4f\xcb\xea\xe4\x25\xad\x72\xbc\x31\x4e\x6c\x02\x15\x02\xe3\x25\x6f\x8d\xf3\x4e\xc4\xb6\xa3\xb0\x6c\x58\xd5\xdc\x1e\x1c\xa6\x77\xe6\x9e\xc7\xc2\xd4\x6b\x33\x8e\x61\x86\xa4\x42\x76\xd2\x9f\x37\x86\xde\x69\x39\xb4\x23\xab\x28\x53\xd3\xba\x8f\xd4\xe8\x3e\x2e\xd1\xe7\x00\xf4\xbe\xb5\x3d\x5f\x5f\xad\x65\xb4\x81\x19\xa4\xa9\x07\x2a\x78\x14\x3c\x1a\x1e\x87\x08\xce\xee\x95\xbe\xac\xad\xd2\x2c\xa4\x84\x37\xf6\xd6\x26\xb0\x92\xcb\xe5\x2c\xc9\x4a\x7a\x62\x83\xc8\x16\xfc\xfb\x5b\xf8\x85\x43\x87\xed\xff\x41\xff\x06\x00\x00\xff\xff\x80\xe9\xa3\x32\xf1\x01\x00\x00")

func test_certsNodeClientKeyBytes() ([]byte, error) {
	return bindataRead(
		_test_certsNodeClientKey,
		"test_certs/node.client.key",
	)
}

func test_certsNodeClientKey() (*asset, error) {
	bytes, err := test_certsNodeClientKeyBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "test_certs/node.client.key", size: 497, mode: os.FileMode(420), modTime: time.Unix(1400000000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _test_certsNodeServerCrt = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\xd6\xcd\x76\xaa\x3a\x14\x07\xf0\x79\x9e\xe2\xce\xbb\xee\x2a\x04\xa8\xcb\x61\x42\xa2\x46\x4d\xda\x08\x5a\x39\x33\xb5\xca\x87\xf6\x70\x2c\xd5\x00\x4f\x7f\x85\x5b\x8f\xad\x6c\x87\x7f\xb2\x93\xfc\x76\x10\xf8\xb7\xf9\x51\x3e\x14\xea\x1f\x9f\xcf\x42\x31\x10\x3e\x09\x79\x9b\x22\x29\xc4\x68\x95\xf9\x3e\x2d\x57\x31\x31\x82\x92\x58\xcc\x88\x3c\xf3\x11\x37\xc3\xc3\xf6\x4f\x15\xd1\xa8\xfc\x70\x3e\x27\xc7\xfd\xc6\xf8\x26\x1a\x4f\xf2\x5f\x22\x39\x6f\x14\xd1\x7c\x2a\xfd\xa2\x44\x3c\x23\x9a\xc6\x6a\x41\x49\x1e\xfa\x73\x75\x8e\xf0\x67\xb5\xc6\x83\x6c\xc5\xf8\x42\x52\x39\x24\xf6\x9c\x93\x52\x4a\x8d\xfb\xd9\xca\x19\x9f\xa3\x57\x95\x0b\xae\xa8\xa4\xee\x92\x85\xdc\x46\x92\x45\xa5\xca\x78\xa5\x42\x69\xab\x45\xde\x84\xb8\xcd\xc2\xbf\x99\x11\x35\x0f\x24\x25\xed\x64\x7e\x22\xc7\xdf\x27\x43\x52\x5b\xcd\xd6\xd8\x42\x6b\xc6\xcd\xec\xb4\xc6\xb3\x83\x1c\x18\xc3\xf4\x8f\xed\x52\xaa\x09\x0b\x0c\x31\x01\x19\x53\xa2\xf4\x7c\x96\xe9\xad\xd0\x95\x85\x5e\x54\x12\x9c\x96\x66\xa2\x97\xfb\x77\xc1\x1e\x92\xde\x69\x9d\x8f\x3f\xbd\xf3\x28\xff\x2d\x8b\x48\x0d\xdc\xa0\x17\x6c\x58\xf5\xc1\x9f\xde\x76\x8b\xb7\x87\xfd\x14\xf7\x1c\x67\x19\x8d\x95\xfe\xa8\x66\xeb\x51\x88\x12\x76\xc4\x59\x11\xaa\xde\xef\xb9\x4f\x0c\x27\x64\xf5\xec\xd3\xc4\x32\x71\x1c\xfd\x92\xc4\x6d\xb6\xfd\xc6\x0c\xa7\x8f\x46\xf3\x4b\x93\x49\xce\x48\xd8\xf4\x6c\x14\xcc\x39\x63\x64\x82\x68\x1c\x7f\xd0\x98\x0f\xa8\xde\x5c\x2e\x5d\x4e\x63\xd7\x5a\xdf\xb8\x0e\x7c\xba\xdd\x5c\x26\x5a\x66\x71\x7c\x28\xd6\x58\x25\xeb\xe1\x05\xe4\x04\x3e\x1b\xfd\xef\x97\xbe\x57\xa0\xf6\x82\x10\x72\xf3\xde\x4f\x22\x1c\x97\xd3\xf7\xf2\x72\x12\x83\x22\x8e\xcb\xaf\xd3\x10\xa7\xf5\xb0\x9f\x45\xaf\xd5\xad\xb0\xf2\x8a\x6b\x1d\xfa\x2a\xb4\xba\x85\xf3\x4e\xa1\x4a\x6f\x85\x5f\x75\x0e\xea\x16\xc6\x9d\xc2\xe7\xe0\x6f\xa1\xba\xee\xb4\xd9\x7e\x93\xa1\x1f\x21\x34\x30\x05\xb2\x1b\x41\x5d\x09\xa5\xf2\xbb\x03\x15\x30\xa1\xba\x9b\x10\xb5\x61\xd5\x1d\xf8\x0c\x4c\x08\x50\x2a\x24\xbb\x03\x2b\x80\x52\x01\x94\xaa\xa1\xa0\xfb\x10\xa0\x54\x00\xa5\xba\x52\xd0\x8f\xb0\x4b\xa9\x00\x4a\xf5\x9d\x72\x6d\x62\x0d\x50\x6a\x80\x52\xdf\x53\x50\x1b\x76\x57\xae\x01\x4a\x0d\x50\x6a\x74\x7f\x2c\xed\x40\x60\x42\x80\x52\x37\x94\xfb\x26\x5a\x00\xc5\x02\x28\x96\x04\x9a\x68\x01\x14\x0b\xa0\x58\x0a\x68\xa2\x05\x50\x2c\x80\x62\xdd\x53\x50\x1b\x76\xb7\x68\x03\x14\x1b\xa0\x5c\x1e\xaa\xdd\x95\x6d\x80\x62\x03\x14\x5b\x01\x4d\xb4\x01\x8a\x0d\x50\xec\x67\xe0\xef\x6c\x03\x14\x0c\x50\xb0\x04\x9a\x88\x01\x0a\x06\x28\xf8\x9e\x82\xda\x10\x58\x19\xa0\x60\x80\x82\x11\x70\x87\x61\x80\xe2\x00\x14\x47\x02\x4d\x74\x00\x8a\x03\x50\x1c\x05\x34\xd1\x01\x28\x0e\x40\x71\x14\xf0\x4c\x74\x00\x8a\x03\x50\xdc\x7b\x0a\x6a\x43\x68\x60\x77\x65\x17\xa0\xb8\x08\xb8\xc3\x5c\x80\xe2\x02\x14\x57\x01\xcf\x44\x17\xa0\xb8\x00\xc5\x83\x5e\x2c\x1e\x40\xf1\x00\x8a\x07\xbd\x58\x3c\x80\xe2\x01\x14\x0f\x7a\xb1\x78\xc0\x1d\xe6\x01\x14\xef\x46\xd9\xf0\x9d\x21\x84\x84\x64\x7a\xf9\x3e\xd8\x1f\x93\x7d\x3a\xec\x1b\x8b\x12\x5d\x30\xad\xe9\x29\xfd\x73\x3c\x3d\xd6\xf3\xbe\x5f\xd8\xbb\x74\x94\xbd\xe4\x65\x7d\xee\x1d\x0f\xdb\xbd\xe8\x99\x24\x4b\x9f\xce\x2f\xd3\xfa\xf4\xf8\xb8\x7d\x91\xd3\x29\x72\x8c\xd0\x02\x0f\x14\x3d\xa4\x34\x18\x4f\xd8\xfb\xd3\x20\x4d\x8f\xf9\x61\x67\xd5\xf6\xc2\x0d\xce\x63\xaf\x3f\x32\xaf\x0f\x25\x6a\xbf\x0a\xb9\x62\xdd\x2f\xc5\xff\x02\x00\x00\xff\xff\x55\x07\x34\x7c\x46\x0a\x00\x00")

func test_certsNodeServerCrtBytes() ([]byte, error) {
	return bindataRead(
		_test_certsNodeServerCrt,
		"test_certs/node.server.crt",
	)
}

func test_certsNodeServerCrt() (*asset, error) {
	bytes, err := test_certsNodeServerCrtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "test_certs/node.server.crt", size: 2630, mode: os.FileMode(420), modTime: time.Unix(1400000000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _test_certsNodeServerKey = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\xd1\xcb\x92\xa2\x30\x00\x85\xe1\x3d\x4f\xd1\x7b\x6b\x4a\xe4\xda\x2c\x03\xe1\x12\xa1\xa3\x01\xc4\x64\x76\x22\xc8\x25\x8a\x9a\x20\x48\x3f\xfd\x54\xf7\x7a\xce\xf6\x6c\xfe\xaa\xef\xcf\xcf\x5c\x3f\x44\xf8\x23\xcd\xc0\xc7\x3e\x45\x05\xc8\xfd\x8f\xd8\x67\xbf\x8f\xf2\x85\x90\xbb\x6b\x90\x0b\xc0\xd6\x05\x98\x1c\xd2\x9e\xd4\x88\x2c\xea\x1e\xb7\xd9\x8b\xce\x31\xa1\xfc\x86\xe0\xaa\xb5\x5f\xe5\x7d\x3b\x9a\x53\x74\x1f\xbe\x24\xc3\x81\x91\xd9\xd9\x19\x2e\x42\xf1\xad\xea\x52\x54\x2b\x9e\x68\xb6\xae\x53\xb6\xc5\x44\x2c\x69\x19\xe5\x2d\x7c\x6a\xbd\xcc\xb1\x3d\x1c\x3c\x30\xfb\x00\x90\x2d\x60\xc7\xf0\xad\xcb\xe7\xc9\x0a\x2b\x4d\xec\x1e\xd1\xc9\x8c\x0b\x65\xed\x34\x37\x33\xf1\x1f\x62\xda\x31\x41\xbb\x5a\x7a\xba\x27\x89\xa7\xa5\xaf\x37\x79\x02\x79\xcd\x9b\xde\xc0\x81\xf9\xe9\x3e\x1b\xb8\xf6\x98\x03\xc0\x50\x3d\x82\x41\x9f\x18\xb3\xb9\x95\x27\x0a\x22\xa8\x05\x3b\x18\xb0\x4e\x2b\xb2\x53\xc8\xdf\x7a\xfd\x4e\xae\x8c\x5f\x23\x61\x95\xea\xf1\x71\x82\xea\x38\xf4\x45\x30\x9b\x28\x98\x54\xe9\x80\xce\x07\x9f\x6c\x34\xd2\xf9\x51\x6c\x64\x28\x95\xe1\x70\x58\xdd\x23\xc4\xfd\xe6\xa2\x0a\x5a\xe0\x03\xdc\x0f\xe6\xde\x2d\xbd\xa5\x75\x88\xc9\x3d\xe4\xeb\xc3\x04\x4e\xb7\xab\xa0\x10\x74\x13\x8d\xef\x38\x5e\x47\xf5\xfa\xac\xad\x43\x21\x37\x8e\xc2\xf3\x6a\x2f\x6a\xf7\x32\x9b\x79\x01\x3a\x37\x73\x9b\x1b\x8d\xfa\xb4\x0e\xa0\x67\x2c\x1b\x9d\x03\x68\x8d\x17\x4c\x3a\xfd\xfb\x78\xb6\x4a\xbc\x8c\xa0\x0d\x47\x61\x3c\xad\x9f\xf2\x44\x22\xe5\xf6\x57\xed\xb5\x64\x57\xd3\x9e\x46\xdb\xa8\x0d\xc5\x37\x9d\xac\x78\x44\x95\xf3\x1a\xf2\x7c\xd5\x73\xc3\xd3\x6d\x3b\x2d\x95\x5f\x34\x1f\xc3\xff\x63\xfe\x0b\x00\x00\xff\xff\x20\xe2\x79\x21\xed\x01\x00\x00")

func test_certsNodeServerKeyBytes() ([]byte, error) {
	return bindataRead(
		_test_certsNodeServerKey,
		"test_certs/node.server.key",
	)
}

func test_certsNodeServerKey() (*asset, error) {
	bytes, err := test_certsNodeServerKeyBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "test_certs/node.server.key", size: 493, mode: os.FileMode(420), modTime: time.Unix(1400000000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _test_certsRootClientCrt = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x64\x92\xcb\x8e\xda\x30\x18\x85\xf7\x79\x8a\xee\x51\x35\xb9\xd0\x00\x8b\x2e\x7e\x5f\x12\x1c\xf2\xbb\x38\xb7\x92\xee\x20\xc3\x38\x0a\x03\xe1\x92\xc6\x64\x9e\xbe\x01\x55\x55\xab\x7a\x63\xe9\x5b\x1c\x7f\x3e\x3a\x9f\x1f\x87\xf0\x50\xc8\x4f\x94\x27\x99\x08\x04\x85\x8c\x3f\xa9\x85\x42\x90\xaa\xa1\x14\x12\x5f\x83\x11\x04\xb4\x48\x60\x15\x2e\xd2\xf3\xd1\x8d\xc7\xcb\x70\x68\x0e\xc5\x40\x4e\x81\x30\xd4\x94\xd1\xaa\xfd\x21\xea\xbe\x92\xa0\x78\x8c\xf4\x76\xb7\x78\x03\x8a\x68\x59\x10\x68\x33\x9a\xcb\xbe\x74\xbb\x61\xe7\x06\xcd\x96\xf1\x02\x09\x86\xe0\xe4\x1c\xee\x88\xca\x5d\x34\x5b\x2f\xea\xcb\xef\xb2\x15\x5c\x12\x24\xd3\x0d\xcb\xb8\x63\x21\x2b\xef\xb2\xe1\x0e\x36\xb9\x23\x8b\xf6\x01\xdd\x27\xcb\xfe\x30\x23\x3e\x78\x8a\x04\x9e\x61\xb4\xc6\xe8\xef\x30\x0b\x95\xfd\x50\x63\x85\x52\x8c\x9b\x64\x7c\x7d\x61\x63\x60\x0c\x53\xff\xe8\x12\xa2\x80\xa5\x06\x4c\x0a\x11\x01\x79\x9a\x1d\x53\x67\xdf\x5d\xaf\xef\xd6\xa9\xfc\xd9\x5f\xbf\x1d\x5e\x66\xd5\xeb\xba\xa7\xc9\x2c\xfe\x32\x49\x17\xf3\x46\xdb\xfc\x03\x9d\xaa\x99\xd5\xd2\xbd\x4d\xe7\xb7\xf8\x5a\xdb\xb0\x7d\xbb\x74\x2e\x0e\xe3\xcf\x76\x55\xf5\x26\xf7\x5d\xeb\xed\x2c\xef\x6a\xa6\x67\xaa\x8e\x4b\x6a\x53\x18\x1b\x83\x2d\x9e\x90\xe6\x86\xe9\xd1\x2a\xb1\xd7\xa0\x96\x2f\x04\x14\x03\x4d\xf5\xef\x4e\x5e\x23\xa5\x10\xa1\x0d\x29\xbd\x85\x16\xa8\x3c\x20\x06\x29\xc2\x2d\xa4\xe9\x25\x4c\xc5\xce\x63\xa3\x32\x35\x92\x00\x0d\x8e\x49\xd6\xdf\x07\x7f\x8f\xa5\x9a\xcf\x2f\x93\xbd\x99\x2e\x6d\xff\xbc\x89\xeb\x49\x12\x9f\xfc\x60\xc8\x2d\xdf\xf3\xeb\x59\x47\x5e\xee\x98\x5d\x0e\x49\xbd\xc7\xa4\x38\x6c\x74\x7e\x99\xe8\xa9\x26\x8a\xd4\x3c\x5a\x78\xed\x9a\xa5\x7a\x15\xa4\x93\x77\xe8\x76\x47\xf8\x6a\x3d\x17\xc0\x25\xfb\x7f\x15\xbf\x02\x00\x00\xff\xff\xc0\x97\x9e\xaf\x32\x02\x00\x00")

func test_certsRootClientCrtBytes() ([]byte, error) {
	return bindataRead(
		_test_certsRootClientCrt,
		"test_certs/root.client.crt",
	)
}

func test_certsRootClientCrt() (*asset, error) {
	bytes, err := test_certsRootClientCrtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "test_certs/root.client.crt", size: 562, mode: os.FileMode(420), modTime: time.Unix(1400000000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _test_certsRootClientKey = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\xd1\x3b\x97\xa2\x30\x00\x05\xe0\x9e\x5f\x61\xcf\xd9\x83\x8a\xbc\x8a\x29\x02\x46\x01\x0d\x48\x02\xbe\xba\x10\x15\x0d\xaf\x15\x08\x20\xbf\x7e\x77\xa6\x9e\xdb\xde\x73\xee\x2d\xbe\x3f\xdf\xb1\xe1\xd6\x0b\x66\x98\x80\xd9\x01\x7b\x47\x10\xc3\xd9\x0e\x5e\x7e\x1a\x09\x79\x9e\x1d\x0e\x9e\x0d\x80\x6f\x83\xa0\x32\x4a\xb2\xb8\x77\x4d\x53\x54\x17\xd1\x37\x61\xae\x18\xec\x76\xe8\x1d\x6c\xec\x35\x99\x58\x26\xcf\xe6\x70\x42\x0b\xc6\x8d\x67\xb0\x6c\x57\x66\xbb\x97\x9a\xe7\x1c\xd0\xc7\xbb\x5b\xa2\x0f\xb2\x51\xca\xd8\x23\xb8\x77\xb5\x9a\xaa\xcd\xb0\xfa\xeb\x44\xa5\xeb\xcc\x1d\x30\x40\x00\xa2\xef\x03\x9a\x31\xd9\xf5\x10\x40\x8b\x96\x97\x55\xba\xe2\x9a\x24\xe3\xe4\xa8\xe2\x2a\x64\xe8\x74\x53\xd5\x8e\x9e\x8d\x95\xdc\x46\x8e\x9d\x28\xf3\x32\x9d\xa6\xe9\xea\xb8\xec\xcc\x50\x4f\x87\xdd\x75\x54\x78\x57\x08\x76\xb7\x97\x5a\xc7\xe4\x08\x87\x28\x85\xd2\xc6\x84\x8e\x17\xad\xf5\x3e\x52\x52\xef\x52\xe5\x57\xf2\xcc\x04\x2d\x88\x1d\x9f\x44\x8d\x8a\x74\xab\x8d\xb9\x40\x8e\xb5\x29\x85\xab\xfb\xd8\x77\x06\xef\x09\x02\x3d\xcc\x35\x76\x32\xb2\xfd\x5a\x0a\x60\xbc\xd7\x73\xc8\xb7\x62\xa3\x1a\xd6\x88\x5e\xb1\xcf\x7a\x58\xbd\x3a\x4b\x2c\x9e\x64\x57\x81\x17\x78\x8d\x38\xee\x26\xce\x23\x73\xba\x94\x48\xe3\x2e\xee\xb7\xef\x91\xc2\x32\x23\x4b\x21\x7d\xee\xca\x40\x76\x22\x6b\xb0\xc2\x41\xf4\x7f\x1d\x59\x80\xb6\x31\xed\x57\x7a\x65\xa7\x8f\x35\x4e\xc8\x23\x29\x5a\x20\x9b\xb7\xc2\x7f\x5f\x8f\x35\xad\x13\x4e\x0c\xd5\x25\x5e\x01\x5e\xb6\x2f\x05\x4e\x74\x08\x47\xd3\x48\x5b\x91\x50\xf2\x19\xd7\x8b\x43\x9d\xf7\xb1\xd9\x94\x53\x6d\xdd\xc6\xf8\x26\xce\xc0\xdd\xe5\x27\x1d\x7c\x7d\x49\x3f\x70\x30\x58\xff\x0e\xfa\x2f\x00\x00\xff\xff\x5f\xdc\x33\xee\xf1\x01\x00\x00")

func test_certsRootClientKeyBytes() ([]byte, error) {
	return bindataRead(
		_test_certsRootClientKey,
		"test_certs/root.client.key",
	)
}

func test_certsRootClientKey() (*asset, error) {
	bytes, err := test_certsRootClientKeyBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "test_certs/root.client.key", size: 497, mode: os.FileMode(420), modTime: time.Unix(1400000000, 0)}
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
	"test_certs/ca.crt":          test_certsCaCrt,
	"test_certs/ca.key":          test_certsCaKey,
	"test_certs/node.client.crt": test_certsNodeClientCrt,
	"test_certs/node.client.key": test_certsNodeClientKey,
	"test_certs/node.server.crt": test_certsNodeServerCrt,
	"test_certs/node.server.key": test_certsNodeServerKey,
	"test_certs/root.client.crt": test_certsRootClientCrt,
	"test_certs/root.client.key": test_certsRootClientKey,
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
	"test_certs": {nil, map[string]*bintree{
		"ca.crt":          {test_certsCaCrt, map[string]*bintree{}},
		"ca.key":          {test_certsCaKey, map[string]*bintree{}},
		"node.client.crt": {test_certsNodeClientCrt, map[string]*bintree{}},
		"node.client.key": {test_certsNodeClientKey, map[string]*bintree{}},
		"node.server.crt": {test_certsNodeServerCrt, map[string]*bintree{}},
		"node.server.key": {test_certsNodeServerKey, map[string]*bintree{}},
		"root.client.crt": {test_certsRootClientCrt, map[string]*bintree{}},
		"root.client.key": {test_certsRootClientKey, map[string]*bintree{}},
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
	err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
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
		err = RestoreAssets(dir, path.Join(name, child))
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
