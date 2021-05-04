package model

// Post just post
type Post struct {
	PostID   int
	UserID   int
	Title    string
	Text     string
	Category string
	Image    string
	Date     string
	Comments []Comment
	Likes    []int
	Dislikes []int
}

//Validate used to check if post is valid
func (post Post) Validate() string {
	if post.Title == "" {
		return "title cannot be empty!"
	}

	if post.Text == "" {
		return "text of the post cannot be empty"
	}

	if post.Category == "" {
		return "select category!"
	}

	return ""
}
