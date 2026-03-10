### Real-time applications
Server side events are real-time but unidirectional while websockets are real-time but bidirectional. 
Both protocols are primarily used in browsers but you can use other protocols like TCP for mobile or 
MQTT for iot devices. 

Event streams are useful for text only data. When large file blobs are converted to base64, event 
streams can be slow due to cpu conversions.


| event stream | websockets |
| :----------- | :--------- |
| uni-directional | bi-directional |
| pure http | starts with http switches protocols |
| text only | text & binary |
| less overhead | more overhead |
| simple | complicated |

Potential applications for websockets
- chat apps
- collaborative apps e.g. Linear

Potential applications for sse
- stock ticks
- football matches
- election results

There's also WebRTC for real time communication apps e.g. zoom or other video call applications.


### Reloading server with Air package
Live server reload
```bash
go get -tool github.com/air-verse/air@latest

# then use it like so:
go tool air -v
```

### Expose local host to the internet
```bash
brew install cloudflared
cloudflared tunnel --url http://localhost:8080
```