<template>
  <ConfigProvider :locale="zh_CN">
    <div id="app">
      <router-view />
    </div>
  </ConfigProvider>
</template>

<script>
import zh_CN from "ant-design-vue/lib/locale-provider/zh_CN";
import { ConfigProvider } from "ant-design-vue";
export default {
  name: "App",
  components: { ConfigProvider },
  setup() {
    return {
      zh_CN,
    };
  },
};
</script>

<style>
html,
body,
#app {
  height: 100%;
  width: 100%;
}

body {
  padding: 0;
  margin: 0;
}
</style>
