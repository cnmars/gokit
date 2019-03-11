/*
 * A toolkit for Golang development
 * https://www.likexian.com/
 *
 * Copyright 2019, Li Kexian
 * Released under the Apache License, Version 2.0
 *
 */

package xtar

import (
	"github.com/likexian/gokit/assert"
	"github.com/likexian/gokit/xfile"
	"os"
	"os/exec"
	"testing"
)

var (
	err error
	dst = "targz.tar.gz"
)

func TestVersion(t *testing.T) {
	assert.NotEqual(t, Version(), "")
	assert.NotEqual(t, Author(), "")
	assert.NotEqual(t, License(), "")
}

func TestCreate(t *testing.T) {
	defer os.Remove(dst)

	err = Create(dst)
	assert.NotNil(t, err)

	err = Create(dst, "no.go")
	assert.NotNil(t, err)

	err = Create(dst, "xtar.go")
	assert.NotNil(t, err)

	os.Remove(dst)
	err = Create(dst, "/dev/null")
	assert.NotNil(t, err)

	os.Remove(dst)
	err = Create(dst, "xtar.go")
	assert.Nil(t, err)

	os.Remove(dst)
	err = Create(dst, "../xtar")
	assert.Nil(t, err)
}

func TestExtract(t *testing.T) {
	defer os.Remove(dst)

	err = Create(dst, "../assert", "../LICENSE")
	assert.Nil(t, err)

	err = Extract("no.tar.gz", "")
	assert.NotNil(t, err)

	err = Extract("targz.go", "")
	assert.NotNil(t, err)

	err = Extract(dst, "")
	assert.Nil(t, err)

	assert.True(t, xfile.IsDir("assert"))
	assert.True(t, xfile.IsFile("LICENSE"))

	os.RemoveAll("assert")
	os.RemoveAll("LICENSE")

	err = Extract(dst, "tmp")
	assert.Nil(t, err)

	assert.True(t, xfile.IsDir("tmp/assert"))
	assert.True(t, xfile.IsFile("tmp/LICENSE"))

	os.RemoveAll("tmp")
}

func TestComdec(t *testing.T) {
	tar := "xtar.tar"
	tgz := "xtar.tar.gz"

	err = Create(tar, "xtar.go")
	assert.Nil(t, err)

	err = Extract(tar, "tmp")
	assert.Nil(t, err)

	os.Remove(tar)
	os.RemoveAll("tmp")

	err = Create(tgz, "xtar.go")
	assert.Nil(t, err)

	err = Extract(tgz, "tmp")
	assert.Nil(t, err)

	os.Remove(tgz)
	os.RemoveAll("tmp")
}

func TestWithSysTar(t *testing.T) {
	tar := "xtar.tar"
	tgz := "xtar.tar.gz"

	exec.Command("tar", "zcvf", tar, "xtar.go").Run()
	assert.True(t, xfile.Exists(tar))

	err = Extract(tar, "tmp")
	assert.NotNil(t, err)

	err = os.Rename(tar, tgz)
	assert.Nil(t, err)

	err = Extract(tgz, "tmp")
	assert.Nil(t, err)

	os.Remove(tgz)
	os.RemoveAll("tmp")

	exec.Command("tar", "cvf", tgz, "xtar.go").Run()
	assert.True(t, xfile.Exists(tgz))

	err = Extract(tgz, "tmp")
	assert.NotNil(t, err)

	err = os.Rename(tgz, tar)
	assert.Nil(t, err)

	err = Extract(tar, "tmp")
	assert.Nil(t, err)

	os.Remove(tar)
	os.RemoveAll("tmp")
}

func TestIsGzName(t *testing.T) {
	assert.True(t, IsGzName("targz.tgz"))
	assert.True(t, IsGzName("targz.tar.gz"))
	assert.False(t, IsGzName("targz.tar"))
	assert.False(t, IsGzName("targz.gz"))
}