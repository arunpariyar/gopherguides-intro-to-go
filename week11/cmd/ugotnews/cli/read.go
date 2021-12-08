type ReadCmd struct {
	IO
	ID int
	DB string	//location of backup
	News *news
}


func (cmd *ReadCmd)Main(ctx context.Context, pwd string, args []string)error{
	if err := cmd.init(pwd); err != nil {
		return err
	}

	return cmd.print(cmd.News)
}

func(cmd *ReadCmd) init(pwd string) error {
	if cmd.News == nil {
		cmd.News = &service.News{}
	}

	if cmd.DB == "" {
		cmd.DB = "news.json"
	}

	err:= OpenNews(filepath.Join(pwd, cmd.DB), cmd.News)
	if err != nil{
		return err
	} 
	
	return nil
}

func OpenNews(path string, news *service.News)error{

}