package repo

import "many_many/model"

// Tìm thành viên, lấy danh sách club mà người đó tham gia
func Get_member_by_name(name string) (member *model.Member, err error) {
	member = new(model.Member)
	err = DB.Model(member).Relation("Clubs").Where("name ILIKE ?", name).Limit(1).Select()
	if err != nil {
		return nil, err
	}
	return member, nil
}

type Club_Membership struct {
	Club   string `json:"club"`
	Status bool   `json:"status"`
}
type Member_Clubs struct {
	Id    string
	Name  string
	Clubs []Club_Membership
}

/*
Khi muốn lấy thêm trường bổ xung ở bảng trung gian. Ví dụ liệt kê tất cả các thành viên.
Trong mỗi thành viên liệt kê club họ tham gia và trạng thái tham gia
John
   Math : false
   Music : true
Anna
   Math : true
Bob
   Sport : true
   Music : false
*/
func Get_active_members_of_club() (members []Member_Clubs, err error) {
	query := `select m.id, m.name, 
	json_agg(json_build_object('club', c.name, 'status', mc.active)) 
	from test."member" m 
	inner join test.member_club mc on m.id = mc.member_id 
	inner join test.club c on mc.club_id = c.id 
	group by m.id`

	_, err = DB.Query(&members, query)
	if err != nil {
		return nil, err
	} else {
		return members, nil
	}
}
