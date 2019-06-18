package models

type Tag struct {
	Model

	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

// func (tag *Tag) BeforeCreate(scope *gorm.Scope) error {
// 	scope.SetColumn("CreatedOn", time.Now().Unix())

// 	return nil
// }

// func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error {
// 	scope.SetColumn("ModifiedOn", time.Now().Unix())

// 	return nil
// }

func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	return
}

func GetTagTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)
	return
}

func ExistTagByName(name string) bool {
	var tag Tag
	db.Select("id").Where("name=?", name).First(&tag)

	if tag.ID > 0 {
		return true
	}
	return false
}

func ExistTagById(id int) bool {
	var tag Tag
	db.Select("id").Where("id=?", id).First(&tag)

	if tag.ID > 0 {
		return true
	}
	return false
}

func AddTag(name string, state int, createBy string) bool {
	db.Create(&Tag{
		Name:      name,
		State:     state,
		CreatedBy: createBy,
	})
	return true
}

func EditTag(id int, name string, state int, modifiedBy string) {
	db.Model(&Tag{}).Where("id = ?", id).Update(&Tag{Name: name, State: state, ModifiedBy: modifiedBy})
}

func DeleteTag(id int) {
	tag := &Tag{}
	tag.ID = id
	db.Delete(&tag)
}

func CleanAllTag() bool {
	db.Unscoped().Where("deleted_on != ? ", 0).Delete(&Tag{})

	return true
}
