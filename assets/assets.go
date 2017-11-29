// Code generated by go-bindata.
// sources:
// markets/market/bitfinex/currencies.json
// currencies/currencies.json
// DO NOT EDIT!

package assets

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

var _marketsMarketBitfinexCurrenciesJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\xd4\xbd\x4e\xc3\x30\x10\x07\xf0\x3d\x4f\x61\x65\xae\x87\x76\x64\x2b\x4d\x04\x48\xb4\x81\x26\x41\x15\x88\xe1\x9a\x9c\x5a\x4b\xfe\x28\xf6\xa5\x12\x20\xde\x1d\x95\xb2\xe5\x4e\x59\x13\xff\x6c\xff\x7d\x67\xbf\x65\x4a\x7d\x67\x4a\x29\x95\x7b\x70\x98\xdf\xa8\x1c\xce\xe8\x69\x48\xf9\xec\xfa\x39\x7d\xba\x7d\xb0\x97\x1f\xcb\x97\x26\xcf\x94\xfa\x99\x8d\xcd\xde\x50\x17\x8c\x1f\x9b\xdb\x66\x35\x61\x74\x07\xe9\xc8\xc0\xd5\xfd\x14\x3c\x04\xdb\x73\x2b\xde\x4d\xc1\xc1\x5b\xe3\x0c\x21\xa7\x57\xad\xa8\x69\xae\xbb\x44\xdc\x8a\x73\xd9\x2c\x24\xb3\x90\x4c\xcf\x9e\x47\x51\x8b\xe7\x81\xa6\x0f\x61\x2c\xca\xa2\x12\x45\x60\xca\x5b\x56\xb5\x38\x9e\x8e\x18\x71\x70\x0c\x6a\xe4\x6d\xfd\x23\xdd\x59\x48\xc9\x74\x1c\x16\x9b\x03\x87\xc8\x45\x6a\xb7\x12\x30\x81\x60\x0c\x1e\x2a\xb1\x65\xad\x21\xe4\x7b\xf6\x51\xde\x96\x43\x82\x33\xc6\x84\x5c\x98\x27\x51\x05\x8f\x5c\x9c\xdd\x5a\x8c\xe3\x91\x19\xbf\x29\xc5\x8a\x06\x67\x12\x1e\x18\x53\xad\xc5\xeb\xf0\x41\x5c\x45\x9f\x9b\xb5\x04\xa2\x39\x9d\x2c\x13\x7d\xb7\x15\xa3\x27\xf0\x64\x1c\x7a\xe6\x0a\xd4\xcb\x8d\xa8\x28\x22\xb8\xa8\x7b\x20\xe0\x6b\x54\x2c\xc5\xc2\xd2\x5f\xe3\x8d\x49\x5b\x17\xa2\x19\xfc\xe5\x39\xd0\x89\x80\x30\xe9\x3e\x58\x0b\xfc\x0c\xd2\x04\x5f\xfc\x2b\xf6\x5a\x5e\x5b\x29\x7b\xcf\x7e\x03\x00\x00\xff\xff\x3a\xe1\xfe\x25\x6a\x05\x00\x00")

func marketsMarketBitfinexCurrenciesJsonBytes() ([]byte, error) {
	return bindataRead(
		_marketsMarketBitfinexCurrenciesJson,
		"markets/market/bitfinex/currencies.json",
	)
}

func marketsMarketBitfinexCurrenciesJson() (*asset, error) {
	bytes, err := marketsMarketBitfinexCurrenciesJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "markets/market/bitfinex/currencies.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _currenciesCurrenciesJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x7d\x5f\x6f\xe3\xb8\x92\xfd\xfb\x7c\x8a\xc6\x3c\x5f\x01\x77\xfa\xde\x1f\x30\xf8\xbd\xf9\x8f\xe2\xe8\xc6\xb6\xdc\xb6\x9c\x38\x59\xec\x03\x25\x95\x65\xb6\x29\x52\x4d\x52\x4e\x9c\xc5\x7e\xf7\x85\x9c\xee\x59\xec\xf6\x39\xea\x7d\x1a\xa0\x31\x27\x14\x8b\x55\xc5\xaa\x53\x55\xf4\xbf\xfd\xf6\xe9\xd3\x7f\xfc\xf6\xe9\xd3\xa7\x4f\xbf\x5b\xd5\xca\xef\xff\xff\xd3\xef\x7f\x7f\xfb\xfd\x6f\x1f\xff\x12\xae\x6d\xe9\xcc\xf0\x6f\x2f\xdb\xc3\xef\xbf\x7d\xfa\xf4\x9f\x7f\xfb\xf9\x7f\xff\xdc\xe8\x8b\xfc\x8c\xf8\xbc\xc8\x1e\x53\x86\xf9\xe7\xe7\xa4\x72\xda\xfe\x8c\xfa\xe7\x67\x06\xf9\xf3\xef\x7f\x62\xc8\x9f\x7f\xff\x93\x61\x54\xf9\x15\x63\x26\xd3\x7f\x51\x4c\x55\x89\x11\xaf\xa2\xf3\x89\x95\xf8\xea\xfc\x19\xe0\x67\x33\x8e\x3f\x29\xb8\xe4\xac\xa0\x90\x5a\x4c\x77\x72\x1a\x80\xe6\x4b\x0e\x8a\xee\x2c\x68\xa1\xf9\xc8\x42\x6f\xc3\x9e\x10\x86\x1e\xaf\x12\x87\x16\x49\xf3\x35\x47\x78\xdd\xb7\x08\xb3\x5d\x8d\x60\xe0\x32\xdb\x91\x55\xa2\x78\xab\xe3\x15\x2d\x44\x41\x8d\xf3\x2a\x24\x37\xc9\x05\x00\x5c\x6c\x77\x1c\xea\xc5\x18\x97\xd4\x62\xa2\xfa\x19\x3a\x5f\x72\xa9\xeb\xda\x85\xe4\xdc\x5b\xc1\xc7\xf5\xc0\x81\x50\x26\xd9\x88\xe8\xb5\x0f\xaf\xaa\x03\xa0\xdd\xc8\xe7\x79\xa6\x49\xd9\x96\x82\xce\xfd\x55\x11\x1b\x9e\x3c\x3c\x53\x98\xd1\x48\xec\xcb\x8c\x8b\xdd\x7c\x73\x08\xf1\x25\xe7\x88\x58\xb9\xb6\xed\x07\xdd\x60\x5f\xb8\x2c\x66\x39\xd7\xc5\xb6\x14\x0f\x40\xab\x29\x45\x58\x7f\x7d\x07\x3e\x70\xfb\xfc\x42\x21\x1d\xf0\xb2\x93\x0d\x37\x43\xaf\x1a\xa8\x0d\x6b\x7e\xb0\xbe\x8a\xba\x22\x12\xd8\x72\x07\xe6\x6b\x87\xb6\xbf\x9d\x73\x6d\x80\x3e\x72\xcb\x15\xdb\xc7\xf2\x1a\x81\xc0\x26\x53\xae\x3a\xa1\x3a\xfd\x0c\x38\x4c\xb8\xe6\x84\xe8\x91\xea\xec\x8a\x2d\xd7\x9d\x58\x12\x79\x15\xfc\xf0\xa3\x51\x16\xf9\xd4\x82\xfb\xee\xd8\xb2\x65\x56\xfc\x5c\x62\xeb\x90\xf1\x14\x2b\x2e\x82\x18\xc5\x46\xed\xec\x87\xc7\x4b\xdc\x31\x69\xa5\xd6\xc0\x7d\x4d\x0a\x6e\x0e\x7d\xd3\x03\x7d\xd8\xa6\x1b\x8e\xf0\xce\x2b\xb2\xc3\x3d\xd7\xa2\x3e\x3a\xab\xc1\x91\xad\x33\x7e\x60\x17\xb1\xb1\x47\x62\x79\xa4\x76\x51\x7e\x2e\x81\xed\x4d\x3f\xd3\x13\x2e\xff\x81\xf7\x32\xfd\x07\x45\x28\x5b\x21\x23\x9a\x72\x6b\x2d\x55\xd0\x55\xf2\xbf\x4e\x0c\xfc\x85\xc9\xc8\x5f\x40\x17\xd3\xb4\x98\x50\x40\x85\xae\x8a\xe9\x6c\x42\x4f\xb6\x14\x43\xce\x75\x9a\x2e\xf9\x3a\xda\x18\xed\xac\xd2\x5e\xd8\xb6\x0e\x53\x6a\x2c\xa5\xb6\xca\x56\x42\x5c\xf9\x74\xcd\x8f\x4d\x87\xb6\x8f\xc0\x69\x4c\xf9\x75\x53\xea\x58\x2a\x10\x52\x4c\x27\xd4\x31\x0d\x10\x51\xe8\xd3\x32\xee\x34\x06\x90\x73\x01\xb8\x8d\xe9\x94\x1f\xb0\x8e\x65\xac\xf0\x42\x05\xf5\x1b\xa5\x8e\x95\x71\x7d\x8d\x54\x83\x87\x7d\x03\xca\x22\x41\x64\xc5\x6c\x3d\x26\x8b\xca\x29\xf3\x33\x6c\x96\x4f\x46\x0e\x38\x92\xb3\x1d\xdf\x94\xd3\x36\xa9\x54\x40\x07\x3c\xbb\xff\x15\xb0\xd6\xaa\x75\x16\x09\x65\x36\xff\x15\xb6\x71\x06\x4a\x73\xf1\x2b\x60\x67\x90\xab\x3a\x4c\x7f\xb9\x4d\x2f\x70\xc1\xd9\x36\xfd\xe5\xc7\x06\x6d\x2e\x28\x9c\x99\x16\xb3\x31\x2b\xb8\x61\x7b\x6b\x74\xab\x23\x5e\x7c\xff\x0b\x74\x0d\xa3\x82\x69\xf1\x4b\xf9\xbe\x43\x14\x8d\xa4\x6e\x28\x6b\xa5\x42\xc6\xc4\xf3\xb4\x1b\xcc\x83\x18\x64\x5a\x8c\x1a\x85\xbf\x86\xa8\x0c\x38\xc6\xe9\x6c\xcc\x2c\x6a\x41\x66\x31\xe5\xf9\xdd\x80\xd1\x15\xf8\xbe\xd9\x6e\x4d\xaf\xc3\x52\xc7\x96\x88\x9d\xde\xef\xa5\x8e\xdf\x90\xa1\x7f\x19\x01\x04\x81\xa6\xb3\x1b\x3b\xda\x20\x5e\x43\x19\x14\x34\x5d\x1b\x50\x27\xe5\x05\x49\xbb\x18\xd3\xde\xe0\x14\x52\xf9\x1d\x8d\x3d\x06\xcc\xab\x3e\x02\x15\xda\x3d\x65\x77\x63\x1e\xb9\x0f\x48\x12\x59\xb1\x1f\x15\xc6\xbb\x20\xe7\xfa\x32\xe2\x59\x8d\xaa\xce\xc4\x4d\x2e\x69\xa0\x7d\x43\xb5\xce\xd9\xa4\xf2\xd7\x2e\x82\xd8\x6a\xca\x03\xce\xd2\xe8\xf8\x4e\x1c\xec\x32\x2b\xb8\x41\x1a\x57\x9d\x2b\x05\x64\x39\x1b\x09\x5e\x06\x50\xab\x02\xca\x71\xa6\xb3\xcd\x38\x0e\x72\x1a\xd3\x65\x3e\x1b\x11\x8c\xab\xce\x9d\x73\x48\x1f\x37\xdc\x26\x07\x54\xd4\x20\x7a\x2c\x32\xee\x34\x5c\x7d\x02\xd4\xce\x34\xe7\x7b\x72\x81\x9c\x74\xce\x95\xde\x8b\x3a\xbb\x1e\xc9\x81\xe7\x61\x3f\x40\x49\x88\xea\x8c\x5c\x21\x67\xfd\x4a\xaf\x59\x24\xb8\xcd\x78\x24\xe8\x75\xdd\x08\xc1\xcd\xb8\x5b\x8b\x7f\x24\x15\x0c\x96\x8a\x3f\x38\xe6\x33\xc3\x50\x8e\xb1\x8c\xd5\x60\x2c\xf0\xfe\xe1\xfe\xb3\xf7\x70\x99\xfd\x96\x73\x2e\x65\xff\xfe\x4e\xa4\xb0\x7f\xe1\x86\x75\x8d\x52\x2a\x03\x74\x76\x31\x7d\x1e\xf1\xa2\xd7\x28\x95\xa0\x14\x75\xfa\xcc\x8d\x7f\x00\x0d\x31\x40\x59\xc1\xb3\xa2\x14\x54\x79\x8d\x0e\x90\x7f\x23\x17\x50\xa5\xac\xaa\xb5\xb2\x49\xed\x8c\x41\x7e\x7b\x36\xa1\xae\xb4\x52\xd6\xaa\x52\x13\x7b\x99\x4d\xd6\xf4\x3b\x2b\xd5\x79\x4d\x60\x1b\x2a\x96\x4a\xf9\x5a\x59\x44\x2b\xcc\xa9\xd2\x57\xca\x47\xf5\xa6\x59\xea\x33\xe3\xf1\x46\xa5\x82\xb6\x8e\x7c\xe4\x8e\x7f\xa4\xd8\xe8\x41\x2a\x38\x2b\xe8\xfd\x37\x40\x7a\x0f\xf9\xc6\x19\xcf\x59\x6f\x54\x37\xf9\xbc\x7b\xfe\x79\x03\xca\x68\x0b\xc2\x94\x65\xb6\xa6\xce\xaa\x3a\x29\xdb\xa0\x78\x68\x42\xe3\xee\xea\xa4\xad\x4a\xae\x28\x39\x1b\x49\x62\xaa\x93\xee\x40\xc4\x31\xbb\xcf\x36\xd4\xfd\x56\x27\xef\xac\x2b\x15\xda\x54\x91\xad\xa8\x6d\x56\xda\xd6\xba\x52\x11\x91\x05\xb3\x35\x57\x7b\xed\xab\x5e\xc7\x90\xb8\x63\x72\x51\xa6\x47\x62\xc9\x1f\x79\xca\x55\xe9\x8b\x06\x89\xe4\xec\x91\x1f\x9a\x51\x2d\x12\xca\x72\xc2\xed\xda\x88\xf2\x9d\x43\xfe\x6a\x93\x2f\xf9\xb7\x19\xa7\x48\xa4\x33\x5b\xe6\x13\xae\x1f\xa6\x27\xe4\xdd\x6c\xb9\xa7\x89\x78\xe5\x4a\x6d\x4f\xce\x81\x20\x6e\x96\x8f\xa0\x8e\xae\xb7\x75\xa2\x51\x90\x73\x97\x71\xd8\x90\x16\xa1\x80\x6a\xc6\x8b\x36\x95\x33\x2e\x84\xfe\xe6\xe2\xde\xd0\x7a\xf9\x92\x3b\x10\x67\x6b\xb1\x41\x21\xae\x75\x3b\xc9\xb8\x73\x74\x9d\xae\x80\x97\x3b\xcc\x36\xf4\x92\xae\x1c\xf1\x1f\x39\xf7\x3a\x2e\x20\x9d\xca\x77\xdc\xce\x5c\x6f\xa3\xf8\x4e\x79\x54\xf8\x39\xcc\x28\x73\x55\x79\xa9\x85\xd3\x48\xb3\xed\x7c\xcd\x75\xdf\x4b\xad\x63\x09\x0f\x7b\xcb\x75\xc4\xbb\x57\xb8\xd2\x13\x47\x0c\x31\x7a\xa5\x7c\x09\xe5\x38\xba\xd6\x80\xb4\x02\x22\xd3\xd9\x9a\x6b\xc7\x07\x4a\x23\xed\x38\xf0\x1b\xfe\x03\xd6\x69\xdb\x00\xc3\xce\xd6\xdc\x1b\x7f\xa4\xce\xc9\xcd\x2f\xa0\xfd\x71\x13\xe8\x3d\x09\x1a\x67\xfb\x2d\xf5\xad\xb5\x72\xc9\xc7\x0d\x8a\xe8\x4a\xba\x58\xad\x7c\xe5\x11\x4f\x33\x9f\x50\x45\xc6\x46\x3d\x9f\xec\x28\x0f\x55\xab\xa8\x86\xa4\x13\xa0\x38\x6b\x58\x57\xce\x03\xc6\x76\xbe\xa5\x6a\x5f\x93\xc8\x6f\xce\x85\xfd\x01\x49\x4a\x94\x53\xcd\xc7\xe4\x26\x1f\x71\x87\x51\x88\x1b\x58\x4d\xd6\x34\x40\xaa\xa5\xd2\x61\x84\xf3\xbe\xe7\x81\x74\x2d\x15\xe4\xc6\xe6\x33\x7e\x56\x22\x9d\xb3\xd0\x55\xe5\xeb\x91\xda\xea\xad\x48\x0f\x9c\xce\x7c\x43\x23\x89\x5a\xac\xf2\xba\x0f\x49\x6d\x81\xc2\xcf\xd7\x23\xdf\x08\x0f\x2d\xe5\xa1\xd8\x80\x20\x09\xd9\x9c\x9b\x72\x2d\x41\x23\x0e\x6c\xce\x99\x92\x5a\x7a\x12\x6d\xcf\xd3\x3d\xf5\xdb\x94\x6e\x9d\xaf\x68\x90\x53\xeb\x46\xe3\x32\xe1\x7c\xc1\x4d\x44\x37\x3a\x2a\x63\x1d\x74\x6a\x73\x2e\x89\x0f\xdc\x90\x13\xa0\x05\xb9\x81\xe9\x46\xbf\xd5\x0a\xb8\x98\xf9\x62\x64\x67\x21\x7a\x5d\x45\xd4\x69\x33\x1f\x39\x63\xc7\x72\xe7\x79\xbe\xe0\x6e\xd0\xb5\x5e\xe9\x1a\x71\xc0\xf3\xed\xc8\x5a\x1d\x5d\x6b\x33\xb2\x16\xa1\xf2\xe7\x9c\xed\xa8\x5d\x5f\x1a\x9c\x75\xe7\x9c\x2f\xaa\xdd\xa5\x47\xcb\x3c\x52\x5a\xba\xf6\xea\xc8\xbe\x8e\x33\x7c\x75\x5f\x2a\xfd\x51\x31\x28\x11\xf1\x33\x9f\x72\xe6\xa7\xbe\x5a\xd5\xa2\x38\x7b\xfe\x4c\xf5\x50\x92\x5a\x5b\xe5\x49\x98\x92\xf2\x1a\xbc\x24\x4d\x6f\x6a\xe4\x3e\xd3\x3b\x1a\x6b\x0b\xeb\xe1\x48\x79\x52\x25\xaa\x31\x44\x35\xd2\xc9\x62\x49\x75\x43\x94\x8f\x27\x86\xa3\xb1\x97\x8c\x57\x7a\xd2\x91\x52\x8f\xdc\x2a\x6c\x64\xc5\x29\xbf\x53\xa4\x8c\x95\x95\x57\x04\xe2\x05\x29\xa9\x48\x9e\x9e\x72\x5a\x4b\xea\x46\x8c\xa0\x10\x38\x9d\x73\xf1\xeb\xda\x01\x6f\x93\xce\xf9\x32\xda\x86\x28\xda\xc2\xa6\xac\x74\x35\xa3\x74\x98\x18\x15\x22\xd2\xdf\x43\xca\x35\xca\x48\x15\xbd\xb3\x02\x17\x2b\xb8\xde\x1b\xfd\xa6\x81\x8b\x4a\x97\xdc\xbe\xc4\x18\xa5\x03\x5a\x67\xc9\x8b\xd1\x62\x98\x4a\x2c\x8b\x59\xce\x53\x22\x69\xc5\x13\x20\xa7\xd0\xa5\x85\x21\x43\xca\x43\x06\xb1\x43\x78\xdd\x20\x3b\x9e\xf3\x08\x4a\xac\xf8\x06\x68\x45\xb1\xe3\x07\x35\x40\xae\x64\x47\xeb\x2d\xd7\x3f\xab\x9b\x56\x25\x9d\x77\x5f\x61\xc9\x2d\xe5\x09\x80\xd8\xaf\x83\x21\x93\x25\x69\x47\xa6\xa0\xb6\x97\x94\x53\xe1\xe2\x22\x8b\x24\x53\x7e\x0d\xc9\xb7\x5e\x47\xaf\xe0\x3d\x99\x7e\xe1\xb0\xd0\x89\x47\xdf\xb7\xa3\x21\x03\x6f\x5c\x1c\x89\xef\x24\x9e\x70\xba\x90\x16\x53\x2e\x89\x78\x12\x4f\x2c\x91\xbb\xcd\xef\xa0\xa4\x84\xbc\xd2\x74\xb9\xe7\x2e\xfe\x07\xb4\x32\x2a\x04\xe4\x39\xd2\x11\x0f\xfa\x03\x8c\x0b\xc7\x69\x71\x4f\xe3\xa9\xbf\xa0\xb8\x32\x9f\xf2\xca\xfc\x0d\x09\x29\xaa\x79\x36\x1b\xdf\x27\xe1\x20\xee\xf6\x23\xce\x31\x9e\xa0\x2e\x17\xf7\x23\xda\xdc\xa3\x2e\xb6\x94\xb7\x51\x0d\x00\x16\xbb\xa5\xbc\xed\x4f\x2e\xe2\x11\x7f\x90\x3e\x72\xff\x7b\x11\xdf\x78\x11\x42\x01\xa7\x8b\xd1\xc5\x50\x8a\x9d\x3e\xf2\x6d\xbd\x7d\x10\xc0\x49\x8f\x73\xb7\xc3\x9e\xaf\xf6\x56\x99\x3e\xe8\x0b\x13\xca\x61\xc6\x4f\xec\xad\x53\x36\x00\x33\x48\x0f\xd4\xbe\x8f\xaa\x82\x25\x90\x3b\x9e\x73\x1f\x95\x26\x97\xcb\xdd\x48\x13\xee\x51\xd9\xe8\x48\x27\xe1\x1d\x4f\xfa\x8e\xca\x2b\x60\x25\x77\x5b\x6a\x5f\x47\xe5\x1b\x12\xde\xdc\x6d\xf9\x29\x1f\x15\x0b\xbf\xee\x78\xf4\x75\x14\x35\x58\x17\xc1\x71\xf7\x71\xd4\x3e\xc4\x21\x91\x00\x5b\xfb\x63\x64\xb9\x01\xc6\xb6\x36\x02\x33\x1a\xf8\xa8\xbb\x65\x46\xe9\xe9\xa3\xd1\x6f\x6f\xc0\x92\xef\x96\xd9\x81\x5a\xd8\xd1\x38\xcf\x2a\x2c\x77\x4b\x1a\xf1\x1d\x9d\xa9\xb5\x6d\x18\x6e\xce\x85\xd8\x7f\xd5\x16\xb5\x09\xac\x0b\xbe\x58\x6f\xeb\xab\xeb\x7d\x10\x73\xb4\x0e\x84\xcd\x77\x3c\xdb\x39\xf6\x76\x50\x7d\xe4\x45\x29\xa6\x51\x2d\xa4\x62\x17\xbc\x12\xd1\xa8\x56\x3e\x48\x5c\xe0\x73\x16\x13\x5e\x9e\x69\x14\x04\x50\x67\xdd\x88\x95\xa0\x43\x72\xb9\xf1\x59\x00\xca\xbb\x52\x1b\x21\x16\xb6\x48\xa9\xe8\x1b\xdd\xa8\xe4\x55\x45\x1a\xef\x3c\x15\x7c\xbd\xe1\x8a\x26\x2b\x2e\xf9\x06\x8d\x2b\x95\x49\xaa\xde\x7b\xb1\xd5\x35\xf1\x12\xc4\xa3\xf1\x9f\x05\x27\xe0\x1a\xeb\x06\x19\x35\x88\x97\x5d\xf0\xde\xa6\xc6\x61\x32\x68\x31\xa5\xe6\x33\x84\x03\x89\x97\x57\xe5\x6b\x26\xa1\x05\x6f\x71\x18\xd0\xb7\x8e\x0f\xa4\x03\x94\x80\x1a\x50\x4c\xac\xd4\xc1\x0e\xa0\xdb\xbd\x46\x90\xfc\x6e\x6b\x9c\x91\xf6\xc7\xa0\x12\x1d\x6c\x59\xf0\xc0\xb2\x71\x06\xc5\x24\x8b\x7c\xc9\x63\x92\x1b\x86\x04\x5b\x8b\x29\x0d\xb6\x1a\xaf\x11\x60\x9b\x71\xb9\x78\xcd\x84\xc9\x23\x99\xc6\x3b\x09\xd1\x30\x1c\xdf\x15\xe1\x4f\xd6\x4b\xbe\xa3\xde\x1c\xf1\x3a\x2b\x7e\x2d\x36\x7d\xd7\x81\xc0\x71\xb1\xa7\x11\x45\xf3\xc6\xda\xe0\x16\x07\xba\x9d\x93\xea\x3a\x92\xdf\xdd\xf3\x06\x87\x93\xa8\x98\x18\xa9\x1b\x94\x07\xdd\xa7\x9c\x90\x3b\x0d\x18\x00\xe1\x4c\xc6\x49\x8c\x11\xcb\x26\x54\xee\x79\x3d\xee\x24\x6d\x87\x41\x05\xef\x3a\x38\x91\x66\x8f\xc3\x3d\x2d\xd7\x9e\xe0\x60\xe3\xfd\x23\xbd\x9a\x4e\xb7\x63\x42\x05\x0c\xea\x0c\x4f\xb1\x25\x9a\x7a\x5f\xac\x96\xff\x8f\xc2\xfa\x56\x59\xfd\x0d\xa0\x56\xb4\x6f\xf3\x74\x2b\x9d\x92\xc5\xb8\x87\x39\xf5\x88\x78\xbb\xdf\xf3\xda\xd6\xe9\xda\x89\x4f\xe2\x05\xa0\x9e\x8b\xc7\x51\x14\x69\x44\xbb\x7f\xa6\xb6\xa1\x2b\x74\xd1\x66\x33\xea\xd6\x07\x80\x6b\x41\x47\x5e\xc6\x43\x66\x5d\x21\x1f\x99\xcd\xb8\x8b\xd4\x3c\xf7\xce\x46\x92\x6f\x6d\x71\xbd\x2e\x5b\x8f\xb4\xe1\x68\x5b\x3b\x1f\xe8\x6c\x45\xc6\xfb\x48\xb4\x3d\x92\x71\xd5\x6c\x7d\x37\x02\xd2\xb7\xe1\x3a\xf9\x90\x64\x85\x7a\xda\x39\x91\xa6\xed\xd1\xf4\x6f\x58\x0d\xb3\xf5\x1d\x3f\x37\xd4\x5b\x33\xd2\x2f\xa4\xad\x75\x17\xd0\x05\x95\xf1\x9e\x30\x3d\x18\x88\x95\x98\xb8\x63\xd2\x89\xeb\x0c\x50\xc5\x2c\xe7\xaa\x78\x83\xa3\x1e\x8d\x6c\x7d\x18\x39\xbd\x38\xe8\xbd\x18\xa3\x7c\x72\xfa\x88\xdc\x81\x44\xef\x73\x1e\x3e\x68\x7b\x91\x10\x8f\x82\x2a\xa1\x23\xdd\xcb\x9a\x44\x9c\x59\x4e\x9d\x01\x8c\x6a\x47\x8a\xa6\xda\xa1\x99\xa4\x55\x96\xf3\xa9\x24\xfd\xc6\xae\x9f\x8c\x0b\xf1\xab\xea\x94\x95\x20\xc9\x15\x59\xc0\xbf\x38\xdf\x7a\x56\xbe\x74\xaf\xca\x0a\x20\xab\x1e\x78\xc3\xc5\x59\x57\x67\xd8\x21\xf3\x90\xf1\xa6\xe3\x33\x12\xf6\x03\xb7\x94\xb3\x33\x50\xde\x0f\xcb\x11\x48\xeb\x6a\xf4\x59\xbc\xcc\x7a\x76\xac\xb9\xe2\x21\xe7\xcd\x15\xe7\xfe\x63\xba\x84\xc4\x25\x0f\x7c\xb8\xe4\x7c\x2d\x65\x64\xb4\xfe\x81\xdf\xfb\x46\xb5\x1d\xaa\xc1\x6d\x78\x89\xc0\x28\xe2\x14\x79\x81\xc0\x48\x23\xb6\x0e\x89\x77\x88\xc4\x59\xf2\xaa\xae\x61\x19\xdc\x92\x67\x70\x46\x97\x5e\xf9\x6b\xf2\x91\x9d\x02\x28\x9f\x0d\x32\xfa\x08\xdc\xd3\x32\xbb\xa3\x67\x66\xb4\x05\xd2\x5b\x66\xbc\x43\xc8\xe8\x80\xba\x34\x77\x54\xc5\x8d\x8e\x42\xe7\xb9\x96\x23\xf5\xb3\x01\x48\x50\x23\x20\xc3\x4e\x97\x4f\xe3\x1b\xd7\xb2\x63\xe2\x75\x1c\xe3\x5c\xe7\x61\xc3\xd3\x92\xe7\x21\xa6\xb7\x57\xe0\xc5\x96\x9c\xda\xa0\xf7\xe2\x72\xcf\x8f\xe8\x7a\x46\x11\xd3\xf2\x81\x9e\x51\xab\x2a\xef\xfe\x00\x5e\x99\xa7\xec\xad\x92\x4a\x2c\x22\x43\x26\xbc\x85\xa0\x55\x0d\x88\xb2\x0e\x2b\x9a\x10\xb4\xaa\x81\x93\x1c\x2b\x5e\x8c\x6e\x95\xae\x83\x3a\x12\xd5\x59\x4d\x78\x6a\xd9\xaa\x30\xdc\xd6\xae\x66\xd8\x82\x3b\xa2\x56\x91\x73\x5a\x4d\xe8\x39\xb5\xea\x9d\xf4\xe7\xac\x5e\xf8\x42\x70\x42\x77\x35\x32\xa1\xdb\x4a\xc3\x96\x49\xf9\x32\x62\xd0\x3d\xb3\xe2\xf7\x4c\x2b\xad\xc0\xd2\xf0\x2a\xe5\x6c\x5a\x2b\xbe\xea\x3d\x48\x7d\x57\x29\xd7\xbc\x0f\x4c\xd2\x79\x17\x5d\x85\xe6\x75\x16\x2b\xae\x81\x12\xd1\xc4\xd9\x8a\x4f\xe6\x0f\x88\x8b\x78\xc8\xf8\x17\x5c\xe8\xda\x92\x58\x7f\xb5\xe6\xd2\x18\x40\xe1\x17\xf4\xd4\x6a\xc4\xbc\xb4\x15\xa6\x86\xdc\xa3\xb7\xda\x12\xef\xbc\xca\x78\x86\xd1\xba\x52\x1b\x41\x15\xe6\xd5\x82\xde\x6d\xad\xab\xa1\x48\x72\x6e\x93\x4e\x6a\x95\x18\x77\x55\x26\x5e\x93\xce\x69\x8b\x08\xdb\x15\x1f\xb5\x68\x9d\x55\x28\x2a\x5b\xf1\x26\x8c\x0f\x08\xe1\x06\x56\x7c\x7e\xf4\x03\x07\x05\x99\xf3\x42\x7d\xeb\xac\xa0\x7a\xde\x61\xc5\x2d\xc0\x59\x89\xca\x5f\x7b\x8b\x02\x84\x15\xaf\xc6\xde\x80\x27\x14\x7e\xf3\xfc\xb3\x75\xf6\x2c\x57\x5e\xdd\x5f\xe5\x3c\xe1\x6a\x1d\xe3\x2f\x57\x39\xcf\x10\x5a\x37\x24\xca\xe1\xa4\x91\xa3\xe3\x45\xf4\x16\x8d\x5d\xae\xf8\xcc\x65\xdb\x07\x42\x00\xad\xf6\xbb\x8c\x9f\x72\x1f\x7b\x65\x48\xcf\xc2\x6a\xcf\x0d\xe6\x5a\x6a\xca\xca\xaf\x9e\x69\x5e\xd1\x5e\xbd\x46\xc5\xb9\xc3\x8a\x66\x30\xed\x75\xb8\xc8\x60\x73\xcf\xea\x99\x57\xb0\x86\xff\xe2\x6d\xad\x79\x1c\x64\xd5\x85\xc8\x62\x3d\xa1\xc4\x8e\x95\xd2\xc0\x67\x3d\x52\xfe\xe8\x83\x15\xb0\x9d\x43\x4a\x2b\x3c\x56\xe0\x02\xd4\xea\xad\x38\xc2\x5f\xaf\xd3\x9c\x46\x1b\x56\xd8\x64\xe6\x7a\xa4\x1f\xc5\x4a\x3c\xc3\xaf\x2b\x1e\x46\xbe\xaf\x8f\xf0\x45\xac\x75\xc1\x9f\xc4\xb2\xf2\x06\xd5\x60\x7d\xe0\x07\x2a\x6f\xa8\x0e\xbf\xe6\xb4\xb2\xd5\x2d\xe2\x1d\xd7\xbc\x31\xdc\xba\xdb\x2b\x02\x44\x70\x4b\xde\x7f\x66\xdd\x85\xb8\xd8\x35\x1f\x54\xb2\x3d\xee\x93\xd9\xef\xd6\xfc\x65\x0d\xdb\xa3\x89\xfe\xf5\x9e\x17\xc2\x6c\xdf\x8a\x57\xb0\x79\x7a\xcd\x7d\x39\x24\x87\x46\xb8\x21\xa7\x40\xc2\x96\xf3\x20\xd3\xe1\xbd\xe7\xd3\x8c\x8f\xcb\xbb\x32\xe8\x5a\xa3\x19\xb9\x9c\xb7\x4b\xbb\x4a\x94\x35\xaa\x04\x20\xde\x4c\xe1\xce\xb8\x89\x33\xa7\x57\x8a\x6b\x75\x80\x61\x47\xce\xb3\x08\xd7\x5a\x90\x76\xe4\xab\x35\xe5\xf9\x9d\x45\xac\x42\xce\x83\x28\xd7\x89\x4d\xa2\x57\xb5\xb6\x0d\x67\x32\x72\xde\xed\xe8\x3a\x64\x75\x39\x1f\x73\x77\xbe\x95\x3e\x10\xdf\x9b\x6f\x79\xcc\xed\xde\x48\x0d\x28\x3f\xd0\x1b\xa5\x53\xe5\xd5\x05\x9d\x7c\x30\x3c\x9d\x54\xf0\xdd\x86\x0d\xa7\x78\x3a\x45\x8a\x3a\x1b\xde\xe8\xdb\xb1\x17\xce\x36\xbc\x17\xb1\x53\xfe\x8c\x8b\xc2\x9b\x07\x7a\xc5\x76\xca\x47\x5d\xa1\x0d\x8d\xe4\xb2\x9d\x0a\x15\x8d\x03\x36\x13\x3e\x8b\xdb\xa9\xa8\xc5\x46\x87\xd2\x9e\x4d\x91\x8f\x9c\xc1\xb5\x43\xed\x58\x9b\xc9\x33\x55\xcb\x01\xa3\x91\x30\x36\x34\x8c\xea\x84\x55\x84\x36\xbc\x46\x38\x80\x3a\xa3\xae\x21\x81\x85\xcc\x0d\x67\x5b\x3b\xe9\x84\x74\x72\x6f\xd2\x4d\x3a\x1b\x19\x98\xea\x24\x08\x9b\x70\xd9\x70\x8a\xa8\x3b\xc1\x47\x5e\x36\xf7\xd4\x47\x77\x1a\x4f\xa1\x6f\x96\x23\x10\x4b\x26\x56\x37\x23\x05\x8a\x4e\x7b\xa4\x85\xd9\x96\x3a\xd0\x4e\x5f\x20\xfd\xc9\x5b\xf4\x3a\xc3\x72\xc4\x0d\x9f\x1a\xed\x1c\x62\x61\x36\x39\x75\x32\x9d\x33\xd7\x44\xa1\xf7\x4b\xa9\xcf\x1d\x20\xa5\x46\x5e\x70\xb3\xe4\xf7\x74\xe7\xba\xde\x38\x88\xe2\xbe\xb3\x73\xe1\x95\x59\x6e\xbe\xa3\xd3\x98\x1d\x9b\x65\xd9\xf0\x26\xe2\xce\xbd\x8a\xa7\xe5\xf3\x4d\xfe\xc4\x55\x88\x99\x21\xe7\x17\x3b\xaf\x5b\x65\x4a\x85\xc8\x8a\xcd\x88\x0c\x07\x18\x92\x20\xcf\x15\x06\x08\x49\x16\x0e\x1b\x1a\x93\x77\x5e\x5f\x14\x7c\x2d\x65\xb3\xe5\x8c\x7d\xe7\xf5\x3b\x88\x61\x37\x2f\x23\xeb\xb8\x1f\xfd\x48\x68\x29\x5e\xcc\xfa\x9e\xe5\x26\xb5\x54\xce\xa3\xc0\x79\xc3\x3b\xe5\x3a\xef\xa0\xef\xe3\x6f\x4b\x76\x7d\x69\x74\x05\xd2\xf1\x0d\x4f\x83\xba\x1e\xbd\xe3\xb0\xd9\x6f\x29\xc1\xd0\xf5\xd0\xdd\x8d\x8c\xc9\x7e\x83\xee\xf8\xcb\x88\x2b\xfe\x16\x91\xb0\xbe\x14\x7b\x7a\x44\xdf\x7a\x65\x63\x88\xaa\x05\x49\xfe\x17\x9e\xe4\xdf\x60\x70\xad\x09\x1d\xf3\xfa\x8e\x49\xbc\x04\x1d\xa2\xb2\xbc\x9b\xe5\x0b\x77\xb4\xdf\x7a\xd8\x83\xfe\x85\x3f\xa5\xf3\xed\x15\x23\x9e\x46\x9e\x41\x1d\x02\x48\xb4\xbb\xed\x64\x4e\x43\x2b\xaf\x34\xeb\x49\x3b\xf0\x72\xe5\x6d\x00\xd0\xfe\xcf\x36\x31\xb0\x2e\x8f\xf7\x3d\x69\x13\xd8\xde\x73\xf3\xf2\xf0\xc9\xb3\x6d\xca\x9f\xa5\xf0\x52\x27\x5d\x6f\x90\x43\xdb\xf2\xa7\x71\xbd\xd4\xa4\x49\x6c\x3b\xa7\x7c\x90\x97\x46\x91\xc6\x9b\x2d\x27\xca\xbd\x58\xd4\x08\xb2\x5d\xf3\xf3\x92\x6f\xbd\x84\xc8\xb3\x84\x6d\x4a\x1b\x76\xbc\x5c\xb0\xd0\x29\x40\x2b\x83\xba\x77\x0f\x5c\xd3\xbd\x26\x8e\x7d\xcb\x49\x31\xaf\x3b\xed\xbe\x97\x2c\x47\x36\xc6\xdb\x69\xbc\xee\x60\x6f\xc5\x81\x0f\xb2\x7b\x0d\xb5\x22\xdb\x51\xbf\xe6\xf5\x45\x22\x78\x4b\x70\xcb\x7b\x70\xbd\x01\x15\x8d\xed\x92\x0b\xc2\xf5\x46\x62\x14\x66\x4f\xbc\x0a\xe9\xdd\x55\x99\xe4\xac\x6d\x53\xbb\x96\x04\x27\xdb\x07\xbe\x70\x5f\x92\xb4\x6e\xcb\x5f\x50\xf6\x7d\x27\x48\x86\xbc\xef\xd0\xf7\x21\x68\x65\x93\x56\xdb\x21\xd3\x25\x2b\x72\xba\xf0\x07\xde\xf7\x25\x3a\xee\x2d\x7f\x96\xc5\xf7\x21\x62\x2e\x63\xbb\xdf\x15\x63\x74\x46\x50\x47\x49\xfe\x1a\x16\xc1\x9f\xbc\x9b\xdc\xa5\xd4\x9d\x04\x65\xd0\x33\x80\x13\x7e\x98\x41\xc1\x77\x3d\x77\x9c\x3d\x0a\xca\xd6\xec\xdb\x78\xb7\x56\x50\x36\xea\x16\xf6\x87\xed\x26\xd4\xd8\x82\x34\xaf\x3a\x7e\x86\x4f\x1e\x73\x21\x88\xad\xf1\x93\xfc\xbb\x39\x27\x20\xc3\xe0\xec\x2c\x1a\xa0\xdf\x71\x2f\x17\x58\x86\xb2\xe3\xb4\xe5\xad\xed\x84\x5c\x49\x3b\xfe\x7a\x4c\x38\x69\x31\x75\xf2\x86\x02\x9e\x03\x8f\x77\xc2\x09\xbf\x0c\x79\x3f\xd2\x5b\x15\xd8\xf3\x78\x9c\x2a\x08\x9a\xbc\x5f\xb4\xcb\xa8\x9d\x04\xdd\xb4\x6c\xa1\x6c\xb1\xa2\x61\x62\xd0\x8d\x55\x30\xb4\xda\x65\x8b\x91\x4d\xd9\xa6\x37\xca\xd7\xa8\xb7\x73\xb7\x1e\x99\x1b\x08\x3a\x80\x40\x76\x97\xed\xf8\x17\x9e\xd9\xe0\xcb\x6e\xa4\x65\x2a\x9c\x89\x5f\xdc\xf1\x1f\x25\x08\xad\xf2\xf1\xfb\x1b\xd5\xc8\x8c\x57\x23\xbc\xd0\x0d\x8b\x09\x8d\xff\x03\x0e\x7f\x29\xf7\xa7\xc1\x29\x92\x26\xee\x72\xfe\x38\x4d\x70\x98\xc0\xdb\x71\xa6\x36\x38\x43\x57\xe2\x34\xc8\x0d\x85\x7e\xe0\xe1\x30\x86\xb1\x48\x09\xd7\x34\x95\x08\x9d\xb2\x67\x66\xfa\x9b\x09\xa7\x5b\x42\x27\x55\x64\xad\x6e\x87\xdd\x86\x47\x7d\xa1\x3b\x09\xca\xaa\x76\x23\x2c\x52\xe8\x4e\x5e\x12\x38\xdb\x74\xe0\x4d\x29\xa1\xf3\xa2\x48\x30\xbb\xdb\xf0\xc5\xa2\xf2\xa8\xcf\x72\xc7\xbb\x2c\x6f\x10\xb2\x50\x31\xa6\xb8\x51\xc1\xf7\xf6\x77\xbc\x79\x20\x44\x51\x86\x3d\xfd\x70\xe0\xac\x43\x88\x82\xaa\x80\xbb\x22\xe5\x75\xc0\x1b\xe6\xfb\x43\x92\xe8\x2b\xa7\x5c\xf4\x1f\x1d\xb8\x48\x79\xc7\x56\x03\xd7\xea\xae\x18\x09\x2e\xa2\xf3\x5f\x11\x24\xdf\xd2\xf9\xf7\x10\x1d\x5c\x65\x64\x11\xaf\x22\xb2\xc3\x5d\xb1\xe5\x03\x1c\x21\x7a\x51\xad\x4f\x6a\x45\x1f\x12\x9a\x8c\xe8\x53\x5f\xde\x96\x85\x57\x0a\x8f\xf3\x42\xdf\x28\xff\x57\xb8\x06\xa0\x0b\xae\xf4\x7d\xeb\xce\x58\x81\xf7\x2b\xca\xc2\x84\xde\x56\xce\x46\xaf\x50\x77\xc3\x8e\xb7\x7c\x85\xbe\x63\x4c\xf9\x6e\xbf\xe1\xfd\x4b\xa1\xf7\x12\x0d\x60\x8a\x76\x7c\x82\x3b\xbc\x2a\xdf\x26\x15\x7c\x22\x60\xf7\xc4\x0f\xf0\x55\xb3\x89\xbe\x91\x0e\x87\x70\xfd\x78\x81\x12\x09\xff\x99\x97\xc0\xc2\xd5\x8a\x47\x45\xf7\xc9\x8a\xa6\x13\x37\x4c\x83\x36\x35\xf2\xde\x44\xb8\xb2\x3d\x3d\xd3\x68\x23\x92\xd8\x7a\xf4\x81\x8f\xa8\x10\x2b\x5a\x4c\xf8\xb4\x67\x44\xcf\x3e\x1d\xf8\xa4\x6c\x94\xea\xc4\x1a\xa6\x8b\x7b\xbe\x8c\xa0\xb6\xdd\x0d\xff\x45\x8a\x28\xde\x2b\x3c\xfc\x50\x70\x4d\x8d\xb7\x89\x95\x9f\x21\xfb\x1d\x7f\x9b\x32\xca\x3b\xe2\x43\x0e\xfc\x71\xf0\x78\x92\xa4\x3a\x29\x3a\xc0\xc5\x1b\xae\xe2\x49\x9a\x8a\xd4\x15\x17\x23\x30\x2d\xf8\x5d\xca\x82\xdf\x58\xd1\x11\x0f\x58\xe4\xd4\x01\xde\x38\x80\x4a\x79\xd0\x52\x52\x3c\x70\xa5\x73\x67\xa8\x0d\x0f\x5c\x1b\xbc\xaa\xce\xe0\x98\x8a\x07\x7e\xb2\x83\xb7\x43\xaf\xc3\x6c\xf9\x8b\x04\xd1\x2b\x1b\x8e\xcc\xe5\xf1\xbb\x27\x7a\x79\x67\xa1\x63\xc1\x1b\x5d\xa3\xd7\x4d\x03\x1f\x52\x29\xb6\x19\xf5\x0d\xb8\x45\xa6\xe0\x23\xbd\xd1\xf7\xe8\xad\xee\x62\x64\xd2\xff\x06\xc1\x3f\xdc\x51\x6c\xf7\x1c\xd7\x97\xa8\x4d\x66\xcf\x7f\x56\xa1\x6f\x18\x2f\xbb\xe7\x19\x59\x6f\x6f\x6f\xb8\xab\x92\x3d\xc9\xb5\xe7\xbf\x8f\xd3\x5b\x7d\x04\xce\x78\xbf\xce\xee\xa8\x6b\xe9\xad\x3e\xf3\x9f\x40\xd9\x3f\xd0\xa3\xea\xad\x8e\x52\x27\x43\xf8\x28\x81\xbe\xf4\x3d\xf2\xa3\x09\xc3\x1f\x40\x87\xb0\xe7\x3f\xe6\xf3\xf1\x13\x70\xda\x36\x0e\x9c\xf9\x9e\x3f\x00\xdd\x5b\x57\x46\x85\xdf\xcb\xda\xf3\x01\xf5\xbe\x3b\x6a\x3c\x24\xb0\xbf\xa3\xb6\x79\xc1\xc9\xe3\x81\x37\x38\x5d\x18\xf1\xf1\x98\x52\x37\x73\x11\x8f\xa2\xab\xc3\x23\x3d\xaf\x8b\xb0\xd7\xd1\x1f\x79\x45\x72\x00\x45\x15\x60\x27\xf4\x63\xba\xa5\xc5\xe0\xcb\xad\x7f\x91\xbe\x25\xf0\xc8\x7f\x37\xf2\x22\x2c\x8b\x79\xe4\x3d\x01\x17\xc6\xcd\x3c\xf2\x5f\x2e\xb8\xe8\x12\x7d\x59\x36\xa5\xfc\xf3\x80\xf0\x30\xae\x7a\xe4\x7c\xce\xed\xe1\x08\x84\xd8\xf1\xdf\x40\xbb\xe8\x0b\x84\x3c\x72\x84\xd3\x41\x2a\x34\xde\xf4\x98\x8f\x30\xea\x17\xc7\x66\x74\x1e\x79\x49\xfc\xe2\xde\x04\x35\xd7\x3d\xe6\xd4\x4b\x5f\x3a\xc2\xfc\x3c\x8e\x14\x25\x2f\xc1\xc0\xa7\x3b\x1f\xf9\x13\x68\x97\xe8\xbc\x87\x74\xea\x23\x7f\x29\xff\x55\x35\xe2\x81\xeb\x7a\xe2\x99\xca\xab\x32\x11\x5d\x54\x4f\x5c\x45\x5f\xd5\x05\xc5\x05\x4f\x93\xc7\x94\xba\xbb\x1b\x26\xf9\xef\xdf\xbe\x64\x6f\x82\xf0\x4b\xff\xf5\xa4\xd8\x3d\xf2\x74\x4f\xc5\xf8\x7a\xa2\xb3\x5b\x87\x27\xbe\x43\x4d\x82\xf4\x27\x9e\xab\xbc\x6a\x53\xd3\x5f\xb2\x79\xca\xf8\x80\xec\x2b\x7d\x77\x6e\xec\x03\xe1\x1c\xee\x53\xb6\x5e\xf0\x23\x70\x8e\xb0\x37\x4b\xde\x55\xfc\xea\x3c\x7b\x32\xe4\x89\x37\x20\xbc\xa9\x1e\xf6\x2c\x1c\x46\x7e\x61\xf1\x4d\x2c\x7c\x3a\x8b\x4f\x44\xbf\x35\x88\x7b\x38\x2c\xb8\xf1\xbe\x69\x98\x11\x64\xbc\x3f\xfa\x0d\x34\x05\xf0\xd7\xb5\xde\xa2\x57\xe5\x35\x22\xeb\x38\xf0\xda\xd7\x95\x4c\x78\x3c\xf3\xf2\xf5\xd5\x5d\xd1\x6b\x46\xcf\xf9\x73\x4e\x7b\x86\xc8\xcf\x25\xbd\x70\x6e\xf1\x86\x20\xe1\xd4\xcb\x8c\xea\xcc\x3b\x7d\x61\xef\x85\xb3\xba\xe4\x27\x67\x0e\x3c\x2c\x7f\x1f\x52\x1a\xbc\x23\xaa\x32\xef\x62\xad\x46\xc7\xf3\x92\xf2\x16\xdc\x77\x38\x15\xf3\xc2\xb3\xd5\x77\xe9\xc3\x47\x39\x07\xc0\x78\x8d\xe5\x1d\x4a\xe0\x25\xe7\x1f\xe6\xde\x89\xea\xbc\x8c\x88\x8d\x24\x40\x2f\xdf\x43\xa7\xdf\xfe\xfd\xb7\xff\x0a\x00\x00\xff\xff\x88\x7a\x9f\xb2\x78\x7d\x00\x00")

func currenciesCurrenciesJsonBytes() ([]byte, error) {
	return bindataRead(
		_currenciesCurrenciesJson,
		"currencies/currencies.json",
	)
}

func currenciesCurrenciesJson() (*asset, error) {
	bytes, err := currenciesCurrenciesJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "currencies/currencies.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"markets/market/bitfinex/currencies.json": marketsMarketBitfinexCurrenciesJson,
	"currencies/currencies.json":              currenciesCurrenciesJson,
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
	"currencies": &bintree{nil, map[string]*bintree{
		"currencies.json": &bintree{currenciesCurrenciesJson, map[string]*bintree{}},
	}},
	"markets": &bintree{nil, map[string]*bintree{
		"market": &bintree{nil, map[string]*bintree{
			"bitfinex": &bintree{nil, map[string]*bintree{
				"currencies.json": &bintree{marketsMarketBitfinexCurrenciesJson, map[string]*bintree{}},
			}},
		}},
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
