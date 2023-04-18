package repository

// import (
// 	"context"
// 	"database/sql"
// 	"regexp"
// 	"testing"
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// 	"gorm.io/gorm/logger"
// 	"github.com/go-test/deep"
// 	"github.com/google/uuid"
// 	"github.com/stretchr/testify/require"
// 	"github.com/stretchr/testify/suite"
// 	"gopkg.in/DATA-DOG/go-sqlmock.v1"

// 	"zoomer/internal/auth"
// 	"zoomer/internal/models"
// )

// type Suite struct {
// 	suite.Suite
// 	DB *gorm.DB
// 	mock sqlmock.Sqlmock

// 	repo auth.UserRepository
// }

// func TestInit(t *testing.T){
// 	suite.Run(t, new(Suite))
// }

// func (s *Suite) AfterTest(_, _ string){
// 	require.NoError(s.T(), s.mock.ExpectationsWereMet())
// }

// func (s *Suite) SetupSuite() {
// 	var (
// 		db  *sql.DB
// 		err error
// 	)

// 	db, mock, err := sqlmock.New()
// 	require.NoError(s.T(), err)

// 	// s.DB, err = gorm.Open("postgres", db)
// 	s.DB, err = gorm.Open(postgres.New(postgres.Config{
// 		Conn: db,
// 	}), &gorm.Config{
// 		Logger: logger.Default.LogMode(logger.Silent),
// 	})
// 	require.NoError(s.T(), err)

// 	// s.mock = mock
// 	// s.DB.LogMode(true)
// 	s.repo = NewUserRepository(s.DB)
// }

// func (s *Suite) TestGetUserByUsername(){
// 	var (
// 		id = uuid.New().String()
// 		username = "test"
// 		password = "testpass"
// 		limit = 100
// 	)

// 	s.mock.ExpectQuery(regexp.QuoteMeta(
// 		`SELECT * FROM "users" WHERE "users"."username" = `
// 	),)
// }
