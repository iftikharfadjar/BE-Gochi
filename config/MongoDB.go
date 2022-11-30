package config

var URL string = "mongodb+srv://admin:UXNIhGeK3Rp655Oi@fainacluster.bu653.mongodb.net/test?retryWrites=true&w=majority"

// func createConnection() (*mongo.Client, error) {
//    credential := options.Credential{
//       AuthMechanism: "SCRAM-SHA-1",
//       Username:      "MongoUser", // your mongodb user
//       Password:      "YourVerySecurePassword", // ...and mongodb
//    }
//    connString := "mongodb+srv://admin:admin@fainacluster.bu653.mongodb.net/test"
//    clientOpts := options.Client().ApplyURI(connString).SetAuth(credential)
//    return mongo.Connect(context.TODO(), clientOpts)
// }