<template>
  <the-user-top :loading="loading" :teachers="teachers" :students="students" />
</template>

<script lang="ts">
import { computed, defineComponent, ref, SetupContext } from '@nuxtjs/composition-api'
import TheUserTop from '~/components/templates/TheUserTop.vue'

export default defineComponent({
  components: {
    TheUserTop,
  },

  setup(_, { root }: SetupContext) {
    const store = root.$store

    const loading = ref<boolean>(false)

    const teachers = computed(() => store.getters['user/getTeachers'])
    const students = computed(() => store.getters['user/getStudents'])

    return {
      loading,
      teachers,
      students,
    }
  },
})
</script>
