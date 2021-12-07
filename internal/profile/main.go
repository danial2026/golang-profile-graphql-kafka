package main

func main() {
	// ctx := context.Background()
	// firestoreClient, err := firestore.NewClient(ctx, os.Getenv("GCP_PROJECT"))
	// if err != nil {
	// 	panic(err)
	// }
	// // firebaseDB := db{firestoreClient}

	// serverType := strings.ToLower(os.Getenv("SERVER_TO_RUN"))
	// switch serverType {
	// case "http":
	// 	go loadFixtures()

	// 	server.RunHTTPServer(func(router chi.Router) http.Handler {
	// 		return HandlerFromMux(HttpServer{firebaseDB}, router)
	// 	})
	// case "grpc":
	// 	server.RunGRPCServer(func(server *grpc.Server) {
	// 		svc := GrpcServer{firebaseDB}
	// 		users.RegisterUsersServiceServer(server, svc)
	// 	})
	// default:
	// 	panic(fmt.Sprintf("server type '%s' is not supported", serverType))
	// }
}
