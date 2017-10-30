package storage

import "github.com/Andiedie/agenda-go/model"

//AddUser 添加用户
func AddUser(user model.User) {
}

//DeleteUser 删除用户
func DeleteUser(username string) {
}

//SetCurrentUser 设置当前登录用户
//nil表示 退出登录
func SetCurrentUser(username string) {
}

//GetCurrentUser 获取当前登录用户的用户名
//nil表示 未登录
func GetCurrentUser() string {
	return ""
}

//GetAllUser 获取所有注册用户
func GetAllUser() []model.User {
	return []model.User{}
}

//AddMeeting 增加会议
func AddMeeting(meeting model.Meeting) {
}

//GetMeeting 获取会议
func GetMeeting(title string) *model.Meeting {
	return nil
}

//UpdateMeeting 修改会议
func UpdateMeeting(meeting model.Meeting) {
}

//GetALLMeeting 获取所有会议
func GetALLMeeting() []model.Meeting {
	return nil
}

//DeleteMeeting 删除指定会议
func DeleteMeeting(title string) {
}
