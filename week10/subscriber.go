package week09

type catagories []catagory

// type subscribers []subscriber //collection type for subscribers
type subscriber struct {
	name string
	cat  catagories
}

func (s subscriber) IsValid() error {
	switch true {
	case s.name == "":
		return ErrSubscriberInvalidName(s.name)
	case len(s.cat) == 0:
		return ErrSubscriberCatgoriesInvalid(0)
	}
	return nil
}

//function to subscriber to add itself to newsservice with catagory
// func (s subscriber) Subscribe(ns *NewsService) error {
// 	//add itsself to the news service subs list
// 	if err := s.IsValid(); err != nil {
// 		return err
// 	}
// 	ns.Subs = append(ns.Subs, s)
// 	return nil
// }

//functon to remove itself from the newsservice
