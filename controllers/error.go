package controllers

type ErrorController struct {
	BaseController
}

func (c *ErrorController) Error404() {
	c.Data["content"] = "page not found"
	c.TplName = "404.html"
}
