<template>
  <v-app-bar app dark color="primary" class="px-2">
    <v-toolbar-title>SHS Web</v-toolbar-title>
    <v-spacer />
    <v-app-bar-nav-icon @click="onClick" />
  </v-app-bar>
</template>

<script lang="ts">
import { defineComponent, SetupContext } from '@nuxtjs/composition-api'

export default defineComponent({
  setup(_, { emit }: SetupContext) {
    const onClick = (): void => {
      emit('click')
    }

    return {
      onClick,
    }
  },
})
</script>
