package httpd

import (
	"fmt"
	"net/http"
	"server/src/api/db"
	"server/src/httpd/handler"
	"server/src/httpd/security"

	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/contrib/static"

	"github.com/gin-gonic/gin"
)

type Server struct {
	db *db.DB

	engine *gin.Engine
}

func Init(DB *db.DB, TM *security.TokenMap) Server {

	r := gin.Default()
	//
	//!testing
	r.GET("/ping_test", func(c *gin.Context) {
		fmt.Printf("Request \n")
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Use(cors.Default())

	r.Use(static.Serve("/", static.LocalFile("src/client", true)))

	// * WEB
	// r.GET("/", handler.TestGet())

	//*User

	//* -- GET
	r.GET("/getUsers", handler.GetUsers(DB))
	r.GET("/validateSessionToken", handler.ValidateUserToken(DB, TM))

	//* -- POST
	r.POST("/addUser", handler.AddUser(DB))
	r.POST("/reqSessionToken", handler.RequestSessionToken(DB, TM))

	return Server{
		engine: r,
		db:     DB,
	}
}

func (s Server) Run(port string) {
	fmt.Println("Running now")
	s.engine.Run("0.0.0.0:" + port) // listen and serve on 0.0.0.0:8080 (for
}
