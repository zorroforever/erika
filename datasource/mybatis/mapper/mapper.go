package mybatis

import userModel "iris/model/user"

// 用户sql集合
type UserMapper struct {
	InsertUserData    func(arg userModel.UserData) (int64, error)
	FetchAllUserData  func() ([]userModel.UserData, error)
}
// 用户sql对象集
var UserMapperImpl = UserMapper{}
