## filecoin-network-stats Deploy
### a note about local development

1.分别修改 `common/frontend/backend package.json`中的`"bignumber.js": "^8.0.1"`    

2.安装

        cd common
        npm i 
        //（如果bignumber.js未到8.1）
        // 运行 npm i npm install bignumber.js
        npm run build
        sudo npm link
        cd ../backend
        npm i
        npm run build
        npm link filecoin-network-stats-common
        cd ../frontend
        npm install
        npm run dev
        npm link filecoin-network-stats-common

3.排错

        若安装过程中有问题
        到common
        执行：
        rm -rf node_modules
        npm i
        npm run build