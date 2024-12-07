#!/bin/zsh
  if [[ $# -ne 1 ]]; then
    echo "Only one arg accepted!"
    return 1
  fi

  local in=$1
  if [[ -e "$in" ]]; then
    if [[ -f "$in" ]]; then
      local out="$in.bak"
      cp "$in" "$out"
    elif [[ -d "$in" ]]; then
      local out="$in.bak"
      cp -r "$in" "$out"
    else
      echo "exists but neither"
    fi
  else
    echo "$in does not exist"
    return 1
  fi
