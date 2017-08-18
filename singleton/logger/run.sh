#!/usr/bin/env bash

CMD="go build"
echo "$CMD :: execute..."
go build log_level.go logger.go
if [ "$?" == 0 ]; then
	echo "$CMD executed successfully"
else
	echo "$CMD failed"
	exit;
fi

CMD="go install"
echo "$CMD :: execute..."
go install
if [ "$?" == 0 ]; then
	echo "$CMD executed successfully"
else
	echo "$CMD failed"
	exit;
fi

CMD="go test"
echo "$CMD :: execute..."
go test log_level.go logger.go logger_test.go
if [ "$?" == 0 ]; then
	echo "$CMD executed successfully"
else
	echo "$CMD failed"
	exit;
fi




