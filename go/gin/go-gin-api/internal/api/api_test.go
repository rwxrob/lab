package api_test

import (
	"encoding/json"
	"example/go-gin-api/internal/api"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

func ExampleGetAlbums() {

	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.GET("/albums", api.GetAlbums)

	req, err := http.NewRequest(http.MethodGet, "/albums", nil)
	if err != nil {
		fmt.Println(err)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	fmt.Println(w.Code)

	first := []api.Album{}
	json.Unmarshal(w.Body.Bytes(), &first)
	fmt.Println(len(first))

	// Output:
	// 200
	// 3
}
