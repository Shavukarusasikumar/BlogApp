package tables

import (
	"fmt"
	"BlogApp/config"
	"log"
	"BlogApp/models"
)

func CreateCommentTable() {
	query := `
	CREATE TABLE IF NOT EXISTS comments (
		id SERIAL PRIMARY KEY,
		post_id INT NOT NULL,
		author VARCHAR(255) NOT NULL,
		email VARCHAR(255) NOT NULL,
		content TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE
	);
	`
	_, err := config.DB.Exec(query)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Comments table created successfully!")
}

func InsertComment(Comment models.Comment) (int, error) {
	var id int
	query := `INSERT INTO comments (post_id, author, email, content, created_at, updated_at) VALUES ($1, $2, $3, $4, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP) RETURNING id`
	err := config.DB.QueryRow(query, Comment.PostID, Comment.Author, Comment.Email, Comment.Content).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func GetCommentsByPostID(postID int) ([]models.Comment, error) {
	query := `SELECT id, post_id, author, email, content, created_at, updated_at FROM comments WHERE post_id = $1`
	rows, err := config.DB.Query(query, postID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	comments:= []models.Comment{}
	for rows.Next() {
		var comment models.Comment
		err := rows.Scan(&comment.ID, &comment.PostID, &comment.Author, &comment.Email, &comment.Content, &comment.CreatedAt, &comment.UpdatedAt)
		
		if err != nil {
			log.Printf("Error scanning comment: %v", err)
			continue
		}
		comments = append(comments, comment)
	}

	return comments, nil
}