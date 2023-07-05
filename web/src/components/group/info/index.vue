<template>
    <layoutVue>
        <TypographyTitle :level="3" style="background: white; padding: 10px; width:100%">{{ info.name }}
        </TypographyTitle>
        <div>
            <Button style="float: right;">
                <PlusOutlined />
                添加
            </Button>
        </div>
        <Add />
    </layoutVue>
</template>

<script>
import { onMounted, reactive, toRefs } from 'vue'
import layoutVue from "@/views/menu/layout.vue";
import router from "@/router";
import { Typography, Button } from "ant-design-vue";
import {
    PlusOutlined
} from "@ant-design/icons-vue";
import Add from './add.vue';


export default {
    name: 'typesInfo',
    components: { layoutVue, TypographyTitle: Typography.Title, Button, PlusOutlined, Add },
    setup() {
        const state = reactive({
            info: 0,
        })
        onMounted(() => {
            state.info = router.currentRoute.value.query;
        })
        return {
            ...toRefs(state),
        }
    }
}
</script>

<style lang="scss" scoped></style>