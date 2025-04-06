package events

import "reflect"

var Topics = []string{
	reflect.TypeOf(AccountAddMoneyEvent{}).Name(),
}

type Event interface {
}
