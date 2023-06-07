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
};
