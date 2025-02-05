// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package io

import (
	"encoding/json"
	"os"
)

func Exist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

type SafeFile struct {
	*os.File
	filename string
	closed   bool
}

func NewSafeFile(filename string) (*SafeFile, error) {
	if ok, err := Exist(filename); ok || err != nil {
		return nil, os.ErrExist
	}
	file, err := os.Create(filename)
	if err != nil {
		return nil, err
	}

	return &SafeFile{filename: filename, File: file}, nil
}

func (f *SafeFile) Write(p []byte) (n int, err error) {
	if f.closed {
		return 0, os.ErrClosed
	}

	return f.File.Write(p)
}

func (f *SafeFile) Close() error {
	if f.closed {
		return os.ErrClosed
	}

	defer func() {
		if r := recover(); r != nil {
			f.Rollback()
		}
	}()

	err := f.File.Close()
	if err != nil {
		return err
	}

	f.closed = true
	return nil
}

func (f *SafeFile) Rollback() {
	if !f.closed {
		f.Close()
		os.Remove(f.filename)
	}
}

func ReadJSON(filename string, v any) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, v)
	if err != nil {
		return err
	}

	return nil
}

func WriteFileFrom(src, dst string) error {
	data, err := os.ReadFile(src)
	if err != nil {
		return err
	}
	err = os.WriteFile(dst, data, 0666)
	if err != nil {
		return err
	}
	return nil
}
