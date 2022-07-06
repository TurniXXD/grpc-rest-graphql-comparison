#!/bin/bash
cd ../client && yarn start && cd ..
cd go && go run main.go && cd ..
cd node && yarn start && cd ..
cd python && python main.py