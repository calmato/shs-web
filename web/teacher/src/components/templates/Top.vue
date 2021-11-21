<template>
  <div>
    <h2>講師用トップページ</h2>
    <v-btn color="secondary" @click="onClick">サインアウト</v-btn>
  </div>
</template>

<script lang="ts">
import { defineComponent, SetupContext } from '@nuxtjs/composition-api'

export default defineComponent({
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

