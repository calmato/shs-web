<template>
  <the-classroom-setting
    :regular-holiday-value.sync="selectedRegularHoliday"
    :booth-value.sync="boothRef"
    :weekday-hour-form="weekdayHourForm"
    :holiday-hour-form="holidayHourForm"
    @click:addWeekdayHourForm="handleWeekdayHourFormAddButton"
    @click:removeWeekdayHourForm="handleWeekdayHourFormRemoveButton"
    @click:addHolidayHourForm="handleHolidayHourFormAddButton"
    @click:removeHolidayHourForm="handleHolidayHourFormRemoveButton"
  />
</template>

<script lang="ts">
import { defineComponent, reactive, ref, Ref } from '@nuxtjs/composition-api'
import customParseFormat from 'dayjs/plugin/customParseFormat'
import dayjs from '~/plugins/dayjs'
import TheClassroomSetting from '~/components/templates/TheClassroomSetting.vue'
import { HourForm } from '~/types/form'
import { CommonStore } from '~/store'

dayjs.extend(customParseFormat)

function showInvalidRemoveOperationMessage() {
  CommonStore.showSnackbar({ color: 'error', message: 'これ以上コマを減らすことはできません。' })
}
const REST_TIME = 10 // TODO 後で変えられるようにする
const BASE_HOUR_TIME = 45 // TODO 後で変えられるようにする

export default defineComponent({
  components: {
    TheClassroomSetting,
  },

  setup() {
    const selectedRegularHoliday = reactive<string[]>([])
    const boothRef: Ref<number> = ref<number>(1)
    const weekdayHourForm = reactive<HourForm[]>([{ startAt: '17:00', endAt: '17:45' }])
    const holidayHourForm = reactive<HourForm[]>([{ startAt: '17:00', endAt: '17:45' }])

    const handleWeekdayHourFormAddButton = () => {
      // 一つ前のコマを基準に開始時刻と終了時刻を設定しておく
      const lastElement = weekdayHourForm.slice(-1)[0]
      const baseTime = dayjs(lastElement.endAt, 'HH:mm')
      const startAt = baseTime.add(REST_TIME, 'minutes')
      const endAt = baseTime.add(BASE_HOUR_TIME, 'minutes')

      const newFormItem: HourForm = {
        startAt: startAt.tz().format('HH:mm'),
        endAt: endAt.tz().format('HH:mm'),
      }
      weekdayHourForm.push(newFormItem)
    }

    const handleWeekdayHourFormRemoveButton = (i: number) => {
      if (weekdayHourForm.length === 1) {
        showInvalidRemoveOperationMessage()
        return
      }
      if (i < weekdayHourForm.length) {
        weekdayHourForm.splice(i, 1)
      }
    }

    const handleHolidayHourFormAddButton = () => {
      // 一つ前のコマを基準に開始時刻と終了時刻を設定しておく
      const lastElement = weekdayHourForm.slice(-1)[0]
      const baseTime = dayjs(lastElement.endAt, 'HH:mm')
      const startAt = baseTime.add(REST_TIME, 'minutes')
      const endAt = baseTime.add(BASE_HOUR_TIME, 'minutes')

      const newFormItem: HourForm = {
        startAt: startAt.tz().format('HH:mm'),
        endAt: endAt.tz().format('HH:mm'),
      }
      holidayHourForm.push(newFormItem)
    }

    const handleHolidayHourFormRemoveButton = (i: number) => {
      if (holidayHourForm.length === 1) {
        showInvalidRemoveOperationMessage()
        return
      }
      if (i < holidayHourForm.length) {
        holidayHourForm.splice(i, 1)
      }
    }

    return {
      selectedRegularHoliday,
      boothRef,
      weekdayHourForm,
      holidayHourForm,
      handleWeekdayHourFormAddButton,
      handleWeekdayHourFormRemoveButton,
      handleHolidayHourFormAddButton,
      handleHolidayHourFormRemoveButton,
    }
  },
})
</script>
