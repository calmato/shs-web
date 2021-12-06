<template>
  <div>
    <!-- PCレイアウト -->
    <pc-shift-top class="hidden-sm-and-down" />
    <!-- スマホレイアウト -->
    <mb-shift-top class="hidden-md-and-up" @click="handleClickTop" />
  </div>
</template>

<script lang="ts">
import { defineComponent, SetupContext } from '@nuxtjs/composition-api'
import MbShiftTop from '~/components/templates/MbShiftTop.vue'
import PcShiftTop from '~/components/templates/PcShiftTop.vue'

export default defineComponent({
  components: {
    MbShiftTop,
    PcShiftTop,
  },

  setup(_, { root }: SetupContext) {
    const router = root.$router

    const handleClickTop = (): void => {
      router.push('/')
    }

    return {
      handleClickTop,
    }
  },
})
</script>
