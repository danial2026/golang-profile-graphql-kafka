module github.com/danial2026/golang-profile-graphql-kafka/internal/profile

go 1.16

require (
	github.com/golang/protobuf v1.5.2
	github.com/sirupsen/logrus v1.5.0
	google.golang.org/grpc v1.42.0
)

require (
	github.com/danial2026/golang-profile-graphql-kafka/internal/common v0.0.0-00010101000000-000000000000
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.0
	go.mongodb.org/mongo-driver v1.8.0
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519 // indirect
	golang.org/x/sync v0.0.0-20200317015054-43a5402ce75a // indirect
	golang.org/x/sys v0.0.0-20211031064116-611d5d643895 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
)

replace github.com/danial2026/golang-profile-graphql-kafka/internal/common => ../common/
