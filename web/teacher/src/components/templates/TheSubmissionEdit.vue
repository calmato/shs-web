<template>
  <v-container class="pl-6">
    <div class="text-h6">2月</div>
    <the-submission-list />
  </v-container>
</template>

<script>
import { defineComponent } from '@vue/composition-api'
import TheSubmissionList from '../organisms/TheSubmissionList.vue'

export default defineComponent({
  components: { TheSubmissionList },
  setup() {},
})
</script>
