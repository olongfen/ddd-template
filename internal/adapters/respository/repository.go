package respository

import (
	"ddd-template/internal/schema"
	"gorm.io/gorm"
)

func findPage(db *gorm.DB, opt schema.QueryOptions, out interface{}) (pagination *schema.Pagination, err error) {
	//if !opt.NotCount {
	//	if opt.Distinct != "" {
	//		_db := db.Session(&gorm.Session{})
	//		if err = _db.Distinct(opt.Distinct).Count(&count).Error; err != nil {
	//			return 0, err
	//		}
	//	} else {
	//		if err = db.Count(&count).Error; err != nil {
	//			return 0, err
	//		}
	//	}
	//
	//	if count == 0 {
	//		return count, nil
	//	}
	//}
	pagination = new(schema.Pagination)
	if err = db.Count(&pagination.TotalCount).Error; err != nil {
		return
	}
	pagination.CurrentPage = opt.CurrentPage
	pagination.PageSize = opt.PageSize
	if pagination.CurrentPage == 0 {
		pagination.CurrentPage = 1
	}
	if pagination.PageSize == 0 {
		pagination.PageSize = 10
	}
	pageNum, pageSize := pagination.CurrentPage, pagination.PageSize
	if pageNum > 0 && pageSize > 0 {
		db = db.Offset((pageNum - 1) * pageSize).Limit(pageSize)
	} else if pageSize > 0 {
		db = db.Limit(pageSize)
	}

	//for _, v := range opt.OrderFields {
	//	db = db.Order(clause.OrderByColumn{Column: clause.Column{Name: v.Column}, Desc: v.Desc})
	//}
	//for _, v := range opt.OrderSQLStr {
	//	db = db.Order(v)
	//}
	if err = db.Find(out).Error; err != nil {
		return
	}

	if pagination.TotalCount%int64(pagination.PageSize) == 0 {
		pagination.TotalPage = int(pagination.TotalCount / int64(pagination.PageSize))
	} else {
		pagination.TotalPage = int(pagination.TotalCount/int64(pagination.PageSize)) + 1
	}
	return

}
