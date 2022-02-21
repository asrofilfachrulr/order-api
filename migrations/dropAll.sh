#!/bin/zsh
migrate -source file://./migrations -database postgres://anya:anya@localhost:5432/order_api drop -f
echo "every table inside order_api has been dropped"
