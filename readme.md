# Multi Datadog Client
Sometimes we need more than one datadog statsd clients when one datadog agent doesn't able to handle all the request.


## How to use
```go
import "github.com/wejick/multi-datadog-client

client := multi-datadog-client.New("host1","host2",...,"hostn")
// do whatever you want just like using datadog statsd
client.Get().Gauge(xxx)
```