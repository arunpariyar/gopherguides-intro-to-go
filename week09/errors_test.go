package week09

import (
	"testing"
)

func Test_Errors_All(t *testing.T) {
	table := []struct {
		name string
		err  error
		exp  string
	}{
		{
			name: "title_invalid",
			err:  ErrTitleInvalid(""),
			exp:  `draft must have a title, got ""`,
		},
		{
			name: "body_invalid",
			err:  ErrBodyInvalid(""),
			exp:  `draft must have a body, got ""`,
		},
		{
			name: "writer_invalid",
			err:  ErrWriterInvalid(""),
			exp:  `draft must have a writer, got ""`,
		},
		{
			name: "story_publisher_invalid",
			err:  ErrStoryPublisherInvalid(""),
			exp:  `story must hava a publisher got ""`,
		},
		{
			name: "story_catagory_invalid",
			err:  ErrStoryCatagoryInvalid(""),
			exp:  `story must hava a catagory got ""`,
		},
		{
			name: "invalid_subscriber_name",
			err:  ErrSubscriberInvalidName(""),
			exp:  `subscriber must hava a valid name got ""`,
		},
		{
			name: "invalid_subscriber_catagories",
			err:  ErrSubscriberCatgoriesInvalid(0),
			exp:  `subscriber must have at least one catagory`,
		},
		{
			name: "sources_empty",
			err:  ErrSourcesEmpty(0),
			exp:  `news service sources is empty`,
		},
	}

	for _, tt := range table {
		act := tt.err.Error()
		if act != tt.exp {
			t.Fatalf("expected %v got %v", tt.exp, act)
		}
	}
}
