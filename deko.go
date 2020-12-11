package deko

import . "github.com/gregoryv/web"

func NewDeko() *Specification {
	n := NewHn(2)

	s := &Specification{
		name: "Deko project specification",
	}
	s.goals = Section(
		P(`The main goal of this project is to "`,
			MainGoal("Simplify time keeping between consultants and customers."),

			`". Time keeping is defined as the process of logging,
			reporting and transforming working hours into invoices.`,
		),
	)

	s.background = Background(
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

		NewCurrentState(n),
		//
	)
	s.changelog = NewChangelog(n)
	return s
}

func NewCurrentState(n *Hn) *Element {
	return Wrap(
		n.H1("Current state"),

		n.H2("Consultant track working hours"),

		P(`Preferits consultants use spreadsheet file for gathering
	    the hours and reporting them to the CEO. The file is prepared
	    in advance by the CEO each year, with expected number of
	    working hours for each day. The file contains monthly sheets
	    and a yearly summary. Consultants register a project per line
	    and enter number of hours for each day. There are also rows
	    for vacation, illness and other non project specific tasks
	    that may need time tracking.`),

		n.H2("CEO prepares spreadsheet file for consultant"),

		P(`Each year the CEO creates a new spreadsheet file based on a
		previous one. He goes through each month and verifies that the
		calendar is correct and changes the expected working hours for
		holidays. Once this is done the file is duplicated for each
		consultant. If the consultant has overtime left from last year
		it added to this years file. The new file is then sent out to
		the consultant.`),

		n.H2("Consultant reports working hours to CEO"),

		P(`Consultant e-mails the spreadsheet file to the CEO`),

		n.H2("Consultant reports working hours to contractor"),

		P(`Consultant opens the spreadsheet file to find the summary for a
		specific project and e-mails that summary to the contractor.`),

		n.H2("CEO creates invoice for customer"),

		P(`CEO opens spreadsheet file for consultant that is
		contracted and finds the summary for the specific project. The
		accumulated hours that are eligible for invoicing are manually
		entered into the financial system Fortnox.`,
		),

		n.H3("Requirements"),
		Requirements(R1),

		//
	)
}
