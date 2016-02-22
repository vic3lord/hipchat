# Hipchat

A little tool to send hipchat short messages from command line  
I use it to send my team a useful links etc. instead of launching hipchat just to send links

```
go get github.com/vic3lord/hipchat

hipchat -help
```

## hipchatrc

use hipchat config file instead of command line flags so settings can be persistent

*by default hipchat-msg checks for $HOME/.hipchatrc*

```javascript
{
  "from": "Or Elimelech",
  "room": "Random room",
  "token": "api token"
}
```
