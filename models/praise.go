package models

type PraiseLog struct {
	Model
	UserId int    //点赞用户id
	Key    string //文章或评论的key
	Table  string // 点赞的表名（评论为messages，文章为notes）
	Flag   bool   //是否点赞
}

//说明：文章表和评论表都有key和praise，他们的key是唯一的，所以我们可以根据表名和key开确定是哪条记录，做具体的更新点赞数量的操作。
//这个结果体是为了方便查询 已经点赞的数量
type TempPraise struct {
	Praise int
}

//核心方法
func UpdatePraise(table, key string, userid int) (pcnt int, err error) {
	//开启事务
	d := db.Begin()
	//判断如果函数返回错误不为空，就事务回滚。
	defer func() {
		if err != nil {
			//回滚事务
			d.Rollback()
		} else {
			//提交事务
			d.Commit()
		}
	}()
	//查询点赞流水表，看是否有记录，并赋值给p，
	//如果有记录，我们判断下Flag是否为true，如果为true就是点赞，就提示已经点赞的错误。
	//如果没有，我们就重新赋值一个flag为false的点赞流水，赋值给p
	var p PraiseLog
	err = d.Model(&PraiseLog{}).Where("`key` = ? and `table` =? and user_id =? ", key, table, userid).Take(&p).Error
	if err == gorm.ErrRecordNotFound {
		// 如果查询不到数据 我们就赋值 Flag为false的点赞流水给p
		p = PraiseLog{
			Key:    key,
			Table:  table,
			UserId: userid,
			Flag:   false,
		}
	} else if err != nil {
		// 如果其他的错误，就返回错误
		return 0, err
	}
	// 如果flag为true，说明已经点赞过，我们就提示已经点赞的错误
	if p.Flag {
		// HasPraiseError是我们定义的错误类型，code为4444，代表已经点赞
		return 0, syserror.HasPraiseError{}
	}
	//更新点赞，为true。
	p.Flag = true
	//保存 流水
	if err = d.Save(&p).Error; err != nil {
		return 0, err
	}
	//更新文章或留言表的点赞数量
	var ppp TempPraise
	err = d.Table(table).Where("key = ?", key).Select("praise").Scan(&ppp).Error
	if err != nil {
		return 0, err
	}
	pcnt = ppp.Praise + 1
	if err = d.Table(table).Where("key = ? ", key).Update("praise", pcnt).Error; err != nil {
		return 0, err
	}
	// 返回点赞数量
	return pcnt, nil
}
