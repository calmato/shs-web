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
      <v-dialog :value.sync="editDialog" width="600px" scrollable @click:outside="toggleEditDialog">
        <the-shift-edit-card
          :form="editForm"
          :loading="loading"
          :delete-dialog="deleteDialog"
          @click:submit="onSubmitEdit"
          @click:delete="onClickDeleteShift"
          @click:delete-accept="onClickDeleteAccept"
          @click:delete-cancel="onClickDeleteCancel"
          @click:close="toggleEditDialog"
        />
      </v-dialog>
      <v-col cols="12" align="center">
        <v-btn color="primary" outlined block @click="onClickNewShift">
          <v-icon>mdi-plus</v-icon>
          <span>シフトを募集する</span>
        </v-btn>
      </v-col>
      <v-col cols="12">
        <h3>募集中のシフト</h3>
        <v-row align="center" class="py-4">
          <v-col v-for="summary in acceptingSummaries" :key="summary.id" cols="4">
            <the-shift-top-card :summary="summary" @click:new="onClickNewLesson" @click:edit="onClickEditShift" />
          </v-col>
        </v-row>
      </v-col>
      <v-col cols="12">
        <h3>募集終了後のシフト</h3>
        <v-row align="center" class="py-4">
          <v-col v-for="summary in finishedSummaries" :key="summary.id" cols="4">
            <the-shift-top-card :summary="summary" @click:new="onClickNewLesson" @click:edit="onClickEditShift" />
          </v-col>
        </v-row>
      </v-col>
      <v-col cols="12">
        <h3>募集開始前のシフト</h3>
        <v-row align="center" class="py-4">
          <v-col v-for="summary in waitingSummaries" :key="summary.id" cols="4">
            <the-shift-top-card :summary="summary" @click:new="onClickNewLesson" @click:edit="onClickEditShift" />
          </v-col>
        </v-row>
      </v-col>
    </v-row>
  </v-container>
</template>

<script lang="ts">
import { defineComponent, PropType, ref, SetupContext } from '@nuxtjs/composition-api'
import TheShiftEditCard from '~/components/organisms/TheShiftEditCard.vue'
import TheShiftNewCard from '~/components/organisms/TheShiftNewCard.vue'
import TheShiftTopCard from '~/components/organisms/TheShiftTopCard.vue'
import {
  ShiftsNewForm,
  ShiftsNewOptions,
  ShiftsNewParams,
  ShiftSummaryEditScheduleForm,
  ShiftSummaryEditScheduleOptions,
  ShiftSummaryEditScheduleParams,
} from '~/types/form'
import { ShiftSummary } from '~/types/store'

export default defineComponent({
  components: {
    TheShiftEditCard,
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
    editDialog: {
      type: Boolean,
      default: false,
    },
    editForm: {
      type: Object as PropType<ShiftSummaryEditScheduleForm>,
      default: () => ({
        params: ShiftSummaryEditScheduleParams,
        options: ShiftSummaryEditScheduleOptions,
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
    const deleteDialog = ref<boolean>(false)

    const toggleNewDialog = (): void => {
      emit('toggle:new-dialog')
    }

    const toggleEditDialog = (): void => {
      deleteDialog.value = false
      emit('toggle:edit-dialog')
    }

    const onClickNewShift = (): void => {
      emit('click:new-shift')
    }

    const onClickEditShift = (shift: ShiftSummary): void => {
      emit('click:edit-shift', shift)
    }

    const onClickDeleteShift = (): void => {
      deleteDialog.value = true
    }

    const onClickDeleteAccept = (): void => {
      emit('submit:delete')
      deleteDialog.value = false
    }

    const onClickDeleteCancel = (): void => {
      deleteDialog.value = false
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

    const onSubmitEdit = (): void => {
      deleteDialog.value = false
      emit('submit:edit')
    }

    return {
      deleteDialog,
      toggleNewDialog,
      toggleEditDialog,
      onClickNewShift,
      onClickEditShift,
      onClickDeleteShift,
      onClickDeleteAccept,
      onClickDeleteCancel,
      onClickNewLesson,
      onClickAddClosedDate,
      onClickRemoveClosedDate,
      onSubmitNew,
      onSubmitEdit,
    }
  },
})
</script>
