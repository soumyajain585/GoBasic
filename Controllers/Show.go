package Controllers

import (
	"encoding/json"
	"root/Models"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func GetHome(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello world",
	})

	successResponse(c, "Good Job", map[string]interface{}{})
}

func ConnectionGet(c *gin.Context) {

	var state []Models.State

	var db, _ = gormDB()
	defer db.Close()

	result := db.Find(&state)
	defer result.Close()

	c.JSON(200, gin.H{
		"data": result,
	})

	/*gorp

	var dbmap = DbConnection()
	defer dbmap.Db.Close()

	result, err := dbmap.Query("SELECT * FROM state")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var state Models.State
		err = result.Scan(&state.Id, &state.State)
		if err != nil {
			panic(err.Error())
		}

		c.JSON(200, gin.H{
			"id":    state.Id,
			"state": state.State,
		})

	}  */
}

func ConnectionGetById(c *gin.Context) {

	var state Models.State

	var db, _ = gormDB()
	defer db.Close()
	json.NewDecoder(c.Request.Body).Decode(&state)

	if state.Id != 0 {
		result := db.Where("id = ?", state.Id).Find(&state)
		defer result.Close()

		c.JSON(200, gin.H{
			"data": result,
		})

	} else {
		NoDataFoundResponse(c, "No Data Found")
	}

}
