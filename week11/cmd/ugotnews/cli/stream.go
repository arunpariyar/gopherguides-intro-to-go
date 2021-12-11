package cli

import (
	"context"
	"flag"
	"week11"
)

type StreamCmd struct {
	IO
	Name string
	Service *week11.Service
	//flags:
	DB 		string //location of the output
	JSON	bool
	Output  string
	flags 	 *flag.FlagSet 
}

func (cmd *StreamCmd) Flags() *flag.FlagSet {
	
	if cmd.flags != nil {
		return cmd.flags
	}

	cmd.flags = flag.NewFlagSet(cmd.Name, flag.ContinueOnError)
	//adding the flags to the flagset
	cmd.flags.BoolVar(&cmd.JSON, "j", cmd.JSON, "output in json format")
	cmd.flags.StringVar(&cmd.DB, "f", cmd.DB, "location of news DB file")
	cmd.flags.StringVar(&cmd.Output,"o", cmd.Output, "output result to a file")
	return cmd.flags
}

func (cmd *StreamCmd)Main(ctx context.Context, pwd string, args []string)error{
	cmd.init(pwd, args)
	return nil
}

func (cmd *StreamCmd) SetIO(oi IO){
	cmd.IO = oi
}

func(cmd *StreamCmd)init(pwd string, args []string) error {

	if err := cmd.Flags().Parse(args); err != nil {
		return err
	}

	if cmd.Service == nil {
		cmd.Service = week11.NewService()
	}
	catagories := cmd.Flags().Args()
	
	cs := make([]week11.Catagory,0)
	for _, v := range catagories { 
		cs = append(cs, week11.Catagory(v))
	}

	cmd.Service.Start(context.Background())
	cmd.Service.Subscribe("terminal", cs...)
	
	nfs := week11.NewFileSource("cmdline")
	ctx := nfs.Start(context.Background())
	cmd.Service.Add(ctx, nfs)

	 nfs.PublishStories()
	return nil 
}