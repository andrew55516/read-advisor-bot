package telegram

const msgHelp = `I can save and keep your pages. Also I can offer you them to read.
In order to save the page, just send me a link to it.
In order to get a random page from your list, send me command /rnd.
Caution! After that, this page will be removes from your list!`

const msgHello = "Hi there! \n\n" + msgHelp

const (
	msgUnknownCommand   = "Unknown command"
	msgNoSavedPages     = "You have no saved pages"
	msgSaved            = "Saved!"
	msgAlreadyExists    = "You have already have this page"
	msgOops             = "Oops... Something went wrong..."
	msgStorageNotExists = "Storage does not exist. Please, send me at least one page"
)
