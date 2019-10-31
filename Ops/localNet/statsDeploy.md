## filecoin-network-stats Deploy
### a note about local development

1.下载与预配置

    git clone https://github.com/filecoin-project/filecoin-network-stats.git
    分别修改 `common/frontend/backend package.json`中的`"bignumber.js": "^8.0.1"`    

2.安装
        
        //super user
        cd common
        npm i 
        npm run build
        sudo npm link
        注意：确保common/node_modules中有bignumber.js
             若没有，运行 npm install bignumber.js
        
        cd ../backend
        npm i
        npm link filecoin-network-stats-common
        npm run build
        
        cd ../frontend
        npm install
        npm link filecoin-network-stats-common
        npm run dev
        
        编译成功后，进行数据库迁移操作
        1.创建数据库用户 stats
        2.创建数据库 filecoin-network-stats
        3.add a connection to a server
        4.到backend下，npm run-script migrate:local


3.排错

        若安装过程中有问题
        1.到common
          执行：
          rm -rf node_modules lib
          npm i
          npm run build
          npm link
          再到F/B
        
        2.bignumber.js
          到common/node_modules查看是否存在 bignumber.js
            
4.远程部署改动
        
          devServer: {
            contentBase: path.join(__dirname, 'dist'),
            historyApiFallback: true,
            overlay: true
            Host:xxx.xxx.x.xx
            Port:xxxx
          },

5.其他链接

    frontend/src/components/GlobalNav.tsx 修改超链接
