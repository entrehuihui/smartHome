import { notification } from "ant-design-vue";

export default {
  login: async (proxy, paprm) => {
    const response = await proxy.$req.post("/v1/login/login", paprm);
    if (response.code) {
      notification["error"]({
        message: "登录失败",
        description: response.message,
      });
      return null;
    }
    return response.data;
  },
  code: async (proxy, paprm) => {
    const response = await proxy.$req.post("/v1/login/user/code", paprm);
    if (response.code) {
      notification["error"]({
        message: "发送验证码失败",
        description: response.message,
      });
      return null;
    }
    notification["success"]({
      message: "发送验证码成功",
      description: "发送验证码成功",
    });
    return response.data;
  },
  update: async (proxy, paprm) => {
    const response = await proxy.$req.put("/v1/login/password", paprm);
    if (response.code) {
      notification["error"]({
        message: "修改失败",
        description: response.message,
      });
      return null;
    }
    notification["success"]({
      message: "修改成功",
      description: "修改成功",
    });
    return true;
  },
};
