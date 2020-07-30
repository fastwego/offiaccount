package type_event

const (
	EventTypeCardPassChecke           = "card_pass_check"              // 卡券通过审核
	EventTypeCardNotPassChecke        = "card_not_pass_check"          // 卡券未通过审核
	EventTypeUserGetCard              = "user_get_card"                // 领取事件推送
	EventTypeUserGiftingCard          = "user_gifting_card"            // 转赠事件推送
	EventTypeUserDelCard              = "user_del_card"                // 删除事件推送
	EventTypeUserConsumeCard          = "user_consume_card"            // 核销事件推送
	EventTypeUserPayFromPayCell       = "user_pay_from_pay_cell"       // 买单事件推送
	EventTypeUserViewCard             = "user_view_card"               // 进入会员卡事件推送
	EventTypeUserEnterSessionFromCard = "user_enter_session_from_card" // 从卡券进入公众号会话事件推送
	EventTypeUpdateMemberCard         = "update_member_card"           // 会员卡内容更新事件
	EventTypeCardSkuRemind            = "card_sku_remind"              // 库存报警事件
	EventTypeCardPayOrder             = "card_pay_order"               // 券点流水详情事件
	EventTypeSubmitMembercardUserInfo = "submit_membercard_user_info"  // 会员卡激活事件推送
)

/*
<xml>
   <ToUserName><![CDATA[toUser]]></ToUserName>
    <FromUserName><![CDATA[FromUser]]></FromUserName>
    <CreateTime>123456789</CreateTime>
    <MsgType><![CDATA[event]]></MsgType>
    <Event><![CDATA[card_pass_check]]></Event> //不通过为card_not_pass_check
    <CardId><![CDATA[cardid]]></CardId>
    <RefuseReason><![CDATA[非法代制]]></RefuseReason>
</xml>

*/
type EventCardPassChecke struct {
	Event
	CardId       string
	RefuseReason string
}

/*
<xml>
   <ToUserName><![CDATA[toUser]]></ToUserName>
    <FromUserName><![CDATA[FromUser]]></FromUserName>
    <CreateTime>123456789</CreateTime>
    <MsgType><![CDATA[event]]></MsgType>
    <Event><![CDATA[card_not_pass_check]]></Event> //不通过为card_not_pass_check
    <CardId><![CDATA[cardid]]></CardId>
    <RefuseReason><![CDATA[非法代制]]></RefuseReason>
</xml>

*/
type EventCardNotPassChecke struct {
	Event
	CardId       string
	RefuseReason string
}

/*
<xml>
   <ToUserName> <![CDATA[gh_fc0a06a20993]]> </ToUserName>
    <FromUserName> <![CDATA[oZI8Fj040-be6rlDohc6gkoPOQTQ]]> </FromUserName>
    <CreateTime>1472551036</CreateTime>
    <MsgType> <![CDATA[event]]> </MsgType>
    <Event> <![CDATA[user_get_card]]> </Event>
    <CardId> <![CDATA[pZI8Fjwsy5fVPRBeD78J4RmqVvBc]]> </CardId>
    <IsGiveByFriend>0</IsGiveByFriend>
    <UserCardCode> <![CDATA[226009850808]]> </UserCardCode>
    <FriendUserName> <![CDATA[]]> </FriendUserName>
    <OuterId>0</OuterId>
    <OldUserCardCode> <![CDATA[]]> </OldUserCardCode>
    <OuterStr> <![CDATA[12b]]> </OuterStr>
    <IsRestoreMemberCard>0</IsRestoreMemberCard>
    <IsRecommendByFriend>0</IsRecommendByFriend>
    <UnionId>o6_bmasdasdsad6_2sgVt7hMZOPfL</UnionId>
</xml>
*/
type EventUserGetCard struct {
	Event
	CardId              string `xml:"CardId"`
	IsGiveByFriend      string `xml:"IsGiveByFriend"`
	UserCardCode        string `xml:"UserCardCode"`
	FriendUserName      string `xml:"FriendUserName"`
	OuterId             string `xml:"OuterId"`
	OldUserCardCode     string `xml:"OldUserCardCode"`
	OuterStr            string `xml:"OuterStr"`
	IsRestoreMemberCard string `xml:"IsRestoreMemberCard"`
	IsRecommendByFriend string `xml:"IsRecommendByFriend"`
	UnionId             string `xml:"UnionId"`
}

/*
<xml>
  <ToUserName><![CDATA[gh_3fcea188bf78]]></ToUserName>
  <FromUserName><![CDATA[obLatjjwDolFjRRd3doGIdwNqRXw]]></FromUserName>
  <CreateTime>1474181868</CreateTime>
  <MsgType><![CDATA[event]]></MsgType>
  <Event><![CDATA[user_gifting_card]]></Event>
  <CardId><![CDATA[pbLatjhU-3pik3d4PsbVzvBxZvJc]]></CardId>
  <UserCardCode><![CDATA[297466945104]]></UserCardCode>
  <IsReturnBack>0</IsReturnBack>
  <FriendUserName><![CDATA[obLatjlNerkb62HtSdQUx66C4NTU]]></FriendUserName>
  <IsChatRoom>0</IsChatRoom>
</xml>
*/
type EventUserGiftingCard struct {
	Event
	CardId         string `xml:"CardId"`
	UserCardCode   string `xml:"UserCardCode"`
	IsReturnBack   string `xml:"IsReturnBack"`
	FriendUserName string `xml:"FriendUserName"`
	IsChatRoom     string `xml:"IsChatRoom"`
}

/*
<xml>
 <ToUserName><![CDATA[toUser]]></ToUserName>
 <FromUserName><![CDATA[FromUser]]></FromUserName>
 <CreateTime>123456789</CreateTime>
 <MsgType><![CDATA[event]]></MsgType>
 <Event><![CDATA[user_del_card]]></Event>
 <CardId><![CDATA[cardid]]></CardId>
 <UserCardCode><![CDATA[12312312]]></UserCardCode>
</xml>
*/
type EventUserDelCard struct {
	Event
	CardId       string `xml:"CardId"`
	UserCardCode string `xml:"UserCardCode"`
}

/*
<xml>
    <ToUserName> <![CDATA[gh_fc0a06a20993]]> </ToUserName>
    <FromUserName> <![CDATA[oZI8Fj040-be6rlDohc6gkoPOQTQ]]> </FromUserName>
    <CreateTime>1472549042</CreateTime>
    <MsgType> <![CDATA[event]]> </MsgType>
    <Event> <![CDATA[user_consume_card]]> </Event>
    <CardId> <![CDATA[pZI8Fj8y-E8hpvho2d1ZvpGwQBvA]]> </CardId>
    <UserCardCode> <![CDATA[452998530302]]> </UserCardCode>
    <ConsumeSource> <![CDATA[FROM_API]]> </ConsumeSource>
    <LocationName> <![CDATA[]]> </LocationName>
    <StaffOpenId> <![CDATA[oZ********nJ3bPJu_Rtjkw4c]]> </StaffOpenId>
    <VerifyCode> <![CDATA[]]> </VerifyCode>
    <RemarkAmount> <![CDATA[]]> </RemarkAmount>
    <OuterStr> <![CDATA[xxxxx]]> </OuterStr>
</xml>
*/
type EventUserConsumeCard struct {
	Event
	CardId        string `xml:"CardId"`
	UserCardCode  string `xml:"UserCardCode"`
	ConsumeSource string `xml:"ConsumeSource"`
	LocationName  string `xml:"LocationName"`
	StaffOpenId   string `xml:"StaffOpenId"`
	VerifyCode    string `xml:"VerifyCode"`
	RemarkAmount  string `xml:"RemarkAmount"`
	OuterStr      string `xml:"OuterStr"`
}

/*
<xml>
    <ToUserName><![CDATA[gh_e2243xxxxxxx]]></ToUserName>
    <FromUserName><![CDATA[oo2VNuOUuZGMxxxxxxxx]]></FromUserName>
    <CreateTime>1442390947</CreateTime>
    <MsgType><![CDATA[event]]></MsgType>
    <Event><![CDATA[user_pay_from_pay_cell]]></Event>
    <CardId><![CDATA[po2VNuCuRo-8sxxxxxxxxxxx]]></CardId>
    <UserCardCode><![CDATA[38050000000]]></UserCardCode>
    <TransId><![CDATA[10022403432015000000000]]></TransId>
    <LocationId>291710000</LocationId>
    <Fee><![CDATA[10000]]></Fee>
    <OriginalFee><![CDATA[10000]]> </OriginalFee>
</xml>
*/
type EventUserPayFromPayCell struct {
	Event
	CardId       string `xml:"CardId"`
	UserCardCode string `xml:"UserCardCode"`
	TransId      string `xml:"TransId"`
	LocationId   string `xml:"LocationId"`
	Fee          string `xml:"Fee"`
	OriginalFee  string `xml:"OriginalFee"`
}

/*
<xml>
    <ToUserName> <![CDATA[gh_fcxxxx6a20993]]> </ToUserName>
    <FromUserName> <![CDATA[oZI8Fj040-xxxxx6gkoPOQTQ]]> </FromUserName>
    <CreateTime>1467811138</CreateTime>
    <MsgType> <![CDATA[event]]> </MsgType>
    <Event> <![CDATA[user_view_card]]> </Event>
    <CardId> <![CDATA[pZI8Fj2ezBbxxxxxT2UbiiWLb7Bg]]> </CardId>
    <UserCardCode> <![CDATA[4xxxxxxxx8558]]> </UserCardCode>
    <OuterStr> <![CDATA[12b]]> </OuterStr>
</xml>
*/
type EventUserViewCard struct {
	Event
	CardId       string `xml:"CardId"`
	UserCardCode string `xml:"UserCardCode"`
	OuterStr     string `xml:"OuterStr"`
}

/*
<xml>
   <ToUserName><![CDATA[toUser]]></ToUserName>
    <FromUserName><![CDATA[FromUser]]></FromUserName>
    <CreateTime>123456789</CreateTime>
    <MsgType><![CDATA[event]]></MsgType>
    <Event><![CDATA[user_enter_session_from_card]]></Event>
    <CardId><![CDATA[cardid]]></CardId>
    <UserCardCode><![CDATA[12312312]]></UserCardCode>
</xml>
*/
type EventUserEnterSessionFromCard struct {
	Event
	CardId       string `xml:"CardId"`
	UserCardCode string `xml:"UserCardCode"`
}

/*
<xml>
  <ToUserName><![CDATA[gh_9e1765b5568e]]></ToUserName>
    <FromUserName><![CDATA[ojZ8YtyVyr30HheH3CM73y7h4jJE]]></FromUserName>
    <CreateTime>1445507140</CreateTime>
    <MsgType><![CDATA[event]]></MsgType>
    <Event><![CDATA[update_member_card]]></Event>
    <CardId><![CDATA[pjZ8Ytx-nwvpCRyQneH3Ncmh6N94]]></CardId>
    <UserCardCode><![CDATA[485027611252]]></UserCardCode>
    <ModifyBonus>3</ModifyBonus>
    <ModifyBalance>0</ModifyBalance>
</xml>
*/
type EventUpdateMemberCard struct {
	Event
	CardId        string `xml:"CardId"`
	UserCardCode  string `xml:"UserCardCode"`
	ModifyBonus   string `xml:"ModifyBonus"`
	ModifyBalance string `xml:"ModifyBalance"`
}

/*
<xml>
   <ToUserName><![CDATA[gh_2d62d*****0]]></ToUserName>
    <FromUserName><![CDATA[oa3LFuBvWb7*********]]></FromUserName>
    <CreateTime>1443838506</CreateTime>
    <MsgType><![CDATA[event]]></MsgType>
    <Event><![CDATA[card_sku_remind]]></Event>
    <CardId><![CDATA[pa3LFuAh2P65**********]]></CardId>
    <Detail><![CDATA[the card's quantity is equal to 0]]></Detail>
</xml>
*/
type EventCardSkuRemind struct {
	Event
	CardId string `xml:"CardId"`
	Detail string `xml:"Detail"`
}

/*
<xml>
   <ToUserName><![CDATA[gh_7223c83d4be5]]></ToUserName>
    <FromUserName><![CDATA[ob5E7s-HoN9tslQY3-0I4qmgluHk]]></FromUserName>
    <CreateTime>1453295737</CreateTime>
    <MsgType><![CDATA[event]]></MsgType>
    <Event><![CDATA[card_pay_order]]></Event>
    <OrderId><![CDATA[404091456]]></OrderId>
    <Status><![CDATA[ORDER_STATUS_FINANCE_SUCC]]></Status>
    <CreateOrderTime>1453295737</CreateOrderTime>
    <PayFinishTime>0</PayFinishTime>
    <Desc><![CDATA[]]></Desc>
    <FreeCoinCount><![CDATA[200]]></FreeCoinCount>
    <PayCoinCount><![CDATA[0]]></PayCoinCount>
    <RefundFreeCoinCount><![CDATA[0]]></RefundFreeCoinCount>
    <RefundPayCoinCount><![CDATA[0]]></RefundPayCoinCount>
    <OrderType><![CDATA[ORDER_TYPE_SYS_ADD]]></OrderType>
    <Memo><![CDATA[开通账户奖励]]></Memo>
    <ReceiptInfo><![CDATA[]]></ReceiptInfo>
</xml>
*/
type EventCardPayOrder struct {
	Event
	OrderId             string `xml:"OrderId"`
	Status              string `xml:"Status"`
	CreateOrderTime     string `xml:"CreateOrderTime"`
	PayFinishTime       string `xml:"PayFinishTime"`
	Desc                string `xml:"Desc"`
	FreeCoinCount       string `xml:"FreeCoinCount"`
	PayCoinCount        string `xml:"PayCoinCount"`
	RefundFreeCoinCount string `xml:"RefundFreeCoinCount"`
	RefundPayCoinCount  string `xml:"RefundPayCoinCount"`
	OrderType           string `xml:"OrderType"`
	Memo                string `xml:"Memo"`
	ReceiptInfo         string `xml:"ReceiptInfo"`
}

/*
<xml>
   <ToUserName> <![CDATA[gh_3fcea188bf78]]></ToUserName>
    <FromUserName><![CDATA[obLatjlaNQKb8FqOvt1M1x1lIBFE]]></FromUserName>
    <CreateTime>1432668700</CreateTime>
    <MsgType><![CDATA[event]]></MsgType>
    <Event><![CDATA[submit_membercard_user_info]]></Event>
    <CardId><![CDATA[pbLatjtZ7v1BG_ZnTjbW85GYc_E8]]></CardId>
    <UserCardCode><![CDATA[018255396048]]></UserCardCode>
</xml>
*/
type EventSubmitMembercardUserInfo struct {
	Event
	CardId       string `xml:"CardId"`
	UserCardCode string `xml:"UserCardCode"`
}
