package tables

import (
	"BlogApp/config"
	"BlogApp/models"
	"fmt"
	"log"
)

func CreatePostTable() {
	query := `
	CREATE TABLE IF NOT EXISTS posts (
		id SERIAL PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
		content TEXT NOT NULL,
		author VARCHAR(255) NOT NULL,
		email VARCHAR(255) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`
	_, err := config.DB.Exec(query)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Posts table created successfully!")
}

func GetPostByID(id int) (models.Post, error) {	
	var post models.Post
	query := `SELECT id, title, content, author, email, created_at, updated_at FROM posts WHERE id = $1`
	err := config.DB.QueryRow(query, id).Scan(&post.ID, &post.Title, &post.Content, &post.Author, &post.Email, &post.CreatedAt, &post.UpdatedAt)

	if err != nil {
		return post, err
	}

	return post, nil
}

func UpdatePost(post models.Post) error {
	query := `
	UPDATE posts 
	SET title = $1,
	    content = $2,
	    updated_at = CURRENT_TIMESTAMP
	WHERE id = $3
	`

	_, err := config.DB.Exec(query,
		post.Title,
		post.Content,
		post.ID,
	)

	return err
}

func DeletePost(id int) error {
	query := `DELETE FROM posts WHERE id = $1`
	_, err := config.DB.Exec(query, id)

	return err
}

func GetAllPosts(limit, offset int, sortBy, sortOrder string) ([]models.Post, error) {

	// whitelist sortable columns
	validColumns := map[string]bool{
		"title":      true,
		"author":     true,
		"created_at": true,
	}

	if !validColumns[sortBy] {
		sortBy = "created_at"
	}

	if sortOrder != "asc" {
		sortOrder = "desc"
	}

	query := fmt.Sprintf(`
		SELECT id, title, content, author, email, created_at, updated_at
		FROM posts
		ORDER BY %s %s
		LIMIT $1 OFFSET $2
	`, sortBy, sortOrder)

	rows, err := config.DB.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var posts []models.Post

	for rows.Next() {
		var post models.Post

		err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.Author,
			&post.Email,
			&post.CreatedAt,
			&post.UpdatedAt,
		)

		if err != nil {
			log.Println("Scan error:", err)
			continue
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func GetPostsCount() (int, error) {
	var count int
	query := `SELECT COUNT(*) FROM posts`
	err := config.DB.QueryRow(query).Scan(&count)

	if err != nil {
		return 0, err
	}

	return count, nil
}