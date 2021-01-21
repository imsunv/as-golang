package info

import (
	"testing"
)

func TestInfo_Visit(t *testing.T) {

	info := Info{}

	var v Visitor = &info

	v = LogVisitor{v}
	v = NameVisitor{v}
	v = OtherThingsVisitor{v}

	loadFile := func(info *Info, err error) error {
		info.Name = "Hao Chen"
		info.Namespace = "MegaEase"
		info.OtherThings = "We are running as remote team."
		return nil
	}
	v.Visit(loadFile)
}

func NameVisitor2(info *Info, err error) error {
	return nil
}

func TestInfo_Visit2(t *testing.T) {

	info := Info{}
	loadFile := func(info *Info, err error) error {
		info.Name = "Hao Chen"
		info.Namespace = "MegaEase"
		info.OtherThings = "We are running as remote team."
		return nil
	}

	var v Visitor = &info
	v = NewDecoratedVisitor(v, NameVisitor2)

	v.Visit(loadFile)
}
