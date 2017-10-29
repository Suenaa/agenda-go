package service


type Service struct {

}


//启动agenda
//传入一个service.Service指针
func StartAgenda(service *Service) (bool, string) {
  return false,""
}

//用户注册
//函数返回一个字符串表示注册信息，
//比如false+“userName already exits”之类的,或者true+"register successful"
func (service *Service) UserRegister(userName string, passworld string,
  email string, phone string) (bool, string) {
  return false,""
}

//用户登录
//返回一个字符串,类似上一个函数
func (service *Service) UserLogin(userName string, passworld string) (bool, string) {
  return false,""
}

//用户登出
//返回用户是否登出成功以及一些登出信息，比如当前没有用户处于登录状态
func (service *Service) UserLogout() (bool, string) {
  return false
}

//查询所有用户
//返回一个slice
func (service *Service) QueryAllUsers() []model.User {
  return []model.User{}
}


//删除用户
//删除该用户的用户信息以及参加和发起的会议，如果该用户是某一个会议的发起者，则删除该会议，
//如果由于删除该用户造成某一个会议的参与者变成0，则删除该会议
//返回的内容同上所述
func (service *Service) DeleteUser(userName string, passworld string) (bool, string) {
  return false,""
}


//创建会议
//返回类型类似上面几个类似函数的返回值
func (service *Service) CreateMeeting(title string, startDate string,
  endDate string, participators []model.User) (bool, string) {
  return false,""
}

//添加自己发起的某一会议的一个参与者
//返回添加成功与否以及相关信息
func (service *Service) AddParticipator(title string, paticipator string) (bool, string) {
  return false,""
}

//删除自己创建的某一会议的一个参与者
//如果会议的参与者因此变成0，则删除该会议
//返回删除成功与否以及相关信息
func (service *Service) DeleteParticipator(title string, participator string) (bool, string) {
  return false,""
}

//查询会议，通过开始时间和结束时间查询当前用户需要参加的所有会议
//返回查询结果的slice
func (service *Service) QueryMeeting(startDate string, endDate string) []model.Metting {
  return []model.Meeting{}
}

//取消会议
//返回取消会议成功与否以及相关信息
func (service *Service) DeleteMeeting(title string) (bool, string) {
  return false,""
}

//退出会议
//如果退出会议之后的参与者为0的会议将会被删除
//返回退出会议成功与否以及相关的信息
func (service *Service) QuitMeeting(title string) (bool, string) {
  return false,""
}

//清空当前用户发起的所有会议安排
//返回清空是否成功以及相关的信息
func (service *Service) DeleteAllMeeting() (bool, string) {
  return false,""
}

/*
//通过主题查询会议
func (service *Service) ListMeetingByTitle(userName string, title string) []model.Meeting {
  return []model.Meeting{}
}
*/
