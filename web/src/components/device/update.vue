<template>
    <Modal :visible="true" title="修改设备型号" :footer="null" @cancel="close">
        <Form :model="formState" name="smartlogin" :labelCol="{ span: 0 }" :wrapperCol="{ span: 24 }" autocomplete="off"
            @finish="onFinish" @finishFailed="onFinishFailed" :rules="rules">
            <FormItem name="deviceName">
                <Input v-model:value="formState.deviceName" size="large" placeholder="输入设备型号名称" :maxlength="30" />
            </FormItem>
            <FormItem name="deviceDetails">
                <Textarea v-model:value="formState.deviceDetails" size="large" placeholder="输入设备型号备注" :maxlength="300" />
            </FormItem>
            <FormItem :wrapperCol="{ offset: 0, span: 24 }">
                <Button type="primary" html-type="submit" style="width:100%">确定</Button>
            </FormItem>
        </Form>
    </Modal>
</template>

<script>
import { reactive, toRefs, getCurrentInstance, onMounted } from 'vue'
import { Modal, Button, Form, Input, Textarea } from "ant-design-vue";
import Que from "./que.js";
export default {
    name: "deviceUpdate",
    components: { Modal, Button, Form, FormItem: Form.Item, Input, Textarea },
    setup(props, { emit }) {
        const { proxy } = getCurrentInstance();
        const state = reactive({
        })
        const formState = reactive({
            deviceName: '',
            deviceDetails: '',
            groupid: -1,
            typeid: -1,
        })
        const validatePass = async (rule, value) => {
            switch (rule.field) {
                case "deviceName":
                    if (value == "") {
                        formState.disable = true;
                        return Promise.reject("设备型号名称不能为空")
                    }
                    formState.disable = false;
                    break;
                case "deviceDetails":
                    break;
                default:
                    break;
            }
            return Promise.resolve()
        }
        const rules = reactive({
            deviceName: [{
                required: true,
                validator: validatePass,
                trigger: 'change',
            }],
            deviceDetails: [{
                validator: validatePass,
                trigger: 'change',
            }],
        })
        const onFinish = async (values) => {
            console.log(values);
            var resp = await Que.put(proxy, {
                "name": values.deviceName,
                "details": values.deviceDetails,
                "id": props.updateInfo.id,
                "groupid": values.groupid,
                "typeid": values.typeid,
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
        onMounted(() => {
            formState.deviceName = props.updateInfo.name
            formState.deviceDetails = props.updateInfo.details
            formState.groupid = props.updateInfo.groupid
            formState.typeid = props.updateInfo.typeid
        })
        return {
            ...toRefs(state), formState, onFinish, onFinishFailed, validatePass, rules, close
        }
    },
    props: {
        updateInfo: Object,
        typeList: Array,
        groupList: Array
    }
}
</script>

<style lang="scss" scoped></style>