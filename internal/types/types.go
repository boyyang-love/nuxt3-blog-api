// Code generated by goctl. DO NOT EDIT.
package types

type ListBlogReq struct {
	Page  int `form:"page"`
	Limit int `form:"limit"`
}

type ListBlogRes struct {
	Name string `json:"name"`
}
