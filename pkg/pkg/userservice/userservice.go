package userservice

import (
	"context"
	"database/sql"
	"fmt"
	"jackbox/pkg/pkg/auth"
	"jackbox/pkg/pkg/db"
	api "jackbox/pkg/pkg/user"
	"log"
	"net/mail"
	"strings"
	"time"
)

type UserService struct {
	authorizer *auth.Authorizor
	querier    *db.Queries
	logger     *log.Logger
	dbDriver   *sql.DB
}

const (
	host     = "postgres"
	port     = 5432
	user     = "jackboxtest"
	password = "jackboxtest"
	dbname   = "jackboxtest"
)

func New(secret []byte, logger *log.Logger) (*UserService, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	postgresqlDbInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	driver, err := sql.Open(
		"postgres",
		postgresqlDbInfo)
	if err != nil {
		return nil, err
	}
	if err = driver.PingContext(ctx); err != nil { // validate connection
		log.Fatal(err)
	}
	querier := db.New(driver)
	authorizor := auth.New(secret, logger)
	return &UserService{
		querier:    querier,
		authorizer: authorizor,
		logger:     logger,
		dbDriver:   driver,
	}, nil
}

func (us *UserService) NewError(ctx context.Context, err error) *api.ErrorStatusCode {
	return &api.ErrorStatusCode{
		StatusCode: 500,
		Response: api.Error{
			Code:    500,
			Message: err.Error(),
		},
	}
}

func (us *UserService) GetUsers(ctx context.Context, params api.GetUsersParams) (resp api.GetUsersRes, err error) {
	claims := ctx.Value("claims").(*auth.JWTClaims)
	if claims.Role != "admin" {
		err = fmt.Errorf("unauthorized")
		return
	}
	var limit, offset int
	if v, ok := params.Limit.Get(); !ok {
		limit = 100
	} else {
		limit = v
	}
	if v, ok := params.Offset.Get(); !ok {
		offset = 0
	} else {
		offset = v
	}
	var users []db.ListUsersRow
	users, err = us.querier.ListUsers(ctx, db.ListUsersParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	})
	if err != nil {
		return nil, err
	}
	var userList api.UserList
	for _, user := range users {
		userList = append(userList, api.User{
			Username:  user.Username,
			Firstname: user.Firstname,
			Lastname:  user.Lastname,
			Email:     user.Email,
		})
	}
	return &userList, nil
}

func (us *UserService) AddUser(ctx context.Context, req *api.User) (resp *api.User, err error) {
	if err = req.Validate(); err != nil {
		return
	}
	if _, err = mail.ParseAddress(req.Email); err != nil {
		return
	}

	var tx *sql.Tx
	tx, err = us.dbDriver.Begin()
	if err != nil {
		return
	}
	defer tx.Rollback()
	qtx := us.querier.WithTx(tx)
	var userId int64
	userId, err = qtx.CreateUser(ctx, db.CreateUserParams{
		Username:  req.Username,
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Email:     req.Email,
		Crypt:     req.Password,
	})
	if err != nil {
		us.logger.Printf("qtx.CreateUser(): %v\n", err)
		return
	}
	role := "user"
	if strings.Contains(req.Email, "@jackboxgames.com") {
		role = "admin"
	}
	err = qtx.AssignUserRole(ctx, db.AssignUserRoleParams{
		UserID: userId,
		Name:   role,
	})
	if err != nil {
		us.logger.Printf("us.querier.AssignUserRole: %v\n", err)
		return nil, err
	}
	if err = tx.Commit(); err != nil {
		us.logger.Printf("tx.Commit(): %v\n", err)
		return
	}

	return &api.User{
		Username:  req.Username,
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Email:     req.Email,
	}, nil
}

func (us *UserService) GetUserByUsername(ctx context.Context, params api.GetUserByUsernameParams) (resp api.GetUserByUsernameRes, err error) {
	claims := ctx.Value("claims").(*auth.JWTClaims)
	if claims.Subject != params.Username {
		err = fmt.Errorf("unauthorized")
		return
	}
	var user db.User
	user, err = us.querier.GetUserByUsername(ctx, params.Username)
	if err != nil {
		us.logger.Printf("us.querier.GetUserByUsername: %v\n", err)
		return nil, err
	}
	return &api.User{
		Username:  user.Username,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Email:     user.Email,
	}, nil
}

func (us *UserService) GetToken(ctx context.Context, req *api.UsernamePassword) (resp *api.LoginResponse, err error) {
	var user db.GetUserByUsernameAndPasswordRow
	user, err = us.querier.GetUserByUsernameAndPassword(ctx, db.GetUserByUsernameAndPasswordParams{
		Username: req.Username,
		Crypt:    req.Password,
	})
	if err != nil {
		us.logger.Printf("us.querier.GetUserByUsernameAndPassword(): %v\n", err)
		return
	}
	var roleName string
	roleName, err = us.querier.GetRoleByUsername(ctx, user.Username)
	if err != nil {
		us.logger.Printf(" us.querier.GetRoleByUsername(): %v\n", err)
		return
	}
	var tokenString string
	tokenString, err = us.authorizer.NewToken(user.Username, roleName)
	if err != nil {
		us.logger.Printf("us.authorizer.NewToken(user.Username): %v\n", err)
		return
	}
	return &api.LoginResponse{
		Token: api.Token{Token: tokenString},
		User: api.User{
			Firstname: user.Firstname,
			Lastname:  user.Lastname,
			Username:  user.Username,
			Email:     user.Email,
		},
	}, nil
}

func (us *UserService) RefreshToken(ctx context.Context, req *api.Username) (token *api.Token, err error) {
	claims := ctx.Value("claims").(*auth.JWTClaims)
	if claims.Subject != req.Username {
		err = fmt.Errorf("unauthorized")
		return
	}
	var roleName string
	roleName, err = us.querier.GetRoleByUsername(ctx, req.Username)
	if err != nil {
		us.logger.Println(" us.querier.GetRoleByUsername(): %v\n", err)
		return
	}
	var tokenString string
	tokenString, err = us.authorizer.NewToken(req.Username, roleName)
	if err != nil {
		us.logger.Printf("us.authorizer.NewToken(user.Username): %v\n", err)
		return
	}
	return &api.Token{
		Token: tokenString,
	}, nil
}

func (us *UserService) HandleBearerAuth(ctx context.Context, _ string, t api.BearerAuth) (context.Context, error) {
	claims, err := us.authorizer.VerifyTokenAndGetClaims(t.GetToken())
	if err != nil {
		us.logger.Printf("us.authorizer.VerifyTokenAndGetClaims(): %v\n", err)
		return ctx, err
	}
	ctx = context.WithValue(ctx, "claims", claims)
	return ctx, nil
}
