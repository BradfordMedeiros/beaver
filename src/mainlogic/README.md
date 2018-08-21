
We have a graph of dependencies, where nodes are only built when 
1) the plugin instance (aka node) signals it is ready to build
and 
2) the builds for all dependent plugin instances/ nodes are in a complete state

Each node has a local and global state:

Local State (per node)
------------
- LOCAL_NOTREADY
- LOCAL_READY
is-building
last-build 

Local ready or not ready indicates whether or not the plugin has signaled that we the node is ready to be built. 
As an example for a github plugin, this might mean a user has pushed code to github, and code is ready to be rebuilt.  
Or, for an ftp excel plugin, this might mean that a plugin has pulled in an excel configuration on a schedule, and wants to 
rebuild it for the weekly excel configuration update.

Local state is simply the plugin saying "i want to rebuild" and has nothing to do with other nodes in the configuration.

We do not rebuild just because the plugin says so, since we honor the dependency tree, which leads into the discussion 
about global state.

Global State (per node)
------------
- NOTREADY
- READY
- QUEUED
- INPROGRESS
- COMPLETE

Each node travels through a state machine where valid transitions are
NOTREADY -> READY      						| node has signaled LOCAL_READY, and all node dependencies are in a COMPLETE state
READY -> QUEUED        						| scheduler adds the node to be queued after it is ready
QUEUED -> INPROGRESS   						| build.sh in the plugin is called, and node put into the inprogress state 
INPROGRESS -> COMPLETE 						| the node signals back that the node build for the node is complete
INPROGRESS -> NOTREADY 						| the node finishes building, but while building the node signaled not ready (or a dep)
COMPLETE -> READY 	   						| the node signals that it is ready
COMPLETE -> NOT READY | dependency becomes NOT_COMPLETE or local node signals NOT_READY

For each node, this the relationship from a parent node can be described as:
Given: S(n) => global state of node n, L(n) => local state of n, and D(n) => dependencies of node n
S(n) = QUEUED | INPROGRESS | COMPLETE |  all(D(n) = COMPLETE) and L(n) = READY
S(n) = NOTREADY |  any(D(n) not = COMPLETE) or L(n) = NOTREADY

Note that if a node is queued, or inprogress, the node will finish building and transitioning through the states, however