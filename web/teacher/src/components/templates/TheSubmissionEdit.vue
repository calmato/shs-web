<template>
  <v-container>
    <div class="text-h6 mt-2">{{ items.summary.year }}年 {{ items.summary.month }}月</div>
    <v-divider class="mt-2" />
    <v-container class="d-flex">
      <v-row>
        <v-col cols="2">授業日</v-col>
        <v-col cols="10">授業希望</v-col>
      </v-row>
    </v-container>
    <v-divider class="mt-1" />
    <the-submission-list :items="items" />
  </v-container>
</template>

<script lang="ts">
import { defineComponent } from '@vue/composition-api'
import TheSubmissionList from '../organisms/TheSubmissionList.vue'
import { SubmissionEditState } from '~/types/props/submission'

export default defineComponent({
  components: { TheSubmissionList },

  props: {
    items: {
      type: Object as PropType<SubmissionEditState>,
      default: () => ({}),
    },
  },

  setup() {},
})
</script>
