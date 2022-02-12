<template>
  <v-container>
    <p class="text-h5 mb-2">授業希望のカスタム設定</p>
    <div class="d-flex align-center my-2">
      <p class="text-subtitle-1 mr-4 mb-0">科目とコマ数を指定</p>
      <v-tooltip right>
        <template #activator="{ on, attrs }">
          <v-btn class="suffix-btn" color="primary" fab elevation="0" v-bind="attrs" v-on="on">
            <v-icon>mdi-plus</v-icon>
          </v-btn>
        </template>
        <span>コマを増やす</span>
      </v-tooltip>
    </div>
    <div class="d-flex h-stack align-center">
      <v-select chips label="科目" />
      <v-text-field type="number" label="コマ数" height="42" />
      <v-tooltip right>
        <template #activator="{ on, attrs }">
          <v-btn class="suffix-btn" color="error" fab elevation="0" v-bind="attrs" v-on="on">
            <v-icon>mdi-minus</v-icon>
          </v-btn>
        </template>
        <span>コマを削除する</span>
      </v-tooltip>
    </div>

    <div class="my-4">
      <v-row v-for="date in dateList" :key="date" class="border ma-0" align="center">
        <v-col cols="2" class="align-center">
          <p class="mb-0">{{ date }}</p>
        </v-col>
        <v-col cols="10" align-self="end">
          <v-chip-group active-class="primary--text" column class="chips">
            <v-chip v-for="time in timeList" :key="time" small>
              {{ time }}
            </v-chip>
          </v-chip-group>
        </v-col>
      </v-row>
    </div>
    <div class="d-flex mt-4">
      <v-spacer />
      <v-btn color="primary">保存</v-btn>
    </div>
  </v-container>
</template>

<script lang="ts">
import { defineComponent } from '@vue/composition-api'

export default defineComponent({
  setup() {
    const dateList: string[] = ['月', '火', '水', '木', '金', '土', '日']
    const timeList: string[] = ['17:00~18:30', '18:30~20:00', '20:00~21:30']
    return { dateList, timeList }
  },
})
</script>

<style lang="scss" scoped>
.suffix-btn {
  height: 24px;
  width: 24px;
}

.h-stack {
  gap: 1rem;
}

::v-deep .chips > .v-slide-group__wrapper > .v-slide-group__content {
  flex-direction: row-reverse;
}

.border {
  border-top: 1px solid #e0e0e0;
  border-bottom: 1px solid #e0e0e0;
}

.border + .border {
  border-top: none;
}
</style>
