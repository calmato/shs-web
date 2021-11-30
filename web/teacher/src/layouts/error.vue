<template>
  <v-app>
    <v-content>
      <v-container>
        <v-layout column align-center>
          <h1>不明なエラーが発生しました</h1>
          <v-img src="/error-500.png" />

          <v-btn @click="handleClick">ホームへ戻る</v-btn>
        </v-layout>
      </v-container>
    </v-content>
  </v-app>
</template>

<script lang="ts">
import { defineComponent, SetupContext } from '@nuxtjs/composition-api'

export default defineComponent({
  layout: 'empty',

  setup(_, { root }: SetupContext) {
    const router = root.$router

    const handleClick = () => {
      router.push('/')
    }

    return {
      handleClick,
    }
  },
})
</script>
