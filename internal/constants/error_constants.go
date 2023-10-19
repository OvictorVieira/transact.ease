package constants

import "errors"

var (
	ErrWhenGetFiles  = errors.New("error when get files name")
	ErrWhenReadFiles = errors.New("error when read files name")

	ErrLoadConfig  = errors.New("failed to load config file")
	ErrParseConfig = errors.New("failed to parse env to config struct")
	ErrEmptyVar    = errors.New("required variable environment is empty")
)
