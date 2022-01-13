<template>
  <v-container class="shift">
    <section class="shift-header">
      <!-- 講師情報一覧 -->
      <div class="shift-summary">
        <table>
          <tr>
            <th class="fixed">講師名</th>
            <td v-for="teacher in teachers" :key="teacher.id">
              {{ teacher.name }}
            </td>
            <th class="fixed-right">合計</th>
          </tr>
          <tr>
            <th class="fixed">担当授業数</th>
            <td v-for="teacher in teachers" :key="teacher.id">
              {{ teacher.lessonTotal }}
            </td>
            <td class="fixed-right">0</td>
          </tr>
        </table>
      </div>
      <!-- 生徒情報一覧 -->
      <div class="shift-summary">
        <table>
          <tr>
            <th class="fixed">生徒名</th>
            <th class="fixed-right">残り</th>
          </tr>
          <tr>
            <th class="fixed">残り授業数</th>
            <td class="fixed-right">0</td>
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
          <v-btn color="primary">授業を確定する</v-btn>
        </v-col>
      </v-row>
    </section>
    <section class="shift-content">
      <!-- 授業タイムテーブル -->
      <div class="shift-lessons">
        <v-container v-for="shift in details" :key="shift.date" tag="v-row" class="shift-lessons-item">
          <v-col cols="1" align="center">{{ getDay(shift.date) }}</v-col>
          <v-col cols="11" class="d-flex flex-column" align="center">
            <v-container
              v-for="num in rooms"
              :key="`${shift.date}-${num}`"
              justify="start"
              class="d-flex align-stretch shift-lessons-detail"
            >
              <v-card tile class="col col-1 shift-lesson-room">
                <v-card-text>ブース{{ num }}</v-card-text>
              </v-card>
              <v-card v-for="lesson in shift.lessons" :key="lesson.id" tile class="shift-lessons-schedule">
                <v-card-text>
                  <h4>{{ getTime(lesson.startTime) }} ~ {{ getTime(lesson.endTime) }}</h4>
                  <v-icon color="primary">mdi-pencil</v-icon>
                </v-card-text>
              </v-card>
            </v-container>
          </v-col>
        </v-container>
      </div>
    </section>
  </v-container>
</template>

<script lang="ts">
import { defineComponent, PropType } from '@nuxtjs/composition-api'
import dayjs from '~/plugins/dayjs'
import { ShiftDetail, ShiftSummary, TeacherShift } from '~/types/store'

export default defineComponent({
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
  },

  setup(props) {
    const getTitle = (): string => {
      return `授業登録 ${props.summary?.year}年${props.summary?.month}月`
    }

    const getDay = (date: string): string => {
      return dayjs(date).tz().format('DD (ddd)')
    }

    const getTime = (time: string): string => {
      return dayjs(`2000-01-01 ${time}`).tz().format('HH:mm')
    }

    return {
      getTitle,
      getDay,
      getTime,
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

.shift-header {
  background: #fafafa;
}

.shift-content {
  flex: 1;
  overflow: auto;
  background: var(--v-primary-base);
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
  border: 1px solid #e5e5e5;
}

.shift-lessons-detail {
  margin: 0;
  padding: 0;
}

.shift-lessons-detail .v-card {
  box-shadow: 0;
  border: 1px solid #f5f5f5;
}

.shift-lessons-room {
  height: 100%;
}

.shift-lessons-schedule {
  width: 100%;
}
</style>
