clear
# 1) clean up previous installation
echo '> Cleaning up previous installation...'
initialDir=$(pwd)
echo '> Entering $GOPATH/bin dir at: '$GOPATH
cd $GOPATH/bin
rm -rf views/ labSoft2_Estoque_client
cd $initialDir
echo '>> Done!'

# 2) Execute go install
echo '> Preparing to execute "go install"...'
go install
echo '>> Done!'

# 4) Copy views/ to bin
echo '> Copying views/ to $GOPATH/bin...'
cp -r views/ $GOPATH/bin
echo '>> Done!'

echo '>> Installation Complete!'

# 5) Run server
echo '>> Executing client'
cd $GOPATH/bin/
./labSoft2_Estoque_client
