package user

import (
	"errors"
	"reflect"

	"server/domain/user/entity"
	"server/domain/user/entity/valueobj"
	"server/domain/user/repository"
	"server/infrastructure/common/random"
	"server/infrastructure/persistence/user/converter"
	"server/infrastructure/persistence/user/po"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &UserRepositoryImpl{db: db}
}

func (ur *UserRepositoryImpl) SaveUser(user *entity.User) (*entity.User, error) {
	poUser := converter.UserEntityToPO(user)
	result := ur.db.Create(poUser)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("unable to save user")
	}

	entityUser := converter.UserPOToEntity(poUser)

	return entityUser, nil
}

func (ur *UserRepositoryImpl) DeleteUser(id int64) error {
	tx := ur.db.Begin()

	// 软删除用户
	if err := tx.Delete(&po.User{}, id).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 真删除关联关系表
	if err := tx.Exec("DELETE FROM t_rbac_user_roles WHERE user_id = ?", id).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (ur *UserRepositoryImpl) UpdateUser(user *entity.User) (*entity.User, error) {
	poUser := converter.UserEntityToPO(user)

	userValue := reflect.ValueOf(user).Elem()
	for i := 0; i < userValue.NumField(); i++ {
		fieldName := userValue.Type().Field(i).Name
		fieldValue := userValue.Field(i)

		if !fieldValue.IsZero() {
			result := ur.db.Model(&poUser).Where("id = ?", poUser.ID).Update(fieldName, fieldValue.Interface())
			if result.Error != nil {
				return nil, result.Error
			}
		}
	}

	entityUser := converter.UserPOToEntity(poUser)

	return entityUser, nil
}

func (ur *UserRepositoryImpl) UpdatesUser(user *entity.User) (*entity.User, error) {
	poUser := converter.UserEntityToPO(user)
	result := ur.db.Model(poUser).Where("id = ?", poUser.ID).Updates(*poUser)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("unable to update user")
	}

	entityUser := converter.UserPOToEntity(poUser)

	return entityUser, nil
}

func (ur *UserRepositoryImpl) GetUser(id int64) (*entity.User, error) {
	user := &po.User{}

	result := ur.db.First(user, id)
	if result.Error != nil {
		return nil, result.Error
	}

	entityUser := converter.UserPOToEntity(user)

	return entityUser, nil
}

func (ur *UserRepositoryImpl) GetUsers() ([]entity.User, error) {
	users := make([]po.User, 0)

	result := ur.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	entityUsers := make([]entity.User, 0, len(users))
	for _, u := range users {
		entityUsers = append(entityUsers, *converter.UserPOToEntity(&u))
	}

	return entityUsers, nil
}

func (ur *UserRepositoryImpl) FindByName(name string) (*entity.User, error) {
	user := &po.User{}

	result := ur.db.Where("name = ?", name).First(user)
	if result.Error != nil {
		return nil, result.Error
	}

	entityUser := converter.UserPOToEntity(user)

	return entityUser, nil
}

func (ur *UserRepositoryImpl) FindByEmpID(empId int) (*entity.User, error) {
	user := &po.User{}

	result := ur.db.Where("employee_id = ?", empId).First(user)
	if result.Error != nil {
		return nil, result.Error
	}

	entityUser := converter.UserPOToEntity(user)

	return entityUser, nil
}

func (ur *UserRepositoryImpl) FindByEmail(email string) (*entity.User, error) {
	user := &po.User{}

	result := ur.db.Where("email = ?", email).First(user)
	if result.Error != nil {
		return nil, result.Error
	}

	entityUser := converter.UserPOToEntity(user)

	return entityUser, nil
}

func (ur *UserRepositoryImpl) FindByUsername(username string) (*entity.User, error) {
	user := &po.User{}

	// manytomany 需要 preload 预加载才可以拿到关联表信息
	result := ur.db.Preload("Roles").Where("username = ?", username).First(user)
	if result.Error != nil {
		return nil, result.Error
	}

	entityUser := converter.UserPOToEntity(user)

	return entityUser, nil
}

func (ur *UserRepositoryImpl) FindByUsernameAndPassword(username, password string) (*entity.User, error) {
	user := &po.User{}

	result := ur.db.Preload("Roles").Where("username = ? AND password = ?", username, password).First(user)
	if result.Error != nil {
		return nil, result.Error
	}

	entityUser := converter.UserPOToEntity(user)

	return entityUser, nil
}

func (ur *UserRepositoryImpl) UpdatePassword(id int64, password ...string) (string, error) {
	var newPassword string
	if len(password) == 0 {
		newPassword = random.GeneratePassword(valueobj.PasswordDigits)
	} else {
		newPassword = password[0]
	}

	user := &po.User{}
	err := ur.db.First(user, id).Error
	if err != nil {
		return "", err
	}

	user.Password = newPassword

	err = ur.db.Save(user).Error
	if err != nil {
		return "", err
	}

	return newPassword, nil
}
