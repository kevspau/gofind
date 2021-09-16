import ( "go.mongodb.org/mongo-driver/mongo" "log")

clientOptions := options.Client().ApplyURI("mongodb+srv://sharpcdf:<password>@gofind.hg63z.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()
client, err := mongo.Connect(ctx, clientOptions)
if err != nil {
    log.Fatal(err)
}