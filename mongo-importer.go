// MongoDB to Couchbase Importer
//
// Author: Ali Hamidi <@ahamidi>

package main

import (
        "fmt"
        "flag"
        "strings"
        "labix.org/v2/mgo"
        "labix.org/v2/mgo/bson"
        "github.com/couchbaselabs/go-couchbase"
)

// Source MongoDB
var mongoHost = flag.String("mhost", "127.0.0.1",
        "URL of Source MongoDB Server")
var mongoPort = flag.String("mport", "27017",
        "MongoDB Server Port")
var mongoUser = flag.String("muser", "",
        "MongoDB Username")
var mongoPass = flag.String("mpass", "",
        "MongoDB Password")
var mongoDb = flag.String("mdb", "",
        "MongoDB Database")
var mongoCollection = flag.String("mcol", "",
        "MongoDB Collection")

// Target Couchbase Server
var cbHost = flag.String("cbhost", "127.0.0.1",
        "URL of Target Couchbase Server")
var cbPort = flag.String("cbport", "8091",
        "Couchbase Server Port")
var cbBucket = flag.String("bucket", "",
        "Couchbase Bucket")
var cbPass = flag.String("cbpass", "",
        "Couchbase Password")
var typeField = flag.String("typeField", "collection",
        "Name of Type field to be used. i.e. collection or type")
var typeName = flag.String("typeName", "",
        "Type Name to be used")

func formatMongoURL(user string, pass string, host string, port string, db string) string {
  s := []string{}

  // Check if username/password is needed
  if len(*mongoUser) > 0{
      s = append(s, []string{"mongodb://", user, ":", pass, "@", host, ":", port, "/", db}...)
  } else {
      s = append(s, []string{"mongodb://", host, ":", port, "/", db}...)
  }

  url := strings.Join(s,"")

  return url
}

func formatCBURL(bucket string, pass string, host string, port string) string {
  s := []string{}

  // Check if username/password is needed
  if len(*cbPass) > 0{
      s = append(s, []string{"http://", bucket, ":", pass, "@", host, ":", port, "/"}...)
  } else {
      s = append(s, []string{"http://", host, ":", port, "/"}...)
  }
  url := strings.Join(s,"")

  return url
}

func main() {
  // Get passed in args
  flag.Parse()

  mongoURL := formatMongoURL(*mongoUser, *mongoPass, *mongoHost, *mongoPort, *mongoDb)

  // Connect to Source MongoDB
  fmt.Println("Connecting to Source MongoDB Server...", mongoURL) // Provide feedback
  session, err := mgo.Dial(mongoURL)
  if err != nil {
    panic(err)
  } else {
    fmt.Println("Connected!")
  }
  defer session.Close()

  // Optional. Switch the session to a monotonic behavior.
  session.SetMode(mgo.Strong, true)

  // Connect to collection
  c := session.DB(*mongoDb).C(*mongoCollection)

  // Count the # of docs
  count, err := c.Count()
  if err != nil {
    panic(err)
  } else {
    fmt.Println("Collection has ", count, " docs.")
  }


  cbURL := formatCBURL(*cbBucket, *cbPass, *cbHost, *cbPort)

  // Connect to Target Couchbase Server
  fmt.Println("Connecting to ", cbURL)
  cb, err := couchbase.Connect(cbURL)
  pool, err := cb.GetPool("default")
  bucket, err := pool.GetBucket(*cbBucket)

  if err != nil {
    panic(err)
  } else {
    fmt.Println("Connected!")
  }

  progress := 0

  //var result []struct{ Value int }
  var obj bson.M
  iter := c.Find(nil).Iter()
  for iter.Next(&obj) {
      progress++
      objID := obj["_id"].(bson.ObjectId).Hex()

      // Inject collection name
      if *typeName == "" {
        *typeName = *mongoCollection
      }
      obj[*typeField] = *typeName

      fmt.Println(progress, "of", count, ": ", "Key: ", objID, " Value: ", obj)
      bucket.Set(objID, 0, obj)
  }
  if err := iter.Close(); err != nil {
      panic(err)
  }

}
