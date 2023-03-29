package bintalk

import "bintalk"

type C_animal_help struct{
	Id uint8
	Title string
	Content string
	Bet_id uint8
	Result_id uint8
}

func (t *C_animal_help) Serialize()[]byte{
	pw:= btk.ProtocolWriter{Bytes:make([]byte,0)}
	pw.WriteData(t.Id)
	pw.WriteData(t.Title)
	pw.WriteData(t.Content)
	pw.WriteData(t.Bet_id)
	pw.WriteData(t.Result_id)
	return pw.Bytes
}

func (t *C_animal_help) DeserializeStruct(buf []byte, startid *int){
	btk.Read(buf,startid,&t.Id)
	btk.Read(buf,startid,&t.Title)
	btk.Read(buf,startid,&t.Content)
	btk.Read(buf,startid,&t.Bet_id)
	btk.Read(buf,startid,&t.Result_id)
}

func (t *C_animal_help) Deserialize(buf []byte){
	startid := 0
	t.DeserializeStruct(buf, &startid)
}


type C_animal_help_table struct{
	Record []*C_animal_help
}

func (t *C_animal_help_table) Serialize()[]byte{
	pw:= btk.ProtocolWriter{Bytes:make([]byte,0)}
	pw.WriteDynSize(len(t.Record))
	for _, v := range t.Record {
		pw.Write(v.Serialize())
	}
	return pw.Bytes
}

func (t *C_animal_help_table) DeserializeStruct(buf []byte, startid *int){
	var s0 int
	btk.ReadDynSize(buf,startid,&s0)
	if s0 < len(buf) {
		t.Record = make([]*C_animal_help,s0)
		for i := 0; i < s0; i++ {
			a := new(C_animal_help)
			a.DeserializeStruct(buf,startid)
			t.Record[i] = a
		}
	}
}

func (t *C_animal_help_table) Deserialize(buf []byte){
	startid := 0
	t.DeserializeStruct(buf, &startid)
}


type C_emoji struct{
	Id uint8
	Name string
	Sort uint8
}

func (t *C_emoji) Serialize()[]byte{
	pw:= btk.ProtocolWriter{Bytes:make([]byte,0)}
	pw.WriteData(t.Id)
	pw.WriteData(t.Name)
	pw.WriteData(t.Sort)
	return pw.Bytes
}

func (t *C_emoji) DeserializeStruct(buf []byte, startid *int){
	btk.Read(buf,startid,&t.Id)
	btk.Read(buf,startid,&t.Name)
	btk.Read(buf,startid,&t.Sort)
}

func (t *C_emoji) Deserialize(buf []byte){
	startid := 0
	t.DeserializeStruct(buf, &startid)
}


type C_emoji_table struct{
	Record []*C_emoji
}

func (t *C_emoji_table) Serialize()[]byte{
	pw:= btk.ProtocolWriter{Bytes:make([]byte,0)}
	pw.WriteDynSize(len(t.Record))
	for _, v := range t.Record {
		pw.Write(v.Serialize())
	}
	return pw.Bytes
}

func (t *C_emoji_table) DeserializeStruct(buf []byte, startid *int){
	var s0 int
	btk.ReadDynSize(buf,startid,&s0)
	if s0 < len(buf) {
		t.Record = make([]*C_emoji,s0)
		for i := 0; i < s0; i++ {
			a := new(C_emoji)
			a.DeserializeStruct(buf,startid)
			t.Record[i] = a
		}
	}
}

func (t *C_emoji_table) Deserialize(buf []byte){
	startid := 0
	t.DeserializeStruct(buf, &startid)
}


type C_fix_text struct{
	Id string
	Chinese string
	Vietnam string
	Korea string
}

func (t *C_fix_text) Serialize()[]byte{
	pw:= btk.ProtocolWriter{Bytes:make([]byte,0)}
	pw.WriteData(t.Id)
	pw.WriteData(t.Chinese)
	pw.WriteData(t.Vietnam)
	pw.WriteData(t.Korea)
	return pw.Bytes
}

func (t *C_fix_text) DeserializeStruct(buf []byte, startid *int){
	btk.Read(buf,startid,&t.Id)
	btk.Read(buf,startid,&t.Chinese)
	btk.Read(buf,startid,&t.Vietnam)
	btk.Read(buf,startid,&t.Korea)
}

func (t *C_fix_text) Deserialize(buf []byte){
	startid := 0
	t.DeserializeStruct(buf, &startid)
}


type C_fix_text_table struct{
	Record []*C_fix_text
}

func (t *C_fix_text_table) Serialize()[]byte{
	pw:= btk.ProtocolWriter{Bytes:make([]byte,0)}
	pw.WriteDynSize(len(t.Record))
	for _, v := range t.Record {
		pw.Write(v.Serialize())
	}
	return pw.Bytes
}

func (t *C_fix_text_table) DeserializeStruct(buf []byte, startid *int){
	var s0 int
	btk.ReadDynSize(buf,startid,&s0)
	if s0 < len(buf) {
		t.Record = make([]*C_fix_text,s0)
		for i := 0; i < s0; i++ {
			a := new(C_fix_text)
			a.DeserializeStruct(buf,startid)
			t.Record[i] = a
		}
	}
}

func (t *C_fix_text_table) Deserialize(buf []byte){
	startid := 0
	t.DeserializeStruct(buf, &startid)
}


type C_icon struct{
	Id uint8
	Name string
	Icon string
}

func (t *C_icon) Serialize()[]byte{
	pw:= btk.ProtocolWriter{Bytes:make([]byte,0)}
	pw.WriteData(t.Id)
	pw.WriteData(t.Name)
	pw.WriteData(t.Icon)
	return pw.Bytes
}

func (t *C_icon) DeserializeStruct(buf []byte, startid *int){
	btk.Read(buf,startid,&t.Id)
	btk.Read(buf,startid,&t.Name)
	btk.Read(buf,startid,&t.Icon)
}

func (t *C_icon) Deserialize(buf []byte){
	startid := 0
	t.DeserializeStruct(buf, &startid)
}


type C_icon_table struct{
	Record []*C_icon
}

func (t *C_icon_table) Serialize()[]byte{
	pw:= btk.ProtocolWriter{Bytes:make([]byte,0)}
	pw.WriteDynSize(len(t.Record))
	for _, v := range t.Record {
		pw.Write(v.Serialize())
	}
	return pw.Bytes
}

func (t *C_icon_table) DeserializeStruct(buf []byte, startid *int){
	var s0 int
	btk.ReadDynSize(buf,startid,&s0)
	if s0 < len(buf) {
		t.Record = make([]*C_icon,s0)
		for i := 0; i < s0; i++ {
			a := new(C_icon)
			a.DeserializeStruct(buf,startid)
			t.Record[i] = a
		}
	}
}

func (t *C_icon_table) Deserialize(buf []byte){
	startid := 0
	t.DeserializeStruct(buf, &startid)
}


type C_loc_text struct{
	Id string
	Chinese string
	Myanmar string
	Korea string
}

func (t *C_loc_text) Serialize()[]byte{
	pw:= btk.ProtocolWriter{Bytes:make([]byte,0)}
	pw.WriteData(t.Id)
	pw.WriteData(t.Chinese)
	pw.WriteData(t.Myanmar)
	pw.WriteData(t.Korea)
	return pw.Bytes
}

func (t *C_loc_text) DeserializeStruct(buf []byte, startid *int){
	btk.Read(buf,startid,&t.Id)
	btk.Read(buf,startid,&t.Chinese)
	btk.Read(buf,startid,&t.Myanmar)
	btk.Read(buf,startid,&t.Korea)
}

func (t *C_loc_text) Deserialize(buf []byte){
	startid := 0
	t.DeserializeStruct(buf, &startid)
}


type C_loc_text_table struct{
	Record []*C_loc_text
}

func (t *C_loc_text_table) Serialize()[]byte{
	pw:= btk.ProtocolWriter{Bytes:make([]byte,0)}
	pw.WriteDynSize(len(t.Record))
	for _, v := range t.Record {
		pw.Write(v.Serialize())
	}
	return pw.Bytes
}

func (t *C_loc_text_table) DeserializeStruct(buf []byte, startid *int){
	var s0 int
	btk.ReadDynSize(buf,startid,&s0)
	if s0 < len(buf) {
		t.Record = make([]*C_loc_text,s0)
		for i := 0; i < s0; i++ {
			a := new(C_loc_text)
			a.DeserializeStruct(buf,startid)
			t.Record[i] = a
		}
	}
}

func (t *C_loc_text_table) Deserialize(buf []byte){
	startid := 0
	t.DeserializeStruct(buf, &startid)
}


type G_animal_bet struct{
	Id uint8
	Name string
	Icon string
	Animal_type uint8
	Color_type uint8
}

func (t *G_animal_bet) Serialize()[]byte{
	pw:= btk.ProtocolWriter{Bytes:make([]byte,0)}
	pw.WriteData(t.Id)
	pw.WriteData(t.Name)
	pw.WriteData(t.Icon)
	pw.WriteData(t.Animal_type)
	pw.WriteData(t.Color_type)
	return pw.Bytes
}

func (t *G_animal_bet) DeserializeStruct(buf []byte, startid *int){
	btk.Read(buf,startid,&t.Id)
	btk.Read(buf,startid,&t.Name)
	btk.Read(buf,startid,&t.Icon)
	btk.Read(buf,startid,&t.Animal_type)
	btk.Read(buf,startid,&t.Color_type)
}

func (t *G_animal_bet) Deserialize(buf []byte){
	startid := 0
	t.DeserializeStruct(buf, &startid)
}


type G_animal_bet_table struct{
	Record []*G_animal_bet
}

func (t *G_animal_bet_table) Serialize()[]byte{
	pw:= btk.ProtocolWriter{Bytes:make([]byte,0)}
	pw.WriteDynSize(len(t.Record))
	for _, v := range t.Record {
		pw.Write(v.Serialize())
	}
	return pw.Bytes
}

func (t *G_animal_bet_table) DeserializeStruct(buf []byte, startid *int){
	var s0 int
	btk.ReadDynSize(buf,startid,&s0)
	if s0 < len(buf) {
		t.Record = make([]*G_animal_bet,s0)
		for i := 0; i < s0; i++ {
			a := new(G_animal_bet)
			a.DeserializeStruct(buf,startid)
			t.Record[i] = a
		}
	}
}

func (t *G_animal_bet_table) Deserialize(buf []byte){
	startid := 0
	t.DeserializeStruct(buf, &startid)
}


type G_animal_color_type struct{
	Id uint8
	CodeName string
	Desc string
}

func (t *G_animal_color_type) Serialize()[]byte{
	pw:= btk.ProtocolWriter{Bytes:make([]byte,0)}
	pw.WriteData(t.Id)
	pw.WriteData(t.CodeName)
	pw.WriteData(t.Desc)
	return pw.Bytes
}

func (t *G_animal_color_type) DeserializeStruct(buf []byte, startid *int){
	btk.Read(buf,startid,&t.Id)
	btk.Read(buf,startid,&t.CodeName)
	btk.Read(buf,startid,&t.Desc)
}

func (t *G_animal_color_type) Deserialize(buf []byte){
	startid := 0
	t.DeserializeStruct(buf, &startid)
}


type G_animal_color_type_table struct{
	Record []*G_animal_color_type
}

func (t *G_animal_color_type_table) Serialize()[]byte{
	pw:= btk.ProtocolWriter{Bytes:make([]byte,0)}
	pw.WriteDynSize(len(t.Record))
	for _, v := range t.Record {
		pw.Write(v.Serialize())
	}
	return pw.Bytes
}

func (t *G_animal_color_type_table) DeserializeStruct(buf []byte, startid *int){
	var s0 int
	btk.ReadDynSize(buf,startid,&s0)
	if s0 < len(buf) {
		t.Record = make([]*G_animal_color_type,s0)
		for i := 0; i < s0; i++ {
			a := new(G_animal_color_type)
			a.DeserializeStruct(buf,startid)
			t.Record[i] = a
		}
	}
}

func (t *G_animal_color_type_table) Deserialize(buf []byte){
	startid := 0
	t.DeserializeStruct(buf, &startid)
}


type G_animal_type struct{
	Id uint8
	CodeName string
	Desc string
}

func (t *G_animal_type) Serialize()[]byte{
	pw:= btk.ProtocolWriter{Bytes:make([]byte,0)}
	pw.WriteData(t.Id)
	pw.WriteData(t.CodeName)
	pw.WriteData(t.Desc)
	return pw.Bytes
}

func (t *G_animal_type) DeserializeStruct(buf []byte, startid *int){
	btk.Read(buf,startid,&t.Id)
	btk.Read(buf,startid,&t.CodeName)
	btk.Read(buf,startid,&t.Desc)
}

func (t *G_animal_type) Deserialize(buf []byte){
	startid := 0
	t.DeserializeStruct(buf, &startid)
}


type G_animal_type_table struct{
	Record []*G_animal_type
}

func (t *G_animal_type_table) Serialize()[]byte{
	pw:= btk.ProtocolWriter{Bytes:make([]byte,0)}
	pw.WriteDynSize(len(t.Record))
	for _, v := range t.Record {
		pw.Write(v.Serialize())
	}
	return pw.Bytes
}

func (t *G_animal_type_table) DeserializeStruct(buf []byte, startid *int){
	var s0 int
	btk.ReadDynSize(buf,startid,&s0)
	if s0 < len(buf) {
		t.Record = make([]*G_animal_type,s0)
		for i := 0; i < s0; i++ {
			a := new(G_animal_type)
			a.DeserializeStruct(buf,startid)
			t.Record[i] = a
		}
	}
}

func (t *G_animal_type_table) Deserialize(buf []byte){
	startid := 0
	t.DeserializeStruct(buf, &startid)
}


type G_card struct{
	Id uint8
	Color uint8
}

func (t *G_card) Serialize()[]byte{
	pw:= btk.ProtocolWriter{Bytes:make([]byte,0)}
	pw.WriteData(t.Id)
	pw.WriteData(t.Color)
	return pw.Bytes
}

func (t *G_card) DeserializeStruct(buf []byte, startid *int){
	btk.Read(buf,startid,&t.Id)
	btk.Read(buf,startid,&t.Color)
}

func (t *G_card) Deserialize(buf []byte){
	startid := 0
	t.DeserializeStruct(buf, &startid)
}


type G_card_table struct{
	Record []*G_card
}

func (t *G_card_table) Serialize()[]byte{
	pw:= btk.ProtocolWriter{Bytes:make([]byte,0)}
	pw.WriteDynSize(len(t.Record))
	for _, v := range t.Record {
		pw.Write(v.Serialize())
	}
	return pw.Bytes
}

func (t *G_card_table) DeserializeStruct(buf []byte, startid *int){
	var s0 int
	btk.ReadDynSize(buf,startid,&s0)
	if s0 < len(buf) {
		t.Record = make([]*G_card,s0)
		for i := 0; i < s0; i++ {
			a := new(G_card)
			a.DeserializeStruct(buf,startid)
			t.Record[i] = a
		}
	}
}

func (t *G_card_table) Deserialize(buf []byte){
	startid := 0
	t.DeserializeStruct(buf, &startid)
}


type G_card_color_type struct{
	Id uint8
	CodeName string
	Desc string
}

func (t *G_card_color_type) Serialize()[]byte{
	pw:= btk.ProtocolWriter{Bytes:make([]byte,0)}
	pw.WriteData(t.Id)
	pw.WriteData(t.CodeName)
	pw.WriteData(t.Desc)
	return pw.Bytes
}

func (t *G_card_color_type) DeserializeStruct(buf []byte, startid *int){
	btk.Read(buf,startid,&t.Id)
	btk.Read(buf,startid,&t.CodeName)
	btk.Read(buf,startid,&t.Desc)
}

func (t *G_card_color_type) Deserialize(buf []byte){
	startid := 0
	t.DeserializeStruct(buf, &startid)
}


type G_card_color_type_table struct{
	Record []*G_card_color_type
}

func (t *G_card_color_type_table) Serialize()[]byte{
	pw:= btk.ProtocolWriter{Bytes:make([]byte,0)}
	pw.WriteDynSize(len(t.Record))
	for _, v := range t.Record {
		pw.Write(v.Serialize())
	}
	return pw.Bytes
}

func (t *G_card_color_type_table) DeserializeStruct(buf []byte, startid *int){
	var s0 int
	btk.ReadDynSize(buf,startid,&s0)
	if s0 < len(buf) {
		t.Record = make([]*G_card_color_type,s0)
		for i := 0; i < s0; i++ {
			a := new(G_card_color_type)
			a.DeserializeStruct(buf,startid)
			t.Record[i] = a
		}
	}
}

func (t *G_card_color_type_table) Deserialize(buf []byte){
	startid := 0
	t.DeserializeStruct(buf, &startid)
}


type G_game_type struct{
	Id uint8
	CodeName string
	Desc string
	Player_max uint8
	No_operate_times uint8
	Watch_times uint8
}

func (t *G_game_type) Serialize()[]byte{
	pw:= btk.ProtocolWriter{Bytes:make([]byte,0)}
	pw.WriteData(t.Id)
	pw.WriteData(t.CodeName)
	pw.WriteData(t.Desc)
	pw.WriteData(t.Player_max)
	pw.WriteData(t.No_operate_times)
	pw.WriteData(t.Watch_times)
	return pw.Bytes
}

func (t *G_game_type) DeserializeStruct(buf []byte, startid *int){
	btk.Read(buf,startid,&t.Id)
	btk.Read(buf,startid,&t.CodeName)
	btk.Read(buf,startid,&t.Desc)
	btk.Read(buf,startid,&t.Player_max)
	btk.Read(buf,startid,&t.No_operate_times)
	btk.Read(buf,startid,&t.Watch_times)
}

func (t *G_game_type) Deserialize(buf []byte){
	startid := 0
	t.DeserializeStruct(buf, &startid)
}


type G_game_type_table struct{
	Record []*G_game_type
}

func (t *G_game_type_table) Serialize()[]byte{
	pw:= btk.ProtocolWriter{Bytes:make([]byte,0)}
	pw.WriteDynSize(len(t.Record))
	for _, v := range t.Record {
		pw.Write(v.Serialize())
	}
	return pw.Bytes
}

func (t *G_game_type_table) DeserializeStruct(buf []byte, startid *int){
	var s0 int
	btk.ReadDynSize(buf,startid,&s0)
	if s0 < len(buf) {
		t.Record = make([]*G_game_type,s0)
		for i := 0; i < s0; i++ {
			a := new(G_game_type)
			a.DeserializeStruct(buf,startid)
			t.Record[i] = a
		}
	}
}

func (t *G_game_type_table) Deserialize(buf []byte){
	startid := 0
	t.DeserializeStruct(buf, &startid)
}


type G_global_var struct{
	Id uint16
	CodeName string
	Desc string
	ValInt int32
	ValFloat float32
	ValString string
}

func (t *G_global_var) Serialize()[]byte{
	pw:= btk.ProtocolWriter{Bytes:make([]byte,0)}
	pw.WriteData(t.Id)
	pw.WriteData(t.CodeName)
	pw.WriteData(t.Desc)
	pw.WriteData(t.ValInt)
	pw.WriteData(t.ValFloat)
	pw.WriteData(t.ValString)
	return pw.Bytes
}

func (t *G_global_var) DeserializeStruct(buf []byte, startid *int){
	btk.Read(buf,startid,&t.Id)
	btk.Read(buf,startid,&t.CodeName)
	btk.Read(buf,startid,&t.Desc)
	btk.Read(buf,startid,&t.ValInt)
	btk.Read(buf,startid,&t.ValFloat)
	btk.Read(buf,startid,&t.ValString)
}

func (t *G_global_var) Deserialize(buf []byte){
	startid := 0
	t.DeserializeStruct(buf, &startid)
}


type G_global_var_table struct{
	Record []*G_global_var
}

func (t *G_global_var_table) Serialize()[]byte{
	pw:= btk.ProtocolWriter{Bytes:make([]byte,0)}
	pw.WriteDynSize(len(t.Record))
	for _, v := range t.Record {
		pw.Write(v.Serialize())
	}
	return pw.Bytes
}

func (t *G_global_var_table) DeserializeStruct(buf []byte, startid *int){
	var s0 int
	btk.ReadDynSize(buf,startid,&s0)
	if s0 < len(buf) {
		t.Record = make([]*G_global_var,s0)
		for i := 0; i < s0; i++ {
			a := new(G_global_var)
			a.DeserializeStruct(buf,startid)
			t.Record[i] = a
		}
	}
}

func (t *G_global_var_table) Deserialize(buf []byte){
	startid := 0
	t.DeserializeStruct(buf, &startid)
}


type G_robot struct{
	Id uint8
	Name string
	Icon uint8
	Gender uint8
}

func (t *G_robot) Serialize()[]byte{
	pw:= btk.ProtocolWriter{Bytes:make([]byte,0)}
	pw.WriteData(t.Id)
	pw.WriteData(t.Name)
	pw.WriteData(t.Icon)
	pw.WriteData(t.Gender)
	return pw.Bytes
}

func (t *G_robot) DeserializeStruct(buf []byte, startid *int){
	btk.Read(buf,startid,&t.Id)
	btk.Read(buf,startid,&t.Name)
	btk.Read(buf,startid,&t.Icon)
	btk.Read(buf,startid,&t.Gender)
}

func (t *G_robot) Deserialize(buf []byte){
	startid := 0
	t.DeserializeStruct(buf, &startid)
}


type G_robot_table struct{
	Record []*G_robot
}

func (t *G_robot_table) Serialize()[]byte{
	pw:= btk.ProtocolWriter{Bytes:make([]byte,0)}
	pw.WriteDynSize(len(t.Record))
	for _, v := range t.Record {
		pw.Write(v.Serialize())
	}
	return pw.Bytes
}

func (t *G_robot_table) DeserializeStruct(buf []byte, startid *int){
	var s0 int
	btk.ReadDynSize(buf,startid,&s0)
	if s0 < len(buf) {
		t.Record = make([]*G_robot,s0)
		for i := 0; i < s0; i++ {
			a := new(G_robot)
			a.DeserializeStruct(buf,startid)
			t.Record[i] = a
		}
	}
}

func (t *G_robot_table) Deserialize(buf []byte){
	startid := 0
	t.DeserializeStruct(buf, &startid)
}


