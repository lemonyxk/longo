
    package main
    
    import (
        "log"
    
        "go.mongodb.org/mongo-driver/bson"
    
        "mongo"
    )
    
    func main() {
        var url = "mongodb://root:123$56@127.0.0.1:27017,127.0.0.1:27018,127.0.0.1:27019"
    
        mgo, _ := mongo.NewMongoClient().Connect(&mongo.Config{Url: url})
    
        err := mgo.RawClient().Ping(nil, mongo.ReadPreference.Primary)
        if err != nil {
            panic(err)
        }
    
        var res = []bson.M{nil}
    
        _ = mgo.DB("QGame").C("GameUser").Find(bson.M{}).All(&res)
    
        for key, value := range res {
            log.Println(key, value)
        }
    }