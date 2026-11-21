#!/bin/zsh
go run ./main.go

line_num=$(grep -E -o -n '<svg width=".*" height=".*"' output.svg | cut -d: -f1)

# MacOS
# sed -i '' "${line_num}s/.*/<svg width=\"100%\" height=\"100%\"/g" output.svg

# Linux
sed -i "${line_num}s/.*/<svg width=\"100%\" height=\"100%\"/g" output.svg

