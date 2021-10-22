package repo

import "one_many/model"

func Get_Foo_Bar() (foos []model.Foo, err error) {
	err = DB.Model(&foos).Relation("Bars").Select()
	if err != nil {
		return nil, err
	}
	return foos, nil
}

func Get_Foo_By_ID(id string) (foo *model.Foo, err error) {
	/*foo = new(model.Foo)
	err = DB.Model(foo).Relation("Bars").Where("id = ?", id).Limit(1).Select()
	*/
	//err = DB.Model(foo).Relation("Bars").Where("id = ?", id).First() --> WHERE (id = 'ox-01') ORDER BY "foo"."id" LIMIT 1

	foo = &model.Foo{
		Id: id,
	}
	err = DB.Model(foo).Relation("Bars").WherePK().Select()
	if err != nil {
		return nil, err
	}
	return foo, nil
}

func Get_Bar_By_ID(id string) (bar *model.Bar, err error) {
	bar = &model.Bar{
		Id: "bar3",
	}

	err = DB.Model(bar).Relation("Foo").WherePK().Select()

	if err != nil {
		return nil, err
	}
	return bar, nil
}
