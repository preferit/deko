package deko

import . "github.com/gregoryv/web"

func NewCurrentState(n *Hn) *Element {
	return Wrap(
		n.H1("Current state of affairs"),

		n.H2("Consultant track working hours"),

		P(`Preferits consultants use excel file for gathering the
	    hours and reporting them to the CEO. The file is prepared in
	    advance by the CEO each year, with expected number of working
	    hours for each day. The file contains monthly
	    sheets and a yearly summary. Consultants register a project
	    per line and enter number of hours for each day. There are
	    also rows for vacation, illness and other non project specific
	    tasks that may need time tracking.`),

		n.H2("CEO prepares excel file for consultant"),

		P(`Each year the CEO creates a new excel file based on a
		previous one. He goes through each month and verifies that the
		calendar is correct and changes the expected working hours for
		holidays. Once this is done the file is duplicated for each
		consultant. If the consultant has overtime left from last year
		it added to this years file. The new file is then sent out to
		the consultant.`),

		n.H2("Consultant reports working hours to CEO"),

		P(`Consultant sends mails the excel file to the CEO`),

		n.H2("Consultant reports working hours to contractor"),

		P(`Consultant opens the excel file to find the summary for a
		specific project and mails that summary to the contractor.`),

		n.H2("CEO creates invoice for customer"),

		P(`CEO opens excelfile for consultant that is contracted and
		finds the summary for the specific project.`,

			Question(`What do you do with the summary? enter it into
		    an invocing system, which?`),
		),

		//
	)
}
