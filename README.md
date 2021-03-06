"# beaver" 
  some high level goals:
 - maybe nuke the current folder in between builds, so we know there are no artifacts
 - stamp each level with the current values, and all the dependencies, so we can trace it. 
 - be able to zip whole folder and transplant to another pc and be able to use this as a cache
 - encryption for build verification? add some encrypted hash or something, if can decrypt is we know we made the build
 - mqtt hooks for everything.  this is cool because we can do the whole dashboard stuff 
 - way to init beaver with system parameters so we can do things like determine if we are building on arm vs x86?  
 -- system parameter ideas:  {{ System.Time , System.OS, System.Arch }} 
 -- also some passed in config values that can be used (config, build.sh overrides, env overrides?)
 
usage:
~~~~
beaver init:
creates the below structure

beaver package 
packages up the tar of the build folder

beaver unpackage 
unzips the folders

beaver verify 
makes sure config is correct

beaver start 
starts build server
.
beaver stop
stops build server

beaver rebuild <resource name>
forces a rebuild of a resource (and cascades the builds upward)

beaver list all
lists the build status of all the resources

beaver list failing			--> should failed build stop the cascade upward? 
lists the failing builds 

beaver list passing
lists the successful builds

beaver list held 
lists the builds that would be done if failing builds were corrected

beaver tree 
prints a graph of the dependencies

beaver grab <artifact name> <destination name> <info destination>

beaver mqtton <servername> <topic-prefix> (defaults to /beaver)
publish data to mqtt server

beaver mqttoff
publish data to mqtt server

beaver start 
beaver stop

~~~~
Topics:

/beaver/status 
	on/off
	
/beaver/<dependency-name>/info    				  // general info about a dependency (name, description, etc)	
/beaver/<dependency-name>/build-status	          // building, done, waiting, error
/beaver/<dependency-name>/last-build-time         // last build time of  the dependency
/beaver/<dependency-name>/last-build-length       // length of last build
~~~~


~~~~

specify build project in a configuration file. 

folder structure needs to look like:

~~~~
build-tools/			// idea of this being separate from plugins is so this could be relocated on dumb build boxes
	github
		/ setup.sh
		/ build.sh 
		/ teardown.sh
	slack 
		/ same as above
	binary-artifacts
		/ same as above
		
		

plugins/
	github
		/ setup.sh
		/ is-valid-option.sh	
	    / copy-output.sh  
		/ add-resource.sh
		/ teardown.sh
	slack 
		/ same as above
	binary-artifacts
		-> add resource might actually pull the binaries?
	
		
			
builds/
	build-output/		--> probably should have a staging area, so we do not copy unless the build passes.
		topresource/
		resource-name1/
		someinlineresourcename/  
		coolbinary/
	additional-data
		binary-versions
			coolbinary:0.2/
				artificats go here
		
		
~~~~	

build-configuration.yaml

which looks like
~~~~
name: topresource
plugin-type: github
prebuild-hook: somescript1.sh
postbuild-hook: somescript2.sh
should-continue: somescript3.sh
should-build: somescript4.sh
time-limit:   1m  // kill the build after this amount of time
after-time: somescript.sh  // script you can call after a certain amount of time (maybe to send a notification)
check-dependencies: somescripttovalidateifdependenciesshapevalid.sh (maybe type thing would be cool here)
labels: 
	- projectName: project name here	// this allows you to do: beaver list projects
						// beaver list {.type="buildThings" } // etc
						// idea is that this is simple metadata to be able to query

options: 
	github-url: someurltogithubproject	
depends-on:	
	- resource-name
	- name: someinlineresourcename
	  plugin-type: slack
	  options: "default channel"
	- resource-name: cool binary
	  plugin-type: binary-versions
	  options:
		url: /filesystempathtobinary

		version: 0.2 (latest?)
	  

~~~~

---

main flow
- parse config
- add depenedencies to main logic, and setup plugins slaves
- respond to events sent by slave (ready/notready/error/buildcomplete)
- control slaves via build, teardown, setup, etc