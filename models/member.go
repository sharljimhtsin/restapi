package models

var (
	Members map[string]*Member
)

type Member struct {
	MemberId    string
	LoginName   string
	LoginPasswd string
}

func init() {
	Members = make(map[string]*Member)
	Members["test"] = &Member{"1111", "2222", "33333333"}
	Members["test2"] = &Member{"1111", "2222", "33333333"}
	Members["test3"] = &Member{"1111", "2222", "33333333"}
}

func GetAllMember() map[string]*Member {
	return Members
}
