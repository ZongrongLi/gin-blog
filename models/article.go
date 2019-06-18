package models

type Article struct {
	Model

	TagID int `json:"tag_id" gorm:"index"`
	Tag   Tag `json:"tag"`

	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

// func (article *Article) BeforeCreate(scope *gorm.Scope) error {
// 	scope.SetColumn("CreatedOn", time.Now().Unix())

// 	return nil
// }

// func (article *Article) BeforeUpdate(scope *gorm.Scope) error {
// 	scope.SetColumn("ModifiedOn", time.Now().Unix())

// 	return nil
// }

func GetArticle(id int) (article Article) {
	db.Where("id = ?", id).First(&article)
	//根据这个实体（article.Tag） 去寻找这个实体的id（tag+ID =tagid）对应的值，然后根据tagid去做关联查询
	db.Model(&article).Related(&article.Tag)
	return
}

func GetArticles(pageNum int, pageSize int, maps interface{}) (articles []Article) {
	db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)

	return
}

func GetArticleTotal(maps interface{}) (count int) {
	db.Model(&Article{}).Where(maps).Count(&count)

	return
}

func ExistArticleByName(name string) bool {
	var article Article
	db.Select("id").Where("name=?", name).First(&article)

	if article.ID > 0 {
		return true
	}
	return false
}

func ExistArticleById(id int) bool {
	var article Article
	db.Select("id").Where("id=?", id).First(&article)

	if article.ID > 0 {
		return true
	}
	return false
}

func AddArticle(data map[string]interface{}) bool {
	db.Create(&Article{
		TagID:     data["tag_id"].(int),
		Title:     data["title"].(string),
		Desc:      data["desc"].(string),
		Content:   data["content"].(string),
		CreatedBy: data["created_by"].(string),
		State:     data["state"].(int),
	})
	return true
}

func EditArticle(id int, data interface{}) bool {
	db.Model(&Article{}).Where("id = ?", id).Updates(data)

	return true
}

func DeleteArticle(id int) {
	article := &Article{}

	article.ID = int(id)
	db.Delete(&article)
}
