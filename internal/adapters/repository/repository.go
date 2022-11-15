package repository

import (
	"context"
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

// repository 增删改查泛型
type repository[T any] struct {
	data *Data
}

// FindOne  get one
func (u repository[T]) FindOne(ctx context.Context, id int) (ret *T, err error) {
	var (
		model T
	)
	if err = u.data.DB(ctx).Model(&model).Where("id = ?", id).First(&ret).Error; err != nil {
		return
	}
	return
}

// FindOneBy  get one
func (u repository[T]) FindOneBy(ctx context.Context, field ...domain.Field) (ret *T, err error) {
	var (
		model T
	)
	db := u.data.DB(ctx).Model(&model)
	fieldsT(field).process(db)
	if err = db.First(&ret).Error; err != nil {
		return
	}
	return
}

// Find get page
func (u repository[T]) Find(ctx context.Context, o domain.OtherCond, fields ...domain.Field) (ret []*T,
	pagination *domain.Pagination, err error) {
	var (
		data  []*T
		model T
		db    = u.data.DB(ctx).Model(&model)
		opt   = newOption(o)
	)
	fieldsT(fields).process(db)
	if pagination, err = findPage(db, opt, &data); err != nil {
		return
	}
	ret = data
	return
}

// Create 往数据库写入user记录
func (u repository[T]) Create(ctx context.Context, stu *T) (err error) {
	var (
		model T
	)
	if err = u.data.DB(ctx).Model(&model).Create(stu).Error; err != nil {
		return
	}
	return
}

// Update update
func (u repository[T]) Update(ctx context.Context, id int, stu *T) (err error) {
	var (
		model T
	)
	if err = u.data.DB(ctx).Model(&model).Where("id = ?", id).Updates(stu).Error; err != nil {
		return
	}
	return
}

// Delete del
func (u repository[T]) Delete(ctx context.Context, id int) (err error) {
	var (
		model T
	)
	if err = u.data.DB(ctx).Where("id = ?", id).Delete(&model).Error; err != nil {
		return
	}
	return
}
