<template>
  <v-card>
    <v-toolbar color="primary" dark elevation="0">新規シフト募集の作成</v-toolbar>

    <v-card-text>
      <the-form-group>
        <v-row class="py-8">
          <v-col cols="12">
            <h4>シフトを募集する月</h4>
            <the-date-picker
              :label="form.options.yearMonth.label"
              :rules="form.options.yearMonth.rules"
              :value.sync="form.params.yearMonth"
              type="month"
            />
          </v-col>
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
          <v-col cols="12">
            <h4>
              <span class="pr-4">休日の設定</span>
              <v-btn color="primary" fab x-small elevation="1" @click="addClosedDate"><v-icon>mdi-plus</v-icon></v-btn>
            </h4>
            <the-date-picker
              v-for="(_, index) in form.params.closedDates"
              :key="index"
              :label="form.options.closedDates.label"
              :rules="form.options.closedDates.rules"
              :value.sync="form.params.closedDates[index]"
              class="d-flex flex-column"
            >
              <v-btn color="error" fab x-small elevation="1" @click="removeClosedDate(index)"
                ><v-icon>mdi-minus</v-icon></v-btn
              >
            </the-date-picker>
          </v-col>
        </v-row>
      </the-form-group>
    </v-card-text>

    <v-card-actions>
      <v-spacer />
      <v-btn color="primary" outlined @click="onClose">閉じる</v-btn>
      <v-btn color="primary" :loading="loading" :disabled="loading" @click="onSubmit">作成する</v-btn>
    </v-card-actions>
  </v-card>
</template>

<script lang="ts">
import { defineComponent, PropType, SetupContext } from '@nuxtjs/composition-api'
import TheFormGroup from '~/components/atoms/TheFormGroup.vue'
import TheDatePicker from '~/components/molecules/TheDatePicker.vue'
import { ShiftsNewForm, ShiftsNewOptions, ShiftsNewParams } from '~/types/form'

export default defineComponent({
  components: {
    TheDatePicker,
    TheFormGroup,
  },

  props: {
    form: {
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
  },

  setup(_, { emit }: SetupContext) {
    const addClosedDate = (): void => {
      emit('click:add')
    }

    const removeClosedDate = (index: number): void => {
      emit('click:remove', index)
    }

    const onSubmit = (): void => {
      emit('click:submit')
    }

    const onClose = (): void => {
      emit('click:close')
    }

    return {
      onSubmit,
      onClose,
      addClosedDate,
      removeClosedDate,
    }
  },
})
</script>

<style lang="scss" scoped>
.h-stack {
  gap: 1rem;
}
</style>
