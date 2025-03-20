#!/bin/sh

IMG="$(realpath $1)"

printf '\033[100S' # Scroll up by 100
printf '\033[1d' # Move cursor to row 0
printf '\033[1G' # Move cursor to column 0

printf '========= Image ============\n'
printf '\033_Gi=1,t=f,q=1,f=100,a=T,r=10,c=20;%s\033\\' "$(printf '%s' "$IMG" | base64 -w0)"
printf '<- image\n'
printf '1\n'
printf '2\n'
printf '3\n'

sleep 1

printf '\033[4T' # Scroll down by 4
printf '\033[1d' # Move cursor to row 0
printf '\033[1G' # Move cursor to column 0

sleep 5