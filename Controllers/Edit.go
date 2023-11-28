package Controllers

import (
	"encoding/json"
	"root/Models"

	"github.com/gin-gonic/gin"
)

func ConnectionUpdate(c *gin.Context) {
	var state Models.State

	var db, _ = gormDB()
	defer db.Close()

	json.NewDecoder(c.Request.Body).Decode(&state)

	if state.State != "" {

		result := db.Where("id = ?", state.Id).Update(&state)
		defer result.Close()
		db.Save(&state)
		successResponse(c, "Successfully updated", map[string]interface{}{"state": state.State, "id": state.Id})

	} else {
		NoDataFoundResponse(c, "No Data Found")
	}

	/*gorp
	var dbmap = DbConnection()
	defer dbmap.Db.Close()

	c.Bind(&state)
	if state.State != "" {
		result, err := dbmap.Query(`UPDATE state SET state=? WHERE id=?`, state.State, state.Id)

		if err != nil {
			panic(err.Error())
		}

		// c.JSON(http.StatusOK, gin.H{
		// 	"message":    fmt.Sprintf("%s Successfully updated at Id %x", state.State, state.Id),
		// 	"body":       state.State,
		// 	"statusCode": http.StatusOK,
		// })
		successResponse(c, "Successfully updated", map[string]interface{}{"state": state.State, "id": state.Id})

		defer result.Close()

	} */
}
