package respository

import (
	"context"
	"ddd-template/internal/domain"
	"ddd-template/pkg/utils"
)

type User struct {
	utils.Model
	Username string `gorm:"type:varchar(16)；uniqueIndex;comment:用户名称"`
	Password string `size:"size:16;comment:密码"`
}

type userRepository struct {
	tx   domain.ITransaction
	data *Data
}

func (u userRepository) marshalUser(in *domain.User) *User {
	user := &User{
		Model: utils.Model{
			UUID: in.UUID(),
		},
		Username: in.Username(),
		Password: in.Password(),
	}
	return user
}

// AddUser 往数据库写入user记录
func (u userRepository) AddUser(ctx context.Context, user *domain.User) (err error) {
	if err = u.data.DB(ctx).Model(&User{}).Create(u.marshalUser(user)).Error; err != nil {
		return
	}
	return
}

// NewUserRepository new user repository
func NewUserRepository(data *Data, tx domain.ITransaction) domain.IUserRepository {
	return &userRepository{
		tx:   tx,
		data: data,
	}
}
