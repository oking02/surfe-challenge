package command

import (
	"context"
	"testing"
	"time"

	"github.com/oking02/surfe-challenge/internal/datasources"
	"github.com/oking02/surfe-challenge/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestNextActionProbabilityCommand_Execute(t *testing.T) {

	var (
		testTime  = time.Now().AddDate(0, 0, -10)
		testUser1 = domain.UserID(1)
		testUser2 = domain.UserID(2)
	)

	tests := []struct {
		name         string
		actionType   domain.ActionType
		actionLister datasources.ActionLister
		result       map[domain.ActionType]float64
		wantErr      assert.ErrorAssertionFunc
	}{
		{
			name:       "success",
			actionType: domain.ADD_CONTACT,
			actionLister: &mockActionLister{
				data: []domain.Action{
					{
						UserID:    testUser1,
						Type:      domain.WELCOME,
						CreatedAt: testTime,
					},
					{
						UserID:    testUser1,
						Type:      domain.ADD_CONTACT,
						CreatedAt: testTime.AddDate(0, 0, 1),
					},
					{
						UserID:    testUser1,
						Type:      domain.VIEW_CONTACTS,
						CreatedAt: testTime.AddDate(0, 0, 2),
					},
					{
						UserID:    testUser1,
						Type:      domain.ADD_CONTACT,
						CreatedAt: testTime.AddDate(0, 0, 3),
					},
					{
						UserID:    testUser1,
						Type:      domain.REFER_USER,
						CreatedAt: testTime.AddDate(0, 0, 4),
					},
					{
						UserID:    testUser1,
						Type:      domain.ADD_CONTACT,
						CreatedAt: testTime.AddDate(0, 0, 5),
					},
					{
						UserID:    testUser1,
						Type:      domain.VIEW_CONTACTS,
						CreatedAt: testTime.AddDate(0, 0, 6),
					},
					{
						UserID:    testUser2,
						Type:      domain.ADD_CONTACT,
						CreatedAt: testTime.AddDate(0, 0, 1),
					},
					{
						UserID:    testUser2,
						Type:      domain.VIEW_CONTACTS,
						CreatedAt: testTime.AddDate(0, 0, 2),
					},
					{
						UserID:    testUser2,
						Type:      domain.ADD_CONTACT,
						CreatedAt: testTime.AddDate(0, 0, 3),
					},
					{
						UserID:    testUser2,
						Type:      domain.ADD_CONTACT,
						CreatedAt: testTime.AddDate(0, 0, 4),
					},
				},
			},
			result: map[domain.ActionType]float64{
				domain.ADD_CONTACT:   0.2,
				domain.REFER_USER:    0.2,
				domain.VIEW_CONTACTS: 0.6,
			},
			wantErr: assert.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := NewNextActionProbabilityCommand(tt.actionLister)

			res, err := cmd.NextActionProbability(t.Context(), "", tt.actionType)
			tt.wantErr(t, err)
			assert.Equal(t, tt.result, res)
		})
	}

}

type mockActionLister struct {
	data []domain.Action
}

func (a *mockActionLister) ListActions(context.Context, domain.ClientID) ([]domain.Action, error) {
	return a.data, nil
}
