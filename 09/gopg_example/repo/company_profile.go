package repo

import (
	"fmt"
	"gopgdemo/model"
)

func Create_company_profile() (err error) {
	qs := []string{
		"CREATE TEMP TABLE companies (id int, name text)",
		"CREATE TEMP TABLE profiles (id int, lang text, user_id int)",
		"INSERT INTO companies VALUES (1, 'Microsoft'), (2, 'Telegram'), (3, 'Honda')",
		"INSERT INTO profiles VALUES (1, 'en', 1), (2, 'ru', 2), (3, 'jp', 2)",
	}
	for _, q := range qs {
		_, err = DB.Exec(q)
		if err != nil {
			return err
		}
	}

	// Select companies joining their profiles with following query:
	//
	// SELECT
	//   "user".*,
	//   "profile"."id" AS "profile__id",
	//   "profile"."lang" AS "profile__lang",
	//   "profile"."user_id" AS "profile__user_id"
	// FROM "companies" AS "user"
	// LEFT JOIN "profiles" AS "profile" ON "profile"."user_id" = "user"."id"

	var companies []model.Company
	err = DB.Model(&companies).
		Column("company.*").
		Relation("Profile").
		Select()

	if err != nil {
		return err
	}

	fmt.Println(len(companies), "results")
	fmt.Println(companies[0].Id, companies[0].Name, companies[0].Profile)
	fmt.Println(companies[1].Id, companies[1].Name, companies[1].Profile)
	return nil
}

/*
CREATE TEMP TABLE companies (id int, name text)
CREATE TEMP TABLE profiles (id int, lang text, user_id int)
INSERT INTO companies VALUES (1, 'Microsoft'), (2, 'Telegram'), (3, 'Honda')
INSERT INTO profiles VALUES (1, 'en', 1), (2, 'ru', 2), (3, 'jp', 2)
SELECT "company".*, "profile"."id" AS "profile__id", "profile"."lang" AS "profile__lang", "profile"."user_id" AS "profile__user_id" FROM "companies" AS "company" LEFT JOIN "profiles" AS "profile" ON "profile"."id" = "company"."id"
3 results
1 Microsoft &{1 en 1}
2 Telegram &{2 ru 2}
*/
