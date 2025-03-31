package command

import (
	"context"
	"fmt"

	"github.com/oking02/surfe-challenge/internal/datasources"
	"github.com/oking02/surfe-challenge/internal/domain"
)

type ReferralIndexCommand struct {
	userLister    datasources.UserLister
	actionFetcher datasources.ActionLister
}

func NewReferralIndexCommand(
	userLister datasources.UserLister,
	actionFetcher datasources.ActionLister) *ReferralIndexCommand {
	return &ReferralIndexCommand{
		userLister:    userLister,
		actionFetcher: actionFetcher,
	}
}

func (cmd *ReferralIndexCommand) ReferralIndex(ctx context.Context, clientID domain.ClientID) (map[domain.UserID]int, error) {

	users, err := cmd.userLister.ListUsers(ctx, clientID)
	if err != nil {
		return nil, fmt.Errorf("failed to list users: %w", err)
	}

	referralActions, err := cmd.getReferralActions(ctx, clientID)
	if err != nil {
		return nil, fmt.Errorf("failed to list actions: %w", err)
	}

	directReferrals := cmd.calculateDirectReferrals(users, referralActions)

	results := make(map[domain.UserID]int)
	for userID, referrals := range directReferrals {
		if len(referrals) == 0 {
			results[userID] = 0
			continue
		}

		total := countReferrals(userID, directReferrals)
		results[userID] = total
	}

	return results, nil

}

func countReferrals(userID domain.UserID, data map[domain.UserID][]domain.UserID) int {
	referrals, ok := data[userID]

	if !ok || len(referrals) == 0 {
		return 0
	}

	total := 0
	for _, referralID := range referrals {

		// can't refer yourself
		// also create infinite loop -> users 147,802 <-
		if referralID == userID {
			continue
		}

		total++

		t := countReferrals(referralID, data)
		total = total + t
	}
	return total

}

func (cmd *ReferralIndexCommand) calculateDirectReferrals(users []domain.User, referActions []domain.Action) map[domain.UserID][]domain.UserID {

	results := make(map[domain.UserID][]domain.UserID)
	for _, user := range users {
		var referrals []domain.UserID
		for _, action := range referActions {
			if action.UserID == user.ID {
				referrals = append(referrals, action.TargetID)
			}
		}
		results[user.ID] = referrals
	}

	return results
}

func (cmd *ReferralIndexCommand) getReferralActions(ctx context.Context, clientID domain.ClientID) ([]domain.Action, error) {

	actions, err := cmd.actionFetcher.ListActions(ctx, clientID)
	if err != nil {
		return nil, fmt.Errorf("failed to list actions: %w", err)
	}
	var results []domain.Action
	for _, action := range actions {
		if action.Type == domain.REFER_USER {

			if action.UserID != action.TargetID {
				results = append(results, action)
			}

		}
	}

	return results, nil
}
