<template>
    <Modal>

    </Modal>
</template>

<script>
import { reactive, toRefs } from 'vue'
import { Modal } from "ant-design-vue";

export default {
    name: "groupInfo",
    components: { Modal },
    setup() {
        const state = reactive({
        })

        return {
            ...toRefs(state),
        }
    }
}
</script>

<style lang="scss" scoped></style>