package Controllers

import (
	"encoding/json"
	"root/Models"

	"github.com/gin-gonic/gin"
)

func ConnectionDelete(c *gin.Context) {
	var state Models.State

	var db, _ = gormDB()
	defer db.Close()

	json.NewDecoder(c.Request.Body).Decode(&state)

	if state.Id != 0 {
		result := db.Where("id = ?", state.Id).Delete(&state)
		defer result.Close()
		successResponse(c, "Successfully deleted", map[string]interface{}{"id": state.Id})

	} else {
		NoDataFoundResponse(c, "No Data Found")
	}

	/*gorp
	var dbmap = DbConnection()
	defer dbmap.Db.Close()

	c.Bind(&state)
	delForm, err := dbmap.Query("DELETE FROM state WHERE id=?", state.Id)
	if err != nil {
		panic(err.Error())
	}

	// c.JSON(http.StatusOK, gin.H{
	// 	"message":    fmt.Sprintf("%d Id Successfully deleted", state.Id),
	// 	"body":       state.Id,
	// 	"statusCode": http.StatusOK,
	// })

	successResponse(c, "Successfully de;leted", map[string]interface{}{"id": state.Id})
	defer delForm.Close()

	*/

}
