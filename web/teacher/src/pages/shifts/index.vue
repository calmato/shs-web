<template>
  <div>
    <!-- PCレイアウト -->
    <pc-shift-top
      class="hidden-sm-and-down"
      :new-dialog="newDialog"
      :new-form="newForm"
      :edit-dialog="editDialog"
      :edit-form="editForm"
      :loading="loading"
      :accepting-summaries="getAcceptingSummaries()"
      :finished-summaries="getFinishedSummaries()"
      :waiting-summaries="getWaitingSummaries()"
      @click:new-lesson="handleClickNewLesson"
      @click:new-shift="handleClickNewShift"
      @click:edit-shift="handleClickEditShift"
      @click:add-closed-date="handleClickAddClosedDate"
      @click:remove-closed-date="handleClickRemoveClosedDate"
      @toggle:new-dialog="toggleNewDialog"
      @toggle:edit-dialog="toggleEditDialog"
      @submit:new="handleSubmitCreateShifts"
      @submit:edit="handleSubmitUpdateShiftSummarySchedule"
      @submit:delete="handleSubmitDeleteShifts"
    />
    <!-- スマホレイアウト -->
    <mb-shift-top class="hidden-md-and-up" @click="handleClickTop" />
  </div>
</template>

<script lang="ts">
import { computed, defineComponent, reactive, ref, SetupContext, useAsync } from '@nuxtjs/composition-api'
import MbShiftTop from '~/components/templates/MbShiftTop.vue'
import PcShiftTop from '~/components/templates/PcShiftTop.vue'
import dayjs from '~/plugins/dayjs'
import { CommonStore, ShiftStore } from '~/store'
import {
  ShiftsNewForm,
  ShiftsNewOptions,
  ShiftsNewParams,
  ShiftSummaryEditScheduleForm,
  ShiftSummaryEditScheduleOptions,
  ShiftSummaryEditScheduleParams,
} from '~/types/form'
import { PromiseState, ShiftStatus, ShiftSummary } from '~/types/store'

export default defineComponent({
  components: {
    MbShiftTop,
    PcShiftTop,
  },

  setup(_, { root }: SetupContext) {
    const router = root.$router
    const store = root.$store

    const newDialog = ref<boolean>(false)
    const editDialog = ref<boolean>(false)

    const newForm = reactive<ShiftsNewForm>({
      params: { ...ShiftsNewParams },
      options: { ...ShiftsNewOptions },
    })
    const editForm = reactive<ShiftSummaryEditScheduleForm>({
      params: { ...ShiftSummaryEditScheduleParams },
      options: { ...ShiftSummaryEditScheduleOptions },
    })

    const summaries = computed<ShiftSummary[]>(() => store.getters['shift/getSummaries'])
    const loading = computed<boolean>(() => {
      return store.getters['common/getPromiseState'] === PromiseState.LOADING
    })

    useAsync(async () => {
      await listShiftSummaries()
    })

    async function listShiftSummaries(): Promise<void> {
      CommonStore.startConnection()

      const limit: number = 0
      const offset: number = 0
      const status: ShiftStatus = ShiftStatus.UNKNOWN

      await ShiftStore.listShiftSummaries({ limit, offset, status })
        .catch((err: Error) => {
          console.log('feiled to list shift summaries', err)
        })
        .finally(() => {
          CommonStore.endConnection()
        })
    }

    const getWaitingSummaries = (): ShiftSummary[] => {
      return summaries.value.filter((summary: ShiftSummary) => {
        return summary.status === ShiftStatus.WAITING
      })
    }

    const getAcceptingSummaries = (): ShiftSummary[] => {
      return summaries.value.filter((summary: ShiftSummary) => {
        return summary.status === ShiftStatus.ACCEPTING
      })
    }

    const getFinishedSummaries = (): ShiftSummary[] => {
      return summaries.value.filter((summary: ShiftSummary) => {
        return summary.status === ShiftStatus.FINISHED
      })
    }

    const toggleNewDialog = (): void => {
      newDialog.value = !newDialog.value
    }

    const toggleEditDialog = (shift?: ShiftSummary): void => {
      const format: string = 'YYYY-MM-DD'
      if (shift) {
        editForm.params = {
          summaryId: shift.id,
          openDate: dayjs(shift.openAt).tz().format(format),
          endDate: dayjs(shift.endAt).tz().format(format),
        }
      } else {
        editForm.params = { ...ShiftSummaryEditScheduleParams }
      }
      editDialog.value = !editDialog.value
    }

    const handleClickNewShift = (): void => {
      toggleNewDialog()
    }

    const handleClickEditShift = (shift: ShiftSummary): void => {
      toggleEditDialog(shift)
    }

    const handleClickNewLesson = (shift: ShiftSummary): void => {
      router.push(`/shifts/${shift.id}`)
    }

    const handleClickAddClosedDate = (): void => {
      newForm.params.closedDates.push('')
    }

    const handleClickRemoveClosedDate = (index: number): void => {
      newForm.params.closedDates.splice(index, 1)
    }

    const handleClickTop = (): void => {
      router.push('/')
    }

    const handleSubmitCreateShifts = async (): Promise<void> => {
      CommonStore.startConnection()
      await ShiftStore.createShifts({ form: newForm })
        .then(() => {
          toggleNewDialog()
          CommonStore.showSnackbar({ color: 'success', message: 'シフト募集を新規登録しました。' })
        })
        .catch((err: Error) => {
          CommonStore.showErrorInSnackbar(err)
        })
        .finally(() => {
          CommonStore.endConnection()
        })
    }

    const handleSubmitUpdateShiftSummarySchedule = async (): Promise<void> => {
      CommonStore.startConnection()
      await ShiftStore.updateShiftSummarySchedule({ form: editForm })
        .then(() => {
          toggleEditDialog()
          CommonStore.showSnackbar({ color: 'success', message: 'シフトの募集期間を更新しました。' })
        })
        .catch((err: Error) => {
          CommonStore.showErrorInSnackbar(err)
        })
        .finally(() => {
          CommonStore.endConnection()
        })
    }

    const handleSubmitDeleteShifts = async (): Promise<void> => {
      CommonStore.startConnection()
      await ShiftStore.deleteShifts({ summaryId: editForm.params.summaryId })
        .then(() => {
          toggleEditDialog()
          CommonStore.showSnackbar({ color: 'success', message: 'シフト募集を削除しました。' })
        })
        .catch((err: Error) => {
          CommonStore.showErrorInSnackbar(err)
        })
        .finally(() => {
          CommonStore.endConnection()
        })
    }

    return {
      newDialog,
      newForm,
      editDialog,
      editForm,
      loading,
      getAcceptingSummaries,
      getFinishedSummaries,
      getWaitingSummaries,
      toggleNewDialog,
      toggleEditDialog,
      handleClickNewShift,
      handleClickEditShift,
      handleClickNewLesson,
      handleClickAddClosedDate,
      handleClickRemoveClosedDate,
      handleClickTop,
      handleSubmitCreateShifts,
      handleSubmitUpdateShiftSummarySchedule,
      handleSubmitDeleteShifts,
    }
  },
})
</script>
