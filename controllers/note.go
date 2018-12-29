package controllers

type NoteController struct {
	BaseController
}

// @router /new [get]
func (ctx *NoteController) NewPage() {
	ctx.Data["key"] = uuid.NewUUID().String()
	ctx.TplName = "note_new.html"
}
func (ctx *NoteController) NestPrepare() {
	ctx.MustLogin()         //用户必须登陆，没有登陆就返回错误
	if ctx.User.Role != 0 { //不是管理员，之前返回错误
		ctx.Abort500(syserrors.NewError("您没有权限修改文章", nil))
	}
}

// @router /edit/:key [get]
func (this *NoteController) EditPage() {
	// 获取页面传过来key
	key := this.Ctx.Input.Param(":key")
	//根据当前文章的key和登陆用户id查询出，文章信息。
	note, err := models.QueryNoteByKeyAndUserId(key, int(this.User.ID))
	if err != nil {
		//查询有问题，就提示文章不存在。
		this.Abort500(syserror.New("文章不存在！", err))
	}
	//将文章的信息以及key传到页面。
	this.Data["note"] = note
	this.Data["key"] = key
	//"note_new.html" 是文章新增的页面，之前实现文章新增功能的时候
	this.TplName = "note_new.html"
}

// @router /save/:key [post]
func (ctx *NoteController) Save() {
	key := ctx.Ctx.Input.Param(":key")
	editor := ctx.GetString("editor", "default")
	title := ctx.GetMustString("title", "标题不能为空！")
	content := ctx.GetMustString("content", "内容不能为空！")
	files := ctx.GetString("files", "")
	summary, _ := getSummary(content)
	note, err := ctx.Dao.QueryNoteByKeyAndUserId(key, int(ctx.User.ID))
	var n models.Note
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			ctx.Abort500(syserrors.NewError("保存失败！", err))
		}
		n = models.Note{
			Key:     key,
			Summary: summary,
			Title:   title,
			Files:   files,
			Content: content,
			UserID:  int(ctx.User.ID),
		}
	} else {
		n = note
		n.Title = title
		n.Content = content
		n.Summary = summary
		n.Files = files
		n.UpdatedAt = time.Now()
	}
	n.Editor = editor
	if strings.EqualFold(editor, "markdown") {
		n.Source = ctx.GetMustString("source", "内容不能为空！")
	}

	if err := ctx.Dao.SaveNote(&n); err != nil {
		ctx.Abort500(syserrors.NewError("保存失败！", err))
	}
	ctx.JSONOk("成功", "/details/"+key)
}

// @router /del/:key [post]
func (this *NoteController) Del() {
	//获取页面传过来的key值
	key := this.Ctx.Input.Param(":key")
	//调用数据库方法删除文章
	if err := models.DeleteNoteByUserIdAndKey(key, int(this.User.ID)); err != nil {
		//删除失败，提示页面删除失败
		this.Abort500(syserror.New("删除失败", err))
	}
	//返回删除成功，并跳转到首页
	this.JSONOk("删除成功", "/")
}
