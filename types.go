package main

type LastModified struct {
	day   string
	month string
	year  string
}

type FileInfo struct {
	name         string
	lastModified LastModified
}

type FileIndex []FileInfo
