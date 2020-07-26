package menu_test

import (
	"fmt"
	"os"

	filecache "github.com/faabiosr/cachego/file"
	"github.com/fastwego/offiaccount"
	"github.com/fastwego/offiaccount/menu"
)

func Example() {
	offiaccount.Init("wx5430528a560c5a03", "5462c323b5d36a44e0007764e715964e")
	offiaccount.SetAccessTokenCache(filecache.New(os.TempDir())) // 开发模式 使用 文件缓存 access_token

	buttons := []menu.Button{
		{
			Type: menu.ButtonTypeClick,
			Name: "ClickMe",
			Key:  "click_key",
		},
		{
			Name: "pic",
			SubButtons: []menu.Button{
				{
					Type: menu.ButtonTypePicWeixin,
					Name: "sub1",
					Key:  "key1",
				},
			},
		},
	}
	create, err2 := menu.Create(buttons)
	fmt.Println(string(create), err2)

	get, err := menu.Get()
	fmt.Println(string(get), err)

	matchRule := menu.MatchRule{
		Sex: "0",
	}
	conditional, err2 := menu.AddConditional(buttons, matchRule)
	fmt.Println(string(conditional), err)

	info, err := menu.GetCurrentSelfMenuInfo()
	fmt.Println(string(info), err)

	match, err2 := menu.TryMatch("weixin")
	fmt.Println(string(match), err)

}
