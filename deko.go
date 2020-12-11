package deko

import . "github.com/gregoryv/web"

func NewDeko() *Specification {
	n := NewHn(2)

	name := "Deko project specification"
	goal := Goal(`Simplify time keeping between consultants and
		customers`)

	background := Background(
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

	return NewSpecification(name, goal, background)
}
