#!/bin/bash

host=bandit.labs.overthewire.org
#host=localhost
pass=UoMYTrfrBFHyQXmg6gzctqAwOmw1IohZ

makefile() {
  file=/tmp/pwn25-tries
  > $file
  for i in {0000..9999}; do
    echo "$pass $i" >> "$file"
  done
}

makefile

#nc "$host" 30002 -w 10

