"# beaver" 
  some high level goals:
 - maybe nuke the current folder in between builds, so we know there are no artifacts
 - stamp each level with the current values, and all the dependencies, so we can trace it. 
 - be able to zip whole folder and transplant to another pc and be able to use this as a cache
 - encryption for build verification? add some encrypted hash or something, if can decrypt is we know we made the build
 - mqtt hooks for everything.  this is cool because we can do the whole dashboard stuff
 
 - question: to include or not to include specific versions?
	-- leaning toward no since i want things to be using up to date dependencies anyway
	-- if you really need it to be, use  proper language  package managers
	-- i  guess....
 
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
	binary-artificacts
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
check-dependencies: somescripttovalidateifdependenciesshapevalid.sh (maybe type thing would be cool here)

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