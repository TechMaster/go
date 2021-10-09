package model

// Company belongs to Corp : công ty thuộc
type Profile struct {
	Id     int
	Lang   string
	UserId int
}

type Company struct {
	Id      int
	Name    string
	Profile *Profile `pg:"rel:belongs-to"`
}

/*
What's the difference between belongs_to and has_one?

They essentially do the same thing, the only difference is what side of the
relationship you are on. If a User has a Profile, then in the User class you'd
have has_one :profile and in the Profile class you'd have belongs_to :user.
To determine who "has" the other object, look at where the foreign key is.
We can say that a User "has" a Profile because the profiles table has a user_id
lumn. If there was a column called profile_id on the users table, however, we
would say that a Profile has a User, and the belongs_to/has_one locations would
apped.
*/
