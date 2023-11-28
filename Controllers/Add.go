package Controllers

import (
	"encoding/json"
	"root/Models"

	"github.com/gin-gonic/gin"
)

func ConnectionPost(c *gin.Context) {
	var state Models.State

	var db, _ = gormDB()
	defer db.Close()

	json.NewDecoder(c.Request.Body).Decode(&state)

	if state.State != "" {
		result := db.Create(&state)
		defer result.Close()
		successResponse(c, "Successfully created", map[string]interface{}{"state": state.State})

	} else {
		NoDataFoundResponse(c, "No Data Found")
	}

	/*gorp

	var dbmap = DbConnection()
	defer dbmap.Db.Close()

	c.Bind(&state)
	if state.State != "" {
	result := db.Raw(`INSERT INTO state (state) VALUES (?)`, state.State)
	fmt.Println("s", result, "error", result.Error)

		if err != nil {
			panic(err.Error())
		}

	c.JSON(http.StatusOK, gin.H{
		"message":    fmt.Sprintf("%s Successfully created", state.State),
		"body":       state.State,
		"statusCode": http.StatusOK,
	})
	defer result.Close() */

}
