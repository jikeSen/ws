package models

import (
    "go/ws/connecter"
)

type Grab struct {
    Id           int `gorm:"primary_key" json:"id"`
    GrabName     string
    GrabPrice    int
    ChargeAssign float32
    AgentAssign  float32
    Cid          int
    Silver       int
    IsLuck       int
    Diamond      int
    Getcoin      int
    VipLevel     int
}

//修改默认表名
func (Grab) TableName() string {
    return "qy_grab"
}

func (grab *Grab) GetGrabInfo(id int) Grab {
    var grabinfo Grab
    // connecter.Db.Debug().First(&grabinfo, id)
    connecter.Db.First(&grabinfo, id)
    return grabinfo
}
