package menu

import (
	"fmt"
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

func TestCreate(t *testing.T) {
	mockResp := map[string][]byte{
		"case1": []byte(`{"errcode":0,"errmsg":"ok"}`),
		"case2": []byte(`{"errcode":40013,"errmsg":"invalid appid"}`),
	}
	var resp []byte
	test.MockSvrHandler.HandleFunc(apiCreate, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(resp))
	})

	type args struct {
		buttons []Button
	}
	tests := []struct {
		name     string
		args     args
		wantResp []byte
		wantErr  bool
	}{
		{name: "case1", args: args{buttons: nil}, wantResp: mockResp["case1"], wantErr: false},
		{name: "case2", args: args{buttons: nil}, wantResp: mockResp["case2"], wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp = mockResp[tt.name]
			gotResp, err := Create(tt.args.buttons)
			fmt.Println(string(gotResp), err)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("Create() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestGet(t *testing.T) {
	mockResp := map[string][]byte{
		"case1": []byte(`{
		"menu": {
			"button": [
				{
					"type": "click", 
					"name": "今日歌曲", 
					"key": "V1001_TODAY_MUSIC", 
					"sub_button": [ ]
				}, 
				{
					"type": "click", 
					"name": "歌手简介", 
					"key": "V1001_TODAY_SINGER", 
					"sub_button": [ ]
				}, 
				{
					"name": "菜单", 
					"sub_button": [
						{
							"type": "view", 
							"name": "搜索", 
							"url": "http://www.soso.com/", 
							"sub_button": [ ]
						}, 
						{
							"type": "view", 
							"name": "视频", 
							"url": "http://v.qq.com/", 
							"sub_button": [ ]
						}, 
						{
							"type": "click", 
							"name": "赞一下我们", 
							"key": "V1001_GOOD", 
							"sub_button": [ ]
						}
					]
				}
			]
		}
	}`),
		"case2": []byte(`{"errcode":40013,"errmsg":"invalid appid"}`),
	}
	var resp []byte
	test.MockSvrHandler.HandleFunc(apiGet, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(resp))
	})

	tests := []struct {
		name     string
		wantResp []byte
		wantErr  bool
	}{
		{name: "case1", wantResp: mockResp["case1"], wantErr: false},
		{name: "case2", wantResp: mockResp["case2"], wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp = mockResp[tt.name]
			gotResp, err := Get()
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("Get() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestGetCurrentSelfMenuInfo(t *testing.T) {
	mockResp := map[string][]byte{
		"case1": []byte(`{ 
	   "is_menu_open": 1, 
	   "selfmenu_info": { 
		   "button": [ 
			   { 
				   "type": "click", 
				   "name": "今日歌曲", 
				   "key": "V1001_TODAY_MUSIC"
			   }, 
			   { 
				   "name": "菜单", 
				   "sub_button": { 
					   "list": [ 
						   { 
							   "type": "view", 
							   "name": "搜索", 
							   "url": "http://www.soso.com/"
						   }, 
						   { 
							   "type": "view", 
							   "name": "视频", 
							   "url": "http://v.qq.com/"
						   }, 
						   { 
							   "type": "click", 
							   "name": "赞一下我们", 
							   "key": "V1001_GOOD"
						   }
					   ]
				   }
			   }
		   ]
	   }
	}`),
		"case2": []byte(`{"errcode":40013,"errmsg":"invalid appid"}`),
	}
	var resp []byte
	test.MockSvrHandler.HandleFunc(apiGetCurrentSelfMenuInfo, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(resp))
	})

	tests := []struct {
		name     string
		wantResp []byte
		wantErr  bool
	}{
		{name: "case1", wantResp: mockResp["case1"], wantErr: false},
		{name: "case2", wantResp: mockResp["case2"], wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp = mockResp[tt.name]
			gotResp, err := GetCurrentSelfMenuInfo()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCurrentSelfMenuInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("GetCurrentSelfMenuInfo() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	mockResp := map[string][]byte{
		"case1": []byte(`{"errcode":0,"errmsg":"ok"}`),
		"case2": []byte(`{"errcode":40013,"errmsg":"invalid appid"}`),
	}
	var resp []byte
	test.MockSvrHandler.HandleFunc(apiDelete, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(resp))
	})

	tests := []struct {
		name     string
		wantResp []byte
		wantErr  bool
	}{
		{name: "case1", wantResp: mockResp["case1"], wantErr: false},
		{name: "case2", wantResp: mockResp["case2"], wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp = mockResp[tt.name]
			gotResp, err := Delete()
			if (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("Delete() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestAddConditional(t *testing.T) {
	mockResp := map[string][]byte{
		"case1": []byte(`{"errcode":0,"errmsg":"ok"}`),
		"case2": []byte(`{"errcode":40013,"errmsg":"invalid appid"}`),
	}
	var resp []byte
	test.MockSvrHandler.HandleFunc(apiAddConditional, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(resp))
	})

	type args struct {
		buttons   []Button
		matchRule MatchRule
	}
	tests := []struct {
		name     string
		args     args
		wantResp []byte
		wantErr  bool
	}{
		{name: "case1", args: args{buttons: nil, matchRule: MatchRule{}}, wantResp: mockResp["case1"], wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp = mockResp[tt.name]
			gotResp, err := AddConditional(tt.args.buttons, tt.args.matchRule)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddConditional() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("AddConditional() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestDelConditional(t *testing.T) {
	mockResp := map[string][]byte{
		"case1": []byte(`{"errcode":0,"errmsg":"ok"}`),
		"case2": []byte(`{"errcode":40013,"errmsg":"invalid appid"}`),
	}
	var resp []byte
	test.MockSvrHandler.HandleFunc(apiDelConditional, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(resp))
	})

	type args struct {
		menuId string
	}
	tests := []struct {
		name     string
		args     args
		wantResp []byte
		wantErr  bool
	}{
		{name: "case1", args: args{menuId: "100"}, wantResp: mockResp["case1"], wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp = mockResp[tt.name]
			gotResp, err := DelConditional(tt.args.menuId)
			if (err != nil) != tt.wantErr {
				t.Errorf("DelConditional() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("DelConditional() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestTryMatch(t *testing.T) {
	mockResp := map[string][]byte{
		"case1": []byte(`{"errcode":0,"errmsg":"ok"}`),
		"case2": []byte(`{"errcode":40013,"errmsg":"invalid appid"}`),
	}
	var resp []byte
	test.MockSvrHandler.HandleFunc(apiTryMatch, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(resp))
	})

	type args struct {
		userId string
	}
	tests := []struct {
		name     string
		args     args
		wantResp []byte
		wantErr  bool
	}{
		{name: "case1", args: args{userId: "100"}, wantResp: mockResp["case1"], wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp = mockResp[tt.name]
			gotResp, err := TryMatch(tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("TryMatch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("TryMatch() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
