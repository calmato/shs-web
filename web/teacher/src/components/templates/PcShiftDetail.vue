<template>
  <v-container class="shift">
    <v-dialog :value.sync="dialog" width="600px" scrollable @click:outside="onCloseDialog">
      <!-- 講師 提出シフト一覧ダイアログ -->
      <v-card v-if="dialogKey == '講師シフト'">
        <v-toolbar color="primary" dark>提出シフト一覧</v-toolbar>
        <v-card-text>{{ teacherSubmission }}</v-card-text>
        <v-card-actions>
          <v-spacer />
          <v-btn color="secondary" @click="onCloseDialog">閉じる</v-btn>
        </v-card-actions>
      </v-card>
      <!-- 講師 授業一覧ダイアログ -->
      <v-card v-if="dialogKey == '講師授業'">
        <v-toolbar color="primary" dark>講師授業一覧</v-toolbar>
        <v-card-text>{{ teacherLessons }}</v-card-text>
        <v-card-actions>
          <v-spacer />
          <v-btn color="secondary" @click="onCloseDialog">閉じる</v-btn>
        </v-card-actions>
      </v-card>
      <!-- 生徒 授業希望一覧ダイアログ -->
      <v-card v-if="dialogKey == '生徒授業希望'">
        <v-toolbar color="primary" dark>授業希望一覧</v-toolbar>
        <v-card-text>{{ studentSubmission }}</v-card-text>
        <v-card-actions>
          <v-spacer />
          <v-btn color="secondary" @click="onCloseDialog">閉じる</v-btn>
        </v-card-actions>
      </v-card>
      <!-- 生徒 授業一覧ダイアログ -->
      <v-card v-if="dialogKey == '生徒授業'">
        <v-toolbar color="primary" dark>生徒授業一覧</v-toolbar>
        <v-card-text>{{ studentLessons }}</v-card-text>
        <v-card-actions>
          <v-spacer />
          <v-btn color="secondary" @click="onCloseDialog">閉じる</v-btn>
        </v-card-actions>
      </v-card>
      <!-- 授業登録/編集ダイアログ -->
      <v-card v-if="dialogKey == '授業登録'">
        <v-toolbar color="primary" dark>授業登録</v-toolbar>
        <v-card-text>{{ lesson }}</v-card-text>
        <v-card-actions>
          <v-spacer />
          <v-btn color="secondary" @click="onCloseDialog">閉じる</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>

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
import {
  ShiftDetail,
  ShiftLessonDetail,
  ShiftSummary,
  ShiftUserLesson,
  StudentShift,
  StudentSubmissionDetail,
  TeacherShift,
  TeacherSubmissionDetail,
} from '~/types/store'
import { LessonDetail, ShiftDialogKey } from '~/types/props/shift'
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
    dialog: {
      type: Boolean,
      default: false,
    },
    dialogKey: {
      type: String as PropType<ShiftDialogKey>,
      default: '未選択',
    },
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
    lesson: {
      type: Object as PropType<ShiftLessonDetail>,
      default: () => {},
    },
    lessons: {
      type: Array as PropType<LessonDetail[]>,
      default: () => [],
    },
    rooms: {
      type: Number,
      default: 0,
    },
    teacherSubmission: {
      type: Object as PropType<TeacherSubmissionDetail>,
      default: () => {},
    },
    teacherLessons: {
      type: Object as PropType<ShiftUserLesson>,
      default: () => {},
    },
    studentSubmission: {
      type: Object as PropType<StudentSubmissionDetail>,
      default: () => {},
    },
    studentLessons: {
      type: Object as PropType<ShiftUserLesson>,
      default: () => {},
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

    const onClickNewLesson = ({ shiftId, room }: { shiftId: number; room: number }): void => {
      emit('click:new-lesson', { shiftId, room })
    }

    const onClickEditLesson = ({
      shiftId,
      lessonId,
      room,
    }: {
      shiftId: number
      lessonId: number
      room: number
    }): void => {
      emit('click:edit-lesson', { shiftId, lessonId, room })
    }

    const onCloseDialog = (): void => {
      emit('click:close')
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
      onCloseDialog,
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
