#!/bin/bash

#PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin

# Start the run once job.
cd /usr/src/app/
source .env.sh
python init.py
