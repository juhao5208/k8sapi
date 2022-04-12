package services

/**
 * @author  巨昊
 * @date  2021/11/11 15:15
 * @version 1.15.3
 */

//@Service
type ImageService struct {
	ImageMap *ImageMap       `inject:"-"`
	Common  *CommonService `inject:"-"`
}

func NewImageService() *ImageService {
	return &ImageService{}
}