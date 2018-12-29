package models

func SaveMessage(message *Message) error {
	return db.Save(message).Error
}
func QueryMessagesByNoteKey(notekey string) (ms []*Message, err error) {
	//Preload("User") 这是预加载，关联出与评论表相关联的用户信息，显示评论需要用户的信息。
	return ms, db.Preload("User").Where("note_key = ? ", notekey).Order("updated_at desc").Find(&ms).Error
}
func QueryMessagesCountByNoteKey(notekey string) (count int, err error) {
	return count, db.Model(&Message{}).Where("note_key = ? ", notekey).Count(&count).Error
}
func QueryPageMessagesByNoteKey(notekey string, pageno, pagesize int) (ms []*Message, err error) {
	return ms, db.Preload("User").Where("note_key = ? ", notekey).Offset((pageno - 1) * pagesize).Limit(pagesize).Order("updated_at desc").Find(&ms).Error
}
