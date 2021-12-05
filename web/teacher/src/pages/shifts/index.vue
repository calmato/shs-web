<template>
  <div>
    <!-- PCレイアウト -->
    <shift-top class="hidden-sm-and-down" />
    <!-- スマホレイアウト -->
    <v-row class="hidden-md-and-up">
      <v-col class="d-flex flex-column text-center px-8">
        <v-img src="/exception.png" />
        <h2 class="my-4">スマホレイアウトに<br />対応していません</h2>
        <v-btn @click="handleClickTop">トップへ戻る</v-btn>
      </v-col>
    </v-row>
  </div>
</template>

<script lang="ts">
import { defineComponent, SetupContext } from '@nuxtjs/composition-api'
import ShiftTop from '~/components/templates/ShiftTop.vue'

export default defineComponent({
  components: {
    ShiftTop,
  },

  setup(_, { root }: SetupContext) {
    const router = root.$router

    const handleClickTop = (): void => {
      router.push('/')
    }

    return {
      handleClickTop,
    }
  },
})
</script>
