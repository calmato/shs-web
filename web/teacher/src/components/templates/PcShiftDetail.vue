<template>
  <v-container class="shift">
    <section class="shift-header">
      <!-- 講師情報一覧 -->
      <the-shift-teacher-table
        :teachers="teachers"
        class="py-2"
        @click:show-submissions="onClickTeacherSubmissions"
        @click:show-lessons="onClickTeacherLessons"
      />
      <!-- 生徒情報一覧 -->
      <the-shift-student-table
        :students="students"
        class="py-2"
        @click:show-submissions="onClickStudentSubmissions"
        @click:show-lessons="onClickStudentLessons"
      />
      <!-- シフトタイトル -->
      <div class="d-flex align-center py-2">
        <h3>{{ getTitle() }}</h3>
        <v-btn color="primary" class="ml-auto" @click="onClickDecidedLesson">授業を確定する</v-btn>
      </div>
    </section>
    <section class="shift-content">
      <!-- 授業タイムテーブル -->
      <the-shift-lesson-list
        :rooms="rooms"
        :shifts="details"
        :lessons="lessons"
        @click:new="onClickNewLesson"
        @click:edit="onClickEditLesson"
      />
    </section>
  </v-container>
</template>

<script lang="ts">
import { defineComponent, PropType, SetupContext } from '@nuxtjs/composition-api'
import { ShiftDetail, ShiftSummary, StudentShift, TeacherShift } from '~/types/store'
import { LessonDetail } from '~/types/props/shift'
import TheShiftLessonList from '~/components/organisms/TheShiftLessonList.vue'
import TheShiftStudentTable from '~/components/organisms/TheShiftStudentTable.vue'
import TheShiftTeacherTable from '~/components/organisms/TheShiftTeacherTable.vue'

export default defineComponent({
  components: {
    TheShiftLessonList,
    TheShiftStudentTable,
    TheShiftTeacherTable,
  },

  props: {
    summary: {
      type: Object as PropType<ShiftSummary>,
      default: () => {},
    },
    details: {
      type: Array as PropType<ShiftDetail[]>,
      default: () => [],
    },
    teachers: {
      type: Array as PropType<TeacherShift[]>,
      default: () => [],
    },
    students: {
      type: Array as PropType<StudentShift[]>,
      default: () => [],
    },
    lessons: {
      type: Array as PropType<LessonDetail[]>,
      default: () => [],
    },
    rooms: {
      type: Number,
      default: 0,
    },
  },

  setup(props, { emit }: SetupContext) {
    const getTitle = (): string => {
      return `授業登録 ${props.summary?.year}年${props.summary?.month}月`
    }

    const onClickTeacherSubmissions = (teacherId: string): void => {
      emit('click:show-teacher-submissions', teacherId)
    }

    const onClickTeacherLessons = (teacherId: string): void => {
      emit('click:show-teacher-lessons', teacherId)
    }

    const onClickStudentSubmissions = (studentId: string): void => {
      emit('click:show-student-submissions', studentId)
    }

    const onClickStudentLessons = (studentId: string): void => {
      emit('click:show-student-lessons', studentId)
    }

    const onClickDecidedLesson = (): void => {
      emit('click:decided-lesson')
    }

    const onClickNewLesson = ({ summaryId }: { summaryId: number }): void => {
      emit('click:new-lesson', { summaryId })
    }

    const onClickEditLesson = ({ summaryId, lessonId }: { summaryId: number; lessonId: number }): void => {
      emit('click:edit-lesson', { summaryId, lessonId })
    }

    return {
      getTitle,
      onClickTeacherSubmissions,
      onClickTeacherLessons,
      onClickStudentSubmissions,
      onClickStudentLessons,
      onClickDecidedLesson,
      onClickNewLesson,
      onClickEditLesson,
    }
  },
})
</script>

<style scoped>
.shift {
  display: flex;
  flex-direction: column;
  height: calc(100vh - 64px); /* header: 64px */
}

.shift-content {
  flex: 1;
  overflow: auto;
}
</style>
