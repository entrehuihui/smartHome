<template>
    <layoutVue>
        <Space direction="vertical" style="width: 100%;">
            <div>
                <Button style="float: right;" @click="addShow = true">
                    <PlusOutlined></PlusOutlined>
                    添加
                </Button>
            </div>
            <div>
                <Input style="width:250px" placeholder="模糊检索" v-model:value="fuzzy" allow-clear :allowClear="true"
                    :maxlength="30"></Input>
                <Button style="margin-left: 10px; width:70px" type="primary" @click="showSizeChange">检索</Button>
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
                        <template v-else-if="column.dataIndex === 'action'">
                            <Space>
                                <Button type="dashed" @click="updateInfo = record; updateShow = true">
                                    修改
                                </Button>
                                <Popconfirm title="是否确定删除该设备组?" ok-text="确定" cancel-text="取消" @confirm=" confirm(record) ">
                                    <Button type="dashed" danger>
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
        <Add v-if=" addShow " @close=" closeChild " />
        <Update v-if=" updateShow " @close=" closeChild " :updateInfo=" updateInfo " />
    </layoutVue>
</template>

<script>
import { reactive, toRefs, onMounted, getCurrentInstance } from 'vue'
import layoutVue from "@/views/menu/layout.vue";
import { Table, Button, Space, Pagination, Popconfirm, Input } from "ant-design-vue";
import {
    PlusOutlined
} from "@ant-design/icons-vue";
import Add from './add.vue';
import Que from "./que.js";
import dayjs from "dayjs";
import Update from './update.vue';


export default {
    name: "groupIndex",
    components: { layoutVue, Table, Button, Space, PlusOutlined, Add, Pagination, Popconfirm, Update, Input },
    setup() {
        const { proxy } = getCurrentInstance();
        const state = reactive({
            dayjs,
            data: [],
            addShow: false,
            limit: 10,
            offset: 1,
            fuzzy: "",
            loading: false,
            total: 0,
            updateShow: false,
            updateInfo: {}
        })

        const columns = reactive([
            {
                title: "序号",
                dataIndex: 'index',
                width: 70,
            },
            {
                title: "名称",
                dataIndex: 'name',
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
            state.addShow = false;
            state.updateShow = false
            getData();
        }
        const getData = async () => {
            state.loading = true
            var resp = await Que.get(proxy, {
                limit: state.limit,
                offset: (state.offset - 1) * state.limit,
                fuzzy: state.fuzzy,
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
        })
        return {
            ...toRefs(state), columns, closeChild, getData, paginationChange, showSizeChange, confirm
        }
    }
}
</script>

<style lang="scss" scoped></style>