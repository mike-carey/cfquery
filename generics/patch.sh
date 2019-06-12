#!/usr/bin/env bash

###
# This file modifies the generated file from ifacemaker to import cfclient and modify the classes that needed to be referenced  from that package.
##

function patch() {
  local pkg="$PWD"
  pkg="${pkg/$GOPATH\//}"
  pkg="${pkg/src\//}"

  local keep_backup=false
  local args=()
  local ignore=()
  local inject_test_imports=false
  while [[ -n "${1:-}" ]]; do
    case $1 in
      --keep-backup )
        keep_backup=true
        ;;
      --ignore )
        ignore+=("$2")
        shift
        ;;
      --inject-test-imports )
        inject_test_imports=true
        shift
        ;;
      -- )
        shift
        args+=($@)
        break
        ;;
      * )
        args+=("$1")
        ;;
    esac
    shift
  done

  set -- ${args[@]}

  for file in ${args[@]}; do
    if [[ " ${ignore[@]} " =~ " $file " ]]; then
      echo "Skipping $file"
      continue
    fi

    echo "Modifying $file"

    sed -i.bak 's/Cfclient//g' $file

    if [[ $inject_test_imports == true ]]; then
      echo "Injecting test imports into $file"
      for import in '. "github.com/onsi/ginkgo"' '. "github.com/onsi/gomega"' '. "'"$pkg"'"'; do
        sed -i.bak 's|\(import (\)|\1\
'$'\t'"$import"'|g' $file
      done
    fi

    if [[ $keep_backup != true ]]; then
      rm -f $file.bak
    else
      echo "Keeping backup $file.bak"
    fi
  done

}

if [[ ${BASH_SOURCE[0]} != $0 ]]; then
export -f patch
else
set -euo pipefail

patch "${@:-}"
exit $?
fi
