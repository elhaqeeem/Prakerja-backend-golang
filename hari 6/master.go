package main
import ( "fmt"
"net/http" "strconv"
"github.com/jinzhu/gorm"
_ "github.com/jinzhu/gorm/dialects/mysql" 
"github.com/labstack/echo"
)
var (
DB *gorm.DB
)
func init() {
  InitDB()
  InitialMigration()
}
type Config struct { 
    DB_Username string 
    DB_Password string
	DB_Port string 
    DB_Host string 
    DB_Name string
}
func InitDB() {
config := Config{ DB_Username: "root", DB_Password: "root123", DB_Port: "3306", DB_Host: "localhost", DB_Name: "crud_go",
}
connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
config.DB_Username, config.DB_Password, config.DB_Host, config.DB_Port, config.DB_Name,
)
var err error
DB, err = gorm.Open("mysql", connectionString) 
if err != nil {
panic(err) 
}
}
type User struct {
gorm.Model
Name string `json:"name" form:"name"`
Email string `json:"email" form:"email"` 
Password string `json:"password" form:"password"`
}
func InitialMigration() { DB.AutoMigrate(&User{})
}
// get all users
func GetUsersController(c echo.Context) error { var users []User
if err := DB.Find(&users).Error; err != nil {
return echo.NewHTTPError(http.StatusBadRequest, err.Error())
}
return c.JSON(http.StatusOK, map[string]interface{}{

	"message": "success get all users",
"users": users, })
}
// get user by id
// Get user by ID
func GetUserController(c echo.Context) error {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
    }

    var user User
    if err := DB.First(&user, id).Error; err != nil {
        if gorm.IsRecordNotFoundError(err) {
            return echo.NewHTTPError(http.StatusNotFound, "User not found")
        }
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
    }

    return c.JSON(http.StatusOK, map[string]interface{}{
        "message": "success get user by ID",
        "user":    user,
    })
}
// create new user
func CreateUserController(c echo.Context) error { user := User{}
c.Bind(&user)
if err := DB.Save(&user).Error; err != nil {
return echo.NewHTTPError(http.StatusBadRequest, err.Error())
}
return c.JSON(http.StatusOK, map[string]interface{}{
"message": "success create new user",
"user": user, })
}
func DeleteUserController(c echo.Context) error {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
    }

    var user User
    if err := DB.First(&user, id).Error; err != nil {
        if gorm.IsRecordNotFoundError(err) {
            return echo.NewHTTPError(http.StatusNotFound, "User not found")
        }
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
    }

    if err := DB.Delete(&user).Error; err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }

    return c.JSON(http.StatusOK, map[string]interface{}{
        "message": "success deleted user",
    })
}
// Update user by ID
func UpdateUserController(c echo.Context) error {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
    }

    var user User
    if err := DB.First(&user, id).Error; err != nil {
        if gorm.IsRecordNotFoundError(err) {
            return echo.NewHTTPError(http.StatusNotFound, "User not found")
        }
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
    }

    c.Bind(&user)
    if err := DB.Save(&user).Error; err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }

    return c.JSON(http.StatusOK, map[string]interface{}{
        "message": "success updated user",
        "user":    user,
    })
}
func main() {
// create a new echo instance
e := echo.New()
// Route / to handler function e.GET("/users", GetUsersController) e.GET("/users/:id", GetUserController) e.POST("/users", CreateUserController) e.DELETE("/users/:id", DeleteUserController) e.PUT("/users/:id", UpdateUserController)
// start the server, and log if it fails
e.Logger.Fatal(e.Start(":8000")) }
