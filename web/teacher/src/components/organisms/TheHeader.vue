<template>
  <v-app-bar app dark color="primary" class="px-2">
    <v-toolbar-title>SHS Web</v-toolbar-title>
    <v-spacer />
    <v-btn icon @click="onClick">
      <v-icon v-if="overlay">mdi-close</v-icon>
      <v-icon v-else>mdi-menu</v-icon>
    </v-btn>
  </v-app-bar>
</template>

<script lang="ts">
import { defineComponent, SetupContext } from '@nuxtjs/composition-api'

export default defineComponent({
  props: {
    overlay: {
      type: Boolean,
      default: false,
    },
  },

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
