<template>
  <top @click="handleClick" />
</template>

<script lang="ts">
import { defineComponent, SetupContext } from '@nuxtjs/composition-api'
import Top from '~/components/templates/Top.vue'

export default defineComponent({
  components: {
    Top,
  },

  setup(_, { root }: SetupContext) {
    const router = root.$router

    const handleClick = () => {
      router.push('/signin')
    }

    return {
      handleClick,
    }
  },
})
</script>
