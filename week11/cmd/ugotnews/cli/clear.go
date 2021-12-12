package cli

import (
	"context"
	"flag"
	"week11"
)

type ClearCmd struct {
	IO
	Name 	string
	Service *week11.Service
	//flags:
	DB 	string //location for the backup file
	flags 	*flag.FlagSet
}


func (cmd *ClearCmd) Flags() *flag.FlagSet {
	if cmd.flags != nil {
		return cmd.flags
	}
	//adding flags to the flagset 
	cmd.flags = flag.NewFlagSet(cmd.Name, flag.ContinueOnError)
	cmd.flags.StringVar(&cmd.DB, "f", cmd.DB, "file location for backup of the backup file")

	return cmd.flags
}

func (cmd *ClearCmd) SetIO(oi IO){
	cmd.IO = oi
}


func (cmd *ClearCmd)Main(ctx context.Context, pwd string, args []string) error {
	cmd.init(pwd, args)
	return nil 
}

func (cmd *ClearCmd) init(pwd string, args []string) error {
	if err := cmd.Flags().Parse(args); err != nil {
		return err
	}

	if cmd.DB == "" {
		cmd.DB = "./tmp/news_service.json"
	}

	if cmd.Service == nil {
		cmd.Service = week11.NewService()
	}

	cmd.Service.Start(context.Background())
	//creating a new file source to publish news
	nfs := week11.NewFileSource("auto")
	// //adding it to the news service
	cmd.Service.Add(context.Background(),nfs)
	// //publishing stories in a go routine
	nfs.PublishStories()
	cmd.Service.Clear()
	// fmt.Println(pwd, args, cmd.DB)
	return nil 
	
}
