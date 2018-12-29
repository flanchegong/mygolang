package models

type Note struct {
	Model          //这儿的Model是我们上面定义的Model不是gorm.Model
	Key     string `gorm:"unique_index;not null;"` //文章唯一标示
	UserID  int    // 用户id
	User    User   //用户
	Title   string //文章标题
	Summary string `gorm:"type:text"` //概要
	Content string `gorm:"type:text"` //文章内容
	Visit   int    `gorm:"default:0"` //浏览次数
	Praise  int    `gorm:"default:0"` // 点赞次数
}

func QueryNotesCount(title string) (count int, err error) {
	//同QueryNotesByPage的修改一致
	return count, db.Model(&Note{}).Where("title like ? ", fmt.Sprintf("%%%s%%", title)).Count(&count).Error
}

func QueryNotesByPage(title string, page, limit int) (note []*Note, err error) {
	// 模糊匹配 title字段
	//.Where("title like ? ",fmt.Sprintf("%%%s%%",title))
	return note, db.Where("title like ? ", fmt.Sprintf("%%%s%%", title)).Offset((page - 1) * limit).Limit(limit).Find(&note).Error
}
func QueryNoteByKeyAndUserId(key string, userid int) (note Note, err error) {
	return note, db.Where("key = ? and user_id = ? ", key, userid).Take(&note).Error
}
func DeleteNoteByUserIdAndKey(key string, userid int) error {
	return db.Delete(&Note{}, "key = ? and user_id = ?", key, userid).Error
}
