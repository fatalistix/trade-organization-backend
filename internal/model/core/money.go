package core

import "fmt"

type Money struct {
	Rubles  int64
	Pennies int8
}

func (m Money) ByteArray() []byte {
	return []byte(fmt.Sprintf("%d%d", m.Rubles, m.Pennies))
}
