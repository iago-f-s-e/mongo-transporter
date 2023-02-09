package domain

type Config struct {
	Receiver ReceiverConfig
	Sender   SenderCofing
}

type Flags struct {
	ConfigFile string
}

func (c Config) IsValid() bool {
	return receiverConfigIsValid(c.Receiver) && senderConfigIsValid(c.Sender)
}

func receiverConfigIsValid(r ReceiverConfig) bool {
	collection := r.CollectionName != ""
	database := r.DatabaseName != ""
	uri := r.Uri != ""

	return collection && database && uri
}

func senderConfigIsValid(s SenderCofing) bool {
	collection := s.CollectionName != ""
	database := s.DatabaseName != ""
	uri := s.Uri != ""

	return collection && database && uri
}
