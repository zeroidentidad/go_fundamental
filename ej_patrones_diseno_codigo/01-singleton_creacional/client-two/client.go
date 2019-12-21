package client_two

import "../singleton"

func IncrementAge() {
	p := singleton.GetInstance()
	p.IncrementAge()
}
