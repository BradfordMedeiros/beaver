
usage:


hooker -e someurlendpoint -p port -s path of script to call when you get the hoo

or 

hooker -c configfilepath

where the config file is url, command pairs
example:

/ touch somefile
/another notify-send hello world
/turn-on-ssh service ssh start
/turn-off-ssh service ssh stop
/somethingelse somebashscript.sh

all actions are synchronous.
