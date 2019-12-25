package models

import "go/ws/connecter"

type User struct {
    Id    int    `json:"id"`
    Token string `json:"token"`
}

//修改默认表名
func (User) TableName() string {
    return "qy_user"
}

func (user *User) GetUserInfoByUid(uid int) (u User) {
    u = User{}
    connecter.Db.First(&u, uid)
    return u
}
