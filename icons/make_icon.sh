#/bin/sh

if [ -z "$GOPATH" ]; then
    echo GOPATH environment variable not set
    exit
fi

if [ ! -e "$GOPATH/bin/2goarray" ]; then
    echo "Installing 2goarray..."
    go get github.com/cratonica/2goarray
    if [ $? -ne 0 ]; then
        echo Failure executing go get github.com/cratonica/2goarray
        exit
    fi
fi

if [ -z "$1" ]; then
    echo Please specify a PNG file
    exit
fi

if [ ! -f "$1" ]; then
    echo $1 is not a valid file
    exit
fi

file="$1"
file_name=${file%.*}
file_name_upper_case="$(tr '[:lower:]' '[:upper:]' <<< ${file_name:0:1})${file_name:1}"

OUTPUT=${file_name}unix.go
echo Generating $OUTPUT
echo "//+build linux darwin" > $OUTPUT
echo >> $OUTPUT
cat "$file" | $GOPATH/bin/2goarray "${file_name_upper_case}Data" icons >> $OUTPUT
if [ $? -ne 0 ]; then
    echo Failure generating $OUTPUT
    exit
fi
echo Finished
