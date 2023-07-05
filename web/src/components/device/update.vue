<template>
    <Modal :visible="true" title="修改设备型号" :footer="null" @cancel="close">
        <Form :model="formState" name="smartlogin" :labelCol="{ span: 4 }" :wrapperCol="{ span: 20 }" autocomplete="off"
            @finish="onFinish" @finishFailed="onFinishFailed" :rules="rules">
            <FormItem name="deviceName" label="设备名称">
                <Input v-model:value="formState.deviceName" size="large" placeholder="输入设备型号名称" :maxlength="30" />
            </FormItem>
            <FormItem name="typeid" label="设备型号">
                <Select v-model:value="formState.typeid" show-search placeholder="设备型号" :options="typeData"
                    :filter-option="filterOption" :fieldNames="{ value: 'id', label: 'name' }">
                </Select>
            </FormItem>
            <FormItem name="groupid" label="设备分组">
                <Select v-model:value="formState.groupid" show-search placeholder="设备型号" :options="groupData"
                    :filter-option="filterOption" :fieldNames="{ value: 'id', label: 'name' }">
                </Select>
            </FormItem>
            <FormItem name="deviceDetails" label="设备备注">
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
import { Modal, Button, Form, Input, Textarea, Select } from "ant-design-vue";
import Que from "./que.js";
export default {
    name: "deviceUpdate",
    components: { Modal, Button, Form, FormItem: Form.Item, Input, Textarea, Select },
    setup(props, { emit }) {
        const { proxy } = getCurrentInstance();
        const state = reactive({
            typeData: [],
            groupData: [],
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
                case "groupid":
                    if (value < -1 || value == 0) {
                        return Promise.reject("设备组不能为空")
                    }
                    break;
                case "typeid":
                    if (value < 1) {
                        return Promise.reject("设备型号不能为空")
                    }
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
            groupid: [{
                validator: validatePass,
                trigger: 'change',
            }],
            typeid: [{
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
            state.typeData = props.typeList.slice(1)
            state.groupData = props.groupList.slice(1)
        })
        const filterOption = (input, option) => {
            return option.name.toLowerCase().indexOf(input.toLowerCase()) >= 0;
        }
        return {
            ...toRefs(state), formState, onFinish, onFinishFailed, validatePass, rules, close, filterOption
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