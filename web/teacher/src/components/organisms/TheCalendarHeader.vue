<template>
  <v-toolbar flat>
    <v-btn outlined class="mr-4" color="grey darken-2" @click="onClickToday">今日</v-btn>
    <v-btn fab text small color="grey darken-2" @click="onClickPrev">
      <v-icon small>mdi-chevron-left</v-icon>
    </v-btn>
    <v-toolbar-title>
      {{ getTitle() }}
    </v-toolbar-title>
    <v-btn fab text small color="grey darken-2" @click="onClickNext">
      <v-icon small>mdi-chevron-right</v-icon>
    </v-btn>
    <v-spacer />
    <v-menu bottom right>
      <template #activator="{ on, attrs }">
        <v-btn outlined color="grey darken-2" v-bind="attrs" v-on="on">
          <span>{{ getTypeName(calendarType) }}</span>
          <v-icon right>mdi-menu-down</v-icon>
        </v-btn>
      </template>
      <v-list v-for="(item, index) in types" :key="index">
        <v-list-item @click="updateType(item)">
          <v-list-item-title>{{ item.name }}</v-list-item-title>
        </v-list-item>
      </v-list>
    </v-menu>
  </v-toolbar>
</template>

<script lang="ts">
import { computed, defineComponent, PropType, SetupContext } from '@nuxtjs/composition-api'
import { Dayjs } from 'dayjs'
import dayjs from '~/plugins/dayjs'
import { Date, CalendarType } from '~/types/props/calendar'

export default defineComponent({
  props: {
    end: {
      type: Object as PropType<Date>,
      default: null,
    },
    now: {
      type: Object as PropType<Dayjs>,
      default: null,
    },
    start: {
      type: Object as PropType<Date>,
      default: null,
    },
    type: {
      type: String,
      default: 'month',
    },
    types: {
      type: Array as PropType<CalendarType[]>,
      default: () => [
        { name: '月', value: 'month' },
        { name: '週', value: 'week' },
        { name: '日', value: 'day' },
      ],
    },
  },

  setup(props, { emit }: SetupContext) {
    const calendarType = computed<string>({
      get: () => props.type,
      set: (val: String) => emit('update:type', val),
    })

    const getTitle = (): string => {
      const format = 'YYYY年MM月'
      if (props.start && props.end) {
        return dayjs(props.start.date).tz().format(format)
      }
      return props.now?.format(format) || ''
    }

    const getTypeName = (val: String): string => {
      const target = props.types.find((type: CalendarType) => {
        return type.value === val
      })
      return target?.name || ''
    }

    const updateType = (type: CalendarType): void => {
      calendarType.value = type.value
    }

    const onClickToday = (): void => {
      emit('click:today')
    }

    const onClickPrev = (): void => {
      emit('click:prev')
    }

    const onClickNext = (): void => {
      emit('click:next')
    }

    return {
      calendarType,
      getTitle,
      getTypeName,
      updateType,
      onClickToday,
      onClickPrev,
      onClickNext,
    }
  },
})
</script>
