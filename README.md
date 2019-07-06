# ctftime-bot
[![CircleCI](https://circleci.com/gh/kg0r0/ctftime-bot.svg?style=svg)](https://circleci.com/gh/kg0r0/ctftime-bot) [![codecov](https://codecov.io/gh/kg0r0/ctftime-bot/branch/master/graph/badge.svg)](https://codecov.io/gh/kg0r0/ctftime-bot)    
:flags:  This is a bot for notifying ctf date to slack.

# Usage
Edit api_token and channel_id in ``ctftime_conf.json``.
```
{
  "slack_config":{
    "api_token":"********************",
    "channel_id":"********************"
  }
}
```
Execute the following commands.
```
$ make
$ ./ctftime-bot
```