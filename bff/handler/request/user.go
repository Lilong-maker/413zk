package request

type Login struct {
	Name     string `form:"name" binding:"required"`
	Password string `form:"password" binding:"required"`
}
type Register struct {
	Name     string `form:"name" binding:"required"`
	Password string `form:"password" binding:"required"`
}
