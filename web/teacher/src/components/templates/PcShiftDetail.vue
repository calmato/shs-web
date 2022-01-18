<template>
  <v-container class="shift">
    <section class="shift-header">
      <!-- 講師情報一覧 -->
      <div class="shift-summary">
        <table>
          <tr>
            <th class="fixed">講師名</th>
            <td
              v-for="teacher in teachers"
              :key="teacher.id"
              class="text-decoration-underline"
              @click="onClickTeacherSubmissions(teacher)"
            >
              {{ teacher.name }}
            </td>
            <th class="fixed-right info--text text--lighten-1 text-decoration-underline">合計</th>
          </tr>
          <tr>
            <th class="fixed">担当授業数</th>
            <td
              v-for="teacher in teachers"
              :key="teacher.id"
              class="text-decoration-underline"
              @click="onClickTeacherLessons(teacher)"
            >
              {{ teacher.lessonTotal }}
            </td>
            <td class="fixed-right">{{ getLessonTotal() }}</td>
          </tr>
        </table>
      </div>
      <!-- 生徒情報一覧 -->
      <div class="shift-summary">
        <table>
          <tr>
            <th class="fixed">生徒名</th>
            <td
              v-for="student in students"
              :key="student.id"
              class="text-decoration-underline"
              @click="onClickStudentSubmissions(student)"
            >
              {{ student.name }}
            </td>
            <th class="fixed-right secondary--text text--accent-4">残り</th>
          </tr>
          <tr>
            <th class="fixed">残り授業数</th>
            <td
              v-for="student in students"
              :key="student.id"
              :class="[student.suggestedClassesTotal > 0 ? 'secondary--text text--accent-4' : '']"
              class="text-decoration-underline"
              @click="onClickStudentLessons(student)"
            >
              {{ student.suggestedClassesTotal }}
            </td>
            <td class="fixed-right" :class="[getRemainingLessonTotal() > 0 ? 'secondary--text text--accent-4' : '']">
              {{ getRemainingLessonTotal() }}
            </td>
          </tr>
        </table>
      </div>
      <!-- シフトタイトル -->
      <v-row class="shift-subheader" align="center">
        <v-col cols="auto">
          <h3>{{ getTitle() }}</h3>
        </v-col>
        <v-spacer />
        <v-col cols="auto">
          <v-btn color="primary" @click="onClickDecidedLesson">授業を確定する</v-btn>
        </v-col>
      </v-row>
    </section>
    <section class="shift-content">
      <!-- 授業タイムテーブル -->
      <div class="shift-lessons">
        <v-container v-for="shift in details" :key="shift.date" tag="v-row" class="shift-lessons-item">
          <v-col cols="1" align="center" style="border: 1px solid #e5e5e5" class="pt-4">{{ getDay(shift.date) }}</v-col>
          <v-col cols="11" class="d-flex flex-column" align="center">
            <div
              v-for="num in rooms"
              :key="`${shift.date}-${num}`"
              justify="start"
              class="d-flex align-stretch shift-lessons-detail"
            >
              <v-card tile outlined class="d-flex align-center col col-1 shift-lesson-room">
                <v-card-text class="text-subtitle-2">ブース{{ num }}</v-card-text>
              </v-card>
              <v-card v-for="lesson in shift.lessons" :key="lesson.id" tile outlined class="shift-lessons-schedule">
                <the-shift-lesson-card
                  :summary="lesson"
                  :detail="getLesson(lesson, num)"
                  @click:new="onClickNewLesson"
                  @click:edit="onClickEditLesson"
                />
              </v-card>
            </div>
          </v-col>
        </v-container>
      </div>
    </section>
  </v-container>
</template>

<script lang="ts">
import { defineComponent, PropType, SetupContext } from '@nuxtjs/composition-api'
import dayjs from '~/plugins/dayjs'
import { ShiftDetail, ShiftDetailLesson, ShiftSummary, StudentShift, TeacherShift } from '~/types/store'
import { LessonDetail } from '~/types/props/shift'
import TheShiftLessonCard from '~/components/organisms/TheShiftLessonCard.vue'

export default defineComponent({
  components: {
    TheShiftLessonCard,
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
    rooms: {
      type: Number,
      default: 0,
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
  },

  setup(props, { emit }: SetupContext) {
    const getTitle = (): string => {
      return `授業登録 ${props.summary?.year}年${props.summary?.month}月`
    }

    const getDay = (date: string): string => {
      return dayjs(date).tz().format('DD (ddd)')
    }

    const getTime = (time: string): string => {
      return dayjs(`2000-01-01 ${time}`).tz().format('HH:mm')
    }

    const getLesson = (shift: ShiftDetailLesson, room: number): LessonDetail | undefined => {
      return props.lessons.find((detail: LessonDetail) => {
        return detail.lesson.shiftId === shift.id && detail.lesson.room === room
      })
    }

    const getLessonTotal = (): number => {
      let total: number = 0
      props.teachers.forEach((teacher: TeacherShift) => {
        total += teacher.lessonTotal
      })
      return total
    }

    const getRemainingLessonTotal = (): number => {
      let total: number = 0
      props.students.forEach((student: StudentShift) => {
        total += student.suggestedClassesTotal
      })
      return total
    }

    const onClickTeacherSubmissions = (teacher: TeacherShift): void => {
      emit('click:show-teacher-submissions', teacher.id)
    }

    const onClickTeacherLessons = (teacher: TeacherShift): void => {
      emit('click:show-teacher-lessons', teacher.id)
    }

    const onClickStudentSubmissions = (student: StudentShift): void => {
      emit('click:show-student-submissions', student.id)
    }

    const onClickStudentLessons = (student: StudentShift): void => {
      emit('click:show-student-lessons', student.id)
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
      getDay,
      getTime,
      getLesson,
      getLessonTotal,
      getRemainingLessonTotal,
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
/* レイアウト */
.shift {
  display: flex;
  flex-direction: column;
  height: calc(100vh - 64px); /* header: 64px */
}

.shift-content {
  flex: 1;
  overflow: auto;
}

.shift-subheader {
  padding: 8px 0;
}

/* 講師/生徒情報一覧 */
.shift-summary {
  padding: 8px 0;
}

.shift-summary table {
  display: block;
  overflow-x: scroll;
  white-space: nowrap;
  -webkit-overflow-scrolling: touch;
  border-collapse: collapse;
  border-spacing: 0;
  width: 100%;
}

.shift-summary th,
.shift-summary td {
  vertical-align: middle;
  padding: 8px 16px;
  border: 1px solid #e5e5e5;
  text-align: center;
}

.shift-summary .fixed {
  position: sticky;
  left: 0;
  background-color: #f5f5f5;
}
.shift-summary .fixed:before {
  content: '';
  position: absolute;
  top: 0px;
  left: -1px;
  width: 100%;
  height: 100%;
  border: 1px solid #f5f5f5;
}

.shift-summary .fixed-right {
  position: sticky;
  right: -1px;
  background-color: #f5f5f5;
}
.shift-summary .fixed:before {
  content: '';
  position: absolute;
  top: 0px;
  right: -1px;
  width: 100%;
  height: 100%;
  border: 1px solid #f5f5f5;
}

/* 授業登録 */
.shift-lessons .col {
  padding: 0;
  margin: 0;
}

.shift-lessons-item {
  margin: 0;
  padding: 0;
}

.shift-lessons-detail {
  margin: 0;
  padding: 0;
}

.shift-lessons-room {
  height: 100%;
}

.shift-lessons-schedule {
  width: 100%;
}
</style>
