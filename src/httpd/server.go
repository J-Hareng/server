package httpd

import (
	"fmt"
	email "server/src/api/Email"
	"server/src/api/db"
	"server/src/httpd/handler"
	"server/src/httpd/security"

	"github.com/gin-contrib/cors"

	// "github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/contrib/static"

	"github.com/gin-gonic/gin"
)

type Server struct {
	db *db.DB

	engine *gin.Engine
}

func Init(DB *db.DB, TM *security.TokenMap, E email.Email, EKM *security.EmailTokenMap) Server {

	r := gin.Default()
	r.Use(cors.Default())

	// * WEB
	r.Use(static.Serve("/", static.LocalFile("client", true)))

	// r.GET("/", handler.TestGet())

	// * User

	// * -- GET
	r.GET("/getUsers", handler.GetUsers(DB))
	r.GET("/validateSessionToken", handler.ValidateUserToken(DB, TM))

	// * -- POST
	r.POST("/addUser", handler.AddUser(DB, EKM))
	r.POST("/reqSessionToken", handler.RequestSessionToken(DB, TM))
	r.POST("/reqEmailKey", handler.RequestEmailKey(E, EKM, DB))

	return Server{
		engine: r,
		db:     DB,
	}
}

func (s Server) Run(port string) {
	fmt.Println("Running now")
	s.engine.Run("0.0.0.0:" + port) // listen and serve on 0.0.0.0:8080 (for
}
