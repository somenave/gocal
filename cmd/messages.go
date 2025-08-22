package cmd

const (
	invalidInputErrMsg       = "Invalid input"
	invalidCommandErrMsg     = "Invalid command, type 'help' for list of supported commands."
	savedCalendarMsg         = "Calendar saved"
	addFormatMsg             = "format: add 'event name' 'date' 'priority'"
	chooseMsg                = "use 'Tab' for choose"
	removeSuccessMsg         = "Event has been removed"
	updateFormatMsg          = "format: update 'event ID' 'name' 'date' 'priority'"
	updateSuccessMsg         = "Event has been updated"
	setReminderFormatMsg     = "format: reminder:set 'event ID' 'message' 'date'"
	setReminderSuccessMsg    = "Reminder has been set"
	removeReminderFormatMsg  = "format: reminder:remove 'event ID'"
	removeReminderSuccessMsg = "Reminder has been removed"
	cancelReminderFormatMsg  = "format: reminder:cancel 'event ID'"
	cancelReminderSuccessMsg = "Reminder has been cancelled"
	noEventsMsg              = "No events found"
)
