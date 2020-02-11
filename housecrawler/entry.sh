#!/bin/bash

# We need env var for cron
printenv | sed 's/^\(.*\)$/export \1/g' > .env.sh

# Setup a cron schedule
echo "0 */1 * * * /usr/src/app/run.sh >> /var/log/cron.log 2>&1
# This extra line makes it a valid cron" > scheduler.txt

crontab scheduler.txt
cron -f
