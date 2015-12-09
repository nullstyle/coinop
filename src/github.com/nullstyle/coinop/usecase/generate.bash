#! /bin/bash
set -e

#

#HACK: in lieu of forking mockery and making it suit my needs, I choose
#      to hack up its output.

echo "package usecase_test" > ports_test.go

mockery -inpkg -all -print | grep import | sort | uniq >> ports_test.go
mockery -inpkg -all -print | grep -v -E "import|package" >> ports_test.go

gofmt -w ports_test.go
