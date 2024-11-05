#!/usr/bin/env bash

CURRENT_WID=$(xdotool getwindowfocus)

WID=$(xdotool search --name "Firefox")
xdotool windowactivate $WID
sleep 0.8
xdotool key F5

xdotool windowactivate $CURRENT_WID
