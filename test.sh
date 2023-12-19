#!/bin/bash

while true; do
    # Replace "YOUR_URL_HERE" with the actual URL you want to call
    curl -s -o /dev/null -w "%{http_code}" http://localhost:4000/7Ppk5blC

    # Sleep for one minute before making the next call
    sleep 60
done
