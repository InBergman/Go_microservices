package main

import (
	"os"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
)

var tmpArticleList = []articleList

func testMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}

func getRouter(withTemplates bool) *gin.Engine {
	r := gin.Default()
	if withTemplates {
		r.LoadHTMLGlob("templates/*")
	}
	return r
}

func testHTTPResponse(test *testing.T, r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder)bool){
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if !f(w) {
		test.Fail()
	}
}

func saveList() {
	tmpArticleList = articleList
}

func restoreLists() {
	articleList = tmpArticleList
}