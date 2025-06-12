package blog

type Service interface {
	CreateBlog()
	QueryBlog()
	DescribeBlog()
	UpdateBlog()
	PublishBlog()
	DeleteBlog()
}
