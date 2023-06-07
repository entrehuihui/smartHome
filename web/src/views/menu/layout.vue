<template>
  <Layout class="homeLayout">
    <LayoutHeader style="
        background: #4CAF50;
        border-bottom: 1px solid black;
        min-width: 600px;
      ">
      <headVue :breadcrumbList="breadcrumbList" />
    </LayoutHeader>
    <Layout>
      <!-- 左侧导航 -->
      <LayoutSider :trigger="null" collapsible v-model:collapsed="collapsed" width="150" style="background: #fff">
        <div class="triggerdiv" @click="() => (collapsed = !collapsed)">
          <Tooltip v-if="collapsed">
            <template #title>展开菜单栏</template>
            <MenuUnfoldOutlined class="trigger" />
          </Tooltip>
          <Tooltip v-else>
            <template #title>收起菜单栏</template>
            <MenuFoldOutlined class="trigger" />
          </Tooltip>
        </div>
        <menuVue @my-emit="cliidclose" />
      </LayoutSider>
      <Layout>
        <LayoutContent :style="{
            padding: '5px',
            margin: 0,
            minHeight: '280px',
            minWidth: '600px',
          }">
          <slot></slot>
        </LayoutContent>
      </Layout>
    </Layout>
  </Layout>
</template>

<script>
// @ is an alias to /src
import { onMounted, reactive, toRefs } from "vue";
import { Layout, Tooltip } from "ant-design-vue";
import headVue from "./head.vue";
import menuVue from "./menu.vue";
import { MenuUnfoldOutlined, MenuFoldOutlined } from "@ant-design/icons-vue";

export default {
  name: "menuHome",
  components: {
    Layout,
    LayoutContent: Layout.Content,
    LayoutSider: Layout.Sider,
    LayoutHeader: Layout.Header,
    Tooltip,
    MenuUnfoldOutlined,
    MenuFoldOutlined,
    headVue,
    menuVue,
  },
  setup() {
    // const { proxy } = getCurrentInstance();
    let data = reactive({
      collapsed: false,
      breadcrumbList: [],
    });

    onMounted(() => { });
    if (document.body.clientWidth <= 600) {
      data.collapsed = true;
    }
    const refData = toRefs(data);
    const cliidclose = (value) => {
      data.breadcrumbList = value;
    };
    return {
      ...refData,
      cliidclose,
    };
  },
};
</script>

<style>
.homeLayout {
  height: 100%;
  width: 100%;
  background: whitesmoke;
}

.trigger {
  font-size: 18px;
  line-height: 64px;
  padding: 0 5px;
  cursor: pointer;
  transition: color 0.3s;
}

.trigger:hover {
  color: #1890ff;
}

.triggerdiv {
  position: absolute;
  left: 100%;
  width: 25px;
  cursor: pointer;
  z-index: 99999;
  top: -22px;
}
</style>
