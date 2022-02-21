#!/bin/zsh
migrate -source file://./migrations -database postgres://anya:anya@localhost:5432/order_api down -all
