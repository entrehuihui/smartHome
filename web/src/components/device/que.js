import { notification } from "ant-design-vue";

export default {
  get: async (proxy, paprm) => {
    const response = await proxy.$req.get("/v1/device", paprm);
    notification
    if (response.code) {
      notification["error"]({
        message: "获取列表失败",
        description: response.message,
      });
      return null;
    }
    return response;
  },
  put: async (proxy, paprm) => {
    const response = await proxy.$req.put("/v1/device", paprm);
    notification
    if (response.code) {
      notification["error"]({
        message: "修改失败",
        description: response.message,
      });
      return null;
    }
    notification["success"]({
      message: "修改成功",
      description: response.message,
    });
    return true;
  },
};
