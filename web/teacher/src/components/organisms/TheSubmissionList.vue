<template>
  <v-row>
    <v-col v-for="item in items.shifts" :key="item.id" cols="12">
      <v-container class="d-flex">
        {{ getDay(item.date) }}
      </v-container>
    </v-col>
  </v-row>
</template>

<script lang="ts">
import { defineComponent } from '@vue/composition-api'
import { SubmissionEditState } from '~/types/props/submission'

export default defineComponent({
  props: {
    items: {
      type: Object as PropType<SubmissionEditState>,
      default: () => ({}),
    },
  },

  setup() {
    const getDay = (date: string): string => {
      const year: number = parseInt(date.slice(0, 4))
      const month: number = parseInt(date.slice(5, 6)) - 1
      const day: number = parseInt(date.slice(7, 8))
      const dayOfWeekStr = ['日', '月', '火', '水', '木', '金', '土']
      const dayName = new Date(year, month, day)
      return dayOfWeekStr[dayName.getDay()]
    }

    return {
      getDay,
    }
  },
})
</script>
