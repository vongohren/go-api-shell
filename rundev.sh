#/bin/bash
#export the needed enviroment variables here

#This is an example of dockerhosted rethink database
export DB_PORT_28015_TCP_ADDR=192.168.99.100
export DB_PORT_28015_TCP_PORT=28015
#My facebook callback hostname, changed when deployed
export DNS_HOSTNAME="localhost"

#This if sentence allows you to run godebug, a go debugging tool for
if [ $# -gt 0 ] ; then
  if [[ $1 -eq "debug" ]] ; then
    godebug run *.go
  else
    p=""
    for var in "$@" ; do
     if [[  $p != "" ]] ; then
        p="$p,$var"
     else
        p=$var
     fi
    done
    echo "go debug with $p"
    godebug run -instrument=$p *.go
  fi
else
  echo "simple go"
  go run *.go
fi
