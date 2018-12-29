package models

//type Model struct {
//	ID        int        `gorm:"primary_key"`
//	CreatedAt time.Time
//	UpdatedAt time.Time
//	DeletedAt *time.Time `sql:"index"`
//}

//用户表
type User struct {
	Model         //这儿的Model是我们上面定义的Model不是gorm.Model
	Name   string `gorm:"unique_index" json:"name"`
	Email  string `gorm:"unique_index" json:"email"`
	Pwd    string `json:"-"` //“-” 代表不输出，密码不能输出到页面。
	Avatar string `json:"avatar"`
	Role   int    `json:"role"`
}

func (db *DB) QueryUserByEmailAndPassword(email, password string) (user User, err error) {
	return user, db.db.Model(&User{}).Where("email = ? and pwd = ?", email, password).Take(&user).Error
}

func (db *DB) QueryUserByName(name string) (user User, err error) {
	return user, db.db.Where("name = ?", name).Take(&user).Error
}

func (db *DB) QueryUserByEmail(email string) (user User, err error) {
	return user, db.db.Where("email = ?", email).Take(&user).Error
}

func (db *DB) UpdateUserEditor(editor string) (err error) {
	return db.db.Model(&User{}).Update("editor", editor).Error
}

func SaveUser(user *User) error {
	return db.Create(user).Error
}
