package trelloClient

import (
	"github.com/VojtechVitek/go-trello"
)

type TrelloCard struct {
	Card trello.Card `json:"trello_card"`
}
