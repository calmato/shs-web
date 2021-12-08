<template>
  <v-app>
    <v-main>
      <v-row column align-center>
        <v-col class="d-flex flex-column text-center px-8">
          <v-img src="/error-500.png" />
          <h2 class="my-4">不明なエラーが発生しました</h2>
          <v-btn @click="handleClick">ホームへ戻る</v-btn>
        </v-col>
      </v-row>
    </v-main>
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
