<template>
  <the-submission :items="submissions" />
</template>

<script lang="ts">
import { computed, defineComponent, SetupContext } from '@nuxtjs/composition-api'
import TheSubmission from '../../components/templates/TheSubmission.vue'
import { Submission } from '~/types/props/submission'

export default defineComponent({
  components: { TheSubmission },

  setup(_, { root }: SetupContext) {
    const store = root.$store
    const submissions = computed<Submission[]>(() => store.getters['submission/getSubmissions'])

    return {
      submissions,
    }
  },
})
</script>
