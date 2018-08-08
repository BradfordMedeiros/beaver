
plugins rely on defining an alerting mechanism, and a build mechanism.

alerting mechanisms ultimately need to call the function alert $1.
Alert $1 causes the plugin to signifiy to beaver that it needs to be rebuilt
setup/teardown can be used for any purpose, but as an example you might want to setup a simple webserver to handle incoming webhooks.

ultimately this causes beaver to invoke the build mechanism.  
the build mechanism passes in the options defined in the beaver config

you can theoretically couple some of these mechanisms.  In general coupling the alerting mechanisms is probably safe if that fits a use case.
do not want to couple the alert with the build mechanisms, b/c want to support master slave mode that may result in these running on different boxes.


also keep in mind that the alerting and build mechanisms should be able to be run  against  more  than a single instance.
for example you may have several build products of a single type.  they should also be designed to run across multiple boxes, although it might make 
sense build against a single box if you know this will be run on a single server.
