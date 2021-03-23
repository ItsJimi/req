#!/usr/bin/env bash

if [ "$@" ]
then
  NAME=$@  
  req run "${NAME}" --silent
  req list
else
  req list
fi
