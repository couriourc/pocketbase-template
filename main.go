package main

import (
	pb_adapter "app/pkg/adapters/pocket-base"
	domain_core "app/pkg/domain/domain-core"
)

func main() {
	//creates a instance/reference to the pocket-base initialization.
	app := pb_adapter.PBNewAPIServer()

	//binds the pocketbase app to the Interface, so that the functions defined in the interface could be accessed with the pocketbase as the reciever
	server := domain_core.BinderPBInterface(app)

	//calls the run server
	server.Run()

}
