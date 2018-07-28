package resolver

import (
	"errors"

	gcontext "github.com/kivutar/chainz/context"
	"github.com/kivutar/chainz/loader"
	"github.com/kivutar/chainz/service"
	"github.com/op/go-logging"
	"golang.org/x/net/context"
)

// User resolves a user graphql query
func (r *Resolver) User(ctx context.Context, args struct {
	Email string
}) (*userResolver, error) {
	//Without using dataloader:
	//user, err := ctx.Value("userService").(*service.UserService).FindByEmail(args.Email)
	userID := ctx.Value("user_id").(*string)
	user, err := loader.LoadUser(ctx, args.Email)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Retrieved user by user_id[%s] : %v", *userID, *user)
	return &userResolver{user}, nil
}

// Users resolves a users graphql query
func (r *Resolver) Users(ctx context.Context, args struct {
	First *int32
	After *string
}) (*usersConnectionResolver, error) {
	if isAuthorized := ctx.Value("is_authorized").(bool); !isAuthorized {
		return nil, errors.New(gcontext.CredentialsError)
	}
	userID := ctx.Value("user_id").(*string)

	users, err := ctx.Value("userService").(*service.UserService).List(args.First, args.After)
	count, err := ctx.Value("userService").(*service.UserService).Count()
	ctx.Value("log").(*logging.Logger).Debugf("Retrieved users by user_id[%s] :", *userID)
	config := ctx.Value("config").(*gcontext.Config)
	if config.DebugMode {
		for _, user := range users {
			ctx.Value("log").(*logging.Logger).Debugf("%v", *user)
		}
	}
	ctx.Value("log").(*logging.Logger).Debugf("Retrieved total users count by user_id[%s] : %v", *userID, count)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	return &usersConnectionResolver{users: users, totalCount: count, from: &(users[0].ID), to: &(users[len(users)-1].ID)}, nil
}
