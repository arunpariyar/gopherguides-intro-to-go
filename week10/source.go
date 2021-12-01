package week10

// type sources []source //collection type for source

type source struct {
	Name string
	// ctx     context.Context
	stories stories
	// draft   drafts
}

//need to create a function that will go through its vault and based on how many drafts there is build them to story and store it in the stories slice within source stories, stories will then be forward to the waiting channel of the news service

func (s source) build(d draft, c catagory) (story, error) {

	if err := d.IsValid(); err != nil {
		return story{}, err //TODO: figure out if return empty story is okay
	}

	if c == "" {
		return story{}, ErrStoryCatagoryInvalid(c)
	}

	return story{
		publisher: s.Name,
		catagory:  c,
		title:     d.title,
		body:      d.body,
		writer:    d.writer,
	}, nil
}

func (s *source) push(st story) error {
	if err := st.IsValid(); err != nil {
		return err
	}
	s.stories = append(s.stories, st)
	return nil
}
