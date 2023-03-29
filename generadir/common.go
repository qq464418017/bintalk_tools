package bintalk

import "bintalk"

const(
	Com_parent_protocol_type_platform = iota
	Com_parent_protocol_type_platform_push
	Com_parent_protocol_type_teen_patti
	Com_parent_protocol_type_teen_patti_push
	Com_parent_protocol_type_rummy
	Com_parent_protocol_type_rummy_push
	Com_parent_protocol_type_andar_bahar
	Com_parent_protocol_type_andar_bahar_push
	Com_parent_protocol_type_poker
	Com_parent_protocol_type_poker_push
)

const(
	Com_error_type_ok = iota
	Com_error_type_error
	Com_error_type_other_login_in
	Com_error_type_not_enough_gold
	Com_error_type_lock_player
	Com_error_type_order_invalid
	Com_error_type_order_failed
	Com_error_type_room_not_found
	Com_error_type_room_number_max
	Com_error_type_game_maintained
	Com_error_type_change_room_error
	Com_error_type_create_friend_room_failed
	Com_error_type_friend_room_max
	Com_error_type_invalid_friend_room
	Com_error_type_already_in_room
	Com_error_type_gold_limit
	Com_error_type_invalid_player
	Com_error_type_invalid_bonus
)

const(
	Com_channel_type_none = iota
	Com_channel_type_android
	Com_channel_type_ios
)

const(
	Com_login_type_google = iota
	Com_login_type_facebook
	Com_login_type_tourist
)

const(
	Com_gender_type_unknown = iota
	Com_gender_type_male
	Com_gender_type_female
)

const(
	Com_order_type_none = iota
	Com_order_type_google
	Com_order_type_rzp
)

const(
	Com_order_state_create = iota
	Com_order_state_verify
	Com_order_state_success
	Com_order_state_fail
)

const(
	Com_mail_type_admin = iota
	Com_mail_type_player
)

const(
	Com_leave_room_state_none = iota
	Com_leave_room_state_not_money
	Com_leave_room_state_not_operation
	Com_leave_room_state_maintain
)

const(
	Com_room_compare_type_none = iota
	Com_room_compare_type_win
	Com_room_compare_type_lose
	Com_room_compare_type_draw
)

type Com_uint8 struct{
	Value uint8
}

func (t *Com_uint8) Serialize()[]byte{
	pw:= btk.ProtocolWriter{Bytes:make([]byte,0)}
	pw.WriteData(t.Value)
	return pw.Bytes
}

func (t *Com_uint8) DeserializeStruct(buf []byte, startid *int){
	btk.Read(buf,startid,&t.Value)
}

func (t *Com_uint8) Deserialize(buf []byte){
	startid := 0
	t.DeserializeStruct(buf, &startid)
}


type Com_uint8_double struct{
	Value1 uint8
	Value2 uint8
}

func (t *Com_uint8_double) Serialize()[]byte{
	pw:= btk.ProtocolWriter{Bytes:make([]byte,0)}
	pw.WriteData(t.Value1)
	pw.WriteData(t.Value2)
	return pw.Bytes
}

func (t *Com_uint8_double) DeserializeStruct(buf []byte, startid *int){
	btk.Read(buf,startid,&t.Value1)
	btk.Read(buf,startid,&t.Value2)
}

func (t *Com_uint8_double) Deserialize(buf []byte){
	startid := 0
	t.DeserializeStruct(buf, &startid)
}


type Com_uint8_list struct{
	Values []uint8
}

func (t *Com_uint8_list) Serialize()[]byte{
	pw:= btk.ProtocolWriter{Bytes:make([]byte,0)}
	pw.WriteData(t.Values)
	return pw.Bytes
}

func (t *Com_uint8_list) DeserializeStruct(buf []byte, startid *int){
	btk.Read(buf,startid,&t.Values)
}

func (t *Com_uint8_list) Deserialize(buf []byte){
	startid := 0
	t.DeserializeStruct(buf, &startid)
}


type Com_uint16 struct{
	Value uint16
}

func (t *Com_uint16) Serialize()[]byte{
	pw:= btk.ProtocolWriter{Bytes:make([]byte,0)}
	pw.WriteData(t.Value)
	return pw.Bytes
}

func (t *Com_uint16) DeserializeStruct(buf []byte, startid *int){
	btk.Read(buf,startid,&t.Value)
}

func (t *Com_uint16) Deserialize(buf []byte){
	startid := 0
	t.DeserializeStruct(buf, &startid)
}


type Com_uint32 struct{
	Value uint32
}

func (t *Com_uint32) Serialize()[]byte{
	pw:= btk.ProtocolWriter{Bytes:make([]byte,0)}
	pw.WriteData(t.Value)
	return pw.Bytes
}

func (t *Com_uint32) DeserializeStruct(buf []byte, startid *int){
	btk.Read(buf,startid,&t.Value)
}

func (t *Com_uint32) Deserialize(buf []byte){
	startid := 0
	t.DeserializeStruct(buf, &startid)
}


type Com_uint32_double struct{
	Value1 uint32
	Value2 uint32
}

func (t *Com_uint32_double) Serialize()[]byte{
	pw:= btk.ProtocolWriter{Bytes:make([]byte,0)}
	pw.WriteData(t.Value1)
	pw.WriteData(t.Value2)
	return pw.Bytes
}

func (t *Com_uint32_double) DeserializeStruct(buf []byte, startid *int){
	btk.Read(buf,startid,&t.Value1)
	btk.Read(buf,startid,&t.Value2)
}

func (t *Com_uint32_double) Deserialize(buf []byte){
	startid := 0
	t.DeserializeStruct(buf, &startid)
}


type Com_uint64 struct{
	Value uint64
}

func (t *Com_uint64) Serialize()[]byte{
	pw:= btk.ProtocolWriter{Bytes:make([]byte,0)}
	pw.WriteData(t.Value)
	return pw.Bytes
}

func (t *Com_uint64) DeserializeStruct(buf []byte, startid *int){
	btk.Read(buf,startid,&t.Value)
}

func (t *Com_uint64) Deserialize(buf []byte){
	startid := 0
	t.DeserializeStruct(buf, &startid)
}


type Com_uint64_double struct{
	Value1 uint64
	Value2 uint64
}

func (t *Com_uint64_double) Serialize()[]byte{
	pw:= btk.ProtocolWriter{Bytes:make([]byte,0)}
	pw.WriteData(t.Value1)
	pw.WriteData(t.Value2)
	return pw.Bytes
}

func (t *Com_uint64_double) DeserializeStruct(buf []byte, startid *int){
	btk.Read(buf,startid,&t.Value1)
	btk.Read(buf,startid,&t.Value2)
}

func (t *Com_uint64_double) Deserialize(buf []byte){
	startid := 0
	t.DeserializeStruct(buf, &startid)
}


type Com_uint32_list struct{
	Values []uint32
}

func (t *Com_uint32_list) Serialize()[]byte{
	pw:= btk.ProtocolWriter{Bytes:make([]byte,0)}
	pw.WriteData(t.Values)
	return pw.Bytes
}

func (t *Com_uint32_list) DeserializeStruct(buf []byte, startid *int){
	btk.Read(buf,startid,&t.Values)
}

func (t *Com_uint32_list) Deserialize(buf []byte){
	startid := 0
	t.DeserializeStruct(buf, &startid)
}


type Com_bool struct{
	Value bool
}

func (t *Com_bool) Serialize()[]byte{
	pw:= btk.ProtocolWriter{Bytes:make([]byte,0)}
	pw.WriteData(t.Value)
	return pw.Bytes
}

func (t *Com_bool) DeserializeStruct(buf []byte, startid *int){
	btk.Read(buf,startid,&t.Value)
}

func (t *Com_bool) Deserialize(buf []byte){
	startid := 0
	t.DeserializeStruct(buf, &startid)
}


type Com_string struct{
	Value string
}

func (t *Com_string) Serialize()[]byte{
	pw:= btk.ProtocolWriter{Bytes:make([]byte,0)}
	pw.WriteData(t.Value)
	return pw.Bytes
}

func (t *Com_string) DeserializeStruct(buf []byte, startid *int){
	btk.Read(buf,startid,&t.Value)
}

func (t *Com_string) Deserialize(buf []byte){
	startid := 0
	t.DeserializeStruct(buf, &startid)
}


type Com_timer struct{
	Ts uint32
	Dur uint32
}

func (t *Com_timer) Serialize()[]byte{
	pw:= btk.ProtocolWriter{Bytes:make([]byte,0)}
	pw.WriteData(t.Ts)
	pw.WriteData(t.Dur)
	return pw.Bytes
}

func (t *Com_timer) DeserializeStruct(buf []byte, startid *int){
	btk.Read(buf,startid,&t.Ts)
	btk.Read(buf,startid,&t.Dur)
}

func (t *Com_timer) Deserialize(buf []byte){
	startid := 0
	t.DeserializeStruct(buf, &startid)
}


type Com_player_info struct{
	Id uint32
	Name string
	Icon uint8
	Platform_icon string
	Icon_frame uint8
	Gender_type uint8
}

func (t *Com_player_info) Serialize()[]byte{
	pw:= btk.ProtocolWriter{Bytes:make([]byte,0)}
	pw.WriteData(t.Id)
	pw.WriteData(t.Name)
	pw.WriteData(t.Icon)
	pw.WriteData(t.Platform_icon)
	pw.WriteData(t.Icon_frame)
	pw.WriteData(t.Gender_type)
	return pw.Bytes
}

func (t *Com_player_info) DeserializeStruct(buf []byte, startid *int){
	btk.Read(buf,startid,&t.Id)
	btk.Read(buf,startid,&t.Name)
	btk.Read(buf,startid,&t.Icon)
	btk.Read(buf,startid,&t.Platform_icon)
	btk.Read(buf,startid,&t.Icon_frame)
	btk.Read(buf,startid,&t.Gender_type)
}

func (t *Com_player_info) Deserialize(buf []byte){
	startid := 0
	t.DeserializeStruct(buf, &startid)
}


type Com_error struct{
	Type uint8
}

func (t *Com_error) Serialize()[]byte{
	pw:= btk.ProtocolWriter{Bytes:make([]byte,0)}
	pw.WriteData(t.Type)
	return pw.Bytes
}

func (t *Com_error) DeserializeStruct(buf []byte, startid *int){
	btk.Read(buf,startid,&t.Type)
}

func (t *Com_error) Deserialize(buf []byte){
	startid := 0
	t.DeserializeStruct(buf, &startid)
}


type Com_notice struct{
	Id uint32
	Start_ts uint32
	End_ts uint32
	Interval uint32
	Play_pos uint8
	Content string
	Is_insert bool
	Times uint8
	Player_name string
	Gold_num uint32
}

func (t *Com_notice) Serialize()[]byte{
	pw:= btk.ProtocolWriter{Bytes:make([]byte,0)}
	pw.WriteData(t.Id)
	pw.WriteData(t.Start_ts)
	pw.WriteData(t.End_ts)
	pw.WriteData(t.Interval)
	pw.WriteData(t.Play_pos)
	pw.WriteData(t.Content)
	pw.WriteData(t.Is_insert)
	pw.WriteData(t.Times)
	pw.WriteData(t.Player_name)
	pw.WriteData(t.Gold_num)
	return pw.Bytes
}

func (t *Com_notice) DeserializeStruct(buf []byte, startid *int){
	btk.Read(buf,startid,&t.Id)
	btk.Read(buf,startid,&t.Start_ts)
	btk.Read(buf,startid,&t.End_ts)
	btk.Read(buf,startid,&t.Interval)
	btk.Read(buf,startid,&t.Play_pos)
	btk.Read(buf,startid,&t.Content)
	btk.Read(buf,startid,&t.Is_insert)
	btk.Read(buf,startid,&t.Times)
	btk.Read(buf,startid,&t.Player_name)
	btk.Read(buf,startid,&t.Gold_num)
}

func (t *Com_notice) Deserialize(buf []byte){
	startid := 0
	t.DeserializeStruct(buf, &startid)
}


type Com_popup_notice struct{
	Id uint16
	Title string
	Content string
	Url string
	Open_interface string
	Sort uint8
}

func (t *Com_popup_notice) Serialize()[]byte{
	pw:= btk.ProtocolWriter{Bytes:make([]byte,0)}
	pw.WriteData(t.Id)
	pw.WriteData(t.Title)
	pw.WriteData(t.Content)
	pw.WriteData(t.Url)
	pw.WriteData(t.Open_interface)
	pw.WriteData(t.Sort)
	return pw.Bytes
}

func (t *Com_popup_notice) DeserializeStruct(buf []byte, startid *int){
	btk.Read(buf,startid,&t.Id)
	btk.Read(buf,startid,&t.Title)
	btk.Read(buf,startid,&t.Content)
	btk.Read(buf,startid,&t.Url)
	btk.Read(buf,startid,&t.Open_interface)
	btk.Read(buf,startid,&t.Sort)
}

func (t *Com_popup_notice) Deserialize(buf []byte){
	startid := 0
	t.DeserializeStruct(buf, &startid)
}


type Com_popup_notice_list struct{
	Notices []*Com_popup_notice
}

func (t *Com_popup_notice_list) Serialize()[]byte{
	pw:= btk.ProtocolWriter{Bytes:make([]byte,0)}
	pw.WriteDynSize(len(t.Notices))
	for _, v := range t.Notices {
		pw.Write(v.Serialize())
	}
	return pw.Bytes
}

func (t *Com_popup_notice_list) DeserializeStruct(buf []byte, startid *int){
	var s0 int
	btk.ReadDynSize(buf,startid,&s0)
	if s0 < len(buf) {
		t.Notices = make([]*Com_popup_notice,s0)
		for i := 0; i < s0; i++ {
			a := new(Com_popup_notice)
			a.DeserializeStruct(buf,startid)
			t.Notices[i] = a
		}
	}
}

func (t *Com_popup_notice_list) Deserialize(buf []byte){
	startid := 0
	t.DeserializeStruct(buf, &startid)
}


type Com_card struct{
	Num uint8
	Color uint8
}

func (t *Com_card) Serialize()[]byte{
	pw:= btk.ProtocolWriter{Bytes:make([]byte,0)}
	pw.WriteData(t.Num)
	pw.WriteData(t.Color)
	return pw.Bytes
}

func (t *Com_card) DeserializeStruct(buf []byte, startid *int){
	btk.Read(buf,startid,&t.Num)
	btk.Read(buf,startid,&t.Color)
}

func (t *Com_card) Deserialize(buf []byte){
	startid := 0
	t.DeserializeStruct(buf, &startid)
}


type Com_emoticon struct{
	Group_id uint8
	Emoji_id uint8
	Player_id uint32
	To_player_id uint32
}

func (t *Com_emoticon) Serialize()[]byte{
	pw:= btk.ProtocolWriter{Bytes:make([]byte,0)}
	pw.WriteData(t.Group_id)
	pw.WriteData(t.Emoji_id)
	pw.WriteData(t.Player_id)
	pw.WriteData(t.To_player_id)
	return pw.Bytes
}

func (t *Com_emoticon) DeserializeStruct(buf []byte, startid *int){
	btk.Read(buf,startid,&t.Group_id)
	btk.Read(buf,startid,&t.Emoji_id)
	btk.Read(buf,startid,&t.Player_id)
	btk.Read(buf,startid,&t.To_player_id)
}

func (t *Com_emoticon) Deserialize(buf []byte){
	startid := 0
	t.DeserializeStruct(buf, &startid)
}


const(
	Com_player_room_state_none = iota
	Com_player_room_state_leave
	Com_player_room_state_change
	Com_player_room_state_watch
)

type Com_update_player_room_state struct{
	Player_id uint32
	State uint8
}

func (t *Com_update_player_room_state) Serialize()[]byte{
	pw:= btk.ProtocolWriter{Bytes:make([]byte,0)}
	pw.WriteData(t.Player_id)
	pw.WriteData(t.State)
	return pw.Bytes
}

func (t *Com_update_player_room_state) DeserializeStruct(buf []byte, startid *int){
	btk.Read(buf,startid,&t.Player_id)
	btk.Read(buf,startid,&t.State)
}

func (t *Com_update_player_room_state) Deserialize(buf []byte){
	startid := 0
	t.DeserializeStruct(buf, &startid)
}


type Com_join_friend_room struct{
	Error uint8
	Num uint32
}

func (t *Com_join_friend_room) Serialize()[]byte{
	pw:= btk.ProtocolWriter{Bytes:make([]byte,0)}
	pw.WriteData(t.Error)
	pw.WriteData(t.Num)
	return pw.Bytes
}

func (t *Com_join_friend_room) DeserializeStruct(buf []byte, startid *int){
	btk.Read(buf,startid,&t.Error)
	btk.Read(buf,startid,&t.Num)
}

func (t *Com_join_friend_room) Deserialize(buf []byte){
	startid := 0
	t.DeserializeStruct(buf, &startid)
}


