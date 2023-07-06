package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

// album 表示有关专辑的数据.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// 专辑切片以填充专辑数据记录.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	// 绑定方法和路径
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/albums/:id", getAlbumByID)
	connectDB()
	router.Run("localhost:8080")
}

var db *sql.DB

// 连接数据库
func connectDB() {
	cfg := mysql.Config{
		User:                 os.Getenv("DBUSER"),
		Passwd:               os.Getenv("DBPASS"),
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "recordings",
		AllowNativePasswords: true,
	}
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected")
}

// getAlbums 以 JSON 格式响应所有专辑的列表.
func getAlbums(c *gin.Context) {
	var albums []album
	rows, err := db.Query("select * from album;")

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, err)
		return
	}

	for rows.Next() {
		var alb album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {

		}
		albums = append(albums, alb)
	}
	if err := rows.Err(); err != nil {
		c.IndentedJSON(http.StatusNotFound, err)
		return
	}
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums 从请求体中收到的JSON中添加一个专辑 .
func postAlbums(c *gin.Context) {
	var newAlbum album
	// 调用BindJson将收到的json绑定到newAlbum
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	result, error := db.Exec("insert into album (title, artist, price) value (?,?,?)", newAlbum.Title, newAlbum.Artist, newAlbum.Price)
	if error != nil {
		c.IndentedJSON(http.StatusNotFound, error)
		return
	}
	_, err := result.LastInsertId()
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, error)
		return
	} else {
		c.IndentedJSON(http.StatusOK, newAlbum)
	}
}

// getAlbumByID 查找 ID 值与客户端发送的 id
// 参数匹配的专辑，然后返回该专辑作为响应.
func getAlbumByID(c *gin.Context) {
	var alb album

	id := c.Param("id")
	row := db.QueryRow("select * from album where id = ?", id)
	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		if err == sql.ErrNoRows {
			c.IndentedJSON(http.StatusNotFound, gin.H{"msg": "album not found"})
			return
		}
	}
	c.IndentedJSON(http.StatusOK, alb)
	// 循环浏览专辑列表，查找
	// ID 值与参数匹配的专辑.
	// for _, a := range albums {
	// 	if a.ID == id {
	// 		c.IndentedJSON(http.StatusOK, a)
	// 		return
	// 	}
	// }
	// c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

/*
获取所有albums
curl http://localhost:8080/albums \
 --header "Content-Type: application/json" \
    --request "GET"

插入album
curl http://localhost:8080/albums \
--include \
--header "Content-Type: application/json" \
--request "POST" \
--data '{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'

根据条件查找album
curl http://localhost:8080/albums/2 \
 --header "Content-Type: application/json" \
    --request "GET"
*/
