#!/bin/zsh
# migrate -source file://./migrations -database postgres://anya:anya@localhost:5432/order_api force 20220221140629

migrate -source file://./migrations -database postgres://anya:anya@localhost:5432/order_api up
