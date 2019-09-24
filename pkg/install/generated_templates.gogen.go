// Code generated by vfsgen; DO NOT EDIT.

package install

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// templates statically implements the virtual filesystem provided to vfsgen.
var templates = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		"/flux-helm-operator-account.yaml.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "flux-helm-operator-account.yaml.tmpl",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 948,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x92\x39\x6f\xdc\x30\x10\x85\x7b\xfe\x8a\x01\x5c\x38\x09\x2c\x05\xee\x02\x75\xb6\x8b\x14\x09\x52\x28\x47\x13\xa4\x18\x92\x4f\x59\xc6\x5c\x8e\x30\x24\x37\x87\xb0\xff\x3d\x90\xb4\x06\xbc\x8e\xed\x34\xdb\x8d\xe6\xd2\x9b\xc7\xaf\x69\x1a\x73\x46\x9f\x36\xa0\x0c\xdd\x05\x07\x62\xe7\xa4\xa6\x72\x41\x2e\xd6\x5c\xa0\xa4\x12\x91\x2f\x88\x93\x3f\x4a\x91\x0d\xc9\x87\xf4\x9d\x58\x61\xce\x48\x52\xfc\x4d\x09\xf0\xf0\x34\x88\xd2\xbb\x6a\xa1\x09\x05\x99\x7e\x86\xb2\x59\x46\x1a\xcb\x19\x7e\xfe\x03\x72\x26\x27\xa9\xa8\x44\x7a\xd1\x5f\x5f\xdd\xbc\x6c\x0d\x8f\xe1\x0b\x34\x07\x49\x1d\xed\x2e\xcd\x6d\x48\xbe\xa3\x8f\xab\xaa\xab\x55\x94\xd9\xa2\xb0\xe7\xc2\x9d\x21\x8a\x6c\x11\xf3\x1c\x11\x25\xde\xa2\xa3\x21\xd6\x5f\xcd\x06\x71\xdb\xc8\x08\xe5\x22\x6a\x9e\x2e\x4d\x13\x85\x81\xda\x0f\xbc\x45\x1e\xd9\x81\xf6\xfb\x43\xf7\xf2\xd9\xd1\x34\x1d\x57\xa7\x89\x90\xfc\x7e\x6f\x66\xcf\xee\x8b\x55\xcb\xae\xe5\x5a\x36\xa2\xe1\x0f\x97\x20\xa9\xbd\x7d\x93\xdb\x20\xaf\x77\x97\x16\x85\xef\x6e\xb9\x59\xdd\xeb\x25\xe2\x94\x87\x18\xad\x11\xcb\x78\x43\x3c\x86\xb7\x2a\x75\xcc\x1d\x7d\x3d\x7f\x75\xfe\x6d\xd9\xa9\xc8\x52\xd5\xe1\x28\xb9\x83\xda\x7b\x89\x86\x92\xa4\xfe\xd0\xf8\xb9\x7f\xff\x74\xef\x09\xae\xbf\x5e\xc9\x39\xad\x09\x12\xd1\x63\x98\x17\xdc\x99\xf0\x8c\x36\x43\xf4\xef\x9b\x3c\xb3\x3d\x57\xfb\x03\xae\x1c\x5c\x7e\x14\xcd\xff\x08\x7f\x88\xd6\x43\xf6\x1e\xa3\x2d\xe6\x39\xf2\x18\xb8\xc6\xb2\xe2\x37\x53\xfa\x37\x00\x00\xff\xff\xad\xec\xff\x2b\xb4\x03\x00\x00"),
		},
		"/flux-helm-release-crd.yaml.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "flux-helm-release-crd.yaml.tmpl",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 7426,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x59\x5f\x6f\xdb\x38\x12\x7f\xf7\xa7\x98\xcb\x1d\x90\xe4\x10\xab\x2d\x7a\x38\xdc\x19\x28\xda\xa2\x41\xaf\xbd\xf4\x4f\x90\x34\x7d\x09\xba\xc0\x58\x1a\x49\xdc\x50\xa4\x76\x48\x3a\xf1\x2e\xf6\xbb\x2f\x48\xfd\xb1\xa5\xc8\xb2\xdd\xa6\x0f\x8b\xae\x5e\x5a\x93\x9c\xe1\x6f\x66\x7e\x33\x43\x32\xd3\xe9\x74\x82\xa5\xf8\x4c\x6c\x84\x56\x33\xc0\x52\xd0\x9d\x25\xe5\x7f\x99\xe8\xe6\x3f\x26\x12\xfa\xd1\xe2\xc9\x9c\x2c\x3e\x99\xdc\x08\x95\xcc\xe0\x95\x33\x56\x17\x17\x64\xb4\xe3\x98\x4e\x29\x15\x4a\x58\xa1\xd5\xa4\x20\x8b\x09\x5a\x9c\x4d\x00\x14\x16\x34\x83\x9c\x64\xc1\x24\x09\x0d\x99\xc8\xff\x88\x52\xe9\xee\xe2\x24\x12\x7a\x62\x4a\x8a\xfd\xca\x8c\xb5\x2b\xab\xa5\x6b\xb3\x95\x06\xe3\x17\x00\x54\xfb\xbe\x21\x59\x5c\x54\xca\xc2\xa8\x14\xc6\x9e\xf5\x67\xde\x09\x63\xc3\x6c\x29\x1d\xa3\xec\x42\x08\x13\x26\xd7\x6c\x3f\xac\x94\x4f\x21\xe7\x09\x80\x89\x75\x49\x33\x08\x13\x25\xc6\x94\x4c\x00\x30\x49\x82\x65\x28\xcf\x59\x28\x4b\xfc\x4a\x4b\x57\xa8\x56\xf0\xff\x97\x1f\x3f\x9c\xa3\xcd\x67\x10\x19\x8b\xd6\x99\xa8\xde\xc9\x6b\x09\x6b\x1a\x47\xac\xe3\x06\xb0\x4b\xbf\x95\xb1\x2c\x54\xb6\x4d\xd5\x65\xf8\xd5\x51\xd6\x19\xda\x49\x57\xac\x55\x65\x89\xb9\x7e\x7e\xf4\x22\xf2\x32\xcf\x9e\x1d\xd4\xa0\x92\x83\xe3\x2f\x51\x41\xc6\x60\xd6\x05\xfd\xbe\x33\x36\xbe\x51\x13\xfb\x28\x66\x42\xbf\xd3\x27\x51\x90\xb1\x58\x94\x1d\x95\x2f\x7b\xea\x12\xb4\x7e\xc0\xb8\x39\xd7\x7c\xaa\x9d\x5b\x01\x9f\xc1\x6f\xbf\x4f\x00\x16\x0d\x3b\x17\x4f\x56\xbf\xda\x28\x54\x9a\xc3\x54\x90\x24\x5e\x50\x32\x03\xcb\xae\xd9\xcb\x58\xcd\x98\x51\x3b\xb6\x40\x29\x92\x80\xb2\xd2\xa1\x4b\x52\x2f\xcf\xdf\x7e\x7e\x7a\x19\xe7\x54\xe0\xac\x16\x2b\x59\x97\xc4\x56\x34\x98\x82\xaa\x9a\xb5\xcd\xc7\xf4\x8b\x13\xec\xf7\xbb\x3e\x8c\x73\x64\x7b\xf8\x65\x6d\x76\x48\x43\x25\xd5\xd2\xa4\x3b\x01\x90\x90\x89\x59\x94\x01\x1c\x7c\xca\x29\x90\xbb\x11\x08\xb6\x46\xf0\x36\x05\xa5\x2d\x18\x57\x96\x52\x50\x72\x02\xc2\xc2\xad\x90\x12\xe6\x04\x19\x29\x62\xb4\x94\xc0\x7c\x09\x98\xa6\xe2\x4e\xa8\x0c\x6c\x4e\xbd\x7d\xea\x04\xf3\x54\x07\xab\xfd\x02\x68\x42\x50\xed\xd2\x5b\x7f\x2f\xfc\x6b\x56\xa2\xb5\xc4\x6a\x06\x07\x3f\x5d\xe3\xf4\xd7\xc7\xd3\xff\x7e\x39\xba\x9e\xd6\xff\xfb\x67\x33\x74\xfc\xfc\x1f\x07\x1d\x41\x8b\x9c\x91\x6d\x13\x6e\x7f\x47\x04\xb1\x01\x6f\x78\x5b\x56\xc6\x35\x8e\xf1\xa3\x66\x95\x97\xab\x0f\xcd\x7d\xeb\x2b\xd5\xdf\xdf\x05\xa2\x20\xed\xec\xa8\xe9\xc1\x6c\xa1\x8c\x45\x29\x41\x33\xb8\x32\x63\x4c\xa8\x91\x05\xa1\xc0\x90\x4f\x70\x33\x08\xd7\xd7\xad\x8c\xb8\x37\x97\x6a\x2e\xd0\x86\xd9\x7f\xff\xab\x47\x4d\x43\xf6\x33\x4a\xd7\xe7\x6c\x0f\xd6\xdb\xb4\xf5\x78\xe5\xe2\x20\xe8\x53\xcb\x91\x01\xad\x42\xe1\x6d\xc0\x0e\x22\x9b\x6b\x2d\x09\xd5\xa4\x87\x2b\xa6\xab\x4a\x68\xbf\xed\x83\x64\xe5\xac\xd6\x43\x39\x6b\x97\xe5\x90\x90\x24\x4b\x8f\x98\x42\x6d\xda\x1d\x0c\x6b\x29\xe7\x18\xdf\xf4\x81\x54\x12\x7a\xfe\x33\xc5\xb6\xcf\x83\x0d\x09\xef\x3f\x52\x38\x97\xf7\xac\x1a\xb4\x8c\xec\x49\x65\x55\x49\xec\x43\xd5\x42\x31\xde\x4e\xb0\xb9\x30\x6d\x26\x68\xd5\x1a\x9c\xa2\x90\x8e\xa9\x4f\x84\x71\x2b\x5b\xb7\xef\x87\xac\xf2\x77\x9b\x34\xae\xf4\xa5\x7c\x93\xcb\x41\xa4\xa0\x88\x92\xd0\x55\xf7\x83\xd6\xa8\x18\x42\x37\x2e\x99\x08\xe3\x1d\xfe\x46\xeb\x9b\x81\x68\x8c\x7a\x9d\x69\x41\xca\x42\xee\x45\x21\x65\x5d\x00\x3b\xa5\x7c\x21\x4d\x9c\xcf\xfe\x36\x1e\x7b\x83\xda\x90\xf1\xf7\xf0\xf8\xf6\xb9\x96\xda\xbe\x44\xdf\xa2\xb0\x21\xfc\xa8\x96\x20\x54\x22\x16\x22\x71\x28\xe1\xcc\xcd\x89\x15\x59\x9f\x76\xa5\xaf\xfd\x42\xab\x93\x01\xfd\x7e\x87\x14\x9d\xb4\x41\xdb\xd3\xc7\x8f\x37\xd4\x0d\xd8\x52\x3b\x60\xb4\x7e\xf8\xcf\x23\xdd\xcf\xe3\xc1\x36\xa7\xac\x90\xa1\x16\x17\x42\x89\xc2\x15\xa0\x5c\x31\x27\x06\x9d\xc2\xb9\x4e\x8c\xff\x17\xe1\x94\x4a\xa9\x97\x05\xa9\x7e\xee\x55\x1f\x72\xf0\x1b\x02\x13\x26\xcb\x70\x8e\x20\x98\x53\xaa\x99\xa0\x40\xbe\xa9\xbb\x61\x9b\x3e\x68\xc0\xb8\x38\x26\x63\x52\x27\xf7\x0a\x67\x28\x75\xaf\x85\xa4\x4b\x4f\x51\x3b\x5e\x2e\x4f\xa9\x64\x8a\x7d\x5f\xfe\x1b\x5c\x19\xaa\xeb\xe4\x6b\xd6\x45\x64\x82\xf8\x19\x2d\x2f\x28\x0d\x85\x9e\xb0\x9f\x26\x15\x08\x64\xc6\x65\x6f\x46\x58\x2a\x06\xd8\x3d\x52\xa2\xba\x67\x16\xdf\xec\x3a\x47\x96\xea\x1b\xab\x63\xf5\x59\x6e\x03\xc3\xd6\x6c\xf6\xad\xdd\xc7\x2c\x74\xde\x60\xe4\x09\x14\xce\x58\xdf\x8d\x85\x6a\x1b\xf2\x5a\xaf\xae\x5b\x71\xff\xa4\x3f\x6c\xdd\x40\x23\x5e\x79\x75\xb8\x66\x3f\x94\x07\xc7\xdd\x13\x6b\x95\x8a\xec\x3d\x96\x55\x4c\x87\x3d\x35\xaa\x1f\x76\x8b\xd2\x76\x28\x30\x1a\x2d\x18\x8b\x58\x65\x45\x81\xe5\x03\x05\x0d\xc6\x4f\x50\xcd\x77\x43\xcb\x1d\xc1\x9e\xd1\xb2\x41\xd4\x62\xf5\x95\x2d\x23\x1b\x06\xeb\xa3\x88\x2f\xdf\x27\x9d\xd2\x57\x4d\x44\x4b\x2c\x86\x12\x7e\x67\xa4\xba\xac\xee\x86\x3b\xc2\x6d\xea\xdd\xaa\xda\x00\x93\x65\x41\x0b\x94\x8d\xcf\x1b\xc8\x42\x12\x08\x03\x4a\x83\xd4\x2a\x23\x86\x02\x55\x82\x56\x73\x9f\xbc\x7d\xc0\x9b\xda\x0e\xc0\x7a\x95\xf9\x93\x32\xf2\x41\x6b\x08\x7c\x47\x3a\x56\x40\xff\xe2\xe2\x26\x2e\xd2\x9d\xbf\x2c\xa1\xbc\x0c\x67\xc7\x87\x21\xa4\x63\xf9\xd5\x7c\x74\xbc\xab\xe3\xae\x2e\xde\x75\xfd\xf3\x83\x45\x2e\x3c\x71\xf8\x33\xcf\xc3\x04\xad\x44\x9b\x7f\x75\xd4\xbc\xf0\x8e\x5e\xf3\x4b\xe1\x56\xd8\xbc\x4e\xd0\x70\x37\x0d\xb6\xc0\x51\x38\xde\x67\xc2\x02\x53\xa9\x8f\xe1\x36\x27\xee\x04\xd7\xbb\x50\xea\x70\x74\xfb\x51\xe2\xac\x15\x7d\x1c\x08\xef\xb4\xfb\xda\xd5\x3d\xe5\x0c\x84\xb1\xbb\x7e\xbd\x01\x6d\x5d\x7c\xaf\x42\x6c\x95\x58\x67\x66\x6f\xf1\x62\xfb\x6b\x46\xac\x95\xf5\x37\x3e\x9d\x8e\xe4\xf5\x46\x6a\x87\xbd\xfb\xfa\x07\x9d\xd8\xc5\x9c\x09\x7b\x78\x02\x9b\xb2\x60\x3c\x03\xb2\xe1\x2b\x56\xcf\xae\xff\x09\x1b\x6a\x16\x45\x59\xe4\x45\x5e\x64\xc2\xe6\x6e\x1e\xc5\xba\x98\x69\xce\x1e\x79\xce\xef\x77\xc4\x6e\xbe\xe6\x06\xe8\x33\xe7\xef\xe1\x01\x2e\xa1\x54\xa8\xea\xdd\xf1\xe3\xcb\xcb\x01\xa1\xcd\x09\xdb\xc1\x7c\xee\x93\x55\x28\x23\x92\xea\xd1\xae\xc9\x4d\x23\x3c\xa5\xeb\x04\x6d\x5a\x7c\x9d\xc5\x62\xe8\x12\xbb\xd5\x0a\xde\x54\xc4\xee\xf9\x70\xce\xa8\xe2\xbc\xdb\xba\x0b\x34\x76\xf0\x72\xbc\x75\x5f\x73\x23\xca\x53\x2a\xaf\xc2\x9b\xc9\x0e\x08\x9a\x62\x90\x68\x32\xc1\xd5\xec\x14\x1c\x26\x54\x1e\x36\xef\x2e\x47\x68\x8c\x2b\xa8\x61\x97\xbf\x1d\xaf\xaa\x17\xca\xea\x2e\x9c\x3a\x99\x0a\x29\x29\x39\x1e\x01\x3d\x5c\x13\xba\xbc\x5d\x45\xc3\xd3\x37\x1c\x05\x4f\xe0\xb0\x7e\x97\xdf\x9b\xc9\x2b\x6d\x3b\xb8\xa2\x7e\x04\x6e\xd9\x70\x75\xf1\xee\xdb\xf8\xeb\x58\xee\xca\xdf\x1d\xaf\xbd\x6b\xb4\x54\x43\xef\xcd\x3b\xc0\x6b\xfe\xe0\xb1\xd7\x66\xb5\xd0\xb7\xb9\xc3\x50\xb1\x20\xde\xd5\x23\x61\xe3\x73\x27\x65\xf5\x04\x32\x8c\xf7\x41\xef\x03\xfd\xf8\xcf\xd1\x88\x18\xd0\xd9\x1c\x8e\x3c\x64\x51\x94\x32\xd0\x7f\x13\xcb\xef\x79\xe3\x8f\x00\x00\x00\xff\xff\x4f\x4e\x49\x9d\x02\x1d\x00\x00"),
		},
		"/helm-operator-deployment.yaml.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "helm-operator-deployment.yaml.tmpl",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 6000,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x58\x4f\x73\xdb\xb8\x15\xbf\xfb\x53\xbc\xb1\x0f\xb9\x58\xa4\x33\x9b\xdd\x03\x33\x39\xb4\x9b\x6e\x92\x99\x24\xf5\xd4\x99\xce\xf4\xb4\xfb\x04\x3e\x89\xa8\x40\x80\x05\x1e\xa5\xb2\x9e\xed\x67\xef\x3c\x80\xa2\x48\x51\x92\xed\xf4\xd0\xf2\x62\x19\xc0\xfb\xff\x7b\x7f\x80\xc5\x62\x71\x85\x8d\xfe\x2b\xf9\xa0\x9d\x2d\x00\x9b\x26\xe4\xdb\xd7\x57\x1b\x6d\xcb\x02\xde\x53\x63\x5c\x57\x93\xe5\xab\x9a\x18\x4b\x64\x2c\xae\x00\x2c\xd6\x54\xc0\xca\xb4\xff\x5c\x54\x64\xea\x85\x6b\xc8\x23\x3b\xff\xf8\x08\x7a\x05\xd9\x57\xac\x29\x34\xa8\x08\x7e\xff\xbd\x3f\x1d\xff\x2d\xe0\xf1\x71\xba\xfb\xf8\x08\x64\x4b\x39\x16\x1a\x52\xc2\xda\x53\x63\xb4\xc2\x50\xc0\xeb\x2b\x80\x40\x86\x14\x3b\x2f\x3b\x00\x35\xb2\xaa\x3e\xe3\x92\x4c\x48\x0b\xe7\x35\x11\x5a\xf6\xc8\xb4\xee\xd2\x51\xee\x1a\x2a\xe0\x2f\xa4\x3c\x21\xd3\x15\x00\x53\xdd\x18\x64\xea\x59\x8f\xac\x93\xcf\x4c\xa4\x5c\x94\x23\x1f\x5a\xeb\x18\x59\x3b\x3b\xa2\x69\xbc\xab\x89\x2b\x6a\x43\xa6\x5d\x1e\x94\x47\x51\xe1\x9a\x7d\x4b\xd7\xf1\xd0\xde\xe6\xf8\x9b\xfc\x56\x2b\xfa\x83\x52\xae\xb5\xfc\xf5\xb2\xb8\xad\x33\x6d\x4d\xa1\xe8\xfd\xfd\x27\x8b\x4b\x43\xdf\xb4\x31\xe4\xbf\x7d\x7e\x48\x5e\x4f\xdf\xa2\xd7\x3c\x72\x61\x13\x16\x0a\x87\x3d\x00\xe5\xec\x4a\xaf\xbf\x60\x53\x8c\x16\xe7\xc6\x26\xba\x45\x3a\x3d\x39\x59\xd2\x0a\x5b\xc3\x5f\x5c\x49\x05\xdc\xfd\x74\x77\x77\x41\x30\x79\x0e\x23\xe2\x20\xa1\xe0\xa9\xe0\xb4\x96\x8c\xef\x6d\x1b\xac\xfa\x99\x3c\x3f\x0c\xfb\x09\x3b\x97\x77\xc9\x04\xf9\x75\x64\x07\x79\x3e\xa0\xee\xbc\x2d\x6f\xee\xee\x0e\x2c\xfa\x73\x37\xfb\xbf\xf0\x37\xd7\xc2\x4e\x1b\x03\x96\xa8\x04\xae\x28\x10\xf0\xce\xed\x03\x23\x9a\x77\x72\x04\x2d\x03\x3b\xa0\xc0\xb8\x34\x3a\x54\xb0\x45\xa3\x4b\x64\x2a\xe1\xdb\xe7\x87\x81\x9d\x72\xd6\x92\x8a\xf0\x01\x5c\xa3\xb6\x81\x21\x99\x36\x93\x7c\x3e\xa0\x37\xa7\x02\x7a\xf3\xec\x80\xde\x5c\x0c\xe8\x0d\x24\xef\xc6\x3c\x82\x4d\xbb\x24\x6f\x89\x29\x22\x9b\x4d\x98\xa9\x37\x77\xfa\x48\xcc\x34\xf4\x37\xff\xcb\xd0\x9f\xb2\xfa\xcd\xc1\xea\xc7\x47\xb2\xe5\xe8\xf0\xb7\x8a\x60\xe5\x8c\x71\x3b\x6d\xd7\x7d\xb4\x41\x07\x58\x39\x0f\x6d\x90\x35\x04\xd5\x06\x76\xb5\x0e\x54\xc2\xc6\xba\x9d\xfd\xb5\x72\x81\x03\xac\xb4\xa1\xdb\x81\xd1\xae\xd2\xaa\x4a\x18\x39\xc0\xc8\x41\xe9\xf6\xd0\x11\x22\xf9\xe1\xc1\xed\x2c\xac\x35\x4b\x65\x74\xe0\x91\xab\x03\x2a\x80\x2b\xb4\xbd\xe0\xb5\xe6\xaa\x5d\x82\xf3\x02\x47\x30\x7a\x43\x99\xc0\xf4\x95\x31\x80\x26\xb8\x41\x44\x2d\xf5\x05\xf4\x21\x1e\xda\xb2\x8b\x34\xca\x59\x46\x6d\xc9\xdf\xc2\x92\x8c\xdb\x65\x27\x61\x5f\x63\x97\x18\xee\x04\xcf\xec\xa4\xcc\x6d\x75\x49\x80\x16\x42\xa8\x7e\x4d\xa0\x3a\x32\x57\x3a\x88\x76\x56\xf4\xac\x9d\xa7\xa4\xb7\xb3\x04\xbf\x7d\x2a\x65\x8b\xbb\x5f\xb4\xa1\xdf\xde\x46\x47\x0a\xfc\xd1\x2a\xba\xed\x7d\xf1\xca\xd3\xc0\x28\xd9\x3a\xe5\xf1\x41\xf3\xc7\x76\x19\xfd\x93\xc1\xd7\x3f\x46\x5b\xc8\xb2\xef\x60\x43\x1d\x84\xca\xb5\xa6\x84\xe5\x81\xc7\x75\x52\xf1\xba\x77\x66\x62\x74\x7d\xd0\xfd\x5a\xe4\x46\x37\x51\x09\xda\xc2\xbf\xf3\x2c\x84\x2a\x9f\xbb\x63\x0f\xf6\x10\xaa\x52\xfb\x17\xa5\x61\x08\xd5\xd3\xe9\x97\x6a\x90\xa4\xc2\xc3\xc3\xc7\x09\xc4\xaf\x0e\x69\xf9\xf0\x31\x9a\xc9\x0e\x50\x29\x0a\x21\x9a\xff\xa1\xc7\x4b\xd0\xec\x7c\x37\x2b\xca\x6b\xcd\x8b\x0d\x75\x2f\xab\xc6\x73\x25\xc6\x87\x67\x9a\x47\x90\x93\x1d\x1c\xe9\x09\xcb\x85\xb3\xa6\xbb\x85\x1d\xc1\xce\xd9\x57\x0c\x4b\x02\xe9\x5c\xa2\xbc\xaa\x6a\x57\x5e\xbd\xa0\xe4\xea\x30\xe4\xdf\x1e\x25\x43\x0a\x0e\xe9\xc2\x15\x1e\x80\x2e\x84\x41\x60\xba\xf7\x99\x80\x2d\x39\xed\x2d\x50\xb6\xce\x6e\x01\xf7\x60\x2a\xe3\xe0\x23\xa7\x32\xf8\xb4\x1a\x58\x4c\xe4\xfc\xbd\x0d\x1c\x01\x18\x5a\x55\x45\x79\xb7\xd1\xf9\xbd\x2b\x46\xd9\x30\xd0\xa3\x11\x37\x74\xd0\x38\x6d\x39\x00\x32\xe4\xc4\x2a\x17\x48\x94\xb9\x80\x4c\xf7\xe9\x00\x18\x00\xf7\xe2\x45\xec\xa1\x72\xf4\x3d\xa5\x0d\x74\x94\x07\x1b\xea\x6e\xa3\x86\xa3\x82\xb2\x4f\xce\x7d\x25\x19\xd8\x8c\x52\x15\x97\x6e\x4b\xb7\xb0\xd3\x5c\x89\x77\xa6\x29\xd9\x67\x52\x1c\xbd\xc4\x68\x42\x55\x0d\x4c\xc4\x89\xda\x46\xa3\x13\x58\xf6\x89\x4e\x25\x54\xe4\xe9\x7c\xca\x4c\x11\xf8\x9c\xa6\x10\xd3\x46\xc8\x52\x68\x2e\xa7\xcd\x77\x81\xef\x7c\xcd\x3f\xea\xee\xe8\x29\x62\x27\x72\x4f\x88\x1b\x92\x4d\x53\xc8\x3a\xac\xcd\x51\x01\x44\x5b\xf6\xb1\xe8\x9b\x04\x2a\x41\x8a\xf6\x71\xbc\xed\xb2\xd8\x59\x0c\x32\x93\x97\x7e\x22\xe1\x23\x29\x5a\x0a\xdb\x70\xa8\x5c\x83\x40\x8e\x7d\xc8\xd7\xe4\x53\x4e\xd4\xb8\xa1\x54\xc3\x85\x6f\x7e\x60\x7c\xb0\xfc\x7c\x2c\xc6\xba\x2f\x44\xf7\x97\x46\x25\x76\xd8\x31\x97\xcb\x22\xa2\x8e\x23\x8e\x54\x37\xdc\xbd\xd7\xbe\x80\xc7\xa1\xb0\x0d\xbd\x68\x98\xa7\xe7\x93\xc5\xd1\x50\xdc\xc7\xca\x53\x8c\x8f\x75\x70\x5d\xc8\x7c\x1f\xf8\x1a\x74\x8d\x6b\x4a\x5d\x7a\x42\x99\xc1\x2f\xda\xc6\xf9\x0d\x6a\xe9\xb7\x9e\x94\x5c\x75\x0e\xfc\x3c\x19\xc2\x40\xd2\x55\x23\x0f\xd8\xa6\x7b\x92\x64\x6e\xc5\xdc\x84\x22\xcf\xab\x76\x99\x95\x4e\x6d\xc8\x67\xca\xd5\xb9\xcf\x77\x84\x5b\xda\x39\xbf\x09\xf9\x44\x5a\xce\xb8\x0e\x23\xe6\x82\x09\xb9\xee\xc8\x55\x48\x54\x60\x5c\x4f\xb2\x06\x92\xcc\x02\x7a\xee\xda\xc5\x42\xa1\xca\x29\xdb\xe2\x75\x76\x97\xdd\x2d\xbc\x7a\x3d\xa5\xbb\x6f\x8d\xb9\x77\x46\xab\xae\x80\x4f\xab\xaf\x8e\xef\x3d\x85\xb1\x79\x8d\xf3\x3c\xba\xae\x0c\x93\x25\x73\x33\xbd\x23\xa4\x48\xdc\x3b\xcf\x05\xfc\x70\xf7\xc3\x61\xc8\x37\x7a\x4b\x96\x42\xb8\xf7\x6e\x49\xe3\xe6\x21\x3c\x3e\x1c\xf7\x93\x66\xce\x20\x2e\x23\x57\x05\xe4\x15\xa1\xe1\xea\x5f\xa3\x2d\x6d\x35\x6b\x34\xef\xc9\x60\xf7\x40\xca\xd9\xb2\xbf\x15\xee\x3f\xd6\x35\xb9\x96\x87\xbd\x1f\x87\x3d\x41\xbd\xfe\x3f\xd5\x2c\xb8\xd6\x2b\x0a\x63\x0d\x3c\xfd\xa3\xa5\xc0\x61\xaa\x95\x6a\xda\x02\x7e\xbc\xab\x27\x8b\x35\xd5\xce\x77\x05\xfc\xf4\xe6\x8b\x1e\x36\x52\x61\xfa\x22\xd5\x61\xc4\xe3\x06\x3e\x59\x65\xda\x92\x52\xb7\xec\x07\xcb\xe9\x1c\x78\x76\x5c\x75\x7e\xde\xbf\x84\xa5\x94\xb5\xb7\xfb\x2e\x33\x1a\x2c\x2b\xda\xb7\xe3\x92\x94\x41\x4f\x65\xea\x2b\xd9\x88\xf6\xe4\xbc\x94\x6a\x40\xd4\xe6\x3e\xf9\xdb\x3b\xc7\x71\xe4\x9a\x9c\x90\x90\xfe\xd9\x9a\xae\x00\xb9\x40\x3f\x31\x17\xc1\xc5\x61\x67\x2a\x6e\xd2\x80\xe7\x13\xc8\xf9\xae\x35\xd7\x7c\xc2\x6a\x7e\xcb\x7c\xba\xe4\xce\x59\x6e\xd1\xf7\x2c\x25\xe9\xf3\x13\x43\xdd\x33\xaa\xec\x0b\xd8\xe6\x91\xee\xfc\xc3\xc2\x09\x07\x9f\xbd\xe2\x9f\x73\x8e\x10\x4c\x8e\x4d\x63\x7b\x51\x04\x3e\x93\xff\xf1\xc9\x23\x11\x17\xc3\x7c\xc6\xa0\xf3\xd1\x9e\x18\x34\x07\xeb\x45\x11\xf8\x4c\xfe\xc7\x27\x67\xf9\x70\x04\x36\xf4\xeb\x49\x2d\xf8\xe8\x76\x92\xf6\x2b\x69\x78\x93\xb7\x05\xf1\xf2\x62\xc1\x71\x69\x31\xbc\xd5\xbd\x9b\x5c\xc2\x8f\xde\xec\x4e\xaf\xf6\x0e\xdd\xb4\x4b\x5a\x84\x2e\x30\xd5\x83\x56\x4f\x3f\x54\xdd\x80\x2c\xa4\x72\xd3\xfa\xf8\x94\x76\x4a\xbf\xfd\xcb\x05\x79\x5e\x48\x79\x7e\x37\xf7\x52\xae\x30\x53\x9e\xcf\x51\x53\xd4\xe1\xdd\x24\x2e\x47\x47\x36\xd4\x9d\x64\x9e\xb3\x09\xd9\x38\xff\x8f\x15\x3b\xa7\x55\x24\xbc\xa0\xd3\x96\xbc\x5e\x75\x17\x75\x7a\x9e\xd1\x27\x61\xfd\xb3\xab\x65\x2c\x07\xd7\xc6\x8b\xc3\xf9\xf7\xa8\xa7\xdf\xa1\x12\x84\xff\x9b\x68\xcc\xe8\x4f\xc5\x63\x76\xe8\xd9\x11\x99\xab\xf7\xec\x98\xcc\x48\x4f\x45\xe5\x3b\xcd\x3f\xa4\xe7\x7f\x02\x00\x00\xff\xff\x9e\xea\x11\xa6\x70\x17\x00\x00"),
		},
		"/tiller-ca-cert-configmap.yaml.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "tiller-ca-cert-configmap.yaml.tmpl",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 226,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8e\xc1\x0a\xc2\x30\x0c\x40\xef\xfd\x8a\xfc\x40\x07\x82\xa7\xde\xa4\x78\x53\x2f\x0e\xef\xb1\xcd\xb4\xd8\x66\xa3\x8b\x22\xd4\xfe\xbb\x6c\x4c\xd1\x63\x78\x2f\x2f\x29\x05\x42\x07\xcd\x96\xf1\x1c\xa9\x0d\x31\x52\x6e\x77\x47\xa8\x55\x6b\xad\x70\x08\x27\xca\x63\xe8\xd9\xc0\x63\xa5\x6e\x81\xbd\x01\xdb\x73\x17\x2e\x7b\x1c\x54\x22\x41\x8f\x82\x46\x01\x30\x26\x32\xd0\xc5\xfb\x53\x5f\x29\x26\x2d\x71\xd4\x0e\xb5\x9b\xe5\xe5\xc8\x01\x13\x8d\x03\x3a\x82\x5a\x97\x95\x79\x34\x50\xca\x3f\x2d\x05\x88\xfd\xa4\x7d\xfa\x0e\x1b\x97\xc5\xc0\x4b\x4d\x31\xf6\xc4\x02\x6b\x68\xbe\x1f\xdb\x8d\xa5\x2c\xb6\x67\x99\xc8\x6f\xe1\x1d\x00\x00\xff\xff\x60\xd7\x1c\xac\xe2\x00\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/flux-helm-operator-account.yaml.tmpl"].(os.FileInfo),
		fs["/flux-helm-release-crd.yaml.tmpl"].(os.FileInfo),
		fs["/helm-operator-deployment.yaml.tmpl"].(os.FileInfo),
		fs["/tiller-ca-cert-configmap.yaml.tmpl"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}
