
plugins rely on defining an alerting mechanism, and a build mechanism.

alerting mechanisms ultimately need to call the function alert $1.
Alert $1 causes the plugin to signifiy to beaver that it needs to be rebuilt
setup/teardown can be used for any purpose, but as an example you might want to setup a simple webserver to handle incoming webhooks.

ultimately this causes beaver to invoke the build mechanism.  
the build mechanism passes in the options defined in the beaver config

you can theoretically couple some of these mechanisms.  do not do this unless you know beaver will only be run on a single machine, and its still probably useful not to if possible (can optimize for performance as a tradeoff by not creating multiple http servers, etc)
 
build mechansism i want to be able to create slaves to process.   alerting mechansism are not worth as much to scale horizontally, but maybe eventually.

also keep in mind that the alerting and build mechanisms should be able to be run  against  more  than a single instance.
for example you may have several build products of a single type.  they should also be designed to run across multiple boxes, although it might make 
sense build against a single box if you know this will be run on a single server.


