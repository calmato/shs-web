<template>
  <v-app>
    <v-main>
      <v-row column align-center>
        <v-col class="d-flex flex-column text-center px-8">
          <v-img :src="getErrorImagePath()" contain max-height="400" />
          <h2 class="my-4">{{ getErrorMessage() }}</h2>
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

  props: {
    error: {
      type: Object,
      default: () => null,
    },
  },

  setup(props, { root }: SetupContext) {
    const router = root.$router

    const isNotFound = (): boolean => {
      return props.error?.statusCode === 404
    }

    const getErrorMessage = (): string => {
      if (isNotFound()) {
        return '指定されたページは存在しません'
      } else {
        return '不明なエラーが発生しました'
      }
    }

    const getErrorImagePath = (): string => {
      return isNotFound() ? '/error-404.png' : '/error-500.png'
    }

    const handleClick = () => {
      router.push('/')
    }

    return {
      getErrorMessage,
      getErrorImagePath,
      handleClick,
    }
  },
})
</script>
