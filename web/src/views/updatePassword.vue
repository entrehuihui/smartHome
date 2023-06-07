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
                <FormItem name="code">
                    <Space>
                        <Input v-model:value="formState.code" size="large" placeholder="输入验证码" :maxlength="6" />
                        <Button size="large" name="code" type="primary" @click="sendCode" :disabled="formState.disable">
                            <Space>
                                <span>发送验证码</span>
                                <span v-if="count > 0">{{ count }}</span>
                            </Space>
                        </Button>
                    </Space>
                </FormItem>
                <FormItem :wrapperCol="{ offset: 0, span: 24 }">
                    <Button type="primary" html-type="submit" style="width:100%" size="large">修改</Button>
                </FormItem>
                <a style="float:right" href="/#/">返回登录</a>
            </Form>
        </div>
    </div>
</template>
  
<script>
import { getCurrentInstance, reactive, toRefs } from 'vue'
import Que from "./que.js";
import { Form, Input, Button, InputPassword, Space } from "ant-design-vue";
import md5 from "js-md5";
import router from "@/router";

export default {
    components: { Form, FormItem: Form.Item, Input, Button, InputPassword, Space },
    setup() {
        const { proxy } = getCurrentInstance();
        const state = reactive({
            count: 0
        })
        const formState = reactive({
            smartName: '',
            smartPwd: '',
            code: '',
            disable: true
        })
        const onFinish = async (values) => {
            const userInfo = await Que.update(proxy, {
                "emails": values.smartName,
                password: md5(values.smartPwd),
                code: values.code,
            })
            if (userInfo) {
                proxy.$req.setCookie("");
                router.push({
                    name: "index",
                });
            }
        }
        const onFinishFailed = () => {
        }
        const validatePass = async (rule, value) => {
            switch (rule.field) {
                case "smartName":
                    if (value == "") {
                        formState.disable = true;
                        return Promise.reject("登录邮箱不能为空")
                    }
                    // 验证登录邮箱
                    if (!proxy.$req.checkEmail(value)) {
                        formState.disable = true;
                        return Promise.reject("登录邮箱有误")
                    }
                    formState.disable = false;
                    break;
                case "smartPwd":
                    if (value == "") {
                        return Promise.reject("登录密码不能为空")
                    }
                    if (!proxy.$req.checPassword(value)) {
                        return Promise.reject("登录密码有误")
                    }
                    break;
                case "code":
                    if (value == "") {
                        return Promise.reject("验证码不能为空")
                    }
                    break
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
            code: [{
                required: true,
                validator: validatePass,
                trigger: 'change',
            }],
        })
        const sendCode = async () => {
            if (!formState.smartName) {
                return
            }
            formState.disable = true;
            state.count = 60
            var timer = setInterval(() => {
                state.count--;
                if (state.count <= 0) {
                    clearInterval(timer)
                    state.count = 0;
                    formState.disable = false;
                }
            }, 1000);
            await Que.code(proxy, {
                "types": "updatepassword",
                "emails": formState.smartName,
                "name": "修改密码",
            })
        }
        return {
            ...toRefs(state), formState, rules, onFinish, onFinishFailed, sendCode
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