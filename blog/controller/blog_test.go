package blogController

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

func TestblogController_NewBlog(t *testing.T) {
	testPostBody := map[string]interface{}{
		"title": "TestblogController_NewBlog TITLE",
		"text":  "TestblogController_NewBlog TEXT",
	}
	body, _ := json.Marshal(testPostBody)
	req, err := http.NewRequest("POST", "/new", bytes.NewReader(body))
	defer req.Body.Close()
	var m map[string]interface{}
	err = json.NewDecoder(req.Body).Decode(&m)

	fmt.Println(err, m)
}
