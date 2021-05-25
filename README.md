# oakenshield
![](./dist/logo.jpg)

Oakenshield is an id generation service based on an algorithm similar to the snowflake, developed to generate unique ids in distributed systems.

We use 16 bit of randomness and a shuffled bit sequence generating hash like ID sequences. We support 256 different machines. We use 63 bit to ensure positive values for an int64 datatype.
Also, we guarantee uniqueness of generated IDs over a time span of 146 years.
#### The id consists of the following four components:
 * 4 bytes of time
 * 1 byte sequence counter
 * 2 random bytes
 * 1 byte machine id

#### You should know that:
* The amount of random bytes will decrease to 1 byte when more than 64 Ids generated within a timespan of one second. The random bytes will turn off when more than 8,224 Ids generated within a timespan of one second.
* Each one second time frame is capable to hold more than 4,000,000 IDs. It's safe to generate unlimited more when stick to a cool down time of id-count / 4,000,000 seconds between program restarts.

### Usage
You can generate new unique id with using following code
```go
 oakenshildId := domain.NewOakenshildID()
 rawId := oakenshildId.Next()
 fmt.Println(fm.Sprintf("Raw ID: %s", rawId))
```

Also, you can encode your raw id to base64, base32, hex and base58
```go
 oakenshildIdEncodeStrategy := domain.NewOakenshieldIDEncodeStrategy() 
 encodedId := oakenshildIdEncodeStrategy.Encode("base64", rawId)
 fmt.Println(fm.Sprintf("Encoded ID: %s", encodedId))
```

### Installation

You can use docker file like:

`docker build . `

`docker run oakenshildId-service`

You can generate new id like:
```curl
curl --location --request GET 'localhost:80/next?encoder=base58' \
--header 'Content-Type: application/json'
```

```json
{
    "data": {
        "raw_id": 13616834933752257,
        "encoded_id": "2qL6AVSEbr"
    },
    "message": ""
}
```

