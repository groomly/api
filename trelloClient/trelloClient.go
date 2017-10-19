package trelloClient

import (
	"fmt"
	"github.com/VojtechVitek/go-trello"
	"log"
	"os"
)

func LinkTrello(token string) (trelloMember trello.Member, trelloBoards []TrelloBoard) {
	var trelloLists []TrelloList
	var trelloCards []TrelloCard
	trelloBoards = make([]TrelloBoard, 0)

	trelloAppKey := os.Getenv("TRELLO_APP_KEY")
	trello, err := trello.NewAuthClient(trelloAppKey, &token)
	if err != nil {
		log.Fatal(err)
	}

	member, err := trello.Member("me")
	trelloMember = *member
	boards, err := trelloMember.Boards()
	for _, board := range boards {
		lists, _ := board.Lists()
		trelloLists = make([]TrelloList, 0)
		for _, list := range lists {
			trelloCards = make([]TrelloCard, 0)
			cards, _ := list.Cards()
			for _, card := range cards {
				trelloCards = append(trelloCards, TrelloCard{Card: card})
			}
			trelloLists = append(trelloLists, TrelloList{List: list, Cards: trelloCards})
		}
		trelloBoards = append(trelloBoards, TrelloBoard{Board: board, Lists: trelloLists})
		fmt.Println("done with board")
		break
	}

	return trelloMember, trelloBoards

}
