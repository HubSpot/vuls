package models

type Exception struct {
	Cve       string
	Reason    string
	Submitter string
}

type ExceptionEntries struct {
	Exceptions map[string]Exception
}
