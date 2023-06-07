<template>
  <div class="home">
    <div class="homeLongin">
      <Form :model="formState" name="smartlogin" :labelCol="{ span: 0 }" :wrapperCol="{ span: 24 }" autocomplete="off"
        @finish="onFinish" @finishFailed="onFinishFailed" :rules="rules">
        <FormItem name="smartName">
          <Input v-model:value="formState.smartName" size="large" placeholder="输入登录邮箱" :maxlength="30" />
        </FormItem>
        <FormItem name="smartPwd">
          <InputPassword v-model:value="formState.smartPwd" size="large" placeholder="输入登录密码" :maxlength="30" />
        </FormItem>
        <FormItem name="remember" :wrapperCol="{ offset: 0, span: 24 }">
          <Checkbox v-model:checked="formState.remember">记住账户密码</Checkbox>
          <a style="float:right" href="/#/updatePassword">忘记密码</a>
        </FormItem>
        <FormItem :wrapper-col="{ offset: 0, span: 24 }">
          <Button type="primary" html-type="submit" style="width:100%" size="large">登录</Button>
        </FormItem>
      </Form>
    </div>
  </div>
</template>

<script>
import { getCurrentInstance, reactive, toRefs } from 'vue'
import Que from "./que.js";
import { Form, Input, Checkbox, Button, InputPassword } from "ant-design-vue";
import md5 from "js-md5";
import router from "@/router";

export default {
  components: { Form, FormItem: Form.Item, Input, Checkbox, Button, InputPassword },
  setup() {
    const { proxy } = getCurrentInstance();
    const state = reactive({
    })
    const formState = reactive({
      smartName: '',
      smartPwd: '',
      remember: true,
    })
    const onFinish = async (values) => {
      var time = await proxy.$req.getTime() + '';
      // const time = '1685065246000'
      const userInfo = await Que.login(proxy, {
        "emails": values.smartName,
        password: md5(md5(values.smartPwd) + time),
        time: time,
      })
      if (userInfo) {
        proxy.$req.setCookie(userInfo);
        router.push({
          name: "device",
        });
      }
    }
    const onFinishFailed = () => {
    }
    const validatePass = async (rule, value) => {
      switch (rule.field) {
        case "smartName":
          if (value == "") {
            return Promise.reject("登录邮箱不能为空")
          }
          // 验证登录邮箱
          if (!proxy.$req.checkEmail(value)) {
            return Promise.reject("登录邮箱有误")
          }
          break;
        case "smartPwd":
          if (value == "") {
            return Promise.reject("登录密码不能为空")
          }
          if (!proxy.$req.checPassword(value)) {
            return Promise.reject("登录密码有误")
          }
          break;
        default:
          break;
      }
      return Promise.resolve()
    }
    const rules = reactive({
      smartName: [{
        required: true,
        validator: validatePass,
        trigger: 'change',
      }],
      smartPwd: [{
        required: true,
        validator: validatePass,
        trigger: 'change',
      }],
    })
    return {
      ...toRefs(state), formState, rules, onFinish, onFinishFailed,
    }
  }
}
</script>

<style>
.home {
  background-image: url("@/assets/beijing.jpg");
  height: 100%;
  widows: 100%;
  -moz-background-size: 100% 100%;
  background-size: 100% 100%;
}

.homeLongin {
  width: 400px;
  background: #FFFFFF;
  margin-right: 20%;
  float: right;
  margin-top: 15%;
  border-radius: 20px;
  padding: 40px;
  opacity: 0.9;
}
</style>