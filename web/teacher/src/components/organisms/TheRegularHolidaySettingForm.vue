<template>
  <div>
    <p class="text-h5 mb-2">定休日の設定</p>
    <p class="text-body-2">定休日とする曜日にチェックをつけてください。</p>
    <v-checkbox
      v-for="(dateString, i) in dateList"
      :key="i"
      v-model="selectedValues"
      class="ma-0 pt-0"
      :label="dateString"
      :value="dateString"
    />
  </div>
</template>

<script lang="ts">
import { defineComponent, PropType, SetupContext, computed } from '@nuxtjs/composition-api'

export default defineComponent({
  props: {
    value: {
      type: Array as PropType<string[]>,
      default: () => [],
    },
  },

  setup(props, { emit }: SetupContext) {
    const dateList: string[] = ['月曜日', '火曜日', '水曜日', '木曜日', '金曜日', '土曜日', '日曜日']

    const selectedValues = computed({
      get: () => props.value,
      set: (val: string[]) => emit('update:value', val),
    })

    return {
      dateList,
      selectedValues,
    }
  },
})
</script>
