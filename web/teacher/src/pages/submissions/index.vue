<template>
  <the-submission :items="submissions" @click:edit="handleClickEdit" />
</template>

<script lang="ts">
import { computed, defineComponent, SetupContext } from '@nuxtjs/composition-api'
import TheSubmission from '../../components/templates/TheSubmission.vue'
import { Submission } from '~/types/props/submission'

export default defineComponent({
  components: { TheSubmission },

  setup(_, { root }: SetupContext) {
    const router = root.$router
    const store = root.$store
    const submissions = computed<Submission[]>(() => store.getters['submission/getSubmissions'])

    const handleClickEdit = (actor: string): void => {
      router.push(`submissions/${actor}/edit`)
    }

    return {
      submissions,
      handleClickEdit,
    }
  },
})
</script>
