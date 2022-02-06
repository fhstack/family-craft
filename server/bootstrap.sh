#!/usr/bin/bash 
nohup python ./output/cat_lock_api.py >cat_lock_api.log 2>&1 &
nohup ./output/server > server.log 2>&1 &
