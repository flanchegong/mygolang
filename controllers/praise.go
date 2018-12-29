package controller

//定义控制器 PraiseController
type PraiseController struct {
	BaseController
}

//实现 NextPrepare接口，每次请求都必须登陆，没有登陆提示错误
func (this *PraiseController) NextPrepare() {
	this.MustLogin()
}

//定义路由
//@router /:type/:key [post]
func (this *PraiseController) Praise() {
	// 获取页面传过来的type，用来区分是文章还是评论
	ttype := this.Ctx.Input.Param(":type")
	// 获取页面传过来的key，文章或评论的key
	key := this.Ctx.Input.Param(":key")

	// 定义table变量，
	table := "notes"
	// 根据ttype的不同，指定不同的表
	switch ttype {
	case "message":
		table = "messages"
	case "note":
		table = "notes"
	default:
		// 不是文章或评论，就是提示 “未知类型”错误
		this.Abort500(errors.New("未知类型"))
	}
	// 调用我们刚才定义的更新点赞的方法。
	pcnt, err := models.UpdatePraise(table, key, int(this.User.ID))
	if err != nil {
		//如果报错，我们得先判断是不是 已经点赞过的错误，如果是，我们放回点赞过的的错误
		if e2, ok := err.(syserror.HasPraiseError); ok {
			this.Abort500(e2)
		}
		//我们重新定义 “点赞失败”，将具体的错误原因显示在日志里面
		this.Abort500(syserror.New("点赞失败", err))
	}
	//点赞成功，返回点赞数量
	this.JSONOkH("点击成功", H{"praise": pcnt})
}
