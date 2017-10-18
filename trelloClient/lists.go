package trelloClient

import (
	"github.com/VojtechVitek/go-trello"
)

type TrelloList struct {
	List  trello.List  `json:"trello_list"`
	Cards []TrelloCard `json:"trello_cards"`
}
