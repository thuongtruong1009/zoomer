package e2e

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
	"github.com/thuongtruong1009/zoomer/infrastructure/app"
	"github.com/thuongtruong1009/zoomer/infrastructure/configs"
	"github.com/thuongtruong1009/zoomer/internal/modules/auth/presenter"
	"github.com/thuongtruong1009/zoomer/pkg/helpers"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"syscall"
	"testing"
	"github.com/thuongtruong1009/zoomer/pkg/constants"
)

type e2eTestSuite struct {
	suite.Suite
	config *configs.Configuration
	dbConn *gorm.DB
}

func TestE2ETestSuite(t *testing.T) {
	suite.Run(t, &e2eTestSuite{})
}

func (s *e2eTestSuite) SetupSuite() {
	s.config = &configs.Configuration{
		AppPort:               "8080",
		SigningKey:            "secret",
		HashSalt:              "salt",
		DatabaseConnectionURL: "postgres://postgres:postgres@localhost:5432/zoomer?sslmode=disable",
		JwtSecret:             "secret",
	}

	dsn := s.config.DatabaseConnectionURL
	dbConn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	s.Require().NoError(err)
	s.dbConn = dbConn

	serverReady := make(chan bool)
	// go server.NewServer(s.config, s.dbConn, serverReady)
	go app.Run()
	<-serverReady
}

func (s *e2eTestSuite) TeaDownSuite() {
	p, _ := os.FindProcess(os.Getpid())
	p.Signal(syscall.SIGINT)
}

func (s *e2eTestSuite) Test_EndToEnd_Register() {
	username := helpers.RandomChain(constants.RandomTypeString, 10)
	pwd := helpers.RandomChain(constants.RandomTypeString, 10)

	reqStr := `{"username":"` + username + `","password":"` + pwd + `", "limit": 10}`
	req, err := http.NewRequest("POST", fmt.Sprintf("http://localhost:%s/api/auth/signup", s.config.AppPort), strings.NewReader(reqStr))
	s.NoError(err)

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	client := http.Client{}
	response, err := client.Do(req)
	s.NoError(err)
	s.Equal(http.StatusOK, response.StatusCode)

	byteBody, err := ioutil.ReadAll(response.Body)
	s.NoError(err)

	expectedResp := `{"username":"` + strings.ToLower(username) + `","limit":10}`
	var expectedUser = presenter.SignUpResponse{}
	err = json.Unmarshal([]byte(expectedResp), &expectedUser)
	s.NoError(err)

	var actualUser = presenter.SignUpResponse{}
	err = json.Unmarshal([]byte(strings.Trim(string(byteBody), "\n")), &actualUser)
	s.NoError(err)

	s.Equal(actualUser.Username, expectedUser.Username)
	s.Equal(actualUser.Limit, expectedUser.Limit)
	response.Body.Close()
}
