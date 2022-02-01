<template>
  <div>
    <p class="text-h5 mb-2">ブース数の設定</p>
    <p class="text-body-2">使用するブース数を選択してください。</p>
    <v-select v-model="formData" type="number" :items="boothList" />
  </div>
</template>

<script lang="ts">
import { computed, defineComponent, PropType, SetupContext } from '@nuxtjs/composition-api'

export default defineComponent({
  props: {
    value: {
      type: Number,
      default: 0,
    },
    boothList: {
      type: Array as PropType<number[]>,
      default: () => [1, 2, 3, 4, 5],
    },
  },

  setup(props, { emit }: SetupContext) {
    const formData = computed({
      get: () => props.value,
      set: (val: number) => {
        emit('update:value', val)
      },
    })

    return {
      formData,
    }
  },
})
</script>
