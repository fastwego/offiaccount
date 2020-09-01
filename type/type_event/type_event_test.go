// Copyright 2020 FastWeGo
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package type_event

import (
	"encoding/xml"
	"testing"
)

func TestEventMenuPicSysPhoto(t *testing.T) {
	s := `
	<xml>
		<ToUserName><![CDATA[gh_e136c6e50636]]></ToUserName>
		<FromUserName><![CDATA[oMgHVjngRipVsoxg6TuX3vz6glDg]]></FromUserName>
		<CreateTime>1408090651</CreateTime>
		<MsgType><![CDATA[event]]></MsgType>
		<Event><![CDATA[pic_sysphoto]]></Event>
		<EventKey><![CDATA[6]]></EventKey>
		<SendPicsInfo>
			<Count>3</Count>
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
		t.Errorf("xml parser error=%+v\n", err)
	}

	if len(photo.SendPicsInfo.PicList.Item) != 3 {
		t.Errorf("xml parser error %+v\n", photo)
	}
}
