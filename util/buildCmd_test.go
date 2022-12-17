package util

import (
	"testing"
)

func TestBuildCmd(t *testing.T) {
	builder := &CmdBuilder{}
	builder.Add("git")
	builder.Add("clone")
	builder.Add("ssh:test.com/aaa.git")
	builder.Add("-b master")
	builder.Add("-o origin")
	command := builder.Build()
	if "git clone ssh:test.com/aaa.git -b master -o origin" != command {
		t.Error("build command result is error")
	}

	builder.Reset()
	builder.Add("git")
	builder.Add("pull")
	builder.Add("-b master")
	builder.Add("-o origin")
	command = builder.Build()
	if "git pull -b master -o origin" != command {
		t.Error("build command result is error")
	}

	builder.Reset()
	builder.Add("git")
	builder.Add("pull")
	builder.Add("-b master")
	builder.AddWithArgAndCond("-o origin", false)
	command = builder.Build()
	if "git pull -b master" != command {
		t.Error("build command result is error")
	}

	builder.Reset()
	builder.Add("git")
	builder.Add("pull")
	builder.Add("-b master")
	builder.AddWithArg("-o %s", "origin")
	command = builder.Build()
	if "git pull -b master -o origin" != command {
		t.Error("build command result is error")
	}

	builder.Reset()
	builder.Add("git")
	builder.Add("pull")
	builder.Add("-b master")
	builder.AddWithArgAndCond("-o %s", false, "origin")
	command = builder.Build()
	if "git pull -b master" != command {
		t.Error("build command result is error")
	}
}

func TestBuildCmdWithArgs(t *testing.T) {
	builder := &CmdBuilder{}
	builder.Add("git")
	builder.AddWithArg("%s %s", "remote", "-v")
	command := builder.Build()
	if "git remote -v" != command {
		t.Error("build command result is error")
	}
}

func TestBuildCmdWithCond(t *testing.T) {
	builder := &CmdBuilder{}
	builder.Add("git")
	builder.Add("clone")
	builder.Add("ssh:test.com/aaa.git")
	builder.Add("-b master")
	builder.AddWithCond("-o origin", false)
	command := builder.Build()
	if "git clone ssh:test.com/aaa.git -b master" != command {
		t.Error("build command result is error")
	}
}
