package view

import (
	"bytes"
	"cmp"
	"encoding/gob"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/wdipax/match/state/group/member"
)

type Poll struct {
	Options []string
}

func (p *Poll) Encode(members []*member.Member) string {
	var options []string

	for _, m := range members {
		options = append(options, formatLabel(m))
	}

	var buf bytes.Buffer

	enc := gob.NewEncoder(&buf)

	err := enc.Encode(options)
	if err != nil {
		log.Printf("encoding poll options: %s", err)
	}

	return buf.String()
}

func (p *Poll) Decode(from string) {
	r := strings.NewReader(from)

	dec := gob.NewDecoder(r)

	err := dec.Decode(&p.Options)
	if err != nil {
		log.Printf("decoding poll options: %s", err)
	}
}

type PollResult struct {
	Choosen []int
}

func (p *PollResult) Encode(choosen []int) string {
	var buf bytes.Buffer

	enc := gob.NewEncoder(&buf)

	err := enc.Encode(choosen)
	if err != nil {
		log.Printf("encoding chosen poll results: %s", err)
	}

	return buf.String()
}

func (p *PollResult) Decode(from string) {
	r := strings.NewReader(from)

	dec := gob.NewDecoder(r)

	err := dec.Decode(&p.Choosen)
	if err != nil {
		log.Printf("decoding poll responses: %s", err)
	}
}

func UserNumberFrom(label string) int {
	s := strings.SplitN(label, " ", 2)

	var n int

	if len(s) > 0 {
		var err error

		n, err = strconv.Atoi(s[0])
		if err != nil {
			log.Printf("parsing label number: %s", err)
		}
	}

	return n
}

func GroupMembers(members []*member.Member) string {
	sortMembers(members)

	var acc []string

	for _, m := range members {
		acc = append(acc, formatLabel(m))
	}

	return strings.Join(acc, "\n")
}

func FormPoll(members []*member.Member) string {
	sortMembers(members)

	var p Poll

	return p.Encode(members)
}

func Matches(members []*member.Member) string {
	var acc []string

	sortMembers(members)

	for _, m := range members {
		acc = append(acc, fmt.Sprintf("%d %s @%s", m.Number, m.Name, m.Contact))
	}

	return strings.Join(acc, "\n")
}

func sortMembers(members []*member.Member) {
	slices.SortFunc(members, func(a, b *member.Member) int {
		if n := cmp.Compare(a.Number, b.Number); n != 0 {
			return n
		}

		return cmp.Compare(a.Name, b.Name)
	})
}

func formatLabel(m *member.Member) string {
	return fmt.Sprintf("%d %s", m.Number, m.Name)
}
