package cli

import (
	"context"
	"week11"
)

type UgotnewsCmd struct {
	IO
	Name    string
	Service *week11.Service
	//no flags
}

func (cmd *UgotnewsCmd)Main(ctx context.Context, pwd string, args []string) error {
	cmd.init(pwd, args)
	return nil 
}

func (cmd *UgotnewsCmd)init(pwd string, args []string) error {
// list of categories in the backup file
// location of the backup file
// number of articles in the backup file
// number of articles per category
// number of articles per source
	return nil 
}