#!/bin/bash

CONGENIAL_OCTO_BASSOON_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

cob () {
    # if no command given force help page
    local OPTION
	if [[ "$1" != "" ]]; then
        OPTION=$1
    else
        OPTION="help"
    fi
	# handle input options
    case "${OPTION}" in
        'help')
echo "Usage: $ ${FUNCNAME} [option]

Options:
- unit: run unit tests
- bench: run bench test
"
        ;;
        'unit')
          # Test main
          go test .

          # Test sub directories
          for directory in $( find . -type d | grep '/test$' ); do
            go test "${directory}" -coverpkg=./...
          done
        ;;
        'bench')
          go test -bench=.
        ;;
        *)
          echo -e "ERROR: invalid option. Try..\n$ ${FUNCNAME} help"
        ;;
    esac
}
