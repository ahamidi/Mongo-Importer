# Mongo Importer
This is a MongoDB to Couchbase import tool. It allows you to configure a source MongoDB Database (and collection) and a Target Couchbase Server (and bucket). It will then connect to both an copy all the docs from the MongoDB collection to the Couchbase Bucket.

**What about Doc IDs?**
This tool automatically sets the Couchbase Document ID to the MongoDB Object ID  (Hex).

Source MongoDB Document from the "people" collection
```
{
  _id: ObjectId("524cca7b7fd28c2a8a0004e7"),
  username: "ahamidi",
  first_name: "Ali",
  last_name: "Hamidi"
}
```

Resulting Couchbase Document  
Document Key: 524cca7b7fd28c2a8a0004e7
```
{
  "_id":"524cca7b7fd28c2a8a0004e7",
  "collection":"people",
  "first_name":"Ali",
  "last_name":"Hamidi",
  "username":"ahamidi"
}
```

## Usage

The tool tries to maintain sensible defaults, and includes full usage instructions which can be access with ```-h``` option.
```
Usage of mongo-importer:
  -bucket="": Couchbase Bucket
  -cbhost="127.0.0.1": URL of Target Couchbase Server
  -cbpass="": Couchbase Password
  -cbport="8091": Couchbase Server Port
  -mcol="": MongoDB Collection
  -mdb="": MongoDB Database
  -mhost="127.0.0.1": URL of Source MongoDB Server
  -mpass="": MongoDB Password
  -mport="27017": MongoDB Server Port
  -muser="": MongoDB Username
  -typeField="collection": Name of Type field to be used. i.e. collection or type
  -typeName="": Type Name to be used
```

## Todo
1. Improve efficiency of transfer.  
Right now it transfers docs one at a time. Ideally it should pull groups of docs (say in groups of 1000) and bulk uploads. Doesn't seem possible with the Go SDK.
1. Add testing.
1. Improve error feedback.
1. Support resuming migrations.
1. Migrate indexes (need to investigate).
1. Allow multiple collections to be specified (one, several or all).
1. Improve progress feedback.

## Building

1. Checkout Code
1. ```go get``` in directory
1. ```go run mongo-importer.go``` _Alternatively ```go build``` instead._


## Contributing

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request

## License

The MIT License (MIT)

Copyright (c) 2014 Ali Hamidi

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
