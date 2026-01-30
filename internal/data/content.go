// Filename: internal/data/content.go
// - Contains content used throughout the program.
package data

// Routes contains a collection of route path and text.
var Routes = map[string]Route{
	"home": {
		Path: "/",
		Text: "Hey, Andres here! Welcome to my CMPS3161 Lab 1! My semester project will be on Hotel Room & Housekeeping Management since I recently engaged with the Westin hotel in the Boston Seaport District. It will be fun investigating how hotels in Belize are similar or different.\n",
	},
	"about": {
		Path: "/about",
		Text: "My name is Andres Hung. I'm currently a 2nd semester bachelor's student in IT at the University of Belize. A hobby that I enjoy is drawing. Currently I use a MacBook Air at school and at home my PC dual boots Windows 11 and NixOS. I like computers.\n",
	},
	"contact": {
		Path: "/contact",
		Text: "My student email can be contacted at 2018118240@ub.edu.bz. My GitHub username is andreshungbz. It is the same username that I typically use for other websites too.\n",
	},
	"calculate": {
		Path: "/calculate",
		Text: "The result of 6 + 7 = 13!\n",
	},
}
