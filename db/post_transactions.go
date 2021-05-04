package db

import (
	"fmt"

	"forum_v2.0/model"
)

//InsertPost inserts post into the database
func InsertPost(post model.Post) error {
	db, err := ConnectDb()
	if err != nil {
		db.Close()
		return err
	}
	p, err := db.Prepare("insert into posts(userid, title, text, category, image, date) values (?, ?, ?, ?, ?, ?)")

	if err != nil {
		db.Close()
		return err
	}

	p.Exec(post.UserID, post.Title, post.Text, post.Category, post.Image, post.Date)

	db.Close()
	return err
}

//GetAllPosts returns the list of posts
func GetAllPosts() []model.Post {
	var posts []model.Post
	db, err := ConnectDb()
	if err != nil {
		db.Close()
		return posts
	}

	rows, err := db.Query("select rowid, userid, title, text, category, image, date from posts")

	if err != nil {
		fmt.Print("post insertion err!: ")
		fmt.Println(err)
		db.Close()
		return posts
	}
	defer rows.Close()

	for rows.Next() {
		var p model.Post
		rows.Scan(&p.PostID, &p.UserID, &p.Title, &p.Text, &p.Category, &p.Image, &p.Date)
		p.Comments = GetCommentsByPostID(p.PostID)
		p.Likes = GetLikesByPostID(p.PostID)
		p.Dislikes = GetDislikesByPostID(p.PostID)
		posts = append(posts, p)
	}

	db.Close()
	return posts
}

//GetPostsByUserID returns the list of posts written by userid
func GetPostsByUserID(userid int) []model.Post {
	var posts []model.Post
	db, err := ConnectDb()
	if err != nil {
		db.Close()
		return posts
	}

	rows, err := db.Query("select rowid, userid, title, text, category, image, date from posts where userid=$1", userid)

	if err != nil {
		fmt.Print("post insertion err!: ")
		fmt.Println(err)
		db.Close()
		return posts
	}
	defer rows.Close()

	for rows.Next() {
		var p model.Post
		rows.Scan(&p.PostID, &p.UserID, &p.Title, &p.Text, &p.Category, &p.Image, &p.Date)
		p.Comments = GetCommentsByPostID(p.PostID)
		p.Likes = GetLikesByPostID(p.PostID)
		p.Dislikes = GetDislikesByPostID(p.PostID)
		posts = append(posts, p)
	}

	db.Close()
	return posts
}

//GetPostByPostID returns the post from its id
func GetPostByPostID(postid int) model.Post {
	var p model.Post
	db, err := ConnectDb()
	if err != nil {
		db.Close()
		return p
	}

	rows, err := db.Query("select rowid, userid, title, text, category, image, date from posts where rowid=$1", postid)

	if err != nil {
		fmt.Print("post insertion err!: ")
		fmt.Println(err)
		db.Close()
		return p
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&p.PostID, &p.UserID, &p.Title, &p.Text, &p.Category, &p.Image, &p.Date)
		break
	}

	db.Close()
	return p
}
