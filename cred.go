package main

import (
	"encoding/csv"
	"os"
	"os/user"
	"path/filepath"

	"golang.org/x/xerrors"
)

var (
	credPath string
)

func init() {
	var user, _ = user.Current()
	credPath = filepath.Join(user.HomeDir, ".twit-go")
}

type cred struct {
	consumerKey    string
	consumerSecret string
	accessToken    string
	accessSecret   string
}

func save(consumerKey, consumerSecret, token, secret string) error {
	file, err := os.Create(credPath)
	if err != nil {
		return xerrors.Errorf("Create file failed: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	if err = writer.Write([]string{consumerKey, consumerSecret, token, secret}); err != nil {
		return xerrors.Errorf("Write file failed: %v", err)
	}
	writer.Flush()
	return nil
}

func load() (*cred, error) {
	file, err := os.Open(credPath)
	if err != nil {
		return nil, xerrors.Errorf("Open file failed: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ','
	record, err := reader.Read()
	if err != nil {
		return nil, xerrors.Errorf("Read file failed: %v", err)
	}
	return &cred{
		consumerKey:    record[0],
		consumerSecret: record[1],
		accessToken:    record[2],
		accessSecret:   record[3],
	}, nil
}
