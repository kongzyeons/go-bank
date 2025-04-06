package events

import "reflect"

var Topics = []string{
	reflect.TypeOf(AccountAddMoneyEvent{}).Name(),
	reflect.TypeOf(AccountWithldrawEvent{}).Name(),
}

type Event interface {
}
