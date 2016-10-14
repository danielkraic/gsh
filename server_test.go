package main

import "testing"

type TestingPair struct {
	server Server
	result string
}

func TestServerString(t *testing.T) {
	list := []TestingPair{
		{Server{"", "", "", 0}, "Name: , @:0"},
		{Server{"name", "", "", 0}, "Name: name, @:0"},
		{Server{"", "user", "", 0}, "Name: , user@:0"},
		{Server{"", "", "host", 0}, "Name: , @host:0"},
		{Server{"", "", "", 2200}, "Name: , @:2200"},
		{Server{"name", "user", "host", 0}, "Name: name, user@host:0"},
		{Server{"name", "user", "host", 22}, "Name: name, user@host:22"},
	}

	for _, pair := range list {
		if pair.server.String() != pair.result {
			t.Error(
				"expected", pair.result,
				"got", pair.server.String(),
			)
		}
	}
}

func TestServerConnectionString(t *testing.T) {
	list := []TestingPair{
		{Server{"", "", "", 0}, "ssh -p 0 @ # "},
		{Server{"name", "", "", 0}, "ssh -p 0 @ # name"},
		{Server{"", "user", "", 0}, "ssh -p 0 user@ # "},
		{Server{"", "", "host", 0}, "ssh -p 0 @host # "},
		{Server{"", "", "", 2200}, "ssh -p 2200 @ # "},
		{Server{"name", "user", "host", 0}, "ssh -p 0 user@host # name"},
		{Server{"name", "user", "host", 22}, "ssh -p 22 user@host # name"},
	}

	for _, pair := range list {
		if pair.server.getConnectionString() != pair.result {
			t.Error(
				"expected", pair.result,
				"got", pair.server.getConnectionString(),
			)
		}
	}
}

func TestServerNormalize(t *testing.T) {
	list := []TestingPair{
		{Server{"", "", "", 0}, "ssh -p 22 testuser@ # "},
		{Server{"", "", "", 22}, "ssh -p 22 testuser@ # "},
		{Server{"", "", "", 44}, "ssh -p 44 testuser@ # "},
	}

	for _, pair := range list {
		pair.server.normalize("testuser", 22)

		if pair.server.getConnectionString() != pair.result {
			t.Error(
				"expected", pair.result,
				"got", pair.server.getConnectionString(),
			)
		}
	}
}

func TestServerMatch(t *testing.T) {
	list := []struct {
		server  Server
		pattern string
		matched bool
	}{
		{Server{"", "", "", 0}, "", false},
		//
		{Server{"", "", "", 0}, "a", false},
		{Server{"a", "", "", 0}, "a", true},
		{Server{"", "a", "", 0}, "a", true},
		{Server{"", "", "a", 0}, "a", true},
		{Server{"", "", "a", 0}, "a", true},
		//
		{Server{"test1", "test2", "test3", 0}, "test1", true},
		{Server{"test1", "test2", "test3", 0}, "test2", true},
		{Server{"test1", "test2", "test3", 0}, "test3", true},
		{Server{"test1", "test2", "test3", 0}, "test4", false},
	}

	for _, pair := range list {
		if pair.server.match(pair.pattern) != pair.matched {
			t.Error(
				"failed to match server:", pair.server.String(),
				",pattern:", pair.pattern,
				",expected:", pair.matched,
				",got:", !pair.matched,
			)
		}
	}
}

func TestServerMatchAll(t *testing.T) {
	list := []struct {
		server   Server
		patterns []string
		matched  bool
	}{
		{Server{"", "", "", 0}, []string{""}, false},
		//
		{Server{"", "", "", 0}, []string{"a"}, false},
		{Server{"a", "", "", 0}, []string{"a"}, true},
		{Server{"", "a", "", 0}, []string{"a"}, true},
		{Server{"", "", "a", 0}, []string{"a"}, true},
		{Server{"", "", "a", 0}, []string{"a"}, true},
		//
		{Server{"", "", "aba", 0}, []string{"a"}, true},
		{Server{"", "", "aba", 0}, []string{"b"}, true},
		//
		{Server{"test1", "test2", "test3", 0}, []string{"test1"}, true},
		{Server{"test1", "test2", "test3", 0}, []string{"test2"}, true},
		{Server{"test1", "test2", "test3", 0}, []string{"test3"}, true},
		{Server{"test1", "test2", "test3", 0}, []string{"test4"}, false},
		//
		{Server{"test1", "test2", "test3", 0}, []string{"test", "1"}, true},
		{Server{"test1", "test2", "test3", 0}, []string{"test", "2"}, true},
		{Server{"test1", "test2", "test3", 0}, []string{"test", "3"}, true},
		{Server{"test1", "test2", "test3", 0}, []string{"test", "4"}, false},
		//
		{Server{"test1", "test2", "test3", 0}, []string{"", "test1"}, false},
		{Server{"test1", "test2", "test3", 0}, []string{"test1", ""}, false},
	}

	for _, pair := range list {
		if pair.server.matchAll(pair.patterns) != pair.matched {
			t.Error(
				"failed to match server:", pair.server.String(),
				",patterns:", pair.patterns,
				",expected:", pair.matched,
				",got:", !pair.matched,
			)
		}
	}
}
