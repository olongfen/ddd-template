package repository

import (
	"context"
	"ddd-template/internal/adapters/repository/db_iface"
	"ddd-template/internal/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// ILike ilike
type ILike clause.Eq

// Build builder sql
func (like ILike) Build(builder clause.Builder) {
	builder.WriteQuoted(like.Column)
	_, err := builder.WriteString(" ILIKE ")
	if err != nil {
		panic(err)
	}
	builder.AddVar(builder, like.Value)
}

// NegationBuild builder sql
func (like ILike) NegationBuild(builder clause.Builder) {
	builder.WriteQuoted(like.Column)
	_, err := builder.WriteString(" NOT LIKE ")
	if err != nil {
		panic(err)
	}
	builder.AddVar(builder, like.Value)
}

// fieldWhere process field symbol
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
		return clause.IN{Column: column, Values: []interface{}{field.Value}}
	default:
		return clause.Eq{Column: column, Value: field.Value}
	}
}

// option query conditions
type option struct {
	order       map[string]bool
	pageSize    int
	currentPage int
	all         bool
	noCount     bool
}

// TFields field array type
type TFields []domain.Field

// process handler db.Where()
func (f TFields) process(db *gorm.DB) *gorm.DB {
	for _, v := range f {
		db = db.Where(fieldWhere(v))
	}
	return db
}

// newPotion new
func newOption(o domain.OtherCond) *option {
	opt := new(option)
	opt.order = map[string]bool{}
	opt.currentPage = o.CurrentPage
	opt.pageSize = o.PageSize
	opt.all = o.All
	opt.noCount = o.NoCount
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

// findPage find page
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
	// 不需要count
	if !opt.noCount {
		if err = db.Count(&pagination.TotalCount).Error; err != nil {
			return
		}
	}
	// 不需要获取全部,默认分页查询
	if !opt.all {
		pagination.CurrentPage = opt.currentPage
		pagination.PageSize = opt.pageSize
		pageNum, pageSize := pagination.CurrentPage, pagination.PageSize
		if pageNum > 0 && pageSize > 0 {
			db = db.Offset((pageNum - 1) * pageSize).Limit(pageSize)
		} else if pageSize > 0 {
			db = db.Limit(pageSize)
		}
		if pagination.TotalCount%int64(pagination.PageSize) == 0 {
			pagination.TotalPage = int(pagination.TotalCount / int64(pagination.PageSize))
		} else {
			pagination.TotalPage = int(pagination.TotalCount/int64(pagination.PageSize)) + 1
		}
	}

	for column, v := range opt.order {
		db = db.Order(clause.OrderByColumn{Column: clause.Column{Name: column}, Desc: v})
	}
	if err = db.Find(out).Error; err != nil {
		return
	}

	return

}

// repository 增删改查泛型
type repository[T any] struct {
	data db_iface.DBData
}

// FindOne  get one
func (u *repository[T]) FindOne(ctx context.Context, id int) (ret *T, err error) {
	var (
		model T
	)
	if err = u.data.DB(ctx).Model(&model).Where("id = ?", id).First(&ret).Error; err != nil {
		return
	}
	return
}

// FindOneBy  get one
func (u *repository[T]) FindOneBy(ctx context.Context, field ...domain.Field) (ret *T, err error) {
	var (
		model T
	)
	db := u.data.DB(ctx).Model(&model)

	if err = TFields(field).process(db).First(&ret).Error; err != nil {
		return
	}
	return
}

// Find get page
func (u *repository[T]) Find(ctx context.Context, o domain.OtherCond, fields ...domain.Field) (ret []*T,
	pagination *domain.Pagination, err error) {
	var (
		data  []*T
		model T
		db    = u.data.DB(ctx).Model(&model)
		opt   = newOption(o)
	)

	if pagination, err = findPage(TFields(fields).process(db), opt, &data); err != nil {
		return
	}
	ret = data
	return
}

// Create 往数据库写入user记录
func (u *repository[T]) Create(ctx context.Context, stu *T) (err error) {
	var (
		model T
	)
	if err = u.data.DB(ctx).Model(&model).Create(stu).Error; err != nil {
		return
	}
	return
}

// Update update
func (u *repository[T]) Update(ctx context.Context, id int, stu *T) (err error) {
	var (
		model T
	)
	if err = u.data.DB(ctx).Model(&model).Where("id = ?", id).Updates(stu).Error; err != nil {
		return
	}
	return
}

// Delete del
func (u *repository[T]) Delete(ctx context.Context, id int) (err error) {
	var (
		model T
	)
	if err = u.data.DB(ctx).Where("id = ?", id).Delete(&model).Error; err != nil {
		return
	}
	return
}

// DeleteBy delete by fields
func (u *repository[T]) DeleteBy(ctx context.Context, fields ...domain.Field) (err error) {
	var (
		model T
		db    = u.data.DB(ctx)
	)

	if err = TFields(fields).process(db).Delete(&model).Error; err != nil {
		return
	}
	return
}

// Count return number
func (u *repository[T]) Count(ctx context.Context, fields ...domain.Field) (count int64, err error) {
	var (
		model T
		db    = u.data.DB(ctx).Model(model)
	)
	if err = TFields(fields).process(db).Count(&count).Error; err != nil {
		return
	}
	return
}
