package handlers

import (
	"backend/database"

	"github.com/gin-gonic/gin"
)

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

func ViewDatabase(gctx *gin.Context) {
	var albums []Album
	rows, _ := database.DB.Query("SELECT * FROM album")
	defer rows.Close()
	for rows.Next() {
		var alb Album
		rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price)
		albums = append(albums, alb)
	}
	rows.Err()

	gctx.JSON(200, albums)
}
