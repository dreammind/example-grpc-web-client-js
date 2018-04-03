export GOPATH=$(echo `pwd` | sed -E 's:(^.*/go).*:\1:')
echo "[info] GOPATH=$GOPATH"

export PATH="$GOPATH/bin:$PATH"

