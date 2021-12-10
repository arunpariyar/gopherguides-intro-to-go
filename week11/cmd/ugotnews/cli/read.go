package cli

import (
	"context"
	"fmt"
	"strconv"
	"week11"
)

type ReadCmd struct {
	IO
	Name 	string
	DB   	string
	Service	*week11.Service
}


func (cmd *ReadCmd)Main(ctx context.Context, pwd string, args []string)error{
	cmd.init(args)
	return nil
}

func(cmd *ReadCmd)init(args []string) error {
	if cmd.Service == nil {
		cmd.Service = week11.NewService()
	}
	keys := make([]int,0)

	for _, v := range args{
		i, _ := strconv.Atoi(v)
	 keys = append(keys, i)
	}
	
	cmd.Service.Start(context.Background())
	
	news, err := cmd.Service.Search(keys...)
	if err != nil {
		return err 
	}
	for _, n := range news {
		fmt.Println(n)
	}
	return nil 

}

