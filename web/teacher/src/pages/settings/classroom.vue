<template>
  <v-container>
    <div class="mb-2">
      <p class="text-h5 mb-2">定休日の設定</p>
      <p class="text-body-2">定休日とする曜日にチェックをつけてください。</p>
      <v-checkbox
        v-for="(dateString, i) in dateList"
        :key="i"
        class="ma-0 pt-0"
        :label="dateString"
        :value="dateString"
      />
    </div>
    <div>
      <p class="text-h5 mb-2">ブース数の設定</p>
      <p class="text-body-2">使用するブース数を入力したください。</p>
      <v-select v-model="boothRef" type="number" :items="boothList" />
    </div>
    <div>
      <p class="text-h5 mb-2">コマ設定</p>
      <div class="d-flex align-center my-2">
        <p class="text-subtitle-1 mr-4 mb-0">平日</p>
        <v-btn class="suffix-btn" color="primary" fab elevation="0"><v-icon>mdi-plus</v-icon></v-btn>
      </div>
      <div>
        <div class="stack d-flex align-center">
          <p>1.</p>
          <v-text-field type="time" label="開始時刻" step="300" min="17:00" max="21:30" />
          <span>~</span>
          <v-text-field type="time" label="終了時刻" step="300" min="17:00" max="21:30" />
          <v-btn class="suffix-btn" color="error" fab elevation="0"><v-icon>mdi-minus</v-icon></v-btn>
        </div>
      </div>
      <div class="d-flex align-center my-2">
        <p class="text-subtitle-1 mr-4 mb-0">休日</p>
        <v-btn class="suffix-btn" color="primary" fab elevation="0"><v-icon>mdi-plus</v-icon></v-btn>
      </div>
      <div>
        <div class="stack d-flex align-center">
          <p>1.</p>
          <v-text-field type="time" label="開始時刻" step="300" min="17:00" max="21:30" />
          <span>~</span>
          <v-text-field type="time" label="終了時刻" step="300" min="17:00" max="21:30" />
          <v-btn class="suffix-btn" color="error" fab elevation="0"><v-icon>mdi-minus</v-icon></v-btn>
        </div>
      </div>
    </div>
  </v-container>
</template>

<script lang="ts">
import { defineComponent, ref, Ref } from '@nuxtjs/composition-api'

export default defineComponent({
  setup() {
    const boothRef: Ref<number> = ref<number>(1)

    const dateList: string[] = ['月曜日', '火曜日', '水曜日', '木曜日', '金曜日', '土曜日', '日曜日']
    const boothList: number[] = [1, 2, 3, 4, 5]

    return {
      boothRef,
      dateList,
      boothList,
    }
  },
})
</script>

<style lang="scss" scoped>
.stack {
  gap: var(--space, 1rem);
}

.suffix-btn {
  height: 24px;
  width: 24px;
}
</style>
