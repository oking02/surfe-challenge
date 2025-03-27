package command

import (
	"context"
	"testing"

	"github.com/oking02/surfe-challenge/internal/datasources"
	"github.com/oking02/surfe-challenge/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestNewReferralIndexCommand(t *testing.T) {

	var (
		testUser1 = domain.UserID(1)
		testUser2 = domain.UserID(2)
		testUser3 = domain.UserID(3)
		testUser4 = domain.UserID(4)
		testUser5 = domain.UserID(5)
		testUser6 = domain.UserID(6)
	)

	tests := []struct {
		name         string
		userLister   datasources.UserLister
		actionLister datasources.ActionLister
		result       map[domain.UserID]int
		wantErr      assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			userLister: &mockUserLister{
				data: []domain.User{
					{ID: testUser1},
					{ID: testUser2},
					{ID: testUser3},
					{ID: testUser4},
					{ID: testUser5},
					{ID: testUser6},
				},
			},
			actionLister: &mockActionLister{
				data: []domain.Action{
					{
						UserID:   testUser1,
						Type:     domain.REFER_USER,
						TargetID: testUser2,
					},
					{
						UserID:   testUser2,
						Type:     domain.REFER_USER,
						TargetID: testUser3,
					},
					{
						UserID:   testUser2,
						Type:     domain.REFER_USER,
						TargetID: testUser4,
					},
					{
						UserID:   testUser4,
						Type:     domain.REFER_USER,
						TargetID: testUser5,
					},
					{
						UserID:   testUser5,
						Type:     domain.REFER_USER,
						TargetID: testUser6,
					},
				},
			},
			result: map[domain.UserID]int{
				testUser1: 5,
				testUser2: 4,
				testUser3: 0,
				testUser4: 2,
				testUser5: 1,
				testUser6: 0,
			},
			wantErr: assert.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := NewReferralIndexCommand(tt.userLister, tt.actionLister)

			res, err := cmd.ReferralIndex(t.Context(), "")
			tt.wantErr(t, err)
			assert.Equal(t, tt.result, res)
		})
	}
}

type mockUserLister struct {
	data []domain.User
}

func (a *mockUserLister) ListUsers(context.Context, domain.ClientID) ([]domain.User, error) {
	return a.data, nil
}
