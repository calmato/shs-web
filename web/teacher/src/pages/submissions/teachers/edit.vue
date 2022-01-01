<template>
  <the-submission-edit :items="submissionsEdit" />
</template>

<script lang="ts">
import { computed, defineComponent, SetupContext } from '@nuxtjs/composition-api'
import TheSubmissionEdit from '~/components/templates/TheSubmissionEdit.vue'
import { SubmissionEditState } from '~/types/props/submission'

export default defineComponent({
  components: { TheSubmissionEdit },

  setup(_, { root }: SetupContext) {
    const store = root.$store
    const submissionsEdit = computed<SubmissionEditState>(() => store.getters['submission/getSubmissionsEdit'])

    return {
      submissionsEdit,
    }
  },
})
</script>
