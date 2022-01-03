package main

import (
	"testing"

	"github.com/onns/zip"
)

/*
@Time : 2022/1/3 15:09
@Author : onns
@File : zip-encrypt/onns_test.go
*/

func TestEncrypt(t *testing.T) {
	type args struct {
		zipFile    string
		src        string
		password   string
		encryption zip.EncryptionMethod
		withRoot   bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "加密文件夹",
			args: args{
				zipFile:    "test-folder.zip",
				src:        "test-dir",
				password:   "pass",
				encryption: zip.AES128Encryption,
				withRoot:   false,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Encrypt(tt.args.zipFile, tt.args.src, tt.args.password, tt.args.encryption, tt.args.withRoot); (err != nil) != tt.wantErr {
				t.Errorf("Encrypt() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
