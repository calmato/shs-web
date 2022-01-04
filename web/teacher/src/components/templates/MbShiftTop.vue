<template>
  <v-container>
    <v-row>
      <v-col class="d-flex flex-column text-center px-8">
        <v-img src="/exception.png" contain max-height="400" />
        <h2 class="my-4">スマホレイアウトに<br />対応していません</h2>
        <v-btn @click="onClick">トップへ戻る</v-btn>
      </v-col>
    </v-row>
  </v-container>
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
