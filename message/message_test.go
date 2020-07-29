package message

import (
	"net/http"
	"os"
	"reflect"
	"testing"

	"github.com/fastwego/offiaccount/test"
)

func TestMain(m *testing.M) {
	test.Setup()
	os.Exit(m.Run())
}
func TestGetCurrentAutoreplyInfo(t *testing.T) {
	mockResp := map[string][]byte{
		"case1": []byte("{\"errcode\":0,\"errmsg\":\"ok\"}"),
	}
	var resp []byte
	test.MockSvrHandler.HandleFunc(apiGetCurrentAutoreplyInfo, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(resp))
	})

	type args struct {
	}
	tests := []struct {
		name     string
		args     args
		wantResp []byte
		wantErr  bool
	}{
		{name: "case1", args: args{}, wantResp: mockResp["case1"], wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp = mockResp[tt.name]
			gotResp, err := GetCurrentAutoreplyInfo()
			//fmt.Println(string(gotResp), err)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCurrentAutoreplyInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("GetCurrentAutoreplyInfo() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
