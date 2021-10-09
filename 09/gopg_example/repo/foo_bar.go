package repo

import "gopgdemo/model"

func GetFooBar() (foos []model.Foo, err error) {
	err = DB.Model(&foos).Relation("Bars").Select()
	if err != nil {
		return nil, err
	}
	return foos, nil
}

func GetFooByID(id string) (foo *model.Foo, err error) {
	foo = new(model.Foo)
	err = DB.Model(foo).Relation("Bars").Where("id = ?", id).Limit(1).Select()
	//err = DB.Model(foo).Relation("Bars").Where("id = ?", id).First() --> WHERE (id = 'ox-01') ORDER BY "foo"."id" LIMIT 1

	if err != nil {
		return nil, err
	}
	return foo, nil
}
