
cd $GOPATH/src/github.com/filecoin-project/filecoin-network-stats/common/
export http_proxy="http://127.0.0.1:12333"
export https_proxy="http://127.0.0.1:12333"

# git clone https://github.com/filecoin-project/filecoin-network-stats.git
rm -rf node_modules lib
wait
npm i
wait
npm run build
wait
sudo npm link
cd ../backend
npm i
wait
sudo npm link filecoin-network-stats-common
npm run build
wait
cd ../frontend
npm install
wait
sudo npm link filecoin-network-stats-common
npm run dev