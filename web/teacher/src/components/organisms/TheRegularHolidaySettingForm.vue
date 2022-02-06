<template>
  <div>
    <p class="text-h5 mb-2">定休日の設定</p>
    <p class="text-body-2">定休日とする曜日にチェックをつけてください。</p>
    <v-checkbox
      v-for="(dateString, i) in dateList"
      :key="i"
      v-model="selectedValues"
      class="ma-0 pt-0"
      :label="dateString.label"
      :value="dateString.value"
    />
  </div>
</template>

<script lang="ts">
import { defineComponent, PropType, SetupContext, computed } from '@nuxtjs/composition-api'

export default defineComponent({
  props: {
    value: {
      type: Array as PropType<number[]>,
      default: () => [],
    },
  },

  setup(props, { emit }: SetupContext) {
    const dateList: { label: string; value: number }[] = [
      { label: '月曜日', value: 1 },
      { label: '火曜日', value: 2 },
      { label: '水曜日', value: 3 },
      { label: '木曜日', value: 4 },
      { label: '金曜日', value: 5 },
      { label: '土曜日', value: 6 },
      { label: '日曜日', value: 7 },
    ]

    const selectedValues = computed({
      get: () => props.value,
      set: (val: number[]) => emit('update:value', val),
    })

    return {
      dateList,
      selectedValues,
    }
  },
})
</script>
