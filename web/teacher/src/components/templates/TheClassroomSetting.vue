<template>
  <v-container>
    <v-btn text class="px-0 mb-4" @click="handleBackButton">
      <v-icon>mdi-chevron-left</v-icon>
      戻る
    </v-btn>
    <the-booth-setting-form :value.sync="boothFormData" />
    <the-regular-holiday-setting-form :value.sync="regularHolidayFormData" />
    <the-hour-setting-form
      :weekday-hour-form="weekdayHourForm"
      :holiday-hour-form="holidayHourForm"
      @click:addWeekdayHourForm="handleWeekdayHourFormAddButton"
      @click:removeWeekdayHourForm="handleWeekdayHourFormRemoveButton"
      @click:addHolidayHourForm="handleHolidayHourFormAddButton"
      @click:removeHolidayHourForm="handleHolidayHourFormRemoveButton"
    />
  </v-container>
</template>

<script lang="ts">
import { computed, defineComponent, PropType, SetupContext } from '@nuxtjs/composition-api'
import TheBoothSettingForm from '../organisms/TheBoothSettingForm.vue'
import TheHourSettingForm from '../organisms/TheHourSettingForm.vue'
import TheRegularHolidaySettingForm from '../organisms/TheRegularHolidaySettingForm.vue'
import { HourForm } from '~/types/form'

export default defineComponent({
  components: { TheRegularHolidaySettingForm, TheBoothSettingForm, TheHourSettingForm },

  props: {
    regularHolidayValue: {
      type: Array as PropType<string[]>,
      default: () => [],
    },
    boothValue: {
      type: Number,
      default: 0,
    },
    weekdayHourForm: {
      type: Array as PropType<HourForm[]>,
      default: () => [],
    },
    holidayHourForm: {
      type: Array as PropType<HourForm[]>,
      default: () => [],
    },
  },

  setup(props, { emit }: SetupContext) {
    const regularHolidayFormData = computed({
      get: () => props.regularHolidayValue,
      set: (val: string[]) => emit('update:regularHolidayValue', val),
    })

    const boothFormData = computed({
      get: () => props.boothValue,
      set: (val: number) => {
        emit('update:boothValue', val)
      },
    })

    const handleBackButton = (): void => {
      emit('onClickBackButton')
    }

    const handleWeekdayHourFormAddButton = () => {
      emit('click:addWeekdayHourForm')
    }

    const handleWeekdayHourFormRemoveButton = (id: number) => {
      emit('click:removeWeekdayHourForm', id)
    }

    const handleHolidayHourFormAddButton = () => {
      emit('click:addHolidayHourForm')
    }

    const handleHolidayHourFormRemoveButton = (id: number) => {
      emit('click:removeHolidayHourForm', id)
    }

    return {
      regularHolidayFormData,
      boothFormData,
      handleBackButton,
      handleWeekdayHourFormAddButton,
      handleWeekdayHourFormRemoveButton,
      handleHolidayHourFormAddButton,
      handleHolidayHourFormRemoveButton,
    }
  },
})
</script>
