package handlers

import "github.com/nxy7/tcp-assignment/internal/storage"

type HandlerCtx struct {
	db storage.Storage
}

func FromEnv() HandlerCtx {
	return HandlerCtx{}
}
