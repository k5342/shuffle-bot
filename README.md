# shuffle-bot
An awesome discord bot to make shuffled teams.

## Usage
After connecting to any voice channel on your server, simply type following command on your text channels to make teams.
```
!!teams <N> [exclude_users]...
```

Parameters:
- `N` -- the number of teams to you are going to make.
- `exclude_users` -- an screen_name of users excluded on team making.

Examples:
- `!!teams` -- shows usage
- `!!teams 2` -- creates `2` teams by voice channel users where you are in
- `!!teams 2 userA` -- creates `2` teams without `userA`
- `!!teams 5 userX userZ` -- creates `5` teams without `userX` and `userZ`

## Installation and Setup

### On Dedicated Host

#### Clone & Install
To run manually, clone this repository and build:
```
go build main.go
```

#### Register an Application on Discord Developer Portal
And create an *Application* (bot feature and some permissions enabled) and issue *Token* from [Discord Developer Portal](https://discordapp.com/developers/applications/)

This bot requires following *Bot Permissions*:
```
- Send Messages
```

#### Run
Set *Token* as `SHUFFLEBOT_TOKEN` environment variable.
```
export SHUFFLEBOT_TOKEN=`XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX`
```

Or you are on Windows:
```
set SHUFFLEBOT_TOKEN=`XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX`
```

Finally run the bot.
```
./main
.\main.exe
```

#### Invite to your server
And to invite your server, generate invitation URL manually using *Client ID* on *Discord Developer Portal*.
You can create Invitation URL manually following this template. More details are available at [discord documentation](https://discordapp.com/developers/docs/topics/oauth2#bots).
```
https://discordapp.com/oauth2/authorize?client_id=<client_id>&scope=bot&permissions=2048
```

### From This Invitation Link (Beta)
**Disclaimer:** We may stop this service at any time on no notice and no warranty. And since this bot can view some messages on your server, please set appropriate permissions to protect from your sensitive messages and unexpected malfunction!!

```
https://discordapp.com/oauth2/authorize?client_id=506729503754682369&scope=bot&permissions=2048
```

## LICENSE
The MIT License

Copyright 2018 k5342

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
