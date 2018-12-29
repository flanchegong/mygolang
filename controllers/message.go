package controllers

type Message struct {
	Model
	Key     string `grom:"unique_index; not null"json:"key"` //评论的key唯一标示
	Content string `json:"content"`                          //评论的内容
	UserId  int    `json:"user_id"`                          //评论人id
	User    User   `json:"user"`                             //评论人
	NoteKey string `json:"note_key"`                         //所属文章的key，可以为空，为空代表系统留言
	Praise  int    `gorm:"default:0" json:"praise"`          //点赞数量
}

// 评论处理的控制器
type MessageController struct {
	BaseController
}

// @router /new/?:key [post]
func (ctx *MessageController) NewMessage() {
	ctx.MustLogin()
	key := ctx.UUID()
	content := ctx.GetMustString("content", "内容不能为空")
	notekey := ctx.Ctx.Input.Param(":key")
	m := &models.Message{
		UserID:  int(ctx.User.ID),
		User:    ctx.User,
		Key:     key,
		NoteKey: notekey,
		Content: content,
	}
	if err := ctx.Dao.SaveMessage(m); err != nil {
		ctx.Abort500(syserrors.NewError("保存失败！", err))
	}
	ctx.JSONOkH("保存成功！", H{
		"data": m,
	})
}

//@router /count [get]
func (this *MessageController) Count() {
	// 查询 留言的总数量
	count, err := models.QueryMessagesCountByNoteKey("")
	if err != nil {
		//查询报错，提示页面失败，并打印错误日志
		this.Abort500(syserror.New("查询失败", err))
	}
	// 将 留言的总数量 返回前台
	this.JSONOkH("查询成功", H{"count": count})
}

//@router /query [get]
func (this *MessageController) Query() {
	//获得第几页，默认第一页
	pageno, err := this.GetInt("pageno", 1)
	if err != nil || pageno < 1 {
		pageno = 1
	}
	//获得每页显示多少条数据，默认10条
	pagesize, err := this.GetInt("pagesize", 10)
	if err != nil {
		pagesize = 10
	}
	//调用数据库方法，查询出留言的数据集
	ms, err := models.QueryPageMessagesByNoteKey("", pageno, pagesize)
	if err != nil {
		//查询报错，提示页面失败，并打印错误日志
		this.Abort500(syserror.New("查询失败", err))
	}
	this.JSONOkH("查询成功", H{"data": ms})
}
