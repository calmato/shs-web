<template>
  <the-submission-edit />
</template>

<script lang="ts">
import { computed, defineComponent } from '@nuxtjs/composition-api'
import TheSubmissionEdit from '~/components/templates/TheSubmissionEdit.vue'
import { SubmissionEditState } from '~/types/props/submission'

export default defineComponent({
  components: {
    TheSubmissionEdit,
  },

  setup(_) {
    const submissionsEdit = computed<SubmissionEditState>(() => store.getters['submission/getSubmissionsEdit'])
    console.log(submissionsEdit)
  },
})
</script>
