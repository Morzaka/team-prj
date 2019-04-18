package train_test

/*
import (
	"team-project/mocks"
	"team-project/services"
	"team-project/services/data"
	"testing"

	"github.com/golang/mock/gomock"
)

var router = services.NewRouter()

type TrainsTestCases struct {
	tcase           string
	url             string
	expected        int
	mockedGetTrains []data.Train
	testTrainID     string
}

func TestGetTrains(t *testing.T) {
	tests := []TrainsTestCases{
		{
			tcase:    "GetTrainsOk",
			url:      "v1/trains",
			expected: 200,
		},
		{
			tcase:    "GetTrainsFailed",
			url:      "v1/trains",
			expected: 400,
		},
	}
	for _, tc := range tests {
		t.Run(tc.tcase, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			testGet := mocks.NewMockITrain(ctrl)

			train := data.Train{
				DepartureCity: "Lviv",
				ArrivalCity:   "Kiev",
			}
			testGet.AddTrain(train)
			testGet.EXPECT().GetAllTrains().Return([]data.Train{
				train,
			})
		})
	}
}
*/
