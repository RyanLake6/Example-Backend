package handlers

import (
	"backend/database"

	"github.com/gin-gonic/gin"
)

func GetDatabase(gctx *gin.Context) {
	// An albums slice to hold data from returned rows.
	var albums []Album

	name := "John Coltrane"

	rows, _ := database.DB.Query("SELECT * FROM album WHERE artist = ?", name)
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var alb Album
		rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price)
		albums = append(albums, alb)
	}
	rows.Err()

	gctx.JSON(200, albums)
}
