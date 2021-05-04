package db

import "fmt"

//LikePost likes or dislikes post based on value of like, true is like, false is dislike
func LikePost(userid int, postid int, like bool) bool {
	db, err := ConnectDb()

	if err != nil {
		db.Close()
		fmt.Println(err)
		return false
	}

	if like {
		userliked := GetUserLikes(userid, postid)

		if len(userliked) != 0 && like {
			return true
		}

		userDisliked := GetUserDislikes(userid, postid)
		if len(userDisliked) != 0 {
			_, err := db.Exec("update likes set like=true where userid=$1 and postid=$2", userid, postid)
			if err != nil {
				fmt.Print("liking disliked post err!: ")
				fmt.Println(err)
				return false
			}
			return true
		}

		likes, err := db.Prepare("insert into likes (userid, postid, like) values (?, ?, ?)")

		if err != nil {
			fmt.Print("liking err!: ")
			fmt.Println(err)
			return false
		}

		likes.Exec(userid, postid, like)
	} else {
		dislikes := GetUserDislikes(userid, postid)
		if len(dislikes) != 0 {
			return true
		}

		userlikes := GetUserLikes(userid, postid)
		if len(userlikes) != 0 {
			_, err := db.Exec("update likes set like=false where userid=$1 and postid=$2", userid, postid)
			if err != nil {
				fmt.Print("disliking liked post err!: ")
				fmt.Println(err)
				return false
			}
			return true
		}

		likes, err := db.Prepare("insert into likes (userid, postid, like) values (?, ?, ?)")

		if err != nil {
			fmt.Println("like insertion err")
			fmt.Println(err)
			return false
		}

		likes.Exec(userid, postid, false)
	}

	db.Close()
	return true
}

//GetUserLikes returns the amount of likes to the post, returns [] or [userid]
func GetUserLikes(postid int, userid int) []int {
	var likes []int
	db, err := ConnectDb()
	if err != nil {
		db.Close()
		return likes
	}

	rows, err := db.Query("select userid from likes where userid=$1 and postid=$2 and like=$3", postid, userid, true)

	if err != nil {
		fmt.Print("like get err!: ")
		fmt.Println(err)
		db.Close()
		return likes
	}

	defer rows.Close()

	for rows.Next() {
		var p int
		rows.Scan(&p)
		likes = append(likes, p)
	}
	db.Close()
	return likes
}

//GetUserDislikes returns the slice of dislikes returns [] or [userid]
func GetUserDislikes(postid int, userid int) []int {
	var likes []int
	db, err := ConnectDb()
	if err != nil {
		db.Close()
		return likes
	}

	rows, err := db.Query("select userid from likes where userid=$1 and postid=$2 and like=$3", postid, userid, false)

	if err != nil {
		fmt.Print("like get err!: ")
		fmt.Println(err)
		db.Close()
		return likes
	}
	defer rows.Close()

	for rows.Next() {
		var p int
		rows.Scan(&p)
		likes = append(likes, p)
	}

	db.Close()
	return likes
}

//GetLikesByPostID returns the list of userids that liked the post
func GetLikesByPostID(postid int) []int {
	var likes []int
	db, err := ConnectDb()
	if err != nil {
		db.Close()
		return likes
	}

	rows, err := db.Query("select userid from likes where postid=$1 and like=$2", postid, true)

	if err != nil {
		fmt.Print("like get err!: ")
		fmt.Println(err)
		db.Close()
		return likes
	}

	defer rows.Close()

	for rows.Next() {
		var p int
		rows.Scan(&p)
		likes = append(likes, p)
	}

	db.Close()
	return likes
}

//GetDislikesByPostID returns the list of userids that disliked the post
func GetDislikesByPostID(postid int) []int {
	var likes []int
	db, err := ConnectDb()
	if err != nil {
		db.Close()
		return likes
	}

	rows, err := db.Query("select userid from likes where postid=$1 and like=$2", postid, false)

	if err != nil {
		fmt.Print("dislike get err!: ")
		fmt.Println(err)
		db.Close()
		return likes
	}

	defer rows.Close()

	for rows.Next() {
		var p int
		rows.Scan(&p)
		likes = append(likes, p)
	}

	db.Close()
	return likes
}
