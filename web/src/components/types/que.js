import { notification } from "ant-design-vue";

export default {
  get: async (proxy, paprm) => {
    const response = await proxy.$req.get("/v1/types", paprm);
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
  post: async (proxy, paprm) => {
    const response = await proxy.$req.post("/v1/types", paprm);
    notification
    if (response.code) {
      notification["error"]({
        message: "添加失败",
        description: response.message,
      });
      return null;
    }
    notification["success"]({
      message: "添加成功",
      description: response.message,
    });
    return true;
  },
  put: async (proxy, paprm) => {
    const response = await proxy.$req.put("/v1/types", paprm);
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
  del: async (proxy, paprm) => {
    const response = await proxy.$req.del("/v1/types", paprm);
    notification
    if (response.code) {
      notification["error"]({
        message: "删除失败",
        description: response.message,
      });
      return null;
    }
    notification["success"]({
      message: "删除成功",
      description: response.message,
    });
    return true;
  },
};
