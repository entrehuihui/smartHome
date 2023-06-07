<template>
  <div>
    <Menu mode="inline" v-model:selectedKeys="selectedKeys" @click="click" v-model:openKeys="openKeys"
      style="user-select: none">
      <MenuItem v-for="v in routes" :key="v.name">
      <span>
        <ZoomInOutlined />
        <span>
          {{ v.details }}
        </span>
      </span>
      </MenuItem>
    </Menu>
  </div>
</template>

<script>
import { Menu } from "ant-design-vue";
import router from "@/router";
import { onMounted, reactive, toRefs, } from "vue";
import {
  ZoomInOutlined
} from "@ant-design/icons-vue";
export default {
  name: "menuMenu",
  components: {
    Menu,
    MenuItem: Menu.Item, ZoomInOutlined
  },
  setup(props, { emit }) {
    let data = reactive({
      selectedKeys: [0],
      openKeys: [],
      keys: {},
      routes: [],
      auth: {},
    });
    const click = ({ keyPath }) => {
      var path = keyPath.reverse().join("/");
      router.push("/" + path);
    };
    onMounted(async () => {
      const routes = router.options.routes;
      var local = router.options.history.location;
      for (const v of routes) {
        if (v.details) {
          data.routes.push(v);
          if (local == v.path) {
            local = v.details
            data.selectedKeys = [v.name];
          }
        }
      }
      emit("my-emit", [local]);
    });
    const refData = toRefs(data);
    return {
      click,
      ...refData,
    };
  },
};
</script>

<style></style>
