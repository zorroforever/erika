package logic

import (
	mapper "iris/datasource/mybatis/mapper"
	userModel "iris/model/user"
)

func FetchAllUser() ([]userModel.UserData, error){
	if !DoCheckDBURI() {
		return nil, nil
	}
	//使用mapper
	var result, err = mapper.UserMapperImpl.FetchAllUserData()
	if err != nil {
		panic(err)
	}
	return result,err
}