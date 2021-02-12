package deko

import (
	. "github.com/gregoryv/web"
	"github.com/preferit/deko/spec"
	. "github.com/preferit/deko/spec"
)

func NewDeko() *Specification {
	n := NewHn(2)

	s := spec.Specification{
		Name:         "Deko project specification",
		LastUpdate:   LastUpdate, // from changelog.go
		Goals:        NewGoals(n),
		CurrentState: NewCurrentState(n),
		Changelog:    NewChangelog(n),
		References:   NewReferences(n),
	}
	return &s
}

func NewGoals(n *Hn) *Element {
	main := MainGoal("Simplify time keeping between consultants and customers.")

	return Section(
		n.H1("Goals"),

		P(`The main goal of this project is to "`, main, `".`),

		n.H2("Background"),
		P(`Working by the hour involves keeping track of
		those working hours and at certain intervals transform the
		accumulated time to an invoice for the customer.`),

		P(`Preferit is a small business with ten consultants and one
		chief executive officer(CEO). Consultants report back to the
		CEO on a monthly basis. He then compiles the times to one or
		more invoices and sends them to customers. Sometimes the
		customer is another consultantcy firm which has contracted
		Preferit. Also at times when consultants are sub-contracted
		this way, working hours need to be reported both to the CEO
		and the other contracting firm. Sometimes it's enough for the
		CEO to do it when sending the invoice, but there are
		contractors who expect the consultant to report the hours
		directly.`),

		//
	)
}

func NewCurrentState(n *Hn) *Element {
	s := Section(
		n.H1("Current state"),

		P(`By describing how time keeping is currently managed we
		define expected features, highlight issues and elicit
		requirements for future solutions.`),

		n.H2("Consultant logs working hours"),

		P(`Preferits consultants use spreadsheet file for logging the
	    hours and reporting them to the CEO. The file is prepared in
	    advance by the CEO each year, with expected number of working
	    hours for each day. The file contains monthly sheets and a
	    yearly summary. Consultants register a project per line and
	    enter number of hours for each day. There are also rows for
	    vacation, illness and other non project specific tasks that
	    may need time keeping.`),

		Question("What features does current log process provide?"),

		n.H3("Issues"),
		Issues(

			Issue(`The monthly sheet is quite large and it's easy to
		    update the wrong cell for any given date.`),

			Issue(`Even if consultant has worked the expected number
		    of hours, the sheet must always be updated.`),

			Issue(`Almost impossible to keep track of time using
			mobile device.`),
			//
		),

		// ----------------------------------------

		n.H2("CEO prepares spreadsheet file for consultant"),

		P(`Each year the CEO creates a new spreadsheet file based on a
		previous one. He goes through each month and verifies that the
		calendar is correct and changes the expected working hours for
		holidays. Once this is done the file is duplicated for each
		consultant. If the consultant has overtime left from last year
		it's added to this years file. The new file is then sent out
		to the consultant.`),

		n.H3("Issues"),
		Issues(

			Issue(`It's tedious to update the holidays with the
			expected working hours. If a mistake is made and later
			found, it must be manually updated for each distributed
			spreadsheet file.`),
			//
		),

		// ----------------------------------------

		n.H2("Consultant reports working hours to CEO"),

		P(`Consultant e-mails the spreadsheet file to the CEO`),

		n.H3("Issues"),
		Issues(

			Issue(`Cannot report working hours unless access to
			spreadsheet file.`),

			Issue(`Easy to send the wrong file.`),

			//
		),

		// ----------------------------------------

		n.H2("Consultant reports working hours to contractor"),

		P(`Consultant opens the spreadsheet file to find the summary for a
		specific project and e-mails that summary to the contractor.`),

		n.H3("Issues"),
		Issues(

			Issue(`Cannot report working hours unless access to
			spreadsheet file.`),

			Issue(`Hard to tell if all expected days are logged.`),
			//
		),

		// ----------------------------------------

		n.H2("CEO creates invoice for customer"),

		P(`CEO opens spreadsheet file for consultant that is
		contracted and finds the summary for the specific project. The
		accumulated hours that are eligible for invoicing are manually
		entered into the financial system.`,
		),

		n.H3("Requirements"),
		Requirements(
			rami73,
			ryre95,
		),

		//
	)
	return s
}

// Keep requirements unique and referencable
var (
	rami73 = &Requirement{
		ID: "rami73",

		Txt: `CEO must easily receive/find the monthly sum eligible
		for invoicing.`,
		//
	}

	ryre95 = &Requirement{
		ID: "ryre95",

		Txt: `Monthly sum should include hours and minutes`,
		//
	}
	//
)

func NewReferences(n *Hn) *Element {
	return Section(
		n.H1("References"),

		Dl(

			Dt("Financial system"),
			Dd(`Fortnox is the financial system used to produce
			invoices.`),

			Dt("Spreadsheet file"),
			Dd("Excel file which works with libreoffice"),

			Dt("Time keeping"),
			Dd(`The process of logging, reporting and transforming
			working hours into invoices.`),
			//
		),
	)
}
