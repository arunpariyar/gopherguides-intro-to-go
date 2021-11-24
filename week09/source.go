package week09

import "context"

type source struct {
	name    string
	ctx     context.Context
	stories stories
}

func (s source) build(d draft, c ...catagory) (story, error) {
	//make a categories
	var cs catagories

	if err := d.IsValid(); err != nil {
		return story{}, err //TODO: figure out if return empty story is okay
	}

	if len(c) == 0 {
		return story{}, ErrStoryCatagoriesInvalid(len(c))
	}

	cs = append(cs, c...)
	return story{
		publisher:  s.name,
		catagories: cs,
		title:      d.title,
		body:       d.body,
		writer:     d.writer,
	}, nil
}

func (s *source) push(st story) error {
	if err := st.IsValid(); err != nil {
		return err
	}
	s.stories = append(s.stories, st)
	return nil
}
