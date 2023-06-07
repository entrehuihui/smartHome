<template>
  <div>
    <div>
      <Breadcrumb style="margin: 16px 0; float: left">
        <BreadcrumbItem>
          <span style="color: #FFFFFF; font-size: large;font-weight: bold;">主页</span>
        </BreadcrumbItem>
        <BreadcrumbItem v-for="(v, k) in breadcrumbList" :key="k">
          <span @click="breadClick(v)" style="color: #FFFFFF; font-size: large;font-weight: bold;">
            {{ v }}
          </span>
        </BreadcrumbItem>
      </Breadcrumb>
    </div>
    <div class="avatarclass" v-if="userInfo.nicename">
      <Dropdown placement="bottomRight">
        <div class="avatarspan">
          <Space>
            <Avatar :size="'large'">
              <!-- <template #icon><UserOutlined /></template> -->
              <template #icon>
                <img src="@/assets/touxiang.jpg" />
              </template>
            </Avatar>
            <span class="avatarspan1">
              {{ userInfo.nicename }}
            </span>
          </Space>
        </div>
        <template #overlay>
          <Menu @click="click">
            <MenuItem :key="0" v-if="false">
            <Space>
              <UserOutlined />
              <span> <a href="javascript:;">账号信息</a></span>
            </Space>
            </MenuItem>
            <MenuItem :key="1">
            <Space>
              <SyncOutlined />
              <span> <a href="javascript:;">修改密码</a></span>
            </Space>
            </MenuItem>
            <MenuItem :key="2">
            <Space>
              <LogoutOutlined />
              <span> <a href="javascript:;">退出登录</a></span>
            </Space>
            </MenuItem>
          </Menu>
        </template>
      </Dropdown>
    </div>
    <updatepasswordVue @my-emit="updateShow = false" v-if="updateShow" />
  </div>
</template>

<script>
// @ is an alias to /src
import { onMounted, reactive, toRefs, getCurrentInstance } from "vue";
import { Menu, Avatar, Dropdown, Breadcrumb, Space } from "ant-design-vue";
import {
  UserOutlined,
  LogoutOutlined,
  SyncOutlined,
} from "@ant-design/icons-vue";
import router from "@/router";
import updatepasswordVue from "./updatepassword.vue";

export default {
  name: "menuHead",
  components: {
    Menu,
    MenuItem: Menu.Item,
    Avatar,
    UserOutlined,
    Dropdown,
    LogoutOutlined,
    SyncOutlined,
    Breadcrumb,
    BreadcrumbItem: Breadcrumb.Item,
    updatepasswordVue, Space
  },
  setup() {
    const { proxy } = getCurrentInstance();
    let data = reactive({
      userInfo: {},
      updateShow: false,
    });
    onMounted(async () => {
      data.userInfo = proxy.$req.getCookie();
    });
    const click = ({ key }) => {
      switch (key) {
        case 0:
          router.push("/info");
          break;
        case 1:
          data.updateShow = true;
          break;
        case 2:
          proxy.$req.setCookie();
          router.push("/");
          break;
      }
    };
    const refData = toRefs(data);
    const breadClick = () => { };
    return { ...refData, click, breadClick };
  },
  props: {
    breadcrumbList: Object,
  },
};
</script>

<style>
.avatarclass {
  position: relative;
  float: right;
  cursor: pointer;
}

.avatarspan {
  height: 100%;
}

.avatarspan1 {
  color: #FFFFFF;
  margin-left: 5px;
  font-size: large;
}
</style>
