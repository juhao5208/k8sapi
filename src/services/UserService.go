package services

/**
 * @author  巨昊
 * @date  2021/11/4 16:18
 * @version 1.15.3
 */

//@Service
type UserService struct {
	UserMap *UserMap       `inject:"-"`
	Common  *CommonService `inject:"-"`
}

func NewUserService() *UserService {
	return &UserService{}
}

/*
func (this *UserService) ListAll() (ret []*models.Users) {
	userList := this.UserMap.GetUserList()
	for _, user := range userList {
		// 加入副本数
		ret = append(ret, &models.Users{
			UserId:   user.UserId,
			UUID:     user.UUID,
			Username: user.Username,
			Password: user.Password,
			UserType: user.UserType,
			UserTel:  user.UserTel,
			UserAdd:  user.UserAdd,
		})
	}
	return
}*/
