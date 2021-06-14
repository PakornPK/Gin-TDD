package getdata_test

import (
	"io"
	"net/http"
	"testing"
)

func TestGetData(t *testing.T) {
	router := "http://127.0.0.1:3000/api/v1/getdata"
	resp, err := http.Get(router)
	if err != nil {
		t.Errorf(err.Error())
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf(err.Error())
	}
	mock := `{"data":"ok"}`

	if v := string(body); v != mock {
		t.Errorf("TestGetData: FAIL: result != mock  RESULT: %v\n",v)
	}
}
