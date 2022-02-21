<template>
  <v-app-bar app dark clipped-left color="primary" class="px-2 header">
    <v-toolbar-title>SHS Web</v-toolbar-title>
    <v-spacer />
    <v-btn v-show="showMenu" class="hidden-md-and-up" icon @click="onClick">
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
    showMenu: {
      type: Boolean,
      default: true,
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

<style lang="scss" scoped>
.header {
  z-index: 10;
}
</style>
