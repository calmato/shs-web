<template>
  <v-container>
    <div class="text-h6 my-2">{{ summary.year }}年 {{ summary.month }}月</div>
    <v-divider />
    <v-container class="d-flex">
      <v-row>
        <v-col cols="4">授業日</v-col>
        <v-col cols="8">授業希望</v-col>
      </v-row>
    </v-container>
    <v-divider />
    <the-submission-list
      :loading="loading"
      :shifts="shifts"
      :selected-items="enabledLessonIds"
      @click:change-items="onChangeItems"
      @click:submit="onClickSave"
    />
  </v-container>
</template>

<script lang="ts">
import { defineComponent, PropType, SetupContext } from '@nuxtjs/composition-api'
import TheSubmissionList from '~/components/organisms/TheSubmissionList.vue'
import { SubmissionDetail, SubmissionLesson, SubmissionSummary } from '~/types/store'

export default defineComponent({
  components: {
    TheSubmissionList,
  },

  props: {
    loading: {
      type: Boolean,
      default: false,
    },
    summary: {
      type: Object as PropType<SubmissionSummary>,
      default: () => {},
    },
    shifts: {
      type: Array as PropType<SubmissionDetail[]>,
      default: () => [],
    },
    lessons: {
      type: Array as PropType<SubmissionLesson[]>,
      default: () => [],
    },
    enabledLessonIds: {
      type: Array as PropType<number[]>,
      default: () => [],
    },
  },

  setup(_, { emit }: SetupContext) {
    const onChangeItems = (lessonId: number): void => {
      emit('click:change-items', lessonId)
    }

    const onClickSave = (): void => {
      emit('click:submit')
    }

    return {
      onChangeItems,
      onClickSave,
    }
  },
})
</script>
