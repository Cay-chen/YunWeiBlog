package test

import (
	_ "YunWeiBlog/routers"
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	_, file, _, _ := runtime.Caller(0)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

// TestBeego is a sample to run an endpoint test
func TestBeego(t *testing.T) {
	fmt.Println("中文字符串包含子串返回：", strings.Contains("蒋春生", "春生"))

	/*	r, _ := http.NewRequest("GET", "/", nil)
		w := httptest.NewRecorder()
		beego.BeeApp.Handlers.ServeHTTP(w, r)

		beego.Trace("testing", "TestBeego", "Code[%d]\n%s", w.Code, w.Body.String())

		Convey("Subject: Test Station Endpoint\n", t, func() {
			Convey("Status Code Should Be 200", func() {
				So(w.Code, ShouldEqual, 200)
			})
			Convey("The Result Should Not Be Empty", func() {
				So(w.Body.Len(), ShouldBeGreaterThan, 0)
			})
		})*/
}
