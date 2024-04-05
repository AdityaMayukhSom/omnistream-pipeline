package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func DeletePost(c echo.Context) error {

	// vars := mux.Vars(r)
	// keys := vars["id"]

	// _, err := repositories.Db.Query(
	// 	"DELETE FROM posts WHERE id = $1;",
	// 	keys,
	// )

	// if err != nil {
	// 	panic(err.Error())
	// }
	return c.JSON(http.StatusOK, map[string]any{"deleted": 100})
}
