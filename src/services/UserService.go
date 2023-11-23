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
