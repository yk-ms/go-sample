# Makefile for building and running the Go application with Docker Compose

.PHONY: build run docker-build docker-run

build:
	docker compose build

up:
	docker compose up
