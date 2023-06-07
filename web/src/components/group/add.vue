<template>
    <Modal :visible="true" title="添加设备组" :footer="null" @cancel="close">
        <Form :model="formState" name="smartlogin" :labelCol="{ span: 0 }" :wrapperCol="{ span: 24 }" autocomplete="off"
            @finish="onFinish" @finishFailed="onFinishFailed" :rules="rules">
            <FormItem name="groupName">
                <Input v-model:value="formState.groupName" size="large" placeholder="输入设备组名称" :maxlength="30" />
            </FormItem>
            <FormItem name="groupDetails">
                <Textarea v-model:value="formState.groupDetails" size="large" placeholder="输入设备组备注" :maxlength="300" />
            </FormItem>
            <FormItem :wrapperCol="{ offset: 0, span: 24 }">
                <Button type="primary" html-type="submit" style="width:100%">确定</Button>
            </FormItem>
        </Form>
    </Modal>
</template>

<script>
import { reactive, toRefs, getCurrentInstance } from 'vue'
import { Modal, Button, Form, Input, Textarea } from "ant-design-vue";
import Que from "./que.js";
export default {
    name: "groupAdd",
    components: { Modal, Button, Form, FormItem: Form.Item, Input, Textarea },
    setup(props, { emit }) {
        const { proxy } = getCurrentInstance();
        const state = reactive({
        })
        const formState = reactive({
            groupName: '',
            groupDetails: '',
        })
        const validatePass = async (rule, value) => {
            switch (rule.field) {
                case "groupName":
                    if (value == "") {
                        formState.disable = true;
                        return Promise.reject("设备组名称不能为空")
                    }
                    formState.disable = false;
                    break;
                case "groupDetails":
                    break;
                default:
                    break;
            }
            return Promise.resolve()
        }
        const rules = reactive({
            groupName: [{
                required: true,
                validator: validatePass,
                trigger: 'change',
            }],
            groupDetails: [{
                validator: validatePass,
                trigger: 'change',
            }],
        })
        const onFinish = async (values) => {
            console.log(values);
            var resp = await Que.post(proxy, {
                "name": values.groupName,
                "details": values.groupDetails,
            })
            if (resp) {
                close()
            }
        }
        const onFinishFailed = () => {
        }
        const close = () => {
            emit("close")
        }
        return {
            ...toRefs(state), formState, onFinish, onFinishFailed, validatePass, rules, close
        }
    },
}
</script>

<style lang="scss" scoped></style>