package week10

import "fmt"

type ErrTitleInvalid string

func (e ErrTitleInvalid) Error() string {
	return fmt.Sprintf("draft must have a title, got %#v", string(e))
}

type ErrBodyInvalid string

func (e ErrBodyInvalid) Error() string {
	return fmt.Sprintf("draft must have a body, got %#v", string(e))
}

type ErrWriterInvalid string

func (e ErrWriterInvalid) Error() string {
	return fmt.Sprintf("draft must have a writer, got %#v", string(e))
}

type ErrStoryPublisherInvalid string

func (e ErrStoryPublisherInvalid) Error() string {
	return fmt.Sprintf("story must hava a publisher got %#v", string(e))
}

type ErrStoryCatagoryInvalid string

func (e ErrStoryCatagoryInvalid) Error() string {
	return fmt.Sprintf("story must hava a catagory got %#v", string(e))
}

type ErrSubscriberInvalidName string

func (e ErrSubscriberInvalidName) Error() string {
	return fmt.Sprintf("subscriber must hava a valid name got %#v", string(e))
}

type ErrSubscriberCatgoriesInvalid int

func (e ErrSubscriberCatgoriesInvalid) Error() string {
	return "subscriber must have at least one catagory"
}

type ErrSourcesEmpty int

func (e ErrSourcesEmpty) Error() string {
	return "news service sources is empty"
}
