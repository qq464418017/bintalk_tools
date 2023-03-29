package bintalk

import "bintalk"

type Cs_pt_session struct{
	Channel uint8
	Login_type uint8
	Device_id string
	User_id string
	Token string
	Valid_ts uint32
}

func (t *Cs_pt_session) Serialize()[]byte{
	pw:= btk.ProtocolWriter{Bytes:make([]byte,0)}
	pw.WriteData(t.Channel)
	pw.WriteData(t.Login_type)
	pw.WriteData(t.Device_id)
	pw.WriteData(t.User_id)
	pw.WriteData(t.Token)
	pw.WriteData(t.Valid_ts)
	return pw.Bytes
}

func (t *Cs_pt_session) DeserializeStruct(buf []byte, startid *int){
	btk.Read(buf,startid,&t.Channel)
	btk.Read(buf,startid,&t.Login_type)
	btk.Read(buf,startid,&t.Device_id)
	btk.Read(buf,startid,&t.User_id)
	btk.Read(buf,startid,&t.Token)
	btk.Read(buf,startid,&t.Valid_ts)
}

func (t *Cs_pt_session) Deserialize(buf []byte){
	startid := 0
	t.DeserializeStruct(buf, &startid)
}


type Cs_pt_login struct{
	Session string
}

func (t *Cs_pt_login) Serialize()[]byte{
	pw:= btk.ProtocolWriter{Bytes:make([]byte,0)}
	pw.WriteData(t.Session)
	return pw.Bytes
}

func (t *Cs_pt_login) DeserializeStruct(buf []byte, startid *int){
	btk.Read(buf,startid,&t.Session)
}

func (t *Cs_pt_login) Deserialize(buf []byte){
	startid := 0
	t.DeserializeStruct(buf, &startid)
}


type Cs_pt_order struct{
	Guid uint32
	Order_id string
}

func (t *Cs_pt_order) Serialize()[]byte{
	pw:= btk.ProtocolWriter{Bytes:make([]byte,0)}
	pw.WriteData(t.Guid)
	pw.WriteData(t.Order_id)
	return pw.Bytes
}

func (t *Cs_pt_order) DeserializeStruct(buf []byte, startid *int){
	btk.Read(buf,startid,&t.Guid)
	btk.Read(buf,startid,&t.Order_id)
}

func (t *Cs_pt_order) Deserialize(buf []byte){
	startid := 0
	t.DeserializeStruct(buf, &startid)
}


type Cs_pt_order_verify struct{
	Guid uint32
	Token string
}

func (t *Cs_pt_order_verify) Serialize()[]byte{
	pw:= btk.ProtocolWriter{Bytes:make([]byte,0)}
	pw.WriteData(t.Guid)
	pw.WriteData(t.Token)
	return pw.Bytes
}

func (t *Cs_pt_order_verify) DeserializeStruct(buf []byte, startid *int){
	btk.Read(buf,startid,&t.Guid)
	btk.Read(buf,startid,&t.Token)
}

func (t *Cs_pt_order_verify) Deserialize(buf []byte){
	startid := 0
	t.DeserializeStruct(buf, &startid)
}


type Cs_pt_mail struct{
	Guid uint32
	Create_ts uint32
	Is_read bool
	Read_ts uint32
	Is_get bool
	Title string
	Content string
	Type uint8
	Gold_num uint32
}

func (t *Cs_pt_mail) Serialize()[]byte{
	pw:= btk.ProtocolWriter{Bytes:make([]byte,0)}
	pw.WriteData(t.Guid)
	pw.WriteData(t.Create_ts)
	pw.WriteData(t.Is_read)
	pw.WriteData(t.Read_ts)
	pw.WriteData(t.Is_get)
	pw.WriteData(t.Title)
	pw.WriteData(t.Content)
	pw.WriteData(t.Type)
	pw.WriteData(t.Gold_num)
	return pw.Bytes
}

func (t *Cs_pt_mail) DeserializeStruct(buf []byte, startid *int){
	btk.Read(buf,startid,&t.Guid)
	btk.Read(buf,startid,&t.Create_ts)
	btk.Read(buf,startid,&t.Is_read)
	btk.Read(buf,startid,&t.Read_ts)
	btk.Read(buf,startid,&t.Is_get)
	btk.Read(buf,startid,&t.Title)
	btk.Read(buf,startid,&t.Content)
	btk.Read(buf,startid,&t.Type)
	btk.Read(buf,startid,&t.Gold_num)
}

func (t *Cs_pt_mail) Deserialize(buf []byte){
	startid := 0
	t.DeserializeStruct(buf, &startid)
}


type Cs_pt_mails struct{
	Mails []*Cs_pt_mail
}

func (t *Cs_pt_mails) Serialize()[]byte{
	pw:= btk.ProtocolWriter{Bytes:make([]byte,0)}
	pw.WriteDynSize(len(t.Mails))
	for _, v := range t.Mails {
		pw.Write(v.Serialize())
	}
	return pw.Bytes
}

func (t *Cs_pt_mails) DeserializeStruct(buf []byte, startid *int){
	var s0 int
	btk.ReadDynSize(buf,startid,&s0)
	if s0 < len(buf) {
		t.Mails = make([]*Cs_pt_mail,s0)
		for i := 0; i < s0; i++ {
			a := new(Cs_pt_mail)
			a.DeserializeStruct(buf,startid)
			t.Mails[i] = a
		}
	}
}

func (t *Cs_pt_mails) Deserialize(buf []byte){
	startid := 0
	t.DeserializeStruct(buf, &startid)
}


type Cs_pt_game_data struct{
	Type uint8
	Win_times uint32
	Lose_times uint32
	High_win_streak uint32
}

func (t *Cs_pt_game_data) Serialize()[]byte{
	pw:= btk.ProtocolWriter{Bytes:make([]byte,0)}
	pw.WriteData(t.Type)
	pw.WriteData(t.Win_times)
	pw.WriteData(t.Lose_times)
	pw.WriteData(t.High_win_streak)
	return pw.Bytes
}

func (t *Cs_pt_game_data) DeserializeStruct(buf []byte, startid *int){
	btk.Read(buf,startid,&t.Type)
	btk.Read(buf,startid,&t.Win_times)
	btk.Read(buf,startid,&t.Lose_times)
	btk.Read(buf,startid,&t.High_win_streak)
}

func (t *Cs_pt_game_data) Deserialize(buf []byte){
	startid := 0
	t.DeserializeStruct(buf, &startid)
}


type Cs_pt_game_info struct{
	Max_gold uint32
	Max_win_gold uint32
	Data_list []*Cs_pt_game_data
}

func (t *Cs_pt_game_info) Serialize()[]byte{
	pw:= btk.ProtocolWriter{Bytes:make([]byte,0)}
	pw.WriteData(t.Max_gold)
	pw.WriteData(t.Max_win_gold)
	pw.WriteDynSize(len(t.Data_list))
	for _, v := range t.Data_list {
		pw.Write(v.Serialize())
	}
	return pw.Bytes
}

func (t *Cs_pt_game_info) DeserializeStruct(buf []byte, startid *int){
	btk.Read(buf,startid,&t.Max_gold)
	btk.Read(buf,startid,&t.Max_win_gold)
	var s2 int
	btk.ReadDynSize(buf,startid,&s2)
	if s2 < len(buf) {
		t.Data_list = make([]*Cs_pt_game_data,s2)
		for i := 0; i < s2; i++ {
			a := new(Cs_pt_game_data)
			a.DeserializeStruct(buf,startid)
			t.Data_list[i] = a
		}
	}
}

func (t *Cs_pt_game_info) Deserialize(buf []byte){
	startid := 0
	t.DeserializeStruct(buf, &startid)
}


type Cs_pt_game_bonus struct{
	Guid uint32
	Player Com_player_info
	Num uint32
	Ts uint32
}

func (t *Cs_pt_game_bonus) Serialize()[]byte{
	pw:= btk.ProtocolWriter{Bytes:make([]byte,0)}
	pw.WriteData(t.Guid)
	pw.Write(t.Player.Serialize())
	pw.WriteData(t.Num)
	pw.WriteData(t.Ts)
	return pw.Bytes
}

func (t *Cs_pt_game_bonus) DeserializeStruct(buf []byte, startid *int){
	btk.Read(buf,startid,&t.Guid)
	t.Player.DeserializeStruct(buf,startid)
	btk.Read(buf,startid,&t.Num)
	btk.Read(buf,startid,&t.Ts)
}

func (t *Cs_pt_game_bonus) Deserialize(buf []byte){
	startid := 0
	t.DeserializeStruct(buf, &startid)
}


type Cs_pt_game_bonuses struct{
	List []*Cs_pt_game_bonus
}

func (t *Cs_pt_game_bonuses) Serialize()[]byte{
	pw:= btk.ProtocolWriter{Bytes:make([]byte,0)}
	pw.WriteDynSize(len(t.List))
	for _, v := range t.List {
		pw.Write(v.Serialize())
	}
	return pw.Bytes
}

func (t *Cs_pt_game_bonuses) DeserializeStruct(buf []byte, startid *int){
	var s0 int
	btk.ReadDynSize(buf,startid,&s0)
	if s0 < len(buf) {
		t.List = make([]*Cs_pt_game_bonus,s0)
		for i := 0; i < s0; i++ {
			a := new(Cs_pt_game_bonus)
			a.DeserializeStruct(buf,startid)
			t.List[i] = a
		}
	}
}

func (t *Cs_pt_game_bonuses) Deserialize(buf []byte){
	startid := 0
	t.DeserializeStruct(buf, &startid)
}


type Cs_pt_player_state struct{
	Player_info Com_player_info
	Room_info Com_uint8_double
	Gold uint64
	Mails []*Cs_pt_mail
	Notices []*Com_notice
	Game_bonuses []*Cs_pt_game_bonus
	Emoji_id_list []uint8
	Emoji_id uint8
	Icon_frame_list []uint8
	Ts uint32
}

func (t *Cs_pt_player_state) Serialize()[]byte{
	pw:= btk.ProtocolWriter{Bytes:make([]byte,0)}
	pw.Write(t.Player_info.Serialize())
	pw.Write(t.Room_info.Serialize())
	pw.WriteData(t.Gold)
	pw.WriteDynSize(len(t.Mails))
	for _, v := range t.Mails {
		pw.Write(v.Serialize())
	}
	pw.WriteDynSize(len(t.Notices))
	for _, v := range t.Notices {
		pw.Write(v.Serialize())
	}
	pw.WriteDynSize(len(t.Game_bonuses))
	for _, v := range t.Game_bonuses {
		pw.Write(v.Serialize())
	}
	pw.WriteData(t.Emoji_id_list)
	pw.WriteData(t.Emoji_id)
	pw.WriteData(t.Icon_frame_list)
	pw.WriteData(t.Ts)
	return pw.Bytes
}

func (t *Cs_pt_player_state) DeserializeStruct(buf []byte, startid *int){
	t.Player_info.DeserializeStruct(buf,startid)
	t.Room_info.DeserializeStruct(buf,startid)
	btk.Read(buf,startid,&t.Gold)
	var s3 int
	btk.ReadDynSize(buf,startid,&s3)
	if s3 < len(buf) {
		t.Mails = make([]*Cs_pt_mail,s3)
		for i := 0; i < s3; i++ {
			a := new(Cs_pt_mail)
			a.DeserializeStruct(buf,startid)
			t.Mails[i] = a
		}
	}
	var s4 int
	btk.ReadDynSize(buf,startid,&s4)
	if s4 < len(buf) {
		t.Notices = make([]*Com_notice,s4)
		for i := 0; i < s4; i++ {
			a := new(Com_notice)
			a.DeserializeStruct(buf,startid)
			t.Notices[i] = a
		}
	}
	var s5 int
	btk.ReadDynSize(buf,startid,&s5)
	if s5 < len(buf) {
		t.Game_bonuses = make([]*Cs_pt_game_bonus,s5)
		for i := 0; i < s5; i++ {
			a := new(Cs_pt_game_bonus)
			a.DeserializeStruct(buf,startid)
			t.Game_bonuses[i] = a
		}
	}
	btk.Read(buf,startid,&t.Emoji_id_list)
	btk.Read(buf,startid,&t.Emoji_id)
	btk.Read(buf,startid,&t.Icon_frame_list)
	btk.Read(buf,startid,&t.Ts)
}

func (t *Cs_pt_player_state) Deserialize(buf []byte){
	startid := 0
	t.DeserializeStruct(buf, &startid)
}


const(
	Cs_pt_request_type_heartbeat = iota
	Cs_pt_request_type_login
	Cs_pt_request_type_mail_read
	Cs_pt_request_type_mail_delete
	Cs_pt_request_type_mail_reward
	Cs_pt_request_type_get_popup_notice
	Cs_pt_request_type_change_icon
	Cs_pt_request_type_change_name
	Cs_pt_request_type_change_gender
	Cs_pt_request_type_icon_frame_buy
	Cs_pt_request_type_emoji_change
	Cs_pt_request_type_emoji_buy
	Cs_pt_request_type_bonus_get
	Cs_pt_request_type_bonus_reward
	Cs_pt_request_type_bonus_get_list
	Cs_pt_request_type_get_game_info
	Cs_pt_request_type_get_customer_service
	Cs_pt_request_type_order_request
	Cs_pt_request_type_order_verify
)

const(
	Cs_pt_push_type_error_push = iota
	Cs_pt_push_type_player_info_update
	Cs_pt_push_type_gold_update
	Cs_pt_push_type_mail_add
	Cs_pt_push_type_mail_update
	Cs_pt_push_type_mail_delete
	Cs_pt_push_type_notice_update
	Cs_pt_push_type_notice_delete
	Cs_pt_push_type_bonus_add
	Cs_pt_push_type_bonus_update
)

const(
)

