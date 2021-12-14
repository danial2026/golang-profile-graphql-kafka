module github.com/danial2026/golang-profile-graphql-kafka/internal/graphql

go 1.16

require (
	github.com/99designs/gqlgen v0.14.0
	github.com/danial2026/golang-profile-graphql-kafka/internal/common v0.0.0-00010101000000-000000000000
	github.com/vektah/gqlparser/v2 v2.2.0
	golang.org/x/net v0.0.0-20201021035429-f5854403a974
	google.golang.org/grpc v1.42.0
)

replace github.com/danial2026/golang-profile-graphql-kafka/internal/common => ../common/
