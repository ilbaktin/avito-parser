package avito

import "univer/avito-parser/model"

type Context struct {
	Regions    map[int]*model.Region
	Categories map[int]*model.Category
}

//var context *Context
//
//func InitContext() error {
//	regions, err := GetAllRegions()
//	if err != nil {
//		return err
//	}
//	categories, err := GetCateroriesWithEnNames()
//}
