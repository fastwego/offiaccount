package type_event

import (
	"encoding/xml"
	"fmt"
	"testing"
)

func TestEventMenuPicSysPhoto(t *testing.T) {
	s := `<xml><ToUserName><![CDATA[gh_e136c6e50636]]></ToUserName>
<FromUserName><![CDATA[oMgHVjngRipVsoxg6TuX3vz6glDg]]></FromUserName>
<CreateTime>1408090651</CreateTime>
<MsgType><![CDATA[event]]></MsgType>
<Event><![CDATA[pic_sysphoto]]></Event>
<EventKey><![CDATA[6]]></EventKey>
<SendPicsInfo><Count>3</Count>
<PicList>
<item><PicMd5Sum><![CDATA[1b5f7c23b5bf75682a53e7b6d163e185]]></PicMd5Sum></item>
<item><PicMd5Sum><![CDATA[2b5f7c23b5bf75682a53e7b6d163e185]]></PicMd5Sum></item>
<item><PicMd5Sum><![CDATA[3b5f7c23b5bf75682a53e7b6d163e185]]></PicMd5Sum></item>
</PicList>
</SendPicsInfo>
</xml>`

	photo := EventMenuPicSysPhoto{}
	err := xml.Unmarshal([]byte(s), &photo)
	if err != nil {
		t.Fatalf("xml parser error=%+v\n", err)
	}
	fmt.Printf("photo=%+v\n", photo)
}
