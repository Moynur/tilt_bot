package messagehandler

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// Service comment
type Service struct {
	discordapi     discordgo.Session
	discordMessage discordgo.MessageCreate
}

// NewConnection establishes a connection to discord server
func NewConnection(Token string) (*Service, error) {
	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return nil, err
	}

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return nil, err
	}
	return &Service{discordapi: *dg}, nil
}

// SendMessage sends string to target channel
func (s *Service) SendMessage(TargetChannelID, messageToSend string) error {
	_, err := s.discordapi.ChannelMessageSend(TargetChannelID, messageToSend)
	if err != nil {
		return fmt.Errorf("error sending message to channel %v", err)
	}
	return nil
}

// Close ends the connection
func (s *Service) Close() {
	s.Close()
}

// Listen for users to message who to annoy
func (s *Service) Listen() {

}
