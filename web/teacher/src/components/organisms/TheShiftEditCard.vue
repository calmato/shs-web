<template>
  <v-card>
    <v-toolbar color="primary" dark elevation="0">
      <span>シフト募集期間の編集</span>
      <v-spacer />
      <v-icon @click="onDelete">mdi-delete</v-icon>
    </v-toolbar>

    <v-card-text>
      <the-form-group>
        <v-row v-if="deleteDialog" class="py-8">
          <h2>本当に削除しますか</h2>
        </v-row>
        <v-row v-else class="py-8">
          <v-col cols="12">
            <h4>シフト提出可能期間</h4>
            <div class="d-flex align-center h-stack">
              <the-date-picker
                :label="form.options.openDate.label"
                :rules="form.options.openDate.rules"
                :value.sync="form.params.openDate"
              />
              <span> ~ </span>
              <the-date-picker
                :label="form.options.endDate.label"
                :rules="form.options.endDate.rules"
                :value.sync="form.params.endDate"
              />
            </div>
          </v-col>
        </v-row>
      </the-form-group>
    </v-card-text>

    <v-card-actions v-if="deleteDialog">
      <v-spacer />
      <v-btn @click="onDeleteCancel">キャンセル</v-btn>
      <v-btn color="error" :disabled="loading" @click="onDeleteAccept">削除する</v-btn>
    </v-card-actions>
    <v-card-actions v-else>
      <v-spacer />
      <v-btn color="primary" outlined @click="onClose">閉じる</v-btn>
      <v-btn color="primary" :disabled="loading" @click="onSubmit">保存する</v-btn>
    </v-card-actions>
  </v-card>
</template>

<script lang="ts">
import { defineComponent, PropType, SetupContext } from '@nuxtjs/composition-api'
import TheFormGroup from '~/components/atoms/TheFormGroup.vue'
import TheDatePicker from '~/components/molecules/TheDatePicker.vue'
import {
  ShiftSummaryEditScheduleForm,
  ShiftSummaryEditScheduleOptions,
  ShiftSummaryEditScheduleParams,
} from '~/types/form'

export default defineComponent({
  components: {
    TheDatePicker,
    TheFormGroup,
  },

  props: {
    form: {
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
    deleteDialog: {
      type: Boolean,
      default: false,
    },
  },

  setup(_, { emit }: SetupContext) {
    const onDelete = (): void => {
      emit('click:delete')
    }

    const onDeleteAccept = (): void => {
      emit('click:delete-accept')
    }

    const onDeleteCancel = (): void => {
      emit('click:delete-cancel')
    }

    const onSubmit = (): void => {
      emit('click:submit')
    }

    const onClose = (): void => {
      emit('click:close')
    }

    return {
      onDelete,
      onDeleteAccept,
      onDeleteCancel,
      onSubmit,
      onClose,
    }
  },
})
</script>

<style lang="scss" scoped>
.h-stack {
  gap: 1rem;
}
</style>
