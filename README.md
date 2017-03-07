# go-bithue
go-bithue is a go program that polls the Bitfinex API and sets the Philips Hue lighting in your crib/office/whatever according to the current bitcoin price against the 24-hour BTC/USD high and low.  if the price is close to the 24h high, your lights will be green, in the middle, yellow, and if you're losing mad crypto notes, red.

you've never partied harder with bitcoin and lights in your life.  get some work off the darknet and go hammers, kemosabe.

send me bitcoins @ **1NiEfjEE35m5UV64K3Gmni2PCthHrr8ucw**.

send me nudes @ **rapper@gmail.com**.

### usage
- clone my repo.
- open up `config.json` in an elite text editor.
- replace `BridgeIP` with the IP address of your Hue bridge.
- replace `Username` with an authorized username for your Hue bridge.
- run `go build`
- run `./go-bithue &`
- `kill -9` that pid when you're ready to get off mr. bitcoin's wild ride.

### todo
a lot.  this is a piece of shit, can't you tell?

### thanks
@Collinux for his gohue project.
