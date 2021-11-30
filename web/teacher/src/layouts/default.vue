<template>
  <v-app>
    <the-header @click="onClick" />
    <v-main>
      <nuxt />
    </v-main>
  </v-app>
</template>

<script lang="ts">
import { defineComponent } from '@nuxtjs/composition-api'
import TheHeader from '~/components/organisms/TheHeader.vue'

export default defineComponent({
  components: {
    TheHeader,
  },

  setup() {
    const onClick = (): void => {
      console.log('debug', 'header')
    }

    return {
      onClick,
    }
  },
})
</script>
