package utils

import (
	"testing"
)

func TestMd5(t *testing.T) {
	if MD5Hash("123456")!= "e10adc3949ba59abbe56e057f20f883e"{
		t.Error("Md5Hash err", MD5Hash("123456"))
	}
	if MD5Hash("apple") != "1f3870be274f6c49b3e31a0c6728957f" {
		t.Error("Md5Hash err", MD5Hash("apple"))
	}
}
