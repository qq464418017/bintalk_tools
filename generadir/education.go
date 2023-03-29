package bintalk


type P_account struct {
    User_id uint32  `gorm:"primary_key"`
    User_name string
    Password string
    Phone string
    Type uint8
    Channel uint8
    Add_ts uint32
    Lock_ts uint32
    Login_account string
    Level int8
    Lang string
    Parent_id int32
}


type P_chat struct {
    Guid uint64  `gorm:"primary_key"`
    Friend_id uint32
    User_id uint32
    Content string
    Picture string
    Emoticon string
    Ts uint32
}


type P_children struct {
    Guid uint32  `gorm:"primary_key"`
    Parent_id uint32
    Name string
    Birthday string
    Add_ts uint32
    Img string
    Gender uint8
    Remark string
}


type P_course struct {
    Guid uint32  `gorm:"primary_key"`
    School_id uint32
    Name string
    Img string
    Remark string
    Cancel_ts uint32
}


type P_course_card struct {
    Guid uint32  `gorm:"primary_key"`
    Course_id uint32
    Teacher_id uint32
    Start_time uint32
    End_time uint32
    Week uint8
    Cancel_ts uint32
}


type P_fcm_token struct {
    Guid uint32  `gorm:"primary_key"`
    Player_id uint32
    Device_id string
    Fcm_token string
    App_language uint8
}


type P_friend struct {
    Guid uint32  `gorm:"primary_key"`
    Teacher_id uint32
    Parent_id uint32
    Add_ts uint32
    Cancel_ts uint32
}


type P_friend_chat struct {
    Guid uint32  `gorm:"primary_key"`
    User_id uint32
    Friend_id uint32
    Unread_count uint32
    Alias string
}


type P_notice struct {
    Guid uint32  `gorm:"primary_key"`
    Title string
    Content string
    Img string
    Must_read uint8
    Is_top uint8
    Add_ts uint32
    Delete_ts uint32
    School_id uint32
    Teacher_id uint32
}


type P_notice_read struct {
    Guid int32  `gorm:"primary_key"`
    User_id uint32
    Notice_id uint32
    Ts uint32
}


type P_player_info struct {
    User_id uint32  `gorm:"primary_key"`
    Name string
    Img string
    Gender uint8
    Remark string
}


type P_scan_code struct {
    User_id uint32  `gorm:"primary_key"`
    School_id uint32
    Path string
    Add_ts uint32
    Cancel_ts uint32
}


type P_school struct {
    Guid uint32  `gorm:"primary_key"`
    Name string
    Img string
    Path string
    Remark string
    Auth_ts uint32
    Expire_ts uint32
    Parent_id uint32
    Cancel_ts uint32
    No_sgin_stat_days uint32
    Course_notify_times uint32
    Date_notify_days uint32
    No_sign_notify_days uint32
}


type P_school_achievement struct {
    Guid uint32  `gorm:"primary_key"`
    School_id uint32
    Name string
    Cancel_ts uint32
    Sort int32
}


type P_sign_in struct {
    Guid uint32  `gorm:"primary_key"`
    School_id uint32
    Teacher_id uint32
    Student_course_id uint32
    Children_id uint32
    Course_id uint32
    Course_card_id uint32
    Scan_code_id uint32
    Add_ts uint32
    State uint8
    Remark string
    Delete_ts uint32
}


type P_student struct {
    Guid uint32  `gorm:"primary_key"`
    Children_id uint32
    Nick string
    School_id uint32
    Add_ts uint32
    Remark string
    Cancel_ts uint32
}


type P_student_achievement struct {
    Guid uint32  `gorm:"primary_key"`
    Children_id uint32
    School_id uint32
    Achieve_ids string
}


type P_student_course struct {
    Guid uint32  `gorm:"primary_key"`
    School_id uint32
    Children_id uint32
    Course_id uint32
    Use_count uint32
    Total_count uint32
    Buy_type uint8
    Start_date string
    End_date string
    Pay_payment uint8
    Pay_amount float32
    Pay_date string
    Pay_remark string
    Pay_ts uint32
    Refund_payment uint8
    Refund_amount float32
    Refund_date string
    Refund_reason string
    Refund_remark string
    Refund_ts uint32
    Cancel_ts uint32
}


type P_teacher struct {
    Guid uint32  `gorm:"primary_key"`
    User_id uint32
    School_id uint32
    Is_master uint8
    Remark string
    Add_ts uint32
    Cancel_ts uint32
}


type S_maintain_notice struct {
    Channel uint8  `gorm:"primary_key"`
    Is_show uint8
    Content string
}


type S_version struct {
    Channel int32  `gorm:"primary_key"`
    Version int32
    Type string
}


type S_white_list struct {
    Id uint16  `gorm:"primary_key"`
    Ip string
}


type W_finance struct {
    Finance_id int32  `gorm:"primary_key"`
    Time string
    Collection float32
    Refund float32
    School_id int32
    New_count int32
    Drop_count int32
}


