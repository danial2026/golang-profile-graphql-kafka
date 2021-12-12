package adopters

import (
	"context"
	"math/rand"
	"testing"
	"time"

	"github.com/danial2026/golang-profile-graphql-kafka/internal/profile/domain"
	user "github.com/danial2026/golang-profile-graphql-kafka/internal/profile/domain"
	"github.com/stretchr/testify/require"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	Name       string
	Repository user.Repository
}

func createRepositories(t *testing.T) []Repository {
	return []Repository{
		{
			Name:       "MongoDB",
			Repository: NewMongoConnection(t),
		},
	}
}

func TestRepository(t *testing.T) {
	t.Parallel()
	rand.Seed(time.Now().UTC().UnixNano())

	repositories := createRepositories(t)

	for i := range repositories {
		// When you are looping over slice and later using iterated value in goroutine (here because of t.Parallel()),
		// you need to always create variable scoped in loop body!
		// More info here: https://github.com/golang/go/wiki/CommonMistakes#using-goroutines-on-loop-iterator-variables
		r := repositories[i]

		t.Run(r.Name, func(t *testing.T) {
			// It's always a good idea to build all non-unit tests to be able to work in parallel.
			// Thanks to that, your tests will be always fast and you will not be afraid to add more tests because of slowdown.
			t.Parallel()

			t.Run("testCreateAccount1", func(t *testing.T) {
				t.Parallel()
				testCreateAccount1(t, r.Repository)
			})
			t.Run("testCreateAccount2", func(t *testing.T) {
				t.Parallel()
				testCreateAccount2(t, r.Repository)
			})

			t.Run("testUnfollow", func(t *testing.T) {
				t.Parallel()
				testUnfollow(t, r.Repository)
			})
			t.Run("testFollow", func(t *testing.T) {
				t.Parallel()
				testFollow(t, r.Repository)
			})
		})
	}
}

func testCreateAccount1(t *testing.T, repository user.Repository) {
	t.Helper()
	ctx := context.Background()

	testCase := domain.User{
		Email: "email2@gmail.com",
		Username: "username1_test",
	}

	t.Run("testFollow", func(t *testing.T) {
		t.Parallel()
		err := repository.CreatUser(ctx, testCases)
		require.NoError(t, err)
	})
}

func testCreateAccount2(t *testing.T, repository user.Repository) {
	t.Helper()
	ctx := context.Background()

	testCase := domain.User{
		Email: "email2@gmail.com",
		Username: "username2_test",
	}

	t.Run("testFollow", func(t *testing.T) {
		t.Parallel()

		err := repository.CreatUser(ctx, testCases)
		require.NoError(t, err)
	})
}

func testFollow(t *testing.T, repository user.Repository) {
	t.Helper()
	ctx := context.Background()

	testUser1 := domain.User{
		Email: "email2@gmail.com",
		Username: "username2_test",
	}

	testUser2 := domain.User{
		Email: "email2@gmail.com",
		Username: "username2_test",
	}

	t.Run("testFollow", func(t *testing.T) {
		t.Parallel()

		err := repository.Follow(ctx, testUser1.Username, testUser2.Username)
		require.NoError(t, err)
	})
}

func testUnfollow(t *testing.T, repository user.Repository) {
	t.Helper()
	ctx := context.Background()

	testUser1 := domain.User{
		Email: "email2@gmail.com",
		Username: "username2_test",
	}

	testUser2 := domain.User{
		Email: "email2@gmail.com",
		Username: "username2_test",
	}

	t.Run("testUnfollow", func(t *testing.T) {
		t.Parallel()

		err := repository.Unfollow(ctx, testUser1.Username, testUser2.Username)
		require.NoError(t, err)
	})
}