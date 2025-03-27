package command

import (
	"context"
	"maps"
	"slices"

	"github.com/oking02/surfe-challenge/internal/datasources"
	"github.com/oking02/surfe-challenge/internal/domain"
)

type actionMap struct {
	base domain.ActionType
	next domain.ActionType
}

type NextActionProbabilityCommand struct {
	actionFetcher datasources.ActionLister
}

func NewNextActionProbabilityCommand(actionFetcher datasources.ActionLister) *NextActionProbabilityCommand {
	return &NextActionProbabilityCommand{
		actionFetcher: actionFetcher,
	}
}

func (na *NextActionProbabilityCommand) NextActionProbability(ctx context.Context, clientID domain.ClientID, actionType domain.ActionType) (map[domain.ActionType]float64, error) {

	actions, err := na.actionFetcher.ListActions(ctx, clientID)
	if err != nil {
		return nil, err
	}
	mapCounts, total := na.buildActionMap(actions, actionType)

	probabilityMap := make(map[domain.ActionType]float64)
	for a, mapCount := range mapCounts {
		probabilityMap[a.next] = float64(mapCount) / float64(total)
	}

	return probabilityMap, nil
}

func (na *NextActionProbabilityCommand) buildActionMap(actions []domain.Action, actionType domain.ActionType) (map[actionMap]int, int) {
	var (
		actionMaps []actionMap
		total      int
	)
	for userActions := range maps.Values(na.bucketActionsByUser(actions)) {
		for i, userAction := range userActions {
			if i == len(userActions)-1 {
				continue
			}
			if userAction.Type != actionType {
				continue
			}
			total++
			actionMaps = append(actionMaps, actionMap{
				base: userAction.Type,
				next: userActions[i+1].Type,
			})
		}
	}

	mapCounts := make(map[actionMap]int)
	for _, a := range actionMaps {
		count, ok := mapCounts[a]
		if ok {
			mapCounts[a] = count + 1
			continue
		}
		mapCounts[a] = 1
	}

	return mapCounts, total
}

func (na *NextActionProbabilityCommand) bucketActionsByUser(actions []domain.Action) map[domain.UserID][]domain.Action {
	results := make(map[domain.UserID][]domain.Action)

	for _, action := range actions {
		if userActions, ok := results[action.UserID]; ok {
			results[action.UserID] = append(userActions, action)
			continue
		}
		results[action.UserID] = []domain.Action{action}
	}

	for userID, userActions := range results {
		slices.SortFunc(userActions, func(a, b domain.Action) int {
			return a.CreatedAt.Compare(b.CreatedAt)
		})
		results[userID] = userActions
	}

	return results
}
