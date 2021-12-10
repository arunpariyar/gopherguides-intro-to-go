package cli

import (
	"context"
	"flag"
	"fmt"
	"strconv"
	"week11"
)

type ReadCmd struct {
	IO
	Name 	string
	
	Service	*week11.Service
	//flags: 
	DB 		string 	//location of the output
	JSON 	bool  	//output in JSON
	Output  string 	//output file
	flags    *flag.FlagSet
}

func (cmd *ReadCmd) Flags() *flag.FlagSet {
	if len(cmd.DB) == 0 {
		cmd.DB = "news.json"
	} 

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



func (cmd *ReadCmd)Main(ctx context.Context, pwd string, args []string)error{
	cmd.init(pwd, args)
	return nil
}

func (cmd *ReadCmd) SetIO(oi IO){
	cmd.IO = oi
}

func(cmd *ReadCmd)init(pwd string, args []string) error {
	
	if cmd.DB == "" {
		cmd.DB = "news.json"
	}

	if err := cmd.Flags().Parse(args); err != nil {
		return err
	}
	
	if cmd.Service == nil {
		cmd.Service = week11.NewService()
	}
	
	keys := make([]int,0)

	for _, v := range args{
		i, _ := strconv.Atoi(v)
	 keys = append(keys, i)
	}
	
	cmd.Service.Start(context.Background())
	//creating a new file source to publish news
	nfs := week11.NewFileSource("auto")
	//adding it to the news service
	cmd.Service.Add(context.Background(),nfs)
	//publishing stories in a go routine
	nfs.PublishStories()
	
	news, err := cmd.Service.Search(keys...)

	if err != nil {
		return err 
	}
	
	// if cmd.JSON {
	// 	return json.NewEncoder(cmd.Stdout()).Encode(news)
	// }
	cmd.print(news)
	//clearing the service
	cmd.Service.Clear()
	//stopping the service
	cmd.Service.Stop()
	return nil 
}

func (cmd *ReadCmd) print(news []week11.News){
	fmt.Println(cmd.Output)

	for _, n := range news {
		fmt.Fprintf(cmd.Stdout(), "%v %v %v \n", n.Id, n.Body, n.Catagory )
	}
}

