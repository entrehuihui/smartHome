const { defineConfig } = require("@vue/cli-service");
module.exports = defineConfig({
  transpileDependencies: true,
  devServer: {
    // 此处开启 https
    https: true,
    host: "0.0.0.0",
    port: 8080,
    client: {
      webSocketURL: 'ws://0.0.0.0:8080/ws',
    },
    headers: {
      'Access-Control-Allow-Origin': '*',
    },
    proxy: {
      // "/websocket": {
      //   target: "https://home.laoyinbi.cn:5800/",
      //   secure: false, // false为http访问，true为https访问
      //   changeOrigin: true, // 跨域访问设置，true代表跨域
      //   wss: true,
      //   // pathRewrite: { // 路径改写规则
      //   //   '^/api': '' // 以/proxy/为开头的改写为''
      //   // }
      // },
      "/v1": {
        target: "https://127.0.0.1:5800/",
        secure: false, // false为http访问，true为https访问
        changeOrigin: true, // 跨域访问设置，true代表跨域
        wss: true,
        // pathRewrite: { // 路径改写规则
        //   '^/api': '' // 以/proxy/为开头的改写为''
        // }
        },
      },
    },
  });
