# Hipchat

A little tool to send hipchat short messages from command line  

```
go get github.com/vic3lord/hipchat

hipchat -help
```

## hipchatrc

use hipchat config file instead of command line flags so settings can be persistent  
*by default hipchat checks for $HOME/.hipchatrc*

```javascript
{
  "from": "Or Elimelech",
  "room": "Random room",
  "token": "api token"
}
```
