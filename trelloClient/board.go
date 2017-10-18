package trelloClient

import (
	"github.com/VojtechVitek/go-trello"
)

type TrelloBoard struct {
	Board trello.Board `json:"trello_board"`
	Lists []TrelloList `json:"trello_lists"`
}
