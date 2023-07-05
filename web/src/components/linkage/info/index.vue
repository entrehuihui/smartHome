<template>
    <layoutVue>
    </layoutVue>
</template>

<script>
import { reactive, toRefs } from 'vue'
import layoutVue from "@/views/menu/layout.vue";

export default {
    name: 'typesInfo',
    components: { layoutVue },
    setup() {
        const state = reactive({
            count: 0,
        })

        return {
            ...toRefs(state),
        }
    }
}
</script>

<style lang="scss" scoped></style>