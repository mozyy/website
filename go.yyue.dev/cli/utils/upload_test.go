package utils

import "testing"

func TestUploadDev(t *testing.T) {

	UploadDev("", "")
}

func TestUploadDeploy(t *testing.T) {
	err := UploadDeploy("", "http://192.168.88.122:9990/upfile.php?pk_id=202")
	if err != nil {
		t.Error(err)
	}
}