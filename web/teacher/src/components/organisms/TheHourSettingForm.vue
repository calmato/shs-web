<template>
  <div>
    <p class="text-h5 mb-2">コマ設定</p>
    <div class="d-flex align-center my-2">
      <p class="text-subtitle-1 mr-4 mb-0">平日</p>
      <v-tooltip right>
        <template #activator="{ on, attrs }">
          <v-btn
            class="suffix-btn"
            color="primary"
            fab
            elevation="0"
            v-bind="attrs"
            v-on="on"
            @click="handleWeekdayHourFormAddButton"
            ><v-icon>mdi-plus</v-icon></v-btn
          >
        </template>
        <span>コマを増やす</span>
      </v-tooltip>
    </div>
    <div>
      <the-hour-form-item
        v-for="(item, i) in weekdayHourForm"
        :id="i"
        :key="i"
        :item="item"
        @click="handleWeekdayHourFormRemoveButton"
      />
    </div>
    <div class="d-flex align-center my-2">
      <p class="text-subtitle-1 mr-4 mb-0">休日</p>
      <v-tooltip right>
        <template #activator="{ on, attrs }">
          <v-btn
            class="suffix-btn"
            color="primary"
            fab
            elevation="0"
            v-bind="attrs"
            v-on="on"
            @click="handleHolidayHourFormAddButton"
          >
            <v-icon>mdi-plus</v-icon></v-btn
          >
        </template>
        <span>コマを増やす</span>
      </v-tooltip>
    </div>
    <div>
      <the-hour-form-item
        v-for="(item, i) in holidayHourForm"
        :id="i"
        :key="i"
        :item="item"
        @click="handleHolidayHourFormRemoveButton"
      />
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, PropType, SetupContext } from '@nuxtjs/composition-api'
import { HourForm } from '~/types/form'
import TheHourFormItem from '~/components/molecules/TheHourFormItem.vue'

export default defineComponent({
  components: {
    TheHourFormItem,
  },

  props: {
    weekdayHourForm: {
      type: Array as PropType<HourForm[]>,
      default: () => [],
    },
    holidayHourForm: {
      type: Array as PropType<HourForm[]>,
      default: () => [],
    },
  },

  setup(_, { emit }: SetupContext) {
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
      handleWeekdayHourFormAddButton,
      handleWeekdayHourFormRemoveButton,
      handleHolidayHourFormAddButton,
      handleHolidayHourFormRemoveButton,
    }
  },
})
</script>

<style lang="scss" scoped>
.suffix-btn {
  height: 24px;
  width: 24px;
}
</style>
