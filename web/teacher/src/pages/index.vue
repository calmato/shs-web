<template>
  <top :events="events" @click="handleClick" />
</template>

<script lang="ts">
import { defineComponent, reactive } from '@nuxtjs/composition-api'
import Top from '~/components/templates/Top.vue'
import { Event } from '~/types/props/calendar'

export default defineComponent({
  components: {
    Top,
  },

  setup() {
    const events = reactive<Event[]>([
      {
        name: '月次報告',
        start: '2021-12-15 09:00',
        end: '2021-12-16 17:00',
        color: 'primary',
      },
      {
        name: 'サンタクロース',
        start: '2021-12-24 18:00:00',
        end: '2021-12-25 05:00:00',
        color: 'primaryLight',
      },
      {
        name: '出張',
        start: '2021-12-26',
        end: '2021-12-28',
        color: 'secondary',
      },
    ])

    const handleClick = (event: Event) => {
      console.log('debug', event)
    }

    return {
      events,
      handleClick,
    }
  },
})
</script>
