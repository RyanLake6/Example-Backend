package handlers

import (
	"github.com/gin-gonic/gin"
)

type User struct {
	ID    int64
	Name  string
	Phone string
}

func ViewDatabase(gctx *gin.Context) {
	// var users []User
	// rows, err := database.DB.Query("SELECT * FROM User")
	// if err != nil {
	// 	log.Fatal("Query failed:", err)
	// 	//gctx.Status(http.StatusInternalServerError)
	// 	//return
	// }
	// defer rows.Close()
	// for rows.Next() {
	// 	var alb User
	// 	rows.Scan(&alb.ID, &alb.Name, &alb.Phone)
	// 	users = append(users, alb)
	// }
	// rows.Err()

	// gctx.JSON(200, users)

	gctx.JSON(200, gin.H{
		"message": "Temporarily commented out this api endpoint",
	})
}
