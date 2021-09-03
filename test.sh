#!/bin/sh

# git clone ...
# check ret value
# cd directory
# check ret value
# npm install
# check ret value
# npm run start
# check ret value

# if something went wrong, curl http webhook to say what step failed and 

retValue=$?
if [ $retValue != 0 ]; then
    echo "something have a problem"
fi