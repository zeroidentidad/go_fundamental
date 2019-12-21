package client_one

import "../singleton"

func IncrementAge() {
	p := singleton.GetInstance()
	p.IncrementAge()
}
