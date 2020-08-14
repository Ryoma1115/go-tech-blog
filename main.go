package main

import (
	"log"
	"os"

	"go-tech-blog/handler"
	"go-tech-blog/repository"

	_ "github.com/go-sql-driver/mysql" // Using MySQL driver
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var db *sqlx.DB

var e = createMux()

func main() {
	db = connectDB()
	repository.SetDB(db)
	// `/` というパス（URL）と `articleIndex` という処理を結びつける
	e.GET("/", handler.ArticleIndex)
	e.GET("/new", handler.ArticleNew)
	e.GET("/:id", handler.ArticleShow)
	e.GET("/:id/edit", handler.ArticleEdit)

	// Webサーバーをポート番号 8080 で起動する
	e.Logger.Fatal(e.Start(":8080"))
}

func createMux() *echo.Echo {
	// アプリケーションインスタンスを生成
	e := echo.New()

	// アプリケーションに各種ミドルウェアを設定
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())

	// `src/css` ディレクトリ配下のファイルに `/css` のパスでアクセスできるようにする
	e.Static("/css", "src/css")
	e.Static("/js", "src/js")

	// アプリケーションインスタンスを返却
	return e
}

func connectDB() *sqlx.DB {
	dsn := os.Getenv("DSN")
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		e.Logger.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		e.Logger.Fatal(err)
	}
	log.Println("db connection succeeded")
	return db
}

// func articleIndex(c echo.Context) error {
// 	// ステータスコード 200 で、"Hello, World!" という文字列をレスポンス
// 	// return c.String(http.StatusOK, "Hello, World!")
// 	data := map[string]interface{}{
// 		"Message": "Article Index",
// 		"Now":     time.Now(),
// 	}
// 	return render(c, "article/index.html", data)
// }

// func articleNew(c echo.Context) error {
// 	data := map[string]interface{}{
// 		"Message": "Article New",
// 		"Now":     time.Now(),
// 	}

// 	return render(c, "article/new.html", data)
// }

// func articleShow(c echo.Context) error {
// 	id, _ := strconv.Atoi(c.Param("id"))

// 	data := map[string]interface{}{
// 		"Message": "Article Show",
// 		"Now":     time.Now(),
// 		"ID":      id,
// 	}

// 	return render(c, "article/show.html", data)
// }

// func articleEdit(c echo.Context) error {
// 	id, _ := strconv.Atoi(c.Param("id"))

// 	data := map[string]interface{}{
// 		"Message": "Article Edit",
// 		"Now":     time.Now(),
// 		"ID":      id,
// 	}

// 	return render(c, "article/edit.html", data)
// }

// func htmlBlob(file string, data map[string]interface{}) ([]byte, error) {
// 	return pongo2.Must(pongo2.FromCache(tmplPath + file)).ExecuteBytes(data)
// }

// func render(c echo.Context, file string, data map[string]interface{}) error {
// 	b, err := htmlBlob(file, data)
// 	if err != nil {
// 		return c.NoContent(http.StatusInternalServerError)
// 	}
// 	return c.HTMLBlob(http.StatusOK, b)
// }
