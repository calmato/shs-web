<template>
  <v-container>
    <v-row class="py-4">
      <v-dialog :value.sync="newDialog" width="600px" scrollable @click:outside="toggleNewDialog">
        <the-shift-new-card
          :form="newForm"
          :loading="loading"
          @click:submit="onSubmitNew"
          @click:close="toggleNewDialog"
          @click:add="onClickAddClosedDate"
          @click:remove="onClickRemoveClosedDate"
        />
      </v-dialog>
      <v-col cols="12" align="center">
        <v-btn color="primary" @click="onClickNewShift">
          <v-icon>mdi-plus</v-icon>
          <span>シフトを募集する</span>
        </v-btn>
      </v-col>
      <v-col cols="12">
        <h3>募集中のシフト</h3>
        <v-row align="center" class="pa-4">
          <v-col v-for="summary in acceptingSummaries" :key="summary.id" cols="4">
            <the-shift-top-card :summary="summary" @click:new="onClickNewLesson" @click:edit="onClickEditShift" />
          </v-col>
        </v-row>
      </v-col>
      <v-col cols="12">
        <h3>募集終了後のシフト</h3>
        <v-row align="center" class="pa-4">
          <v-col v-for="summary in finishedSummaries" :key="summary.id" cols="4">
            <the-shift-top-card :summary="summary" @click:new="onClickNewLesson" @click:edit="onClickEditShift" />
          </v-col>
        </v-row>
      </v-col>
      <v-col cols="12">
        <h3>募集開始前のシフト</h3>
        <v-row align="center" class="pa-4">
          <v-col v-for="summary in waitingSummaries" :key="summary.id" cols="4">
            <the-shift-top-card :summary="summary" @click:new="onClickNewLesson" @click:edit="onClickEditShift" />
          </v-col>
        </v-row>
      </v-col>
    </v-row>
  </v-container>
</template>

<script lang="ts">
import { defineComponent, PropType, SetupContext } from '@nuxtjs/composition-api'
import TheShiftNewCard from '~/components/organisms/TheShiftNewCard.vue'
import TheShiftTopCard from '~/components/organisms/TheShiftTopCard.vue'
import { ShiftsNewForm, ShiftsNewOptions, ShiftsNewParams } from '~/types/form'
import { ShiftSummary } from '~/types/store'

export default defineComponent({
  components: {
    TheShiftNewCard,
    TheShiftTopCard,
  },

  props: {
    newDialog: {
      type: Boolean,
      default: false,
    },
    newForm: {
      type: Object as PropType<ShiftsNewForm>,
      default: () => ({
        params: ShiftsNewParams,
        options: ShiftsNewOptions,
      }),
    },
    loading: {
      type: Boolean,
      default: false,
    },
    acceptingSummaries: {
      type: Array as PropType<ShiftSummary[]>,
      default: () => [],
    },
    finishedSummaries: {
      type: Array as PropType<ShiftSummary[]>,
      default: () => [],
    },
    waitingSummaries: {
      type: Array as PropType<ShiftSummary[]>,
      default: () => [],
    },
  },

  setup(_, { emit }: SetupContext) {
    const toggleNewDialog = (): void => {
      emit('toggle:new-dialog')
    }

    const onClickNewShift = (): void => {
      emit('click:new-shift')
    }

    const onClickEditShift = (shift: ShiftSummary): void => {
      emit('click:edit-shift', shift)
    }

    const onClickNewLesson = (shift: ShiftSummary): void => {
      emit('click:new-lesson', shift)
    }

    const onClickAddClosedDate = (): void => {
      emit('click:add-closed-date')
    }

    const onClickRemoveClosedDate = (index: number): void => {
      emit('click:remove-closed-date', index)
    }

    const onSubmitNew = (): void => {
      emit('submit:new')
    }

    return {
      toggleNewDialog,
      onClickNewShift,
      onClickEditShift,
      onClickNewLesson,
      onClickAddClosedDate,
      onClickRemoveClosedDate,
      onSubmitNew,
    }
  },
})
</script>
