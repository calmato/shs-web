<template>
  <v-container>
    <div class="text-h6 my-2">{{ summary.year }}年 {{ summary.month }}月</div>
    <the-submission-lesson-list
      :subjects="subjects"
      :lessons="lessons"
      @click:add-lesson="onClickAddLesson"
      @click:remove-lesson="onClickRemoveLesson"
    />
    <the-submission-shift-list
      :loading="loading"
      :shifts="shifts"
      :selected-items="enabledLessonIds"
      @click:change-items="onChangeItems"
    />
    <v-container>
      <v-row class="justify-end">
        <v-btn color="primary" class="right mt-4" large :disabled="loading" @click="onClickSubmit">提出</v-btn>
      </v-row>
    </v-container>
  </v-container>
</template>

<script lang="ts">
import { defineComponent, PropType, SetupContext } from '@nuxtjs/composition-api'
import TheSubmissionLessonList from '~/components/organisms/TheSubmissionLessonList.vue'
import TheSubmissionShiftList from '~/components/organisms/TheSubmissionShiftList.vue'
import { ISubmissionSuggestedLesson } from '~/types/form'
import { Subject, SubmissionDetail, SubmissionSummary } from '~/types/store'

export default defineComponent({
  components: {
    TheSubmissionLessonList,
    TheSubmissionShiftList,
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
      type: Array as PropType<ISubmissionSuggestedLesson[]>,
      default: () => [],
    },
    subjects: {
      type: Array as PropType<Subject[]>,
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

    const onClickAddLesson = (): void => {
      emit('click:add-lesson')
    }

    const onClickRemoveLesson = (index: number): void => {
      emit('click:remove-lesson', index)
    }

    const onClickSubmit = (): void => {
      emit('click:submit')
    }

    return {
      onChangeItems,
      onClickAddLesson,
      onClickRemoveLesson,
      onClickSubmit,
    }
  },
})
</script>
