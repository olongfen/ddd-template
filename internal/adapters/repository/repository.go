package repository

import (
	"ddd-template/internal/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ILike clause.Eq

func (like ILike) Build(builder clause.Builder) {
	builder.WriteQuoted(like.Column)
	builder.WriteString(" ILIKE ")
	builder.AddVar(builder, like.Value)
}

func (like ILike) NegationBuild(builder clause.Builder) {
	builder.WriteQuoted(like.Column)
	builder.WriteString(" NOT LIKE ")
	builder.AddVar(builder, like.Value)
}

func fieldWhere(field domain.Field) clause.Expression {
	column := snakeString(field.Column)
	switch field.Symbol {
	case ">":
		return clause.Gt{Column: column, Value: field.Value}
	case ">=":
		return clause.Gte{Column: column, Value: field.Value}
	case "<":
		return clause.Lt{Column: column, Value: field.Value}
	case "<=":
		return clause.Lte{Column: column, Value: field.Value}
	case "like":
		return clause.Like{Column: column, Value: field.Value}
	case "ilike":
		return ILike{Column: column, Value: field.Value}
	case "in":
		return clause.IN{Column: column, Values: field.Value.([]interface{})}
	default:
		return clause.Eq{Column: column, Value: field.Value}
	}
}

type option struct {
	order       map[string]bool
	pageSize    int
	currentPage int
}

type fieldsT []domain.Field

func (f fieldsT) process(db *gorm.DB) {
	for _, v := range f {
		db = db.Where(fieldWhere(v))
	}
	return
}

func newOption(o domain.OtherCond) *option {
	opt := new(option)
	opt.order = map[string]bool{}
	opt.currentPage = o.CurrentPage
	opt.pageSize = o.PageSize
	if opt.currentPage == 0 {
		opt.currentPage = 1
	}
	if opt.pageSize == 0 {
		opt.pageSize = 10
	}
	for i := 0; i < len(o.Sort) && i < len(o.Order); i++ {
		switch o.Order[i] {
		case "asc":
			opt.order[snakeString(o.Sort[i])] = false
		default:
			opt.order[snakeString(o.Sort[i])] = true

		}
	}
	return opt
}

func findPage(db *gorm.DB, opt *option, out interface{}) (pagination *domain.Pagination, err error) {
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
	pagination = new(domain.Pagination)
	if err = db.Count(&pagination.TotalCount).Error; err != nil {
		return
	}
	pagination.CurrentPage = opt.currentPage
	pagination.PageSize = opt.pageSize
	pageNum, pageSize := pagination.CurrentPage, pagination.PageSize
	if pageNum > 0 && pageSize > 0 {
		db = db.Offset((pageNum - 1) * pageSize).Limit(pageSize)
	} else if pageSize > 0 {
		db = db.Limit(pageSize)
	}

	for column, v := range opt.order {
		db = db.Order(clause.OrderByColumn{Column: clause.Column{Name: column}, Desc: v})
	}
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
