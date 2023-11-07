package memory

import (
	"testing"

	"github.com/gschei/cyclecm/domain/club"
)

func TestMemory_GetClub(t *testing.T) {
	type testCase struct {
		name        string
		id          int64
		expectedErr error
	}

	c, err := club.NewClub("G1")
	if err != nil {
		t.Fatal(err)
	}
	id := c.ID

	repo := MemoryRepository{
		clubs: map[int64]club.Club{
			id: c,
		},
	}

	testCases := []testCase{
		{
			name:        "No Customer By ID",
			id:          55,
			expectedErr: club.ErrClubNotFound,
		}, {
			name:        "Customer By ID",
			id:          id,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			_, err := repo.Get(tc.id)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}

func TestMemory_AddClub(t *testing.T) {
	type testCase struct {
		name        string
		club        club.Club
		expectedErr error
	}

	repo := MemoryRepository{
		clubs: make(map[int64]club.Club),
	}

	testCases := []testCase{
		{
			name:        "Add Club",
			club:        club.Club{ID: int64(1), Name: "Test1"},
			expectedErr: nil,
		}, {
			name:        "Add Club 2",
			club:        club.Club{ID: int64(2), Name: "Test2"},
			expectedErr: nil,
		}, {
			name:        "Empty Name",
			club:        club.Club{ID: int64(3), Name: ""},
			expectedErr: club.ErrInvalidClub,
		},
	}

	for _, tc := range testCases {
		c, err := club.NewClub(tc.club.Name)

		var res club.Club
		if err == nil {
			t.Logf("created club %v,%v", c.ID, c.Name)
			res, err = repo.Add(c)
			t.Logf("added club %v,%v", res.ID, res.Name)
		}
		if err != tc.expectedErr {
			t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
		}
		if err == nil && res.ID != tc.club.ID {
			t.Errorf("Expected id %v, got %v", tc.club.ID, res.ID)
		}
	}
}

func TestMemory_UpdateClub(t *testing.T) {

	type testCase struct {
		name        string
		club        club.Club
		newName     string
		expectedErr error
	}

	c, _ := club.NewClub("XXX")

	repo := MemoryRepository{
		clubs: map[int64]club.Club{
			c.ID: c,
		},
	}

	testCases := []testCase{
		{
			name:        "Update Club",
			club:        club.Club{ID: int64(1), Name: "Test1"},
			expectedErr: nil,
		}, {
			name:        "Update Nonexisting Club",
			club:        club.Club{ID: int64(2), Name: "Test2"},
			expectedErr: club.ErrClubNotFound,
		},
	}

	for _, tc := range testCases {
		err := repo.Update(tc.club)
		if err != tc.expectedErr {
			t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
		}
		if err == nil {
			c, _ := repo.Get(tc.club.ID)
			t.Logf("new name: %v\n", c.Name)
			if c.Name != tc.club.Name {
				t.Errorf("Expected Name %v, got %v", tc.club.Name, c.Name)
			}
		}
	}

}

func TestMemory_Singleton(t *testing.T) {
	repo1 := New()

	c1, err := club.NewClub("G1")
	if err != nil {
		c1, _ = repo1.Add(c1)
	}

	repo2 := New()

	c2, err2 := club.NewClub("G1")
	if err2 != nil {
		c2, _ = repo2.Add(c2)
	}

	if c1.ID != c2.ID || c1.ID != 1 {
		t.Errorf("Ids are wrong, should be equal to 1: %v %v ", c1.ID, c2.ID)
	}

}
