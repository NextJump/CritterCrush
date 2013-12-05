# Critter Crush

Crush critters for honor and glory

## Setup

**This is still very much in flux, if you notice it's out of date please update these instructions.**

Clone the repository into your gopath

    mkdir -p $GOPATH/src/github.com/nextjump
    cd $GOPATH/src/github.com/nextjump
    git clone https://github.com/NextJump/CritterCrush.git
    
After you make changes to the code, you can recompile the code by typing
    
    go install github.com/nextjump/CritterCrush
    
and start the game server by executing 
    
    $GOPATH/bin/CritterCrush
    
You should see something like this:

    ready player one...
    
This lets you know that the server is running and listening on port 8080. If you see the "ready player one..." prompt but executable exits and returns you to the prompt, you most likely already have a process running on port 8080. You can run 

    netstat -tulpn | grep :8080
    
to see which process is running and kill it if necessary. You can also change the port numbers the server listens on inside of the base level critter.go file. 

Once the server is running, you should be able to see the game at locahost:8080/game

Still needs more tweaking before the PLAY button works with this new layout...
