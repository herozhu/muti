package models

import "github.com/jinzhu/gorm"

type Tag struct {
	Model

	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

var (
	//切换本地数据库连接时使用
	//db = GetSelfDB()
	db = GetDockerDB()
)

func GetTags(pageNum, pageSize int, maps interface{}) (tags []Tag) {
	//用于本地数据库和docker数据库切换
	//DB.Self.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	DB.docker.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	return
}

func GetTagTotal(maps interface{}) (count int) {
	//用于本地数据库和docker数据库切换
	//DB.Self.Model(&Tag{}).Where(maps).Count(&count)
	DB.docker.Model(&Tag{}).Where(maps).Count(&count)
	return
}

func ExistTagByName(name string) (bool, error) {
	var tag Tag
	err := db.Select("id").Where("name = ? AND deleted_on = ? ", name, 0).First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err

	}
	if tag.ID > 0 {
		return true, nil
	}

	return false, nil
}

func ExistTagByID(id int) (bool, error) {
	var tag Tag
	err := db.Select("id").Where("id = ? AND deleted_on = ? ", id, 0).First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if tag.ID > 0 {
		return true, nil
	}

	return false, nil
}

func CreateTag(name string, state int, createdBy string) error {
	tag := Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	}
	if err := db.Create(&tag).Error; err != nil {
		return err
	}

	return nil
}

func DeleteTag(id int) error {
	if err := db.Where("id = ? ", id).Delete(&Tag{}).Error; err != nil {
		return err
	}
	return nil
}

func EditTag(id int, data interface{}) error {
	if err := db.Model(&Tag{}).Where("id = ? AND deleted_on = ?", id, 0).Update(data).Error; err != nil {
		return err
	}
	return nil
}

func CleanAllTag() (bool, error) {
	if err := db.Unscoped().Where("deleted_on = ?", 0).Delete(&Tag{}).Error; err != nil {
		return false, err
	}
	return true, nil
}
