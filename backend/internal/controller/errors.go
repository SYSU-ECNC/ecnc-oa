package controller

import "errors"

var (
	// req
	ErrUsernameEmpty = errors.New("用户名为空")
	ErrPasswordEmpty = errors.New("密码为空")
	// sql.DB
	ErrUsernameNotExists = errors.New("用户名不存在")
	// bcrypt
	ErrWrongPassword = errors.New("密码错误")
)
