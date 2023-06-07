import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";
import store from "./store";
import "ant-design-vue/dist/antd.css";
// import Antd from "ant-design-vue";

var app = createApp(App);

import req from "./global/req.js";
app.config.globalProperties.$req = req;
app.use(store).use(router).mount("#app");
