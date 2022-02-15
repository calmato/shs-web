<template>
  <v-container>
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
    <div class="d-flex">
      <v-spacer />
      <v-btn color="primary" @click="handleSubmitButton">保存</v-btn>
    </div>
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
      type: Array as PropType<number[]>,
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
      set: (val: number[]) => emit('change:regular-holiday-value', val),
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

    const handleSubmitButton = () => {
      emit('click:submit')
    }

    return {
      regularHolidayFormData,
      boothFormData,
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
