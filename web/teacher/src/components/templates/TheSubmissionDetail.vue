<template>
  <v-container>
    <div class="text-h6 my-2">{{ summary.year }}年 {{ summary.month }}月</div>
    <v-divider />
    <v-container class="d-flex">
      <v-row>
        <v-col cols="3">授業日</v-col>
        <v-col cols="9">授業希望</v-col>
      </v-row>
    </v-container>
    <v-divider />
    <the-submission-list :shifts="shifts" />
  </v-container>
</template>

<script lang="ts">
import { defineComponent, PropType } from '@nuxtjs/composition-api'
import TheSubmissionList from '../organisms/TheSubmissionList.vue'
import { TeacherShiftDetail, TeacherShiftSummary } from '~/types/store'

export default defineComponent({
  components: {
    TheSubmissionList,
  },

  props: {
    summary: {
      type: Object as PropType<TeacherShiftSummary>,
      default: () => {},
    },
    shifts: {
      type: Array as PropType<TeacherShiftDetail[]>,
      default: () => [],
    },
  },
})
</script>
