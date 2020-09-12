The purpose of this bot is to monitor games and let people know when someone loses 

TODOS - build the following and test 

-- Riot package && Riot Struct package, should consider joining these later 
perform HTTP request to riot API
- Get Summoner info from name - done
- Be able to check if summoner is in game - not necessary for now 
- Be able to check summoners w/l  - done 

-- Riot package - maybe make a seperate polling package for reuseability 
Poll the api
20 requests every 1 seconds(s) limit
100 requests every 2 minutes(s) limit
- Build out a polling function with configurable tick - Can poll API need to figure out how to extract information, as we can overwrite a single struct a db may not be necessary initially 
- Channel built to send information after performing request, need to look into turning my main.go into a function now and encapsulate everything it does

-- Psql/Mongo package
Store info from API
- poll api and store info
- upon finding changes we want to see if player has just lost a game

-- Insults package
Flame player
- If we find a player has lost we want to flame them 
- Build a set of remarks to annoy someone

-- Discord package
Connect to Discord server
- Be able to connect to a discord server
- Send a your flame in a channel 
=======
Functionality coming 