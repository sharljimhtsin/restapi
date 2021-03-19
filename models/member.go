package models

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
)

var (
	Members map[string]*Member
)

type Member struct {
	MemberId    int64 `orm:"pk"`
	LoginName   string
	LoginPasswd string
}

func (m *Member) TableName() string {
	return "t_member"
}

func init() {
	Members = make(map[string]*Member)
	Members["test"] = &Member{1111, "2222", "33333333"}
	Members["test2"] = &Member{1111, "2222", "33333333"}
	Members["test3"] = &Member{1111, "2222", "33333333"}
	//DB
	_ = orm.RegisterDriver("mysql", orm.DRMySQL)
	_ = orm.RegisterDataBase("default", "mysql", "imdeal:123456@tcp(192.168.134.52:3306)/imdeal?charset=utf8")
	orm.RegisterModel(new(Member))
	orm.Debug = true
}

func GetAllMember() []*Member {
	o := orm.NewOrm()
	m := new(Member)
	query := o.QueryTable(m)
	var members []*Member
	_, _ = query.Limit(10).Filter("login_name", "admin").All(&members)
	return members
}

func GetDBMember() Member {
	o := orm.NewOrm()
	m := Member{MemberId: 1346635616758009822}
	err := o.Read(&m)
	if err != nil {
		fmt.Println(m.LoginName)
	}
	return m
}
