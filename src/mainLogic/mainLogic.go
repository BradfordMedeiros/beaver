/*
	parse config and get config
	add it to mainlogic here
	open server to start listening for alert statuses
	when we get the status we call the relevant command hooks here
	command hooks here trigger rebuilds or whatever outside
*/
// add all resources here
// add resource id
// remove resource id
// get resource status
// on resources status change
/*

	someid: Resource,
	anotherid: Resource,
	etc

	and maybe the dependency graph right here
		somideid
		/		\
	anotherid	something here

	tree has someething like this:?
	

	// maybe ready with depndencies vs ready without deps? 
	-> onReady() 	
		ready
	ready  not-ready

	goes to
		ready
	build not-ready

		ready
	complete complete
	
		build
	complete complete

		complete	
	complete complete	



*/