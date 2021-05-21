package controllers

type ProductController struct {
	usecase interface{}
}

func NewProductController() error {
	handler := &ProductController{}
}
