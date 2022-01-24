<template>
  <v-container class="shift">
    <!-- 講師 提出シフト一覧ダイアログ -->
    <v-dialog
      :value.sync="teacherSubmissionsDialog"
      width="600px"
      scrollable
      @click:outside="onCloseTeacherSubmissionsDialog"
    >
      <v-card>
        <v-toolbar color="primary" dark>提出シフト一覧</v-toolbar>
        <v-card-text>{{ teacherSubmission }}</v-card-text>
        <v-card-actions>
          <v-spacer />
          <v-btn color="secondary" @click="onCloseTeacherSubmissionsDialog">閉じる</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
    <!-- 講師 授業一覧ダイアログ -->
    <v-dialog :value.sync="teacherLessonsDialog" width="600px" scrollable @click:outside="onCloseTeacherLessonsDialog">
      <v-card>
        <v-toolbar color="primary" dark>講師授業一覧</v-toolbar>
        <v-card-actions>
          <v-spacer />
          <v-btn color="secondary" @click="onCloseTeacherLessonsDialog">閉じる</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
    <!-- 生徒 授業希望一覧ダイアログ -->
    <v-dialog
      :value.sync="studentSubmissionsDialog"
      width="600px"
      scrollable
      @click:outside="onCloseStudentSubmissionsDialog"
    >
      <v-card>
        <v-toolbar color="primary" dark>授業希望一覧</v-toolbar>
        <v-card-text>{{ studentSubmission }}</v-card-text>
        <v-card-actions>
          <v-spacer />
          <v-btn color="secondary" @click="onCloseStudentSubmissionsDialog">閉じる</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
    <!-- 生徒 授業一覧ダイアログ -->
    <v-dialog :value.sync="studentLessonsDialog" width="600px" scrollable @click:outside="onCloseStudentLessonsDialog">
      <v-card>
        <v-toolbar color="primary" dark>生徒授業一覧</v-toolbar>
        <v-card-actions>
          <v-spacer />
          <v-btn color="secondary" @click="onCloseStudentLessonsDialog">閉じる</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
    <!-- 授業登録ダイアログ -->
    <v-dialog :value.sync="newLessonDialog" width="600px" scrollable @click:outside="onCloseNewLessonDialog">
      <v-card>
        <v-toolbar color="primary" dark>授業登録</v-toolbar>
      </v-card>
    </v-dialog>
    <!-- 授業編集ダイアログ -->
    <v-dialog :value.sync="editLessonDialog" width="600px" scrollable @click:outside="onCloseEditLessonDialog">
      <v-card>
        <v-toolbar color="primary" dark>授業登録</v-toolbar>
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
import { ShiftDetail, ShiftSummary, StudentShift, StudentSubmissionDetail, TeacherShift, TeacherSubmissionDetail } from '~/types/store'
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
    teacherSubmission: {
      type: Object as PropType<TeacherSubmissionDetail>,
      default: () => {},
    },
    studentSubmission: {
      type: Object as PropType<StudentSubmissionDetail>,
      default: () => {},
    },
    teacherSubmissionsDialog: {
      type: Boolean,
      default: false,
    },
    teacherLessonsDialog: {
      type: Boolean,
      default: false,
    },
    studentSubmissionsDialog: {
      type: Boolean,
      default: false,
    },
    studentLessonsDialog: {
      type: Boolean,
      default: false,
    },
    newLessonDialog: {
      type: Boolean,
      default: false,
    },
    editLessonDialog: {
      type: Boolean,
      default: false,
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

    const onClickNewLesson = ({ summaryId, room }: { summaryId: number; room: number }): void => {
      emit('click:new-lesson', { summaryId, room })
    }

    const onClickEditLesson = ({
      summaryId,
      lessonId,
      room,
    }: {
      summaryId: number
      lessonId: number
      room: number
    }): void => {
      emit('click:edit-lesson', { summaryId, lessonId, room })
    }

    const onCloseTeacherSubmissionsDialog = (): void => {
      emit('click:close-teacher-submissions')
    }

    const onCloseTeacherLessonsDialog = (): void => {
      emit('click:close-teacher-lessons')
    }

    const onCloseStudentSubmissionsDialog = (): void => {
      emit('click:close-student-submissions')
    }

    const onCloseStudentLessonsDialog = (): void => {
      emit('click:close-student-lessons')
    }

    const onCloseNewLessonDialog = (): void => {
      emit('click:close-new-lesson')
    }

    const onCloseEditLessonDialog = (): void => {
      emit('click:close-edit-lesson')
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
      onCloseTeacherSubmissionsDialog,
      onCloseTeacherLessonsDialog,
      onCloseStudentSubmissionsDialog,
      onCloseStudentLessonsDialog,
      onCloseNewLessonDialog,
      onCloseEditLessonDialog,
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
