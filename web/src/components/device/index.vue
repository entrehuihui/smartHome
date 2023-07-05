<template>
    <layoutVue>
        <Space direction="vertical" style="width: 100%;">
            <div>
                <div style="float:left; padding:5px">
                    设备型号:
                    <Select v-model:value="typeid" show-search placeholder="设备型号" style="width: 250px;" :options="typeList"
                        :filter-option="filterOption" :fieldNames="{ value: 'id', label: 'name' }">
                    </Select>
                </div>
                <div style="float:left; padding:5px">
                    设备组:
                    <Select v-model:value="groupid" show-search placeholder="设备型号" style="width: 250px;"
                        :options="groupList" :filter-option="filterOption" :fieldNames="{ value: 'id', label: 'name' }">
                    </Select>
                </div>
                <div style="float:left; padding:5px">
                    <Input style="width:250px" placeholder="模糊检索" v-model:value="fuzzy" allow-clear :allowClear="true"
                        :maxlength="30"></Input>
                </div>
                <div style="float:left; padding:5px">
                    <Button style="width:70px; padding:5px" type="primary" @click="showSizeChange">检索</Button>
                </div>
            </div>
            <div>
                <Table :columns="columns" :data-source="data" :rowKey="(record) => record.ID" :pagination="false"
                    :expandRowByClick="false" :defaultExpandAllRows="true" :loading="loading" size="small">
                    <template #bodyCell="{ column, record, index }">
                        <template v-if="column.dataIndex === 'index'">
                            <span>
                                {{ index + (offset - 1) * limit + 1 }}
                            </span>
                        </template>
                        <template v-else-if="column.dataIndex === 'createtime'">
                            {{ dayjs(parseInt(record.createtime)).format("YYYY-MM-DD HH:mm:ss") }}
                        </template>
                        <template v-else-if="column.dataIndex === 'typeid'">
                            {{ searchName(record.typeid, typeList) }}
                        </template>
                        <template v-else-if="column.dataIndex === 'groupid'">
                            {{ searchName(record.groupid, groupList) }}
                        </template>
                        <template v-else-if="column.dataIndex === 'action'">
                            <Space>
                                <Button type="dashed" @click="updateInfo = record; updateShow = true">
                                    修改
                                </Button>
                                <Popconfirm title="是否确定删除该设备组?" ok-text="确定" cancel-text="取消" @confirm=" confirm(record) ">
                                    <Button type="dashed" danger :disabled=" true ">
                                        删除
                                    </Button>
                                </Popconfirm>
                            </Space>
                        </template>
                    </template>
                </Table>
                <Pagination :total=" total " :showTotal=" (total) => `一共 ${total}条` " :pageSize=" limit "
                    @change=" paginationChange " @showSizeChange=" showSizeChange " v-model:current=" offset "
                    style="float: right" :showLessItems=" true " :showSizeChanger=" true ">

                </Pagination>
            </div>
        </Space>
        <Update v-if=" updateShow " @close=" closeChild " :updateInfo=" updateInfo " :typeList=" typeList "
            :groupList=" groupList " />
    </layoutVue>
</template>

<script>
import { reactive, toRefs, onMounted, getCurrentInstance } from 'vue'
import layoutVue from "@/views/menu/layout.vue";
import { Table, Button, Space, Pagination, Popconfirm, Input, Select } from "ant-design-vue";
import Que from "./que.js";
import dayjs from "dayjs";
import groupQue from '../group/que';
import typeQue from '../types/que';
import Update from './update.vue';

export default {
    name: "groupIndex",
    components: { layoutVue, Table, Button, Space, Pagination, Popconfirm, Input, Select, Update },
    setup() {
        const { proxy } = getCurrentInstance();
        const state = reactive({
            dayjs,
            data: [],
            limit: 10,
            offset: 1,
            fuzzy: "",
            loading: false,
            total: 0,
            typeList: [],
            groupList: [],
            typeid: 0,
            groupid: 0,
            updateShow: false,
            updateInfo: []
        })
        const columns = reactive([
            {
                title: "序号",
                dataIndex: 'index',
                width: 70,
            },
            {
                title: "SN",
                dataIndex: 'sn',
            },
            {
                title: "名称",
                dataIndex: 'name',
            },
            {
                title: "设备型号",
                dataIndex: 'typeid',
            },
            {
                title: "设备组",
                dataIndex: 'groupid',
            },
            {
                title: "备注",
                dataIndex: 'details',
            },
            {
                title: "创建时间",
                dataIndex: 'createtime',
            },
            {
                title: '操作',
                dataIndex: 'action',
            }
        ])
        const closeChild = () => {
            state.updateShow = false;
            state.updateInfo = {}
            getData();
        }
        const getData = async () => {
            state.loading = true
            var resp = await Que.get(proxy, {
                limit: state.limit,
                offset: (state.offset - 1) * state.limit,
                fuzzy: state.fuzzy,
                typeid: state.typeid,
                groupid: state.groupid,
            })
            state.loading = false
            if (resp) {
                state.data = resp.data
                state.total = parseInt(resp.total)
            }
        }
        const paginationChange = () => {
            getData();
        }
        const showSizeChange = (current, size) => {
            state.offset = 1;
            if (size) {
                state.limit = size;
            }
            getData();
        }
        const confirm = async (record) => {
            await Que.del(proxy, record.id)
            getData()
        }
        onMounted(() => {
            getData()
            getTypeList()
            getGroupList()
        })
        const getTypeList = async () => {
            state.typeList = [
                {
                    id: 0,
                    name: "全部"
                },
            ]
            const resp = await typeQue.get(proxy, { limit: 999 })
            if (resp) {
                state.typeList = state.typeList.concat(resp.data)
            }
        }
        const getGroupList = async () => {
            state.groupList = [
                {
                    id: 0,
                    name: "全部"
                },
                {
                    id: -1,
                    name: "未分组"
                },
            ]
            const resp = await groupQue.get(proxy, { limit: 999 })
            if (resp) {
                state.groupList = state.groupList.concat(resp.data)
            }
        }
        const filterOption = (input, option) => {
            return option.name.toLowerCase().indexOf(input.toLowerCase()) >= 0;
        }
        const searchName = (id, listInfo) => {
            for (const iterator of listInfo) {
                if (id == iterator.id) {
                    return iterator.name
                }
            }
            return ""
        }
        return {
            ...toRefs(state), columns, closeChild, getData, paginationChange, showSizeChange, confirm, filterOption, searchName
        }
    }
}
</script>

<style lang="scss" scoped></style>