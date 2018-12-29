package controllers

type IndexController struct {
	BaseController
}

//首页
// @router / [get]
func (c *IndexController) Get() {
	c.TplName = "index.html"
}

//留言
// @router /message [get]
func (c *IndexController) GetMessage() {
	c.TplName = "message.html"
}

//关于
// @router /about [get]
func (c *IndexController) GetAbout() {
	c.TplName = "about.html"
}

// @router / [get]
func (this *IndexController) Get() {
	//每页显示10条数据
	limit := 10
	// 得到页面传过来的参数，没有就默认为1
	page, err := this.GetInt("page", 1)
	if err != nil || page <= 0 {
		page = 1
	}
	title := this.GetString("title")
	//根据 当前页 和 每页显示的行数 得到文章列表数据集
	notes, err := models.QueryNotesByPage(title, page, limit)
	if err != nil {
		this.Abort500(err)
	}
	//将数据传到模版页面index.html，等待渲染
	this.Data["notes"] = notes
	//得到文章的总行数
	count, err := models.QueryNotesCount(title)
	if err != nil {
		this.Abort500(err)
	}
	//这儿计算总页数，如果“文章的总数量”不是“每页显示的行数”的倍数，就要多显示一页
	totpage := count / limit
	if count%limit != 0 {
		//取余数，不为0。那就要多加一页显示这些数据
		totpage = totpage + 1
	}
	// 将总页数 当前页 传到模版页面。等待渲染
	this.Data["totpage"] = totpage
	this.Data["page"] = page
	this.Data["title"] = title
	this.TplName = "index.html"
}

//@router /details/:key [get]
func (this *IndexController) GetDetails() {
	//得到页面传过来的文章的key
	key := this.Ctx.Input.Param(":key")
	/// 查询文章评论
	ms, err := models.QueryMessagesByNoteKey(key)
	if err != nil {
		//查询出错提示“文章不存在”
		this.Abort500(syserror.New("文章不存在", err))
	}
	this.Data["messages"] = ms
	this.TplName = "details.html"
}

//@router /comment/:key
func (this *IndexController) GetComment() {
	//得到页面的key
	key := this.Ctx.Input.Param(":key")
	//根据可以 从数据库中查询出文章
	note, err := models.QueryNoteByKey(key)
	// 如果 查询时候报错，就提示文章不存在，并将错误原因打印到日志。
	if err != nil {
		this.Abort500(syserror.New("文章不存在", err))
	}
	this.Data["note"] = note
	this.TplName = "comment.html"
}
