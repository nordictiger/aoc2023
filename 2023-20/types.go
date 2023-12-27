package main

import (
	"fmt"
	"slices"
	"strings"
)

type moduleType int

const (
	Broadcaster moduleType = iota
	FlipFlop
	Conjunction
)

func (mt moduleType) String() string {
	return [...]string{"bc ", "%", "&"}[mt]
}

type state int

const (
	Low state = iota
	High
)

func (s state) getReverse() state {
	if s == Low {
		return High
	} else {
		return Low
	}
}

func (s state) String() string {
	return [...]string{"0", "1"}[s]
}

type incoming map[string]state

func (i incoming) String() string {
	var keys []string
	for key := range i {
		keys = append(keys, key)
	}
	slices.Sort(keys)

	var builder strings.Builder
	for _, k := range keys {
		builder.WriteString(fmt.Sprintf("%s:%v ", k, i[k]))
	}
	return builder.String()
}

type node struct {
	name       string
	moduleType moduleType
	state      state
	incoming   incoming
	outgoing   []string
}

func (n node) String() string {
	return fmt.Sprintf("%s:%v-%v, in: %v, out: %v", n.name, n.moduleType, n.state, n.incoming, n.outgoing)
}

type moduleConfiguration map[string]node

func (mc moduleConfiguration) String() string {
	var keys []string
	for key := range mc {
		keys = append(keys, key)
	}
	slices.Sort(keys)

	var builder strings.Builder
	for _, k := range keys {
		builder.WriteString(fmt.Sprintf("%s\n", mc[k]))
	}
	return builder.String()
}

type signal struct {
	source      string
	level       state
	destination string
}

func (s signal) String() string {
	return fmt.Sprintf("%s -%v-> %s", s.source, s.level, s.destination)
}
