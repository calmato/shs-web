<template>
  <div>
    <h2>{{ message }}</h2>
    <v-btn color="primary" @click="onClick">サインイン</v-btn>
  </div>
</template>

<script lang="ts">
import { defineComponent, SetupContext } from '@vue/composition-api'

export default defineComponent({
  props: {
    message: {
      type: String,
      default: '',
    },
  },

  setup(_, { emit }: SetupContext) {
    const onClick = () => {
      emit('click')
    }

    return {
      onClick,
    }
  },
})
</script>
