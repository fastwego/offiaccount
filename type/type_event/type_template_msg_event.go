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

const (
	EventTypeTemplateSendJobFinish = "TEMPLATESENDJOBFINISH" // 模版消息发送任务完成
)

/*
<xml>
  <ToUserName><![CDATA[gh_7f083739789a]]></ToUserName>
  <FromUserName><![CDATA[oia2TjuEGTNoeX76QEjQNrcURxG8]]></FromUserName>
  <CreateTime>1395658920</CreateTime>
  <MsgType><![CDATA[event]]></MsgType>
  <Event><![CDATA[TEMPLATESENDJOBFINISH]]></Event>
  <MsgID>200163836</MsgID>
  <Status><![CDATA[success]]></Status>
</xml>

*/
type EventTemplateSendJobFinish struct {
	Event
	MsgID  string
	Status string
}
