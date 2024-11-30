#!/bin/bash
for i in {1..4}; do curl -is -w "Request $i: %{http_code}\n" -o /dev/null "http://localhost:8080/01001001"; done
echo "wait for block duration: 5s" && sleep 5
curl -is -w "status: %{http_code}\n" -o /dev/null http://localhost:8080/01001001

## Limitação por token
for i in {1..6}; do curl -is -w "Request $i: %{http_code}\n" -o /dev/null -H "API_KEY: 3851b1ae73ca0ca6e3c24a0256a80ace" http://localhost:8080/01001001; done
echo "wait for block duration: 5s" && sleep 5
curl -is -w "status: %{http_code}\n" -o /dev/null -H "API_KEY: 3851b1ae73ca0ca6e3c24a0256a80ace" "http://localhost:8080/01001001"