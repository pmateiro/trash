import (
    "github.com/streadway/amqp"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/bson"
)

conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
if err != nil {
    // handle error
}
defer conn.Close()

ch, err := conn.Channel()
if err != nil {
    // handle error
}
defer ch.Close()

client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
if err != nil {
    // handle error
}
defer client.Disconnect(context.Background())

db := client.Database("mydb")
collection := db.Collection("mycollection")

q, err := ch.QueueDeclare(
    "myqueue", // queue name
    false,    // durable
    false,    // delete when unused
    false,    // exclusive
    false,    // no-wait
    nil,      // arguments
)
if err != nil {
    // handle error
}

err = ch.QueueBind(
    q.Name,       // queue name
    "mykey",      // routing key
    "myexchange", // exchange
    false,
    nil,
)
if err != nil {
    // handle error
}

msgs, err := ch.Consume(
    q.Name, // queue name
    "",     // consumer
    true,   // auto-ack
    false,  // exclusive
    false,  // no-local
    false,  // no-wait
    nil,    // arguments
)
if err != nil {
    // handle error
}

for msg := range msgs {
    // handle message
}
var data MyDataStruct // replace MyDataStruct with the name of your struct type
err = json.Unmarshal(msg.Body, &data)
if err != nil {
    // handle error
}

_, err = collection.InsertOne(context.Background(), bson.M{
    "field1": data.Field1,
    "field2": data.Field2,
    // insert additional fields as needed
})
if err != nil {
    // handle error
}


