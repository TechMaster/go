package model

type User struct { //Tên Entity --> biên thành snake_case trong Postgresql
	tableName  struct{} `pg:"test.users"` //--> trỏ đến schema.table
	Id         string   `pg:"id,pk"`      //pk--> đây là primary key
	Name       string   //-> name, kiểu string --> text
	Email      string
	Mobile     string
	Int_roles  []int    `pg:"int_roles,array"`  //Quy ước IntRoles --> int_roles snake case
	Enum_roles []string `pg:"enum_role,array"` //kiểu cột là array lưu string
}

type User_Role struct {
	tableName struct{} `pg:"test.user_role"`
	User_id   string
	Role_id   int
}
