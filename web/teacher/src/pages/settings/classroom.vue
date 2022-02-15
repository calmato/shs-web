<template>
  <the-classroom-setting
    class="mt-4"
    :regular-holiday-value="selectedRegularHoliday"
    :booth-value.sync="boothRef"
    :weekday-hour-form="weekdayHourForm"
    :holiday-hour-form="holidayHourForm"
    @change:regular-holiday-value="changeSelectedRegularHoliday"
    @click:addWeekdayHourForm="handleWeekdayHourFormAddButton"
    @click:removeWeekdayHourForm="handleWeekdayHourFormRemoveButton"
    @click:addHolidayHourForm="handleHolidayHourFormAddButton"
    @click:removeHolidayHourForm="handleHolidayHourFormRemoveButton"
    @click:submit="handleSubmitButton"
    @onClickBackButton="handleBackButton"
  />
</template>

<script lang="ts">
import {
  computed,
  ComputedRef,
  defineComponent,
  reactive,
  ref,
  Ref,
  useAsync,
  useRouter,
} from '@nuxtjs/composition-api'
import customParseFormat from 'dayjs/plugin/customParseFormat'
import dayjs from '~/plugins/dayjs'
import TheClassroomSetting from '~/components/templates/TheClassroomSetting.vue'
import { HourForm } from '~/types/form'
import { ClassroomStore, CommonStore } from '~/store'

dayjs.extend(customParseFormat)

function showInvalidRemoveOperationMessage() {
  CommonStore.showSnackbar({ color: 'error', message: 'これ以上コマを減らすことはできません。' })
}
const REST_TIME = 0 // TODO 後で変えられるようにする
const BASE_HOUR_TIME = 90 // TODO 後で変えられるようにする

export default defineComponent({
  components: {
    TheClassroomSetting,
  },

  beforeRouteLeave(_to, _from, next) {
    if (!this.$data.isChanged) {
      next()
      return
    }
    const answer = window.confirm('保存されていない変更がありますが、よろしいですか？')
    answer ? next() : next(false)
  },

  setup() {
    const router = useRouter()

    const boothRef: Ref<number> = ref<number>(0)
    const selectedRegularHoliday = reactive<number[]>([])
    const weekdayHourForm = reactive<HourForm[]>([])
    const holidayHourForm = reactive<HourForm[]>([])

    const isChanged: ComputedRef<boolean> = computed(() => {
      return boothRef.value !== ClassroomStore.getTotalRooms
    })

    useAsync(async () => {
      await initialize()
    })

    async function initialize(): Promise<void> {
      await Promise.all([ClassroomStore.getTotalRoomsByApi(), ClassroomStore.getSchedulesByApi()]).then(() => {
        boothRef.value = ClassroomStore.getTotalRooms
        ClassroomStore.regularHoliday.forEach((val) => selectedRegularHoliday.push(val))
        ClassroomStore.weekdayHourFormValue.forEach((val) => weekdayHourForm.push(val))
        ClassroomStore.holidayHourFormValue.forEach((val) => holidayHourForm.push(val))
      })
    }

    const changeSelectedRegularHoliday = (items: number[]): void => {
      selectedRegularHoliday.splice(0, selectedRegularHoliday.length, ...items)
    }

    const handleWeekdayHourFormAddButton = () => {
      // 一つ前のコマを基準に開始時刻と終了時刻を設定しておく
      const lastElement = weekdayHourForm.slice(-1)[0]
      const baseTime = dayjs(lastElement?.endAt || '00:00', 'HH:mm')
      const startAt = baseTime.add(REST_TIME, 'minutes')
      const endAt = baseTime.add(BASE_HOUR_TIME, 'minutes')

      const newFormItem: HourForm = {
        startAt: startAt.tz().format('HH:mm'),
        endAt: endAt.tz().format('HH:mm'),
      }
      weekdayHourForm.push(newFormItem)
    }

    const handleBackButton = (): void => {
      router.back()
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
      const lastElement = holidayHourForm.slice(-1)[0]
      const baseTime = dayjs(lastElement?.endAt || '00:00', 'HH:mm')
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

    const handleSubmitButton = async () => {
      try {
        await ClassroomStore.updateTotalRooms({ total: boothRef.value })
        await ClassroomStore.updateSchedules({
          regularHoliday: selectedRegularHoliday,
          weekdayHourForm,
          holidayHourForm,
        })
        CommonStore.showSuccessInSnackbar('更新しました。')
      } catch (err) {
        CommonStore.showErrorInSnackbar(err)
      }
    }

    return {
      selectedRegularHoliday,
      boothRef,
      weekdayHourForm,
      holidayHourForm,
      isChanged,
      changeSelectedRegularHoliday,
      handleBackButton,
      handleWeekdayHourFormAddButton,
      handleWeekdayHourFormRemoveButton,
      handleHolidayHourFormAddButton,
      handleHolidayHourFormRemoveButton,
      handleSubmitButton,
    }
  },
})
</script>
