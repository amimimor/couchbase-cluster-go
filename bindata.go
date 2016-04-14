package cbcluster

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

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindata_file_info struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindata_file_info) Name() string {
	return fi.name
}
func (fi bindata_file_info) Size() int64 {
	return fi.size
}
func (fi bindata_file_info) Mode() os.FileMode {
	return fi.mode
}
func (fi bindata_file_info) ModTime() time.Time {
	return fi.modTime
}
func (fi bindata_file_info) IsDir() bool {
	return false
}
func (fi bindata_file_info) Sys() interface{} {
	return nil
}

var _data_confd_service = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x53\x4d\x6b\xdb\x40\x10\xbd\xeb\x57\x0c\xb4\xb9\x65\xa3\x43\xdb\x4b\xc1\x87\x90\x3a\xd4\x87\x62\x61\x37\xa1\x90\x86\xb0\x91\x46\xf2\x10\x69\xd6\x9d\x1d\xc9\x16\xa5\xff\xbd\xbb\x2b\x43\x48\x14\x68\x2f\x5a\x76\xbe\xf6\xbd\x37\x4f\x77\x37\x4c\x7a\x9f\x7d\x41\x5f\x0a\xed\x95\x1c\x2f\xae\x1c\xd7\xd4\xf4\x62\xe3\x0d\xb6\x28\x03\x95\x98\x65\xef\xc0\xf5\x02\x95\x55\x0b\x83\x6b\xfb\x0e\xa1\xeb\xbd\xc2\x23\x82\xa0\xad\xc6\xec\xb2\x56\x94\x45\x19\x9a\x63\xcd\x85\x3f\xf5\x6d\xf0\x57\x4f\x82\x7e\x9e\xc9\xee\x4e\xb3\xef\xb3\x25\x0f\x24\x8e\x3b\x64\xbd\xa6\x16\x17\x39\x6a\x99\xe3\x73\x30\xbe\xfe\x44\x6d\x0b\x96\x47\xc0\x23\x79\x25\x6e\x20\x4d\xcc\x96\x47\x2c\xb7\x6a\x45\x0b\xc1\x85\xc9\x7b\x2f\xf9\x23\x71\x5e\xb9\xf2\x09\x65\xea\x3a\xe3\x7f\x54\x49\x17\x6b\x5e\x16\xbd\xae\xd9\xf7\x61\x92\xb6\x38\x56\xc8\x9f\xe8\x70\xcc\xfd\xc8\xa5\x69\xac\xe2\xc1\x8e\x86\x1b\xe2\xa3\x99\x10\x05\xb0\x07\x04\x46\xac\x40\x1d\xec\xc5\x0d\x54\x61\x52\x2f\xe5\xe3\x57\x2d\x71\x98\x79\x20\xdd\x81\xee\x10\x56\x05\x90\x42\x69\x39\x8a\x59\xee\x20\xf0\xaf\xa2\xe2\x7c\x9e\xd2\x27\x08\x3e\x1e\x1a\x8e\x58\xec\x91\x2b\xf8\x7a\x53\x80\xa7\x86\x6d\xeb\xe3\x5b\x09\xc5\x79\x50\xa9\x7a\xbd\xac\x67\x6e\x33\x62\xd2\x33\x18\x13\x24\xf8\x99\x01\x18\x84\xab\xf5\x66\xb9\xde\x3e\x14\x9b\xd5\xed\xe5\xf7\xe5\xc3\xaa\xb8\xfd\xb8\x78\xff\xfb\x8d\xe8\x9f\xa9\x63\x80\x7c\xb0\x92\x87\x31\xa7\x89\x17\x11\xe7\xe7\xb7\x82\x53\x83\x99\x20\x79\x53\x8b\xeb\x92\x2f\x4c\x42\x3a\x25\xd9\x06\x6b\x9d\x71\xba\xfd\x9f\xdc\x13\x35\xb7\x9f\x31\xf3\x21\x08\x46\xe1\x43\xdc\xee\x06\x7d\xa2\xef\xd8\xd4\x96\xda\x5e\xa2\x03\x7f\x98\xeb\x16\x31\xfc\x00\x2f\x56\x16\x3c\x1d\x9c\x1f\x85\xf7\x11\x4c\x17\x36\x12\xd6\x05\xd6\xc3\xcc\xc4\xdf\xa6\xdc\xba\x9e\xf9\xfb\x6f\x00\x00\x00\xff\xff\x33\xd6\xa5\x80\x5a\x03\x00\x00")

func data_confd_service_bytes() ([]byte, error) {
	return bindata_read(
		_data_confd_service,
		"data/confd.service",
	)
}

func data_confd_service() (*asset, error) {
	bytes, err := data_confd_service_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "data/confd.service", size: 858, mode: os.FileMode(420), modTime: time.Unix(1427466163, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _data_confdata_service = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x64\x90\xb1\x72\xea\x30\x10\x45\x7b\x7d\xc5\x0e\x14\xaf\xc1\xcf\x5f\xe0\xe2\xcd\x83\xd4\x99\x90\xa4\x61\x28\x14\xf9\x02\x9b\xe0\x15\x59\xad\x8c\xf9\xfb\xc8\x98\x09\x33\x49\xe9\xeb\xbb\xda\x73\x76\xf3\x22\x6c\x5b\xb7\x44\x0a\xca\x27\xe3\x28\xcd\xff\x28\x3b\xde\x67\xf5\xe3\x17\x2d\xbd\x79\x7a\x8d\xc7\xdc\x81\xd6\xd0\x9e\x03\xdc\xbf\x9d\x41\x9b\x36\x86\x0f\xe8\xdf\x74\x0b\x9f\xf0\x99\x59\x91\x7e\xe6\x6e\x73\x1b\xdb\xba\x95\xf4\xac\x51\x3a\x88\x3d\xf0\x11\x4d\x0d\x0b\x35\xee\xa1\x73\xf3\x33\xc8\x2b\xe4\x8f\x91\x27\x89\xda\xf9\x23\xdd\x1e\x5a\x50\xf9\xf7\x9e\x93\x91\x00\x2d\x59\x24\x48\xca\x0a\xb2\x83\x1f\xdb\xed\x08\xda\x5f\x41\xdd\x1c\x03\x27\x4b\x0b\xf2\xd2\x52\x50\x78\x03\x45\x01\xf1\x8e\xd8\xa8\x8d\x48\x65\x83\x7b\xbe\x9c\xd0\x94\x38\x1d\xa2\x15\xfc\xce\xb3\x5c\xcd\x56\x03\x5b\x73\x41\x72\x6e\x35\x20\xac\xcd\xab\x3d\x2a\x9a\xaa\xce\x49\xeb\x37\x96\x7a\x32\x24\xed\x28\x94\x5b\x55\xe3\xe6\x7b\xb5\xf9\x55\xcb\x42\x55\x4f\x57\x59\xd9\xb3\x0c\x54\x55\xe2\xcb\x39\xbf\x87\x69\x8a\x11\x0e\x91\x66\x13\x6e\x5b\x2c\xcf\x93\x53\xa9\x59\x41\x83\xce\xdc\x57\x00\x00\x00\xff\xff\x3a\x3d\x3f\xc3\xae\x01\x00\x00")

func data_confdata_service_bytes() ([]byte, error) {
	return bindata_read(
		_data_confdata_service,
		"data/confdata.service",
	)
}

func data_confdata_service() (*asset, error) {
	bytes, err := data_confdata_service_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "data/confdata.service", size: 430, mode: os.FileMode(420), modTime: time.Unix(1427465921, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _data_couchbase_node_service_template = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x93\x4f\xab\xda\x40\x14\xc5\xf7\xf9\x14\xb3\x28\x08\x85\x79\xd3\x45\xbb\x79\x25\x8b\xd4\xe6\x95\x6c\x8c\x24\xa9\x14\x1e\x12\xc6\xc9\xb5\x0e\x4e\x66\xa6\xf3\x27\x5a\xc4\xef\xde\xd1\x58\x53\x0d\xa5\xd2\x5d\x38\x9c\xfb\xcb\x39\xc3\xbd\xaf\x5f\x25\x77\xcb\xe8\x33\x58\x66\xb8\x76\x5c\xc9\x98\x29\xcf\x36\x2b\x6a\xa1\x96\xaa\x81\x28\x59\x3b\x30\x71\xa3\xd8\x16\xcc\x93\x05\xd3\x71\x06\x51\x01\x3f\x3c\x37\x60\xef\xf5\xde\x0c\x8e\x35\x63\xeb\x8d\x1a\xbd\x96\xfd\xd7\x32\xaa\x78\x0b\xca\xbb\xd2\x51\xe3\x4a\x60\xf1\xbb\x41\x51\xba\x17\x52\xd9\x71\xa3\x64\x0b\xd2\xbd\x70\x01\x31\x09\x2c\x02\x83\x18\xa5\x7b\x60\x67\xc0\xdc\x40\x8c\x89\xb7\x86\xac\xb8\x24\x7d\x3a\xb4\xe5\x42\xa0\x6b\xad\x7f\x98\x4d\xfb\x37\xeb\xbd\x53\xfb\x3f\xb1\xe4\x54\x0d\xcc\xf3\xe1\x80\x9e\xa6\x9f\xea\x45\x5a\x94\x59\x3e\x43\xc7\xe3\x03\x10\x27\xe0\x67\x03\xf2\x03\xdf\xed\xc9\x15\x88\x99\xf0\x36\x3c\x27\xfe\xae\x7a\x68\x3e\xab\x92\x6c\x96\x16\x75\x95\x7c\xb9\xe1\xc6\x67\x60\x98\xd9\x20\xcc\xd0\x64\xd4\xc8\x4b\x84\xb1\xa4\x2d\x0c\x69\x11\xee\x10\x51\xda\x0d\xbf\x23\x1d\x35\xcf\x63\xe9\x34\x09\x2e\xde\x28\xeb\x1e\xe8\x3a\xb9\x84\x52\xfa\xb1\x4c\xbf\xc9\xff\xf5\x00\xc8\xeb\x86\x3a\xc0\x3b\x43\xb5\x0e\xcc\xd1\x20\x32\xd0\xaa\x0e\x30\x95\x0d\x36\xb0\xa2\x82\x4a\x16\x9a\x63\xa1\x18\x15\x98\x6b\xf4\x66\x9a\x17\x69\x5e\xd6\xf3\x22\x5b\x24\x55\x5a\x67\xf3\xc5\xfb\x8f\xc8\xfa\x46\xa1\x4b\x4e\x1b\xaa\x0c\xe0\x49\xd8\xdb\x6f\xf8\x45\x00\x84\x9b\x99\x2a\xb9\x16\x9c\x39\x7b\x77\x31\x6f\xaf\x4b\xfe\x2b\x00\x00\xff\xff\x30\x7b\x9b\x8e\x5d\x03\x00\x00")

func data_couchbase_node_service_template_bytes() ([]byte, error) {
	return bindata_read(
		_data_couchbase_node_service_template,
		"data/couchbase_node@.service.template",
	)
}

func data_couchbase_node_service_template() (*asset, error) {
	bytes, err := data_couchbase_node_service_template_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "data/couchbase_node@.service.template", size: 861, mode: os.FileMode(420), modTime: time.Unix(1432311931, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _data_couchbase_sidekick_service_template = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x52\x4d\x6b\xdb\x40\x14\xbc\xef\xaf\xd8\x43\x21\xa7\x8d\x7a\x68\x2f\x85\x85\x2a\xa9\x52\x74\x88\x64\x24\x39\x14\x8c\x11\xca\xea\xb9\x7a\x68\xb5\xab\xee\x47\x92\x62\xfc\xdf\x2b\x59\xc5\xae\x2d\xd3\x9a\xde\x96\xd9\x99\x79\xf3\x86\xb7\x5a\x2a\x74\x6b\xf2\x05\xac\x30\xd8\x3b\xd4\x8a\x0b\xed\x45\xf3\x5c\x59\x28\x2d\xd6\xd0\xa2\x68\x49\xb8\x71\x60\x78\xad\x45\x0b\xe6\xd6\x82\x79\x41\x01\x24\x83\x1f\x1e\x0d\xd8\x73\x7c\x22\x83\x13\xf5\x9c\x7a\x82\x4e\xc4\x8d\x04\x70\x73\xe6\x29\x7c\x87\xaa\xb6\x85\xfe\x23\x9b\xd2\x35\x7c\xde\x6e\xe9\xed\x32\x89\x8b\x32\x59\x3e\xde\x45\x19\xdd\xed\xce\xcc\xaf\xe7\x93\x55\x3e\xbd\xd6\xa4\xc0\x0e\xb4\x77\xb9\xab\x8c\xcb\x41\xf0\xf7\x24\x52\x2f\x68\xb4\xea\x40\xb9\x07\x94\xc0\x83\x61\x8f\x00\x8e\x20\x89\xde\x40\xec\xf9\x0b\x03\x9c\x05\xde\x9a\xe0\x19\x55\x30\x35\x43\x5b\x94\x92\x1e\xa2\xb0\x43\xad\x7f\x57\x99\xee\x9f\x9a\x73\x49\xef\x87\x41\x4e\xc2\xcf\x1a\xd4\x47\x7c\x7d\x0b\x8e\x06\x42\x7a\x3b\x34\xc2\xbe\xeb\x4f\x63\x0b\xf7\x69\x52\x84\x71\x12\x65\x65\x11\x7e\x1d\x7a\x38\xfa\xf2\xbd\xe1\xa0\x69\x28\x13\xf4\x66\x96\xca\x2b\xca\x98\xaa\x3a\xb8\x90\x6e\xfc\x01\xc7\x1b\x6d\xdd\xff\xc5\xa0\xbe\xaf\x2b\x07\xec\xd5\x54\x7d\x3f\x4c\x9b\x09\xa9\x1d\x33\xb2\x8b\xa3\xa5\x16\x95\x64\xd8\xf3\x77\xf7\x69\x16\xa5\x79\xb9\xc8\xe2\xa7\xb0\x88\xca\x78\xf1\xf4\xe1\xe6\xf7\x86\xba\x9f\xb5\x66\x07\xf0\x52\xd5\x64\xf5\x8d\x3d\x8c\x67\xb8\x26\x8f\x95\x68\x50\x41\xba\xb9\xfe\xa2\x7e\x05\x00\x00\xff\xff\x28\xcb\x1c\xaa\x5a\x03\x00\x00")

func data_couchbase_sidekick_service_template_bytes() ([]byte, error) {
	return bindata_read(
		_data_couchbase_sidekick_service_template,
		"data/couchbase_sidekick@.service.template",
	)
}

func data_couchbase_sidekick_service_template() (*asset, error) {
	bytes, err := data_couchbase_sidekick_service_template_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "data/couchbase_sidekick@.service.template", size: 858, mode: os.FileMode(420), modTime: time.Unix(1426775948, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _data_nginx_service = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\x91\x4f\x6b\xeb\x30\x10\xc4\xef\xfa\x14\x0b\xe1\xf1\x2e\x4f\xcf\x81\x5e\x42\xc1\x87\x42\x93\x5b\xff\x90\x50\x28\x84\x1c\x14\x79\x1d\x2f\xb1\x57\xae\xb4\x8a\xf3\xf1\x2b\xc5\x09\xa1\xc9\xa1\x47\x0d\x3f\x66\x47\x33\xeb\x0f\x26\xd9\xa8\x67\x0c\xd6\x53\x2f\xe4\xb8\x7c\xdd\x11\x1f\x61\x85\xfe\x40\x16\xd5\x53\x2d\xe8\x4b\xeb\xb8\xae\xfe\x87\xb3\xa6\x26\x03\xc2\xe0\xf8\xaf\xc0\x60\x58\x80\x04\xc4\x81\xc7\xaf\x48\x1e\x41\x1a\x84\x33\x09\x3a\xbd\x4c\xa2\x5c\x6c\x2b\x08\xe2\x7a\x88\x21\x81\x41\x8c\x17\xe2\x9d\x9a\x90\xfc\x83\xa1\x21\xdb\x00\x05\x08\xa6\x46\x35\x59\x8e\x3e\xe1\xf6\xe8\xfa\x1c\x69\xa3\xe6\x7c\x20\xef\xb8\x43\x96\x05\xb5\x58\x16\x28\xb6\xc0\xab\xa8\xe6\x47\xb4\xab\x7c\xe3\xdd\x63\xa9\x8b\x18\x7c\xb1\x25\x2e\x2a\x67\xf7\xe8\x61\x4f\x6d\x0b\x7f\xf8\x17\xca\x77\x77\xcc\x2d\xd2\xc7\x64\xc4\xb9\xae\x2b\x77\x07\xf9\xc8\xa0\x35\x9b\x0e\x93\x1f\xe8\x1e\x66\xd3\xc7\xd9\x34\x49\x07\xd7\xc6\x0e\x83\xae\xbd\xeb\x4e\x7f\xd5\x95\x11\xf3\xc3\xcf\xf5\x77\x76\xa7\x12\xb5\xc0\x43\x4e\xb7\x1c\x9b\x2c\x1d\xeb\xda\x50\x1b\x7d\xae\xe9\x53\x2f\x5a\xc4\x34\x6a\x5e\x89\x11\xab\x3c\xce\x16\xc1\xf1\x38\x4d\x4e\xd2\x19\xdb\x10\x23\x98\x00\xa7\x96\xd3\x61\xf5\x32\x6a\x6f\x75\x79\x91\x2e\xdd\x7f\x07\x00\x00\xff\xff\xfa\x80\xca\x3a\x25\x02\x00\x00")

func data_nginx_service_bytes() ([]byte, error) {
	return bindata_read(
		_data_nginx_service,
		"data/nginx.service",
	)
}

func data_nginx_service() (*asset, error) {
	bytes, err := data_nginx_service_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "data/nginx.service", size: 549, mode: os.FileMode(420), modTime: time.Unix(1427465946, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _data_sync_gw_node_service_template = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x53\x51\x6b\xdb\x30\x10\x7e\xd7\xaf\xd0\x5b\x61\x20\x6b\x2f\x7b\x29\xf8\xc1\x74\xe9\x30\x0c\x6f\x24\x0e\x1b\x84\x60\x5c\xf9\x62\x6b\xb5\x25\x4f\x3a\xc5\x0d\xa5\xff\x7d\x72\xdd\x26\x73\x9c\x2d\xd9\xe8\x9b\xf8\xf8\xbe\xbb\xef\x3b\xdd\xad\x96\x4a\xe2\x9a\x7c\x04\x2b\x8c\x6c\x51\x6a\x15\xda\x9d\x12\x59\xd9\x65\x4a\x17\x40\xa2\x0d\x82\x09\x0b\x2d\xee\xc1\x04\x16\xcc\x56\x0a\x20\x73\xf8\xe9\xa4\x01\x7b\x8c\x0f\x64\x40\x51\x4c\xa9\x23\x74\x20\x6e\x6a\x00\x9c\x32\xc7\x30\x59\x2d\x86\xd7\x9a\xa4\xb2\x01\xed\x70\x81\xb9\xc1\x05\x88\xf0\x3d\x99\xa9\xad\x34\x5a\x35\xa0\xf0\x56\xd6\x10\x72\xdf\x85\xc3\x01\x24\xb3\x07\x10\xcf\xfc\xaf\x06\x42\xc6\x9d\x35\xfc\x4e\x2a\x3e\xf8\xa6\xf7\xb2\xae\xe9\x4b\xdc\x33\x54\xd3\x9c\x26\x1e\xf3\x5a\xe7\x4b\x0a\xed\x44\x75\x97\x5b\xe0\xbd\x86\x95\x39\x42\x97\xef\x2e\x10\x62\x0d\xbb\x02\xd4\x07\xd9\x3d\xf0\x7d\x11\x26\x6a\x67\xfd\xbc\x58\xa9\xaf\x1f\x1f\x69\x70\xf3\x25\x49\xa3\x38\x99\xcd\xb3\x34\xfa\x44\x9f\x9e\xfe\x5e\xd7\x38\x45\x19\x53\x80\x61\xa5\x2d\xfe\x5f\x07\xea\xda\xc2\x67\x60\x9d\xc9\xdb\xd6\xd7\x9c\x08\x69\xaf\xfa\x16\xc5\x69\xb6\x4c\xd2\xf8\x73\x36\x5f\x26\x49\x9c\xfc\xb3\x39\xb6\xa5\xbc\xd2\x0d\x78\x67\x06\xae\x0f\xcf\xb7\x71\x3d\xfc\x45\xc7\x84\x56\x1b\x59\x52\x03\x9d\x91\x08\xbe\x7f\x01\x16\xa5\xca\xfb\xe5\xff\xad\x3f\x0f\xc6\x82\xe0\x87\xd5\xea\x10\x27\x7c\xce\xe1\x7d\x54\x94\x09\x7a\xf5\x87\x64\x79\x03\xaf\x7b\x73\x51\xd0\xd3\x9b\x73\xce\xd5\xd5\x8b\x2d\xdd\x4e\x26\x6c\x3d\xb8\xdf\x5c\xb2\xfa\xce\x6e\xfb\xeb\x5a\x93\x1b\xaf\xae\xa5\x40\x3b\x3a\xf7\x77\xaf\x67\xf7\x2b\x00\x00\xff\xff\xfe\xa2\xd3\xef\x17\x04\x00\x00")

func data_sync_gw_node_service_template_bytes() ([]byte, error) {
	return bindata_read(
		_data_sync_gw_node_service_template,
		"data/sync_gw_node@.service.template",
	)
}

func data_sync_gw_node_service_template() (*asset, error) {
	bytes, err := data_sync_gw_node_service_template_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "data/sync_gw_node@.service.template", size: 1047, mode: os.FileMode(420), modTime: time.Unix(1432768089, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _data_sync_gw_sidekick_service_template = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x52\xcb\x6a\xdc\x30\x14\xdd\xeb\x2b\xb4\x28\x64\xa5\xb8\x8b\x76\x53\x10\xd4\x49\x9d\xe2\x45\xec\xc1\xf6\x84\xc2\x10\x8c\x22\xdf\x8c\x2f\x96\x25\x57\x8f\x38\x21\xe4\xdf\xab\xc9\xa4\x49\x3b\x86\x76\xc8\x4e\x1c\xce\x4b\x87\xbb\x59\x6b\xf4\xd7\xe4\x1b\x38\x69\x71\xf2\x68\x34\x77\x0f\x5a\xb6\xdb\xb9\x75\xd8\xc1\x80\x72\x20\xe9\xad\x07\xcb\x3b\x23\x07\xb0\xa7\x0e\xec\x1d\x4a\x20\x15\xfc\x0c\x68\xc1\x1d\xe2\x7b\x32\x78\xd9\x2d\xa9\x7f\xa1\x67\xa8\x3b\xd7\x98\xd7\x38\x6d\x3a\xf8\xfa\xf8\x48\x4f\xd7\x45\xde\xb4\xc5\xfa\xf2\x2c\xab\xe8\xd3\xd3\x81\xf1\xb1\x6c\xb2\xa9\xf7\xaf\x6b\xd2\xe0\x08\x26\xf8\xda\x0b\xeb\x6b\x90\xfc\x23\xc9\xf4\x1d\x5a\xa3\x47\xd0\xfe\x02\x15\xf0\x24\x16\x4b\xe0\x0d\x24\xd9\x3d\xc8\x67\xfe\xca\x02\x67\x49\x70\x36\xb9\x41\x9d\xec\xbf\x4a\x07\x54\x8a\xee\x8a\xb0\xed\xcc\x5e\x57\xfa\xb7\xc6\x8e\xff\x51\x1c\x0a\xa6\x10\x43\xbc\x82\x87\x0e\xf4\x67\x9c\xef\x13\x69\x82\xec\x6f\x84\x03\x26\x55\x70\x71\x0b\xb6\x35\x5f\x76\x0b\x9c\x97\x45\x93\xe6\x45\x56\xb5\x4d\xfa\x3d\x6e\xf0\xe6\xcb\x9f\x0d\xa3\xa6\xa7\x4c\xd2\x93\x45\xa7\xa0\x29\x63\x5a\x8c\xb0\xe8\xb6\xc3\xc1\xf3\xde\x38\xff\xbe\x12\x34\x4c\x9d\xf0\xc0\x66\x2b\xa6\x29\x66\xfd\x0e\x78\x91\x51\x25\x82\x96\xfd\x9f\x79\xca\x48\xa1\x18\x4e\xfc\xc3\x79\x59\x65\x65\xdd\xae\xaa\xfc\x2a\x6d\xb2\x36\x5f\x5d\x7d\x3a\x79\xf9\x94\x99\x16\x43\xb9\x08\x2e\xb7\x25\x9b\x1f\xec\x42\x01\xc4\xdb\xbe\x14\xb2\x47\x0d\xe5\xed\xd1\xc7\xf3\x2b\x00\x00\xff\xff\x90\x04\x86\x17\x15\x03\x00\x00")

func data_sync_gw_sidekick_service_template_bytes() ([]byte, error) {
	return bindata_read(
		_data_sync_gw_sidekick_service_template,
		"data/sync_gw_sidekick@.service.template",
	)
}

func data_sync_gw_sidekick_service_template() (*asset, error) {
	bytes, err := data_sync_gw_sidekick_service_template_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "data/sync_gw_sidekick@.service.template", size: 789, mode: os.FileMode(420), modTime: time.Unix(1426775948, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _data_test_fleet_api_units_json = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x9d\xdb\x73\xda\x30\x16\xc6\xdf\xf3\x57\x68\x98\x74\xd2\xee\xac\x63\x6e\x81\x40\x86\x99\xa6\x69\xba\xcd\x43\x37\x99\x90\x76\x77\x67\xd9\xc9\x18\x5b\x06\x6d\x8c\xcd\xfa\x02\xcd\x64\xf2\xbf\xaf\x64\x02\x81\x70\x4b\x7d\x69\x25\xf2\xe5\x21\x0d\x20\xec\xef\x1c\x89\xa3\xef\x57\x4b\xe6\x61\x8f\xf0\x9f\x42\xe4\xb2\x30\x28\x34\xff\x1d\x3f\x12\x3f\x0f\xb3\xbf\xe2\xd7\xcd\xc8\xf7\xa9\x1b\xb6\x43\x23\xa4\x85\x66\x81\xb9\x86\x19\xb2\x11\x2d\xfc\x75\xb1\x99\x45\x03\xe6\x53\x6b\x5b\x33\xd7\x18\x88\x97\xcd\xae\x1d\xdc\x1a\xae\xeb\x45\xae\x49\x3f\x1e\x06\xd4\x1f\x31\x73\xa9\xb1\x37\x0c\x99\xe7\x2e\x88\x5b\x2d\xf2\xe5\xe1\x3f\xd3\xc0\xf4\x59\xfc\xee\x17\xc7\x9c\xb5\x0c\xa8\x19\xbf\xdc\x2c\x7c\xe7\x09\x58\xd7\x6a\x64\x38\x91\x38\xe0\xe9\x93\x54\x72\xf6\xe9\x4b\xbb\xb0\xd4\xf6\x71\xf9\xed\x9b\x05\x7e\x62\xae\x15\xdc\x78\xd9\x88\x8b\x93\xe9\x7a\x16\xfd\xf8\x8e\xcd\x52\x99\x5e\xe3\xa9\x1d\x52\x5f\x6a\x85\xe7\x3f\xa9\xc9\x07\x9c\xbf\xf6\xfc\xcf\x2a\xdb\x2b\x47\xd8\x92\x50\xbd\xcb\x5c\x3d\xe8\x13\xcd\x24\x9d\xc2\xb8\xcf\x1c\x4a\x42\x3f\xa2\x27\xc4\xf2\x08\x0d\x4d\xcb\x0c\x1d\x12\xd0\x90\xe8\xa6\x17\x99\xfd\xae\x11\xd0\x43\xd3\x1b\xe8\x22\x3e\x7d\x3e\x48\x72\xf0\x40\x3a\x9d\x4e\xa1\xef\x05\xa1\xf8\xb7\x19\x3f\x7a\xf7\x55\xfc\x7e\x3c\x20\x9a\x16\xf2\x03\xd5\x8a\x27\x81\x43\xe9\x90\x54\x8f\x4e\x2c\xcf\xa5\x9d\x42\x66\x39\xf1\x86\xd9\xa5\x24\x0a\xfc\x38\x2d\xd3\xf8\xfd\xc1\xd6\xf0\x33\x88\xe3\x9b\x61\xf6\x99\x4b\x2f\xed\xed\x81\xfc\x53\xfb\xc2\xd3\x98\xd5\x20\x5c\x78\xe6\x3f\x7b\x2b\x02\xd8\x5c\x20\x1d\x83\x17\x8a\x3e\xb5\xb6\x14\xc8\x75\xcd\x06\x93\xb8\x2f\x3e\xf3\x36\xc5\x52\x95\xd6\x4a\x15\xc3\x3e\xae\x1f\x55\xeb\xd5\x8a\x51\x2b\xd6\x1b\x8d\x4a\xdd\x6a\xd4\xed\xb2\x59\xa9\xbd\xaa\xb8\x96\x50\x5d\x53\x89\x43\x75\x45\x75\x45\x75\xdd\xc1\xea\x5a\xad\x96\x4a\x8d\x7a\xa3\x5a\xaa\xd4\xba\xd5\x7a\xc3\x6e\x1c\x19\x5d\xbb\x4c\xcb\xb4\x58\xb2\x8b\xbc\xf0\xbe\xaa\xba\x96\x51\x5d\x53\x89\x43\x75\x45\x75\x45\x75\xdd\xc1\xea\x6a\xd7\xa8\x69\x1f\x57\x8f\xbb\x75\xc3\xa8\x5a\xb5\xda\x71\xb9\x5a\x2f\x1b\xc7\xb4\xd4\x30\x4a\x75\xb3\x56\x7a\x55\x75\xad\xa0\xba\xa6\x12\x87\xea\x8a\xea\x8a\xea\x2a\x4b\x75\xcd\xf0\xbf\x4e\x63\xb1\xf2\x17\xc7\x99\xd6\x0c\xba\xfd\x86\x0d\xa8\x17\x85\xf1\xa7\xba\x4d\xcd\xcc\x46\x71\x31\x8b\x8f\x96\x3b\x62\xbe\xe7\x0e\x78\x67\x7f\xe1\x95\x21\xbb\x4f\x18\xff\x64\xe9\xf4\xf9\xe0\x59\x56\xc6\x2b\x3f\x3b\x9d\xda\xac\x14\x58\x9e\x79\x47\x7d\x72\xc7\x1c\x87\x88\xde\x57\x46\x31\xaf\x5d\x4a\xe9\x1d\x46\x3c\xc3\xa1\x43\xef\x2d\xea\x1e\xb1\xf1\x4f\x5d\x66\xf5\xb1\x70\x3e\x29\xc4\xd3\xe7\x01\x35\xfb\x1e\x89\xcf\xc1\xdc\x5e\x9c\x75\x22\x6a\x04\xb3\x9a\x64\xff\xec\xf2\xfa\xfc\xb2\x7d\x7b\x75\x7d\xf1\xe3\xf4\xe6\xfc\xf6\xe2\xea\x47\xf5\x40\x5e\x3f\x30\x0b\xe9\xec\xf2\xfb\xd9\xd7\x4f\xa7\xed\xf3\xdb\xf6\xf9\xf5\x8f\xf3\x6b\x2e\xbc\xb5\xff\x5e\xe7\x95\x78\x7e\x82\xd4\x7a\x34\xd4\x6c\xe6\x07\xcb\x36\x61\xfa\x48\x13\x99\xd0\x02\x31\x15\x7c\x20\x9d\xa8\x58\x2c\xd7\x26\xbf\xc9\x74\x9c\x46\x2e\x37\x09\x22\xac\x49\xe6\x34\x8d\xd7\x07\x7e\x50\xee\x47\x5e\x56\x0b\xa2\x8d\x88\x3e\x32\x7c\xdd\x61\xdd\xc9\x4c\x6c\x19\xa1\xd1\x5c\x7e\x4a\x1c\x90\x86\x2d\x61\x49\x96\x86\x14\x99\xc6\xd8\x29\xf8\xd4\xf6\x69\xd0\xd7\xe2\xa7\x6d\x6a\x84\x91\x4f\xf5\xa7\xf6\x27\x4f\x6a\x84\xfc\x8b\xcf\xad\xfd\x87\x15\x1d\xf9\x48\xb4\x6e\xc4\x83\x08\x5b\x93\xb6\xb3\xa0\x5b\xfd\x30\x1c\x36\x75\x5d\xbc\x6b\x29\x8f\x8f\xcd\xe3\x62\xa3\xa4\x13\xcd\xf7\xbc\xb0\xb5\x4a\xfc\x88\xd1\xf1\x95\xef\xfd\xbc\xef\x14\xb2\x1b\x2b\x79\xf8\xa4\xa7\x2e\x0c\xf8\xc1\xb3\x2a\x36\x67\x9e\x6b\x3b\xcc\x0c\x83\xec\xcd\xd0\x5f\xd4\x00\xcd\x34\x17\x49\x62\x1b\xa5\xc0\x05\x12\xf8\xa8\x74\xda\xe0\xa3\xe0\xa3\xe0\xa3\xe0\xa3\xe0\xa3\xe0\xa3\x56\x8b\x85\x8f\x4a\x73\x39\x34\xf6\x51\x0a\x5c\x0a\x85\x8f\x4a\xa7\x0d\x3e\x0a\x3e\x0a\x3e\x0a\x3e\x0a\x3e\x0a\x3e\x6a\xb5\x58\xf8\xa8\x34\x0b\x1f\x62\x1f\xa5\xc0\xa2\x07\xf8\xa8\x74\xda\xe0\xa3\xe0\xa3\xe0\xa3\xe0\xa3\xe0\xa3\xe0\xa3\x56\x8b\x85\x8f\x4a\x7a\x5d\x6f\x3a\xf8\xd4\xb9\xb8\xb7\xa0\x58\xae\xe5\x99\x93\x91\x9d\xe1\xc2\xcc\x6b\xfa\xbf\x88\x8f\x84\x57\x0c\xeb\x3f\x22\x2f\xc3\xd4\x89\xb9\x41\xda\xc4\x65\x2c\x4e\x66\x0f\x3f\xd3\xe6\x0d\x65\x93\x06\xbc\x58\x89\x17\xd3\x7a\xa8\x8c\x6c\xc1\x18\xb2\x8b\xde\xce\x19\x33\xab\x2a\xea\x02\xf5\xb5\xca\x61\xf1\xb0\xd4\x74\xf8\x8c\x1d\x48\x3b\x82\x7e\x21\x2a\xd3\x89\x02\x5e\xdc\xb5\x9e\x97\x43\x4c\x39\xe1\xc7\xd2\x40\x9b\x83\x84\x69\x60\x31\x0d\x08\x20\x99\x3d\x23\xbc\x74\x73\xf9\xa9\xb5\x34\xb0\xa9\xe3\xe5\x34\xde\xaf\x49\xd2\xe6\x50\x97\x46\x03\x89\x86\x9c\x3a\xa8\x36\xf6\x8d\xe1\x90\x1f\x65\xa9\x29\xf1\xe9\xc0\x1b\x51\xcd\x70\x2d\xcd\xa7\x5d\xc3\x31\xc4\xae\x1f\x4d\x73\x3c\xd3\x70\x34\x36\x5c\xc9\xb8\x27\x24\x88\x2c\x8f\x2c\x00\xc2\xf4\xc0\x59\xa4\x36\x07\x4c\x58\x70\xa3\x8a\xb0\x42\xd2\x6b\xd7\x8b\xac\xa0\xc2\x05\x6c\xb0\x82\x34\xf2\xc0\x0a\x89\xc4\x81\x15\x12\x49\x03\x2b\x80\x15\x7e\x97\x68\xb0\x02\x58\x01\xac\x00\x56\xd8\x49\x56\x48\xba\x3e\x63\x91\x15\x54\x58\xa4\x01\x56\x90\x46\x1e\x58\x21\x91\x38\xb0\x42\x22\x69\x60\x05\xb0\xc2\xef\x12\x0d\x56\x00\x2b\x80\x15\xc0\x0a\x3b\xc9\x0a\xa9\xd7\x20\x05\xcc\xa2\x77\xcc\xbc\x53\x6b\x1d\xd2\x54\xb5\x5c\xc6\x17\xcc\x90\x58\xdb\x5b\x62\x86\x0c\xd3\x66\x8b\xc2\x26\x6d\xde\xb2\x56\x97\xed\x9d\x27\xd7\x2d\xc3\x94\xaa\x83\x73\x54\x29\x33\xba\x82\x0f\x37\xf2\xa1\x96\xe1\xfc\xf7\xfb\x41\x51\x7a\xf5\x60\xab\xe7\x3e\xca\x05\x28\x02\x11\x9d\xb6\xf2\x64\x53\xa8\x68\xe5\xbb\x21\x27\xf7\x4d\x16\x79\x0c\xf6\x1c\xee\x2b\xfb\xfa\xe9\x45\x22\xe8\x49\xbd\x98\x6a\x06\x3d\x4a\x2d\xa8\x02\xf4\xc8\x20\x0f\xd0\x93\x48\x1c\xa0\x27\xa1\xba\x3c\xa1\xa7\x2c\x67\x07\xe7\xa8\x12\xd0\x03\xe8\x01\xf4\xac\x52\x0f\xe8\x01\xf4\xbc\x49\xe8\xd9\x34\xbd\x48\x04\x3d\xa9\x57\x85\xcd\xa0\x47\xa9\x95\x61\x80\x1e\x19\xe4\x01\x7a\x12\x89\x03\xf4\x24\x54\x97\x27\xf4\x54\xe4\xec\xe0\x1c\x55\x02\x7a\x00\x3d\x80\x9e\x55\xea\x01\x3d\x80\x9e\x37\x09\x3d\x9b\xa6\x17\xf5\xa1\x27\xb8\x77\xcd\xdb\xde\x58\x99\x1b\x6c\xcd\xeb\x95\xcb\x95\x00\x74\x12\x6b\x03\xe8\x24\x52\xf6\xb6\x40\x07\xd6\x5c\x35\x6b\xfe\x54\xab\x95\x11\xcd\xfd\xb8\xe4\x92\xb7\x9a\x70\xa1\x5f\xeb\x71\xd7\x30\x36\xee\xb9\x97\xf4\xa9\x17\x60\x77\xce\x1f\x8f\x69\xc3\x6e\x93\xf5\x1d\xb6\x9d\x18\xc6\x06\x0b\xb5\xc8\x0d\x99\xa3\xf1\x33\xb8\xcc\xed\x29\x9a\x10\xb1\x3b\xa9\xef\x0d\xa8\x2e\x12\xd0\x7c\xfe\x33\x29\x5c\x4d\x92\x3a\xe6\xf9\x74\x6d\xd6\x23\x3e\x1d\xfb\x2c\x14\xfb\x70\xb8\xaf\x0e\x99\x6b\x08\xf1\x73\x67\xd4\x0f\x17\xdf\x70\xf8\xdf\x80\x07\x27\x3f\xb3\xb6\xff\xf5\xf7\xb3\xdb\xbf\xfd\xe3\xf6\xec\xf2\xdb\xb7\x8b\x9b\xd6\xfe\xfb\xe9\x77\xb7\xf7\x96\xbf\xbb\x7e\x7e\x9c\xf1\x97\x06\x03\x16\x7e\x38\x21\x64\x03\xf6\x3e\x95\xc2\x5f\xee\xa7\x15\x23\x7a\xd6\x21\x31\xe3\x0a\xed\xfb\x8b\xda\x89\xd6\x23\xdb\x3a\x44\x1d\xce\xcd\x6e\x12\xc9\x7e\xeb\xd6\x3c\x3f\xed\xf6\xc6\xad\x05\xb2\x55\x60\xf5\x22\xc8\x56\x16\x79\x20\xdb\x44\xe2\x40\xb6\x09\xd5\x81\x6c\x41\xb6\x20\x5b\x90\x2d\xc8\x16\x64\x0b\xb2\x05\xd9\x26\xef\x11\x90\xad\x52\x64\x9b\x70\x77\xde\x02\xd9\x2a\xb0\x44\x15\x64\x2b\x8b\x3c\x90\x6d\x22\x71\x20\xdb\x84\xea\x40\xb6\x20\x5b\x90\x2d\xc8\x16\x64\x0b\xb2\x05\xd9\x82\x6c\x93\xf7\x08\xc8\x56\x29\xb2\x4d\xb9\x1a\x59\xa5\x5b\x6d\xbe\xd4\x2c\x17\x6f\x80\x70\x13\x6b\x7b\x4b\x84\x9b\xe9\x2e\xc2\x35\x7b\x0a\xa4\xea\xdc\xdc\x34\x02\x79\x55\x44\xde\xd8\xe8\x48\xbe\x03\x6f\x1d\xfb\xaa\xa0\x5d\x16\x62\xcc\xc9\xd5\x6f\xb1\xe4\xf3\x3d\x94\xc1\xde\xc1\x99\xd1\x7e\x62\xca\x89\x21\xdb\xc1\x0d\x83\x39\x8c\xee\xec\xb7\x0b\xbe\x7a\x26\x91\xc8\x9e\xa7\x5c\x52\xa9\xd2\x4d\x21\x61\xcf\x65\x92\x07\x7b\x9e\x48\x5c\x7e\xf6\x5c\xd2\xfb\x1a\xe6\xa6\x11\xf6\x1c\xf6\x1c\xf6\xfc\xa5\x76\xd8\x73\xd8\xf3\xcd\xf9\xdc\x0d\x7b\xae\xc8\x0d\x0c\x53\xae\x0b\x53\xe9\xf6\x85\xb0\xe7\x32\xc9\x83\x3d\x4f\x24\x2e\x3f\x7b\x2e\xe9\x1d\xf8\x72\xd3\x08\x7b\x0e\x7b\x0e\x7b\xfe\x52\x3b\xec\x39\xec\xf9\xe6\x7c\xee\x86\x3d\x4f\x72\xab\xbd\xbd\xc9\xe3\xc7\xbd\xbd\xff\x07\x00\x00\xff\xff\xa8\x19\xae\x06\x04\xd5\x00\x00")

func data_test_fleet_api_units_json_bytes() ([]byte, error) {
	return bindata_read(
		_data_test_fleet_api_units_json,
		"data-test/fleet_api_units.json",
	)
}

func data_test_fleet_api_units_json() (*asset, error) {
	bytes, err := data_test_fleet_api_units_json_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "data-test/fleet_api_units.json", size: 54532, mode: os.FileMode(420), modTime: time.Unix(1432312052, 0)}
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
	"data/confd.service":                        data_confd_service,
	"data/confdata.service":                     data_confdata_service,
	"data/couchbase_node@.service.template":     data_couchbase_node_service_template,
	"data/couchbase_sidekick@.service.template": data_couchbase_sidekick_service_template,
	"data/nginx.service":                        data_nginx_service,
	"data/sync_gw_node@.service.template":       data_sync_gw_node_service_template,
	"data/sync_gw_sidekick@.service.template":   data_sync_gw_sidekick_service_template,
	"data-test/fleet_api_units.json":            data_test_fleet_api_units_json,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func     func() (*asset, error)
	Children map[string]*_bintree_t
}

var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"data": &_bintree_t{nil, map[string]*_bintree_t{
		"confd.service":                        &_bintree_t{data_confd_service, map[string]*_bintree_t{}},
		"confdata.service":                     &_bintree_t{data_confdata_service, map[string]*_bintree_t{}},
		"couchbase_node@.service.template":     &_bintree_t{data_couchbase_node_service_template, map[string]*_bintree_t{}},
		"couchbase_sidekick@.service.template": &_bintree_t{data_couchbase_sidekick_service_template, map[string]*_bintree_t{}},
		"nginx.service":                        &_bintree_t{data_nginx_service, map[string]*_bintree_t{}},
		"sync_gw_node@.service.template":       &_bintree_t{data_sync_gw_node_service_template, map[string]*_bintree_t{}},
		"sync_gw_sidekick@.service.template":   &_bintree_t{data_sync_gw_sidekick_service_template, map[string]*_bintree_t{}},
	}},
	"data-test": &_bintree_t{nil, map[string]*_bintree_t{
		"fleet_api_units.json": &_bintree_t{data_test_fleet_api_units_json, map[string]*_bintree_t{}},
	}},
}}

// Restore an asset under the given directory
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

// Restore assets under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	if err != nil { // File
		return RestoreAsset(dir, name)
	} else { // Dir
		for _, child := range children {
			err = RestoreAssets(dir, path.Join(name, child))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
