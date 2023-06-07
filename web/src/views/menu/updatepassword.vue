<template>
  <Modal :visible="true" title="修改密码" @cancel="close" @ok="postdata" :footer="null">
    <Form name="updateuserpassword" ref="ruleForm" @finish="postdata" :model="form">
      <FormItem name="password" :rules="{ required: true, validator: validator }">
        <InputPassword :disabled="closeValue" size="large" style="width: 470px" addonBefore="旧密码"
          v-model:value="form.password" :maxlength="20" />
      </FormItem>
      <FormItem name="newpassword1" :rules="{ required: true, validator: validator }">
        <InputPassword :disabled="closeValue" size="large" style="width: 470px" addonBefore="新密码"
          v-model:value="form.newpassword1" :maxlength="20" />
      </FormItem>
      <FormItem name="newpassword2" :rules="{ required: true, validator: validator }">
        <InputPassword :disabled="closeValue" size="large" style="width: 470px" addonBefore="确认密码"
          v-model:value="form.newpassword2" :maxlength="20" />
      </FormItem>

      <FormItem :wrapper-col="{ span: 16, offset: 16 }">
        <Button @click="close"> 取消 </Button>
        <Button style="margin-left: 10px" type="primary" html-type="submit">
          确定
        </Button>
      </FormItem>
    </Form>
  </Modal>
</template>

<script>
import { reactive, toRefs, getCurrentInstance } from "vue";
import router from "@/router";
import md5 from "js-md5";
import { Modal, Form, Button, Input, notification } from "ant-design-vue";
export default {
  name: "menuUpdate",
  components: {
    Modal,
    FormItem: Form.Item,
    Form: Form,
    Button: Button,
    InputPassword: Input.Password,
  },
  setup(props, { emit }) {
    const { proxy } = getCurrentInstance();
    let data = reactive({
      closeValue: false,
    });
    const refData = toRefs(data);
    let form = reactive({
      password: "",
      newpassword1: "",
      newpassword2: "",
    });
    function close() {
      if (!data.closeValue) {
        emit("my-emit", "我是子组件值");
      }
    }
    async function validator(rule, value) {
      switch (rule.field) {
        case "password":
          if (!value) {
            return Promise.reject("旧密码不能为空");
          }
          if (!proxy.$req.checPassword(value)) {
            return Promise.reject("密码为6-20位包含数组字母(大小写)");
          }
          break;
        case "newpassword1":
          if (!value) {
            return Promise.reject("新密码不能为空");
          }
          if (!proxy.$req.checPassword(value)) {
            return Promise.reject("密码为6-20位包含数组字母(大小写)");
          }
          break;
        case "newpassword2":
          if (!value) {
            return Promise.reject("新密码不能为空");
          }
          if (!proxy.$req.checPassword(value)) {
            return Promise.reject("密码为6-20位包含数组字母(大小写)");
          }
          if (form.newpassword1 != value) {
            return Promise.reject("两次新密码不相同");
          }
          break;
      }
      return Promise.resolve();
    }
    async function postdata() {
      var time = await proxy.$req.getTime();
      data.closeValue = true;
      proxy.$req
        .put("/v1/login/password/login", {
          password: md5(md5(form.password) + time),
          newPassword: md5(form.newpassword1),
          time: time,
        })
        .then((response) => {
          data.closeValue = false;
          if (response.code) {
            notification["error"]({
              message: "登录失败",
              description: response.message,
            });
            return;
          }
          router.push("/");
        });
    }
    return {
      close,
      postdata,
      form,
      validator,
      ...refData,
    };
  },
};
</script>

<style></style>
