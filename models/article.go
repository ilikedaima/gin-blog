package models

type Article struct {
	Model

	TagID         int    `json:"tag_id" gorm:"index"`
	Tag           Tag    `json:"tag"`
	CoverImageUrl string `json:"cover_image_url"`
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CreatedBy     string `json:"created_by"`
	ModifiedBy    string `json:"modified_by"`
	State         int    `json:"state"`
}

func ExistArticleByID(id int) bool {
	var article Article
	db.Select("id").Where("id = ?", id).First(&article)

	return article.ID > 0
}

func GetArticleTotal(maps interface{}) (count int64) {
	db.Model(&Article{}).Where(maps).Count(&count)
	return
}

func GetArticles(pageNum int, pageSize int, maps interface{}) (articles []Article) {
	db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)
	return
}
func GetArticle(id int) (article Article) {
	db.Preload("Tag").Where("id = ?", id).First(&article)
	// db.Where("id = ?", id).First(&article)

	return
}

func EditArticle(id int, data interface{}) bool {
	db.Model(&Article{}).Where("id = ?", id).Updates(data)
	return true
}

func AddArticle(data map[string]interface{}) bool {
	db.Create(&Article{
		TagID:         data["tag_id"].(int),
		Title:         data["title"].(string),
		Desc:          data["desc"].(string),
		Content:       data["content"].(string),
		CreatedBy:     data["created_by"].(string),
		State:         data["state"].(int),
		CoverImageUrl: data["cover_image_url"].(string),
	})

	return true
}
func DeleteArticle(id int) bool {
	db.Where("id = ?", id).Delete(Article{})

	return true
}

// func (article *Article) BeforeCreate(tx *gorm.DB) error {
//     tx.Statement.SetColumn("CreatedOn", time.Now().Unix())

//     return nil
// }

// func (article *Article) BeforeUpdate(tx *gorm.DB) error {
//     tx.Statement.SetColumn("ModifiedOn", time.Now().Unix())

//     return nil
// }