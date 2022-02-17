// Package controllers
/*
	Package controller implements control over views from user input.

	Each struct defines a controller over a model and hosts the current context and the database implementation.
	For most cases, each controller also defines a NewController() method (constructor), and various methods to interact
	and validate user input and therefore correctly send data to the database by a call to the corresponding view.

	The principal use case of the controllers will be as Gin Handlers
*/
package controllers
