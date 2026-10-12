package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func args(s string) []string {
	return strings.Split(s, " ")
}

func TestState_parseArgs(t *testing.T) {

	t.Run("pass in nil", func(t *testing.T) {
		s := new(state)
		assert.NotPanics(t, func() {
			s.parseArgs(nil)
		})
		assert.Equal(t, ".", s.rootDir)
		assert.Nil(t, s.excludedOrgs)
		assert.Nil(t, s.selectedOrgs)
	})

	t.Run("pass in emptiness", func(t *testing.T) {
		s := new(state)
		assert.NotPanics(t, func() {
			s.parseArgs([]string{})
		})
		assert.Equal(t, ".", s.rootDir)
		assert.Nil(t, s.excludedOrgs)
		assert.Nil(t, s.selectedOrgs)
	})

	t.Run("custom root", func(t *testing.T) {
		s := new(state)
		assert.NotPanics(t, func() {
			s.parseArgs(args("/home/billgates/Documents"))
		})
		assert.Equal(t, "/home/billgates/Documents", s.rootDir)
		assert.Nil(t, s.excludedOrgs)
		assert.Nil(t, s.selectedOrgs)
	})

	t.Run("selected orgs with unspecified root", func(t *testing.T) {
		s := new(state)
		assert.NotPanics(t, func() {
			s.parseArgs(args("--with-org suckless,hashicorp"))
		})
		assert.Equal(t, ".", s.rootDir)
		assert.EqualValues(t, []string{"suckless", "hashicorp"}, s.selectedOrgs)
	})

	t.Run("selected orgs with specified root", func(t *testing.T) {
		s := new(state)
		assert.NotPanics(t, func() {
			s.parseArgs(args("--with-org suckless,hashicorp /my/mirror"))
		})
		assert.Equal(t, "/my/mirror", s.rootDir)
		assert.EqualValues(t, []string{"suckless", "hashicorp"}, s.selectedOrgs)
	})

	t.Run("selected orgs by including flag twice", func(t *testing.T) {
		s := new(state)
		assert.NotPanics(t, func() {
			s.parseArgs(args("--with-org suckless --with-org hashicorp"))
		})
		assert.EqualValues(t, []string{"suckless", "hashicorp"}, s.selectedOrgs)
	})

	t.Run("without some orgs", func(t *testing.T) {
		s := new(state)
		assert.NotPanics(t, func() {
			s.parseArgs(args("--without-org ibm,redhat,oracle,microsoft,salesforce,cisco,citrix"))
		})
		assert.Contains(t, s.excludedOrgs, "redhat")
		assert.Contains(t, s.excludedOrgs, "ibm")
		assert.Contains(t, s.excludedOrgs, "salesforce")
	})

}
