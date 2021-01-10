package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/team11/app/controllers"
	"github.com/team11/app/ent"
	"github.com/team11/app/ent/author"
	"github.com/team11/app/ent/book"
	"github.com/team11/app/ent/category"
	"github.com/team11/app/ent/role"
	"github.com/team11/app/ent/servicepoint"
	"github.com/team11/app/ent/status"
	"github.com/team11/app/ent/user"
)

//User struct
type User struct {
	Name     string
	Email    string
	Password string
	Role     int
}

//Location struct
type Location struct {
	Name string
}

//ServicePoint struct
type ServicePoint struct {
	CounterNumber string
}

//ClientEntity struct
type ClientEntity struct {
	Name  string
	State int
}

//Book struct
type Book struct {
	BookName string
	Category int
	Author   int
	User     int
	Status   int
}

// Author struct
type Author struct {
	NameAuthor string
}

// Researchtype struct
type Researchtype struct {
	NameResearchtype string
}

//Researchs struct
type Researchs struct {
	Research []Research
}

//Research struct
type Research struct {
	NameResearch string
}

//Bookborrow struct
type Bookborrow struct {
	UserID         int
	BookID         int
	ServicePointID int
	BorrowDate     string
}

// @title SUT SA Example API Playlist Vidoe
// @version 1.0
// @description This is a sample server for SUT SE 2563
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information
func main() {
	router := gin.Default()
	router.Use(cors.Default())

	client, err := ent.Open("sqlite3", "file:ent.db?cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("fail to open sqlite3: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	v1 := router.Group("/api/v1")
	controllers.NewAuthorController(v1, client)
	controllers.NewBookController(v1, client)
	controllers.NewBookborrowController(v1, client)
	controllers.NewBookingController(v1, client)
	controllers.NewBookreturnController(v1, client)

	controllers.NewCategoryController(v1, client)
	controllers.NewClientEntityController(v1, client)
	controllers.NewLocationController(v1, client)
	controllers.NewPreemptionController(v1, client)
	controllers.NewPurposeController(v1, client)

	controllers.NewResearchController(v1, client)
	controllers.NewResearchTypeController(v1, client)
	controllers.NewRoleController(v1, client)
	controllers.NewRoominfoController(v1, client)
	controllers.NewServicePointController(v1, client)

	controllers.NewStatusController(v1, client)
	controllers.NewUserController(v1, client)

	Role := []string{"Library Member", "Librarian"}
	for _, r := range Role {
		client.Role.
			Create().
			SetROLENAME(r).
			Save(context.Background())
	}
	// Set Users Data
	Users := []User{
		{"Suphachai Phetthamrong", "B6111427@gmail.com", "B6111427", 1},
		{"Name Surname", "me@example.com", "pass", 1},
		{"root", "root@gmail.com", "root", 2},
	}
	for _, u := range Users {

		r, err := client.Role.
			Query().
			Where(role.IDEQ(int(u.Role))).
			Only(context.Background())
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		client.User.
			Create().
			SetUSEREMAIL(u.Email).
			SetUSERNAME(u.Name).
			SetPASSWORD(u.Password).
			SetPosition(r).
			Save(context.Background())
	}

	Status := []string{"Available", "In Use", "No Service", "Borrowed"}
	for _, st := range Status {
		client.Status.
			Create().
			SetSTATUSNAME(st).
			Save(context.Background())
	}

	ClientEntity := []ClientEntity{
		{"Video 0n Demand - 01", 1},
		{"Video 0n Demand - 02", 1},
		{"Video 0n Demand - 02", 1},
		{"Video 0n Demand - 03", 1},
		{"Video 0n Demand - 04", 1},
		{"Video 0n Demand - 05", 1},
		{"Video 0n Demand - 06", 1},
		{"Video 0n Demand - 07", 1},
		{"Video 0n Demand - 08", 1},
		{"Video 0n Demand - 09", 1}}
	for _, cl := range ClientEntity {

		s, err := client.Status.
			Query().
			Where(status.IDEQ(int(cl.State))).
			Only(context.Background())
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		client.ClientEntity.
			Create().
			SetCLIENTNAME(cl.Name).
			SetState(s).
			Save(context.Background())
	}

	ServicePoint := []ServicePoint{
		{"เค้าเตอร์ 1"},
		{"เค้าเตอร์ 2"}}
	for _, se := range ServicePoint {
		client.ServicePoint.
			Create().
			SetCOUNTERNUMBER(se.CounterNumber).
			Save(context.Background())
	}

	Author := []string{"โยชิฮิโระ โทงาชิ", "เออิจิโร โอดะ", "เจ.อาร์.อาร์ โทลคีน"}
	for _, a := range Author {
		client.Author.
			Create().
			SetName(a).
			Save(context.Background())
	}

	Researchtype := []string{"การวิจัยเชิงประวัติศาสตร์", "การวิจัยเชิงบรรยาย", "การวิจัยเชิงทดลอง", "การวิจัยเชิงปริมาณ", "การวิจัยพื้นฐาน", "การวิจัยประยุกต์", ""}
	for _, rt := range Researchtype {
		client.Researchtype.
			Create().
			SetTYPENAME(rt).
			Save(context.Background())
	}

	Category := []string{"นิยาย", "เรื่องสั้น", "การ์ตูน", "ศิลปกรรม", "ความรู้ทั้วไป", "ปรัชญา จิตวิทยา ศาสนา", "ประวัติศาสตร์", "ภูมิศาสตร์", "เทคโนโลยี", "การศึกษา"}
	for _, ca := range Category {
		client.Category.
			Create().
			SetCategoryName(ca).
			Save(context.Background())
	}

	Location := []string{"Building A", "Building B", "Building C"}
	for _, l := range Location {
		client.Location.
			Create().
			SetLOCATIONNAME(l).
			Save(context.Background())
	}

	Book := []Book{
		{"หนูน้อยผจญภัยในโลกไดโนเสาร์", 3, 1, 1, 1},
		{"lord of the ring", 1, 3, 1, 1}}
	for _, b := range Book {

		ca, err := client.Category.
			Query().
			Where(category.IDEQ(int(b.Category))).
			Only(context.Background())
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		au, err := client.Author.
			Query().
			Where(author.IDEQ(int(b.Author))).
			Only(context.Background())
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		us, err := client.User.
			Query().
			Where(user.IDEQ(int(b.User))).
			Only(context.Background())
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		st, err := client.Status.
			Query().
			Where(status.IDEQ(int(b.Status))).
			Only(context.Background())
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		client.Book.
			Create().
			SetBookName(b.BookName).
			SetCategory(ca).
			SetAuthor(au).
			SetUser(us).
			SetStatus(st).
			Save(context.Background())
	}

	Bookborrow := []Bookborrow{
		{1, 2, 1,"2021-01-10 23:48:00+07:00"},
		{2, 1, 2,"2021-01-10 23:48:00+07:00"}}
	for _, bb := range Bookborrow {

		b, err := client.Book.
			Query().
			Where(book.IDEQ(int(bb.BookID))).
			Only(context.Background())
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		u, err := client.User.
			Query().
			Where(user.IDEQ(int(bb.UserID))).
			Only(context.Background())
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		sp, err := client.ServicePoint.
			Query().
			Where(servicepoint.IDEQ(int(bb.ServicePointID))).
			Only(context.Background())
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		times, err := time.Parse(time.RFC3339, bb.BorrowDate)

		client.Bookborrow.
			Create().
			SetBOOK(b).
			SetUSER(u).
			SetSERVICEPOINT(sp).
			SetBORROWDATE(times).
			Save(context.Background())
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}
