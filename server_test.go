package main

import "testing"

type TestingPair struct {
	server Server
	result string
}

func TestServerString(t *testing.T) {
	list := []TestingPair{
		{Server{"", "", "", 0}, ", @:0"},
		{Server{"name", "", "", 0}, "name, @:0"},
		{Server{"", "user", "", 0}, ", user@:0"},
		{Server{"", "", "host", 0}, ", @host:0"},
		{Server{"", "", "", 2200}, ", @:2200"},
		{Server{"name", "user", "host", 0}, "name, user@host:0"},
		{Server{"name", "user", "host", 22}, "name, user@host:22"},
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
		{Server{"", "", "", 0}, "ssh -p 0 @"},
		{Server{"name", "", "", 0}, "ssh -p 0 @"},
		{Server{"", "user", "", 0}, "ssh -p 0 user@"},
		{Server{"", "", "host", 0}, "ssh -p 0 @host"},
		{Server{"", "", "", 2200}, "ssh -p 2200 @"},
		{Server{"name", "user", "host", 0}, "ssh -p 0 user@host"},
		{Server{"name", "user", "host", 22}, "ssh -p 22 user@host"},
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
		{Server{"", "", "", 0}, "ssh -p 22 testuser@"},
		{Server{"", "", "", 22}, "ssh -p 22 testuser@"},
		{Server{"", "", "", 44}, "ssh -p 44 testuser@"},
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

func TestServerValidate(t *testing.T) {
	list := []struct {
		server  Server
		isValid bool
	}{
		{Server{"", "", "", 0}, false},
		{Server{"a", "", "", 0}, false},
		{Server{"", "a", "", 0}, false},
		{Server{"", "", "a", 0}, false},
		//
		{Server{"", "test2", "test3", 0}, false},
		{Server{"test1", "", "test3", 0}, false},
		{Server{"test1", "test2", "", 0}, false},
		//
		{Server{"test1", "test2", "test3", 0}, true},
	}

	for _, pair := range list {
		if (pair.server.validate() == nil) != pair.isValid {
			t.Error(
				"failed to validate server:", pair.server.String(),
				",expected:", pair.isValid,
				",got:", pair.server.validate() == nil,
				",err:", pair.server.validate(),
			)
		}
	}
}

