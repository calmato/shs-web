<template>
  <v-container class="pl-6">
    <div class="text-h6">{{ items.summary.month }}月</div>
    <the-submission-list />
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
