package deko

import . "github.com/gregoryv/web"

func NewDeko() *Element {
	p := Project("Deko project specification",

		Goal(`Simplify time keeping between consultants and
		customers`),

		Background(`Working by the hour involves keeping track of
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
		CEO to do it when sending the invoice and there are
		contractors who expect the consultant to report the hours
		directly.`),

		H2("Current state of affairs"),

		P(`Preferits consultants use excel spreadsheets for gathering
		the hours and reporting them to the CEO. The file contains
		monthly sheets and a yearly summary. Consultants register a
		project per line and enter number of hours for each day. There
		are also rows for vacation, illness and other non project
		specific tasks that may need time tracking. The file is
		prepared in advance by the CEO each year, with expected number
		of working hours for each day.`),

		//
	)

	return p
}
